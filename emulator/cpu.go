package emulator

import (
	"fmt"
	"sync"
)

var (
	prgRomPageSize int = 16 * 1024 // プログラムROMのページサイズ
	chrRomPageSize int = 8 * 1024  // キャラクタROMのページサイズ
	maxHistory         = 64
)

// CPURegister CPUのレジスタです。
type cpuRegister struct {
	A  byte   // アキュムレータ
	X  byte   // インデックスレジスタ
	Y  byte   // インデックスレジスタ
	S  byte   // スタックポインタ
	P  byte   // ステータスレジスタ
	PC uint16 // プログラムカウンタ
}

// CPU Central Processing Unit
type CPU struct {
	Reg     cpuRegister
	RAM     [0x10000]byte
	PPU     PPU
	joypad1 Joypad
	mutex   sync.Mutex
	history []string
}

// InitReg レジスタの初期化
func (cpu *CPU) InitReg() {
	cpu.Reg.S = 0xfd
	cpu.Reg.P = 0x34

	lower := uint16(cpu.FetchMemory8(0xfffc))
	upper := uint16(cpu.FetchMemory8(0xfffd))
	cpu.Reg.PC = (upper << 8) | lower
}

// LoadROM ROMのバイトデータからプログラムROMとページROMを取り出してRAMにロードする
func (cpu *CPU) LoadROM(rom []byte) {
	// for i := 0; i < 16; i++ {
	// 	fmt.Printf("iNES[%d]: 0x%x\n", i, rom[i])
	// }

	// ミラーをセットする
	mirrorFlag := rom[6]
	cpu.PPU.mirror = mirrorFlag > 0

	prgAddr := 0x0010
	prgPage := rom[4]

	chrAddr := prgAddr + int(prgPage)*prgRomPageSize
	chrPage := rom[5]

	prgBytes := rom[prgAddr : prgAddr+int(prgPage)*prgRomPageSize]
	chrBytes := rom[chrAddr : chrAddr+int(chrPage)*chrRomPageSize]

	// プログラムROMを0x8000~に配置
	for i := 0; i < len(prgBytes); i++ {
		cpu.RAM[0x8000+i] = prgBytes[i]

		// ページサイズが1のときは0x8000-0xbfffを0xc000-0xffffにミラーする
		if prgPage == 1 {
			cpu.RAM[0x8000+i+0x4000] = prgBytes[i]
		}
	}

	// キャラクタROMをPPUの0x0000~に配置
	for i := 0; i < len(chrBytes); i++ {
		cpu.PPU.RAM[i] = chrBytes[i]
	}

	// 割り込みベクタデバッグ
	// for i := 0x0a; i <= 0x0f; i++ {
	// 	addr := 0xfff0 + i
	// 	fmt.Printf("IRQVector[0x%x]: %x\n", addr, cpu.RAM[addr])
	// }
}

