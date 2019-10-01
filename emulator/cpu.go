package emulator

import (
	"fmt"
	"time"
)

var (
	prgRomPageSize int = 16 * 1024 // プログラムROMのページサイズ
	chrRomPageSize int = 8 * 1024  // キャラクタROMのページサイズ
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
	Reg cpuRegister
	RAM [0x10000]byte
	PPU PPU
}

// InitIRQVector 割り込みベクタの初期化
func (cpu *CPU) InitIRQVector() {
	cpu.RAM[0xfffc] = 0x00
	cpu.RAM[0xfffd] = 0x80
}

// InitReg レジスタの初期化
func (cpu *CPU) InitReg() {
	cpu.Reg.S = 0x00 // スタックは0x100から下に伸びる
	cpu.Reg.PC = 0x8000
}

// LoadROM ROMのバイトデータからプログラムROMとページROMを取り出してRAMにロードする
func (cpu *CPU) LoadROM(rom []byte) {
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
	}

	// キャラクタROMをPPUの0x0000~に配置
	for i := 0; i < len(chrBytes); i++ {
		cpu.PPU.RAM[i] = chrBytes[i]
	}
}

// MainLoop CPUのメインサイクル
func (cpu *CPU) MainLoop() {
	for {
		time.Sleep(time.Millisecond * 1)

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
		case "Ind":
			addr = cpu.AbsoluteIndirectAddressing()
		default:
			fmt.Printf("addressing is not found: %d\n", opcode)
		}

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
			fmt.Printf("instruction is not found: %d\n", opcode)
		}
	}
}

// FetchCode8 メモリから次の命令をフェッチする PCのインクリメントはしない
func (cpu *CPU) FetchCode8(index uint) byte {
	code := cpu.RAM[(uint)(cpu.Reg.PC)+index]
	return code
}

// FetchMemory8 引数で指定したアドレスから値を取得する
func (cpu *CPU) FetchMemory8(addr uint16) byte {
	value := cpu.RAM[addr]
	return value
}

// SetMemory8 引数で指定したアドレスにvalueを書き込む
func (cpu *CPU) SetMemory8(addr uint16, value byte) {
	cpu.RAM[addr] = value
}

// FlagN Nフラグを立てるか判定する
func (cpu *CPU) FlagN(b byte) {
	if (b & 0x80) != 0 {
		cpu.Reg.P = cpu.Reg.P | 0x80 // 0b1000_0000
	} else {
		cpu.Reg.P = cpu.Reg.P & 0x7f // 0b0111_1111
	}
}

// FlagV Vフラグを立てるか判定する
func (cpu *CPU) FlagV(b0, b1 byte, u16 uint16) {
	if ((b0>>7)^(b1>>7) != 0) && (uint16(b1) != u16) {
		cpu.Reg.P = cpu.Reg.P | 0x40 // 0b0100_0000
	} else {
		cpu.Reg.P = cpu.Reg.P & 0xbf // 0b1011_1111
	}
}

// FlagZ Zフラグを立てるか判定する
func (cpu *CPU) FlagZ(b byte) {
	if b == 0 {
		cpu.Reg.P = cpu.Reg.P | 0x02 // 0b0000_0010
	} else {
		cpu.Reg.P = cpu.Reg.P & 0xfd // 0b1111_1101
	}
}

// FlagC Cフラグを立てるか判定する
func (cpu *CPU) FlagC(u16 uint16) {
	if (u16 >> 8) != 0 {
		cpu.Reg.P = cpu.Reg.P | 0x01 // 0b0000_0001
	} else {
		cpu.Reg.P = cpu.Reg.P & 0xfe // 0b1111_1110
	}
}