// Cycle CPUのメインサイクル
func (cpu *CPU) exec() {
	cpu.mutex.Lock()
	prevPC := cpu.Reg.PC
	opcode := cpu.FetchCode8(0)
	instruction, addressing := instructions[opcode][0], instructions[opcode][1]

	var addr uint16
	switch addressing {
	case "impl":
		addr = cpu.ImpliedAddressing()
	case "A":
		addr = cpu.AccumulatorAddressing()
	case "#":
		addr = cpu.ImmediateAddressing()
	case "zpg":
		addr = cpu.ZeroPageAddressing()
	case "zpg,X":
		addr = cpu.ZeroPageXAddressing()
	case "zpg,Y":
		addr = cpu.ZeroPageYAddressing()
	case "abs":
		addr = cpu.AbsoluteAddressing()
	case "abs,X":
		addr = cpu.AbsoluteXAddressing()
	case "abs,Y":
		addr = cpu.AbsoluteYAddressing()
	case "rel":
		addr = cpu.RelativeAddressing()
	case "X,ind":
		addr = cpu.IndexedIndirectAddressing()
	case "ind,Y":
		addr = cpu.IndirectIndexedAddressing()
	case "ind":
		addr = cpu.AbsoluteIndirectAddressing()
	default:
		cpu.writeHistory()

		panicMsg := fmt.Sprintf("addressing is not found => { eip: %x, opcode: %x }\n", cpu.Reg.PC, opcode)
		panic(panicMsg)
	}

	// fmt.Printf("eip: 0x%x %s:%s:%x\n", prevPC, instruction, addressing, addr)

	switch instruction {
	case "ADC":
		cpu.ADC(addr)
	case "SBC":
		cpu.SBC(addr)
	case "AND":
		cpu.AND(addr)
	case "ORA":
		cpu.ORA(addr)
	case "EOR":
		cpu.EOR(addr)
	case "ASL":
		cpu.ASL(addr)
	case "LSR":
		cpu.LSR(addr)
	case "ROL":
		cpu.ROL(addr)
	case "ROR":
		cpu.ROR(addr)
	case "BCC":
		cpu.BCC(addr)
	case "BCS":
		cpu.BCS(addr)
	case "BEQ":
		cpu.BEQ(addr)
	case "BNE":
		cpu.BNE(addr)
	case "BVC":
		cpu.BVC(addr)
	case "BVS":
		cpu.BVS(addr)
	case "BPL":
		cpu.BPL(addr)
	case "BMI":
		cpu.BMI(addr)
	case "BIT":
		cpu.BIT(addr)
	case "JMP":
		cpu.JMP(addr)
	case "JSR":
		cpu.JSR(addr)
	case "RTS":
		cpu.RTS(addr)
	case "BRK":
		cpu.BRK(addr)
	case "RTI":
		cpu.RTI(addr)
	case "CMP":
		cpu.CMP(addr)
	case "CPX":
		cpu.CPX(addr)
	case "CPY":
		cpu.CPY(addr)
	case "INC":
		cpu.INC(addr)
	case "DEC":
		cpu.DEC(addr)
	case "INX":
		cpu.INX(addr)
	case "DEX":
		cpu.DEX(addr)
	case "INY":
		cpu.INY(addr)
	case "DEY":
		cpu.DEY(addr)
	case "CLC":
		cpu.CLC(addr)
	case "SEC":
		cpu.SEC(addr)
	case "CLI":
		cpu.CLI(addr)
	case "SEI":
		cpu.SEI(addr)
	case "CLD":
		cpu.CLD(addr)
	case "SED":
		cpu.SED(addr)
	case "CLV":
		cpu.CLV(addr)
	case "LDA":
		cpu.LDA(addr)
	case "LDX":
		cpu.LDX(addr)
	case "LDY":
		cpu.LDY(addr)
	case "STA":
		cpu.STA(addr)
	case "STX":
		cpu.STX(addr)
	case "STY":
		cpu.STY(addr)
	case "TAX":
		cpu.TAX(addr)
	case "TAY":
		cpu.TAY(addr)
	case "TXA":
		cpu.TXA(addr)
	case "TYA":
		cpu.TYA(addr)
	case "TSX":
		cpu.TSX(addr)
	case "TXS":
		cpu.TXS(addr)
	case "PHA":
		cpu.PHA(addr)
	case "PLA":
		cpu.PLA(addr)
	case "PHP":
		cpu.PHP(addr)
	case "PLP":
		cpu.PLP(addr)
	case "NOP":
		cpu.NOP(addr)
	default:
		cpu.writeHistory()

		panicMsg := fmt.Sprintf("instruction is not found: %d=0x%x\n", opcode, opcode)
		panic(panicMsg)
	}

	cpu.pushHistory(prevPC, opcode, instruction, addressing, addr)

	cpu.mutex.Unlock()
}

// FetchCode8 メモリから次の命令をフェッチする PCのインクリメントはしない
func (cpu *CPU) FetchCode8(index uint) byte {
	code := cpu.RAM[(uint)(cpu.Reg.PC)+index]
	return code
}

// getVRAMDelta CPUのVRAMアクセス時のポインタの増加量を返す
func (cpu *CPU) getVRAMDelta() (delta uint16) {
	value := cpu.RAM[0x2000]
	if (value & 0x04) > 0 {
		return 32
	}
	return 1
}

// pushHistory CPUのログを追加する
func (cpu *CPU) pushHistory(eip uint16, opcode byte, instruction, addressing string, addr uint16) {
	log := fmt.Sprintf("eip:0x%x   opcode:%x   %s:%s 0x%x", eip, opcode, instruction, addressing, addr)
	cpu.history = append(cpu.history, log)

	if len(cpu.history) > maxHistory {
		cpu.history = cpu.history[1:]
	}
}

// writeHistory CPUのログを書き出す
func (cpu *CPU) writeHistory() {
	for i, log := range cpu.history {
		fmt.Printf("%d: %s\n", i, log)
	}
}
