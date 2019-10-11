package emulator

import (
	"sync"
	"time"
)

// ==============================================  演算  ==============================================

// ADC Add M to A with C (A + M + C -> A)
func (cpu *CPU) ADC(addr uint16) {
	cFlag := cpu.Reg.P & 0x01
	aFlag := cpu.Reg.A
	value := (cpu.Reg.A + cpu.FetchMemory8(addr) + cFlag) & (0xff)                // キャリーオーバー対策のため
	value16 := uint16(cpu.Reg.A) + uint16(cpu.FetchMemory8(addr)) + uint16(cFlag) // Cフラグ判定のため
	cpu.Reg.A = value

	cpu.FlagN(value)
	cpu.FlagV(aFlag, value, value16)
	cpu.FlagZ(value)
	cpu.FlagC("ADC", value16)
}

// SBC Subtract M from A with C (A - M - not C -> A)
func (cpu *CPU) SBC(addr uint16) {
	notCFlag := ^(cpu.Reg.P) & 0x01
	aFlag := cpu.Reg.A
	value := (cpu.Reg.A - cpu.FetchMemory8(addr) - notCFlag) & (0xff)                // キャリーオーバー対策のため
	value16 := uint16(cpu.Reg.A) - uint16(cpu.FetchMemory8(addr)) - uint16(notCFlag) // Cフラグ判定のため
	cpu.Reg.A = value

	cpu.FlagN(value)
	cpu.FlagV(aFlag, value, value16)
	cpu.FlagZ(value)
	cpu.FlagC("SBC", value16)
}

// ============================================== 論理演算 ==============================================

// AND "AND" M with A (A and M -> A)
func (cpu *CPU) AND(addr uint16) {
	value := cpu.Reg.A & cpu.FetchMemory8(addr)
	cpu.Reg.A = value

	cpu.FlagN(value)
	cpu.FlagZ(value)
}

// ORA "OR" M with A (A or M -> A)
func (cpu *CPU) ORA(addr uint16) {
	value := cpu.Reg.A | cpu.FetchMemory8(addr)
	cpu.Reg.A = value

	cpu.FlagN(value)
	cpu.FlagZ(value)
}

// EOR "Exclusive-OR" M with A (A eor M -> A)
func (cpu *CPU) EOR(addr uint16) {
	value := cpu.Reg.A ^ cpu.FetchMemory8(addr)
	cpu.Reg.A = value

	cpu.FlagN(value)
	cpu.FlagZ(value)
}

// ============================================== シフト ==============================================

// ASL Arithmetic shift left one bit
func (cpu *CPU) ASL(addr uint16) {
	if addr == null {
		// ASLAccumulator
		cpu.Reg.P = cpu.Reg.P | ((cpu.Reg.A & 0x80) >> 7) // Aのbit7をcにセット
		cpu.Reg.A = cpu.Reg.A << 1
		cpu.Reg.A = cpu.Reg.A | 0 // Aのbit0に0をセット
		cpu.FlagN(cpu.Reg.A)
		cpu.FlagZ(cpu.Reg.A)
	} else {
		value := cpu.FetchMemory8(addr)
		cpu.Reg.P = cpu.Reg.P | ((value & 0x80) >> 7) // valueのbit7をcにセット
		value = value << 1
		value = value | 0 // valueのbit0に0をセット
		cpu.SetMemory8(addr, value)
		cpu.FlagN(value)
		cpu.FlagZ(value)
	}
}

// LSR Logical shift right one bit
func (cpu *CPU) LSR(addr uint16) {
	if addr == null {
		// Aのbit0をcにセット
		if cpu.Reg.A&0x01 > 0 {
			cpu.Reg.P = cpu.Reg.P | 0x01
		} else {
			cpu.Reg.P = cpu.Reg.P & 0xfe
		}

		cpu.Reg.A = cpu.Reg.A >> 1
		cpu.FlagN(cpu.Reg.A)
		cpu.FlagZ(cpu.Reg.A)
	} else {
		value := cpu.FetchMemory8(addr)

		// valueのbit0をcにセット
		if value&0x01 > 0 {
			cpu.Reg.P = cpu.Reg.P | 0x01
		} else {
			cpu.Reg.P = cpu.Reg.P & 0xfe
		}

		value = value >> 1
		cpu.SetMemory8(addr, value)
		cpu.FlagN(value)
		cpu.FlagZ(value)
	}
}

// ROL Rotate left one bit
func (cpu *CPU) ROL(addr uint16) {
	if addr == null {
		cFlag := cpu.Reg.P & 0x01
		cpu.Reg.P = cpu.Reg.P | ((cpu.Reg.A & 0x80) >> 7) // Aのbit7をcにセット
		cpu.Reg.A = cpu.Reg.A << 1
		cpu.Reg.A = cpu.Reg.A | cFlag // Aのbit0にcをセット
		cpu.FlagN(cpu.Reg.A)
		cpu.FlagZ(cpu.Reg.A)
	} else {
		value := cpu.FetchMemory8(addr)
		cFlag := cpu.Reg.P & 0x01
		cpu.Reg.P = cpu.Reg.P | ((value & 0x80) >> 7) // valueのbit7をcにセット
		value = value << 1
		value = value | cFlag // valueのbit0にcをセット
		cpu.SetMemory8(addr, value)
		cpu.FlagN(value)
		cpu.FlagZ(value)
	}
}

// ROR Rotate right one bit
func (cpu *CPU) ROR(addr uint16) {
	if addr == null {
		cFlag := cpu.Reg.P & 0x01
		cpu.Reg.P = cpu.Reg.P | (cpu.Reg.A & 0x01) // valueのbit0をcにセット
		cpu.Reg.A = cpu.Reg.A >> 1
		cpu.Reg.A = cpu.Reg.A | (cFlag << 7) // valueのbit7にcをセット
		cpu.FlagN(cpu.Reg.A)
		cpu.FlagZ(cpu.Reg.A)
	} else {
		value := cpu.FetchMemory8(addr)
		cFlag := cpu.Reg.P & 0x01
		cpu.Reg.P = cpu.Reg.P | (value & 0x01) // valueのbit0をcにセット
		value = value >> 1
		value = value | (cFlag << 7) // valueのbit7にcをセット
		cpu.SetMemory8(addr, value)
		cpu.FlagN(value)
		cpu.FlagZ(value)
	}
}

// ============================================== 条件分岐 ==============================================

// BCC Branch on C clear
func (cpu *CPU) BCC(addr uint16) {
	cFlag := uint8(cpu.Reg.P & 0x01)
	if cFlag == 0 {
		cpu.Reg.PC = addr
	}
}

// BCS Branch on C set
func (cpu *CPU) BCS(addr uint16) {
	cFlag := uint8(cpu.Reg.P & 0x01)
	if cFlag > 0 {
		cpu.Reg.PC = addr
	}
}

// BEQ Branch on Z set
func (cpu *CPU) BEQ(addr uint16) {
	zFlag := uint8(cpu.Reg.P & 0x02)
	if zFlag > 0 {
		cpu.Reg.PC = addr
	}
}

// BNE Branch on Z clear
func (cpu *CPU) BNE(addr uint16) {
	zFlag := uint8(cpu.Reg.P & 0x02) // 0b0000_0010
	if zFlag == 0 {
		cpu.Reg.PC = addr
	}
}

// BVC Branch on V clear
func (cpu *CPU) BVC(addr uint16) {
	vFlag := uint8(cpu.Reg.P & 0x40)
	if vFlag == 0 {
		cpu.Reg.PC = addr
	}
}

// BVS Branch on V set
func (cpu *CPU) BVS(addr uint16) {
	vFlag := uint8(cpu.Reg.P & 0x40)
	if vFlag > 0 {
		cpu.Reg.PC = addr
	}
}

// BPL Branch on N clear
func (cpu *CPU) BPL(addr uint16) {
	nFlag := uint8(cpu.Reg.P & 0x80)
	if nFlag == 0 {
		cpu.Reg.PC = addr
	}
}

// BMI Branch on N set
func (cpu *CPU) BMI(addr uint16) {
	nFlag := uint8(cpu.Reg.P & 0x80)
	if nFlag > 0 {
		cpu.Reg.PC = addr
	}
}

// ============================================== ビットテスト ==============================================

// BIT Test Bits in M with A
func (cpu *CPU) BIT(addr uint16) {
	value := cpu.FetchMemory8(addr)

	// NZフラグを立てる
	cpu.FlagZ(value & cpu.Reg.A)
	cpu.FlagN(value)
	// bit6をVフラグに転送
	if (value & 0x40) != 0 {
		cpu.Reg.P = cpu.Reg.P | 0x40 // 0b0100_0000
	} else {
		cpu.Reg.P = cpu.Reg.P & 0xbf // 0b1011_1111
	}
}

// ============================================== ジャンプ ==============================================

// JMP Jump to new location
func (cpu *CPU) JMP(addr uint16) {
	cpu.Reg.PC = addr
}

// JSR Jump to new location saving return address
func (cpu *CPU) JSR(addr uint16) {
	upper := byte((cpu.Reg.PC - 1) >> 8)
	lower := byte((cpu.Reg.PC - 1))
	cpu.SetMemory8((0x100 + uint16(cpu.Reg.S)), upper)
	cpu.SetMemory8((0x100 + uint16(cpu.Reg.S) - 1), lower)
	cpu.Reg.S -= 2
	cpu.Reg.PC = addr
}

// RTS Return from Subroutine
func (cpu *CPU) RTS(addr uint16) {
	if addr == null {
		lower := uint16(cpu.FetchMemory8((0x100 + uint16(cpu.Reg.S) + 1)))
		cpu.Reg.S++
		upper := uint16(cpu.FetchMemory8((0x100 + uint16(cpu.Reg.S) + 1)))
		cpu.Reg.S++
		cpu.Reg.PC = (upper << 8) | lower
		cpu.Reg.PC++
	}
}

// ============================================== 割り込み ==============================================

// BRK Break Interrupt
func (cpu *CPU) BRK(addr uint16) {
	if addr == null {
		iFlag := cpu.Reg.P & 0x04
		if iFlag == 0 {
			// allow BRK
			cpu.Reg.P = cpu.Reg.P | 0x10 // set B Flag
			cpu.Reg.PC++

			// push PC & SR
			upper0 := byte((cpu.Reg.PC) >> 8)
			lower0 := byte((cpu.Reg.PC))
			cpu.SetMemory8((0x100 + uint16(cpu.Reg.S)), upper0)
			cpu.SetMemory8((0x100 + uint16(cpu.Reg.S) - 1), lower0)
			cpu.SetMemory8((0x100 + uint16(cpu.Reg.S) - 2), cpu.Reg.P)
			cpu.Reg.S -= 3

			cpu.Reg.P = cpu.Reg.P | 0x04 // set I Flag

			upper1 := uint16(cpu.FetchMemory8(0xffff))
			lower1 := uint16(cpu.FetchMemory8(0xfffe))
			cpu.Reg.PC = (upper1 << 8) | lower1
		}
	}
}

// RTI Return from Interrupt
func (cpu *CPU) RTI(addr uint16) {
	if addr == null {
		// ステータスレジスタをpop
		SR := cpu.FetchMemory8((0x100 + uint16(cpu.Reg.S) + 1))
		cpu.Reg.S++
		cpu.Reg.P = SR
		// PCをpop
		lower := uint16(cpu.FetchMemory8((0x100 + uint16(cpu.Reg.S) + 1)))
		cpu.Reg.S++
		upper := uint16(cpu.FetchMemory8((0x100 + uint16(cpu.Reg.S) + 1)))
		cpu.Reg.S++
		cpu.Reg.PC = (upper << 8) | lower // ここでPCにリターンしているかつ割り込みなのでインクリメントは必要ない

		cpu.RAM[0x2000] = cpu.RAM[0x2000] | 0x80 // allow NMI
	}
}

// NMI Non-maskable Interrupt
func (cpu *CPU) NMI(addr uint16) {
	if addr == null && (cpu.RAM[0x2000]>>7) > 0 {
		cpu.RAM[0x2000] = cpu.RAM[0x2000] & 0x7f // block NMI

		cpu.Reg.P = cpu.Reg.P & 0xef // clear B Flag

		// push PC & SR
		lower0 := byte((cpu.Reg.PC))
		upper0 := byte((cpu.Reg.PC) >> 8)
		cpu.SetMemory8((0x100 + uint16(cpu.Reg.S)), upper0)
		cpu.SetMemory8((0x100 + uint16(cpu.Reg.S) - 1), lower0)
		cpu.SetMemory8((0x100 + uint16(cpu.Reg.S) - 2), cpu.Reg.P)
		cpu.Reg.S -= 3

		cpu.Reg.P = cpu.Reg.P | 0x04 // set I Flag

		lower1 := uint16(cpu.FetchMemory8(0xfffa))
		upper1 := uint16(cpu.FetchMemory8(0xfffb))
		cpu.Reg.PC = (upper1 << 8) | lower1
	}
}

// ============================================== 比較 ==============================================

// CMP Compare M and A (A - M)
func (cpu *CPU) CMP(addr uint16) {
	value := cpu.Reg.A - cpu.FetchMemory8(addr)
	value16 := uint16(cpu.Reg.A) - uint16(cpu.FetchMemory8(addr))

	cpu.FlagN(value)
	cpu.FlagZ(value)
	cpu.FlagC("CMP", value16)
}

// CPX Compare M and X (X - M)
func (cpu *CPU) CPX(addr uint16) {
	value := cpu.Reg.X - cpu.FetchMemory8(addr)
	value16 := uint16(cpu.Reg.X) - uint16(cpu.FetchMemory8(addr))

	cpu.FlagN(value)
	cpu.FlagZ(value)
	cpu.FlagC("CPX", value16)
}

// CPY Compare M and Y (Y - M)
func (cpu *CPU) CPY(addr uint16) {
	value := cpu.Reg.Y - cpu.FetchMemory8(addr)
	value16 := uint16(cpu.Reg.Y) - uint16(cpu.FetchMemory8(addr))

	cpu.FlagN(value)
	cpu.FlagZ(value)
	cpu.FlagC("CPY", value16)
}

// =========================================== インクリメント、デクリメント ===========================================

// INC Increment M by one (M + 1 -> M)
func (cpu *CPU) INC(addr uint16) {
	value := cpu.FetchMemory8(addr) + 1
	cpu.SetMemory8(addr, value)

	cpu.FlagN(value)
	cpu.FlagZ(value)
}

// DEC Decrement M by one (M - 1 -> M)
func (cpu *CPU) DEC(addr uint16) {
	value := cpu.FetchMemory8(addr) - 1
	cpu.SetMemory8(addr, value)

	cpu.FlagN(value)
	cpu.FlagZ(value)
}

// INX Increment X by one (X + 1 -> X)
func (cpu *CPU) INX(addr uint16) {
	cpu.Reg.X++
	cpu.FlagN(cpu.Reg.X)
	cpu.FlagZ(cpu.Reg.X)
}

// DEX Decrement X by one (X - 1 -> X)
func (cpu *CPU) DEX(addr uint16) {
	cpu.Reg.X--
	cpu.FlagN(cpu.Reg.X)
	cpu.FlagZ(cpu.Reg.X)
}

// INY Increment Y by one (Y + 1 -> Y)
func (cpu *CPU) INY(addr uint16) {
	cpu.Reg.Y++
	cpu.FlagN(cpu.Reg.Y)
	cpu.FlagZ(cpu.Reg.Y)
}

// DEY Decrement Y by one (Y - 1 -> Y)
func (cpu *CPU) DEY(addr uint16) {
	cpu.Reg.Y--
	cpu.FlagN(cpu.Reg.Y)
	cpu.FlagZ(cpu.Reg.Y)
}

// ============================================ フラグ操作 ============================================

// CLC Clear C Flag
func (cpu *CPU) CLC(addr uint16) {
	if addr == null {
		cpu.Reg.P = cpu.Reg.P & 0xfe
	}
}

// SEC Set C Flag
func (cpu *CPU) SEC(addr uint16) {
	if addr == null {
		cpu.Reg.P = cpu.Reg.P | 0x01
	}
}

// CLI Clear I Flag
func (cpu *CPU) CLI(addr uint16) {
	if addr == null {
		cpu.Reg.P = cpu.Reg.P & 0xfb // 0b11111011
	}
}

// SEI Set I Flag
func (cpu *CPU) SEI(addr uint16) {
	if addr == null {
		cpu.Reg.P = cpu.Reg.P | 0x04
	}
}

// CLD Clear D Flag
func (cpu *CPU) CLD(addr uint16) {
	if addr == null {
		cpu.Reg.P = cpu.Reg.P & 0xfb // 0b1111_1011
	}
}

// SED Set D Flag
func (cpu *CPU) SED(addr uint16) {
	if addr == null {
		cpu.Reg.P = cpu.Reg.P | 0x08
	}
}

// CLV Clear V Flag
func (cpu *CPU) CLV(addr uint16) {
	if addr == null {
		cpu.Reg.P = cpu.Reg.P & 0xbf // 0b1011_1111
	}
}

// ============================================ ロード ============================================

// LDA Load A from M (M -> A)
func (cpu *CPU) LDA(addr uint16) {
	switch addr {
	case 0x2002:
		cpu.Reg.A = cpu.FetchMemory8(addr)
		cpu.PPU.scrollFlag = false
		cpu.clearVBlank()
	case 0x2007:
		cpu.Reg.A = cpu.PPU.RAM[cpu.PPU.ptr]
		cpu.PPU.ptr += cpu.getVRAMDelta()
	case joypad1:
		time.Sleep(time.Nanosecond)
		cpu.Reg.A = cpu.joypad1.cmd[cpu.joypad1.ctr]
		cpu.joypad1.ctr++
	default:
		cpu.Reg.A = cpu.FetchMemory8(addr)
	}

	cpu.FlagN(cpu.Reg.A)
	cpu.FlagZ(cpu.Reg.A)
}

// LDX Load X from M (M -> X)
func (cpu *CPU) LDX(addr uint16) {
	switch addr {
	case 0x2002:
		cpu.Reg.X = cpu.FetchMemory8(addr)
		cpu.PPU.scrollFlag = false
		cpu.clearVBlank()
	case 0x2007:
		cpu.Reg.X = cpu.PPU.RAM[cpu.PPU.ptr]
		cpu.PPU.ptr += cpu.getVRAMDelta()
	case joypad1:
		cpu.Reg.X = cpu.joypad1.cmd[cpu.joypad1.ctr]
		cpu.joypad1.ctr++
	default:
		cpu.Reg.X = cpu.FetchMemory8(addr)
	}

	cpu.FlagN(cpu.Reg.X)
	cpu.FlagZ(cpu.Reg.X)
}

// LDY Load Y from M (M -> Y)
func (cpu *CPU) LDY(addr uint16) {
	switch addr {
	case 0x2002:
		cpu.Reg.Y = cpu.FetchMemory8(addr)
		cpu.PPU.scrollFlag = false
		cpu.clearVBlank()
	case 0x2007:
		cpu.Reg.Y = cpu.PPU.RAM[cpu.PPU.ptr]
		cpu.PPU.ptr += cpu.getVRAMDelta()
	case joypad1:
		cpu.Reg.Y = cpu.joypad1.cmd[cpu.joypad1.ctr]
		cpu.joypad1.ctr++
	default:
		cpu.Reg.Y = cpu.FetchMemory8(addr)
	}

	cpu.FlagN(cpu.Reg.Y)
	cpu.FlagZ(cpu.Reg.Y)
}

// ============================================ ストア ==================================================

// STA Store A to M (A -> M)
func (cpu *CPU) STA(addr uint16) {
	switch addr {
	case 0x2004:
		cpu.PPU.sRAM[cpu.RAM[0x2003]] = cpu.Reg.A
		cpu.RAM[0x2003]++
	case 0x2005:
		if cpu.PPU.scrollFlag {
			cpu.PPU.scroll[1] = cpu.Reg.A
		} else {
			cpu.PPU.scroll[0] = cpu.Reg.A
			cpu.PPU.scrollFlag = true
		}
	case 0x2006:
		cpu.PPU.ptr = (cpu.PPU.ptr<<8 | uint16(cpu.Reg.A))
	case 0x2007:
		cpu.setVRAM(cpu.Reg.A)
	case spriteDMA:
		start := uint16(cpu.Reg.A) << 8
		var wait sync.WaitGroup
		wait.Add(256)
		for i := 0; i < 256; i++ {
			go func(i int) {
				cpu.PPU.sRAM[i] = cpu.FetchMemory8(start + uint16(i))
				wait.Done()
			}(i)
		}
		wait.Wait()
	case joypad1:
		cpu.joypad1.ctr = 0
		for i := 0; i < 8; i++ {
			cpu.joypad1.cmd[i] = 0
		}
	}
	cpu.SetMemory8(addr, cpu.Reg.A)
}

// STX Store X to M (X -> M)
func (cpu *CPU) STX(addr uint16) {
	switch addr {
	case 0x2004:
		cpu.PPU.sRAM[cpu.RAM[0x2003]] = cpu.Reg.X
		cpu.RAM[0x2003]++
	case 0x2005:
		if cpu.PPU.scrollFlag {
			cpu.PPU.scroll[1] = cpu.Reg.X
		} else {
			cpu.PPU.scroll[0] = cpu.Reg.X
			cpu.PPU.scrollFlag = true
		}
	case 0x2006:
		cpu.PPU.ptr = (cpu.PPU.ptr<<8 | uint16(cpu.Reg.X))
	case 0x2007:
		cpu.setVRAM(cpu.Reg.X)
	case spriteDMA:
		start := uint16(cpu.Reg.X) << 8
		var wait sync.WaitGroup
		wait.Add(256)
		for i := 0; i < 256; i++ {
			go func(i int) {
				cpu.PPU.sRAM[i] = cpu.FetchMemory8(start + uint16(i))
				wait.Done()
			}(i)
		}
		wait.Wait()
	case joypad1:
		cpu.joypad1.ctr = 0
		for i := 0; i < 8; i++ {
			cpu.joypad1.cmd[i] = 0
		}
	}
	cpu.SetMemory8(addr, cpu.Reg.X)
}

// STY Store Y to M (Y -> M)
func (cpu *CPU) STY(addr uint16) {
	switch addr {
	case 0x2004:
		cpu.PPU.sRAM[cpu.RAM[0x2003]] = cpu.Reg.Y
		cpu.RAM[0x2003]++
	case 0x2005:
		if cpu.PPU.scrollFlag {
			cpu.PPU.scroll[1] = cpu.Reg.Y
		} else {
			cpu.PPU.scroll[0] = cpu.Reg.Y
			cpu.PPU.scrollFlag = true
		}
	case 0x2006:
		cpu.PPU.ptr = (cpu.PPU.ptr<<8 | uint16(cpu.Reg.Y))
	case 0x2007:
		cpu.setVRAM(cpu.Reg.Y)
	case spriteDMA:
		start := uint16(cpu.Reg.Y) << 8
		var wait sync.WaitGroup
		wait.Add(256)
		for i := 0; i < 256; i++ {
			go func(i int) {
				cpu.PPU.sRAM[i] = cpu.FetchMemory8(start + uint16(i))
				wait.Done()
			}(i)
		}
		wait.Wait()
	case joypad1:
		cpu.joypad1.ctr = 0
		for i := 0; i < 8; i++ {
			cpu.joypad1.cmd[i] = 0
		}
	}
	cpu.SetMemory8(addr, cpu.Reg.Y)
}

// setVRAM VRAM(0x2007)に書き込む処理を共通化したもの
func (cpu *CPU) setVRAM(value byte) {
	cpu.PPU.RAM[cpu.PPU.ptr] = value
	if cpu.PPU.ptr == 0x3f0f {
		cpu.PPU.BGPalleteOK = false
	} else if cpu.PPU.ptr == 0x3f1f {
		cpu.PPU.SPRPalleteOK = false
	} else if !cpu.PPU.mirror && ((0x2000 <= cpu.PPU.ptr && cpu.PPU.ptr < 0x2400) || (0x2800 <= cpu.PPU.ptr && cpu.PPU.ptr < 0x2c00)) {
		cpu.PPU.RAM[cpu.PPU.ptr+0x0400] = value
	}
	cpu.PPU.ptr += cpu.getVRAMDelta()
}

// ============================================ レジスタ間転送 ============================================

// TAX Transfer A to X
func (cpu *CPU) TAX(addr uint16) {
	cpu.Reg.X = cpu.Reg.A
	cpu.FlagN(cpu.Reg.X)
	cpu.FlagZ(cpu.Reg.X)
}

// TAY Transfer A to Y
func (cpu *CPU) TAY(addr uint16) {
	cpu.Reg.Y = cpu.Reg.A
	cpu.FlagN(cpu.Reg.Y)
	cpu.FlagZ(cpu.Reg.Y)
}

// TXA Transfer X to A
func (cpu *CPU) TXA(addr uint16) {
	cpu.Reg.A = cpu.Reg.X
	cpu.FlagN(cpu.Reg.A)
	cpu.FlagZ(cpu.Reg.A)
}

// TYA Transfer Y to A
func (cpu *CPU) TYA(addr uint16) {
	cpu.Reg.A = cpu.Reg.Y
	cpu.FlagN(cpu.Reg.A)
	cpu.FlagZ(cpu.Reg.A)
}

// TSX Transfer S to X
func (cpu *CPU) TSX(addr uint16) {
	cpu.Reg.X = cpu.Reg.S
	cpu.FlagN(cpu.Reg.X)
	cpu.FlagZ(cpu.Reg.X)
}

// TXS Transfer X to S
func (cpu *CPU) TXS(addr uint16) {
	cpu.Reg.S = cpu.Reg.X
	cpu.FlagN(cpu.Reg.S)
	cpu.FlagZ(cpu.Reg.S)
}

// ============================================ スタック ============================================

// PHA Push A on stack
func (cpu *CPU) PHA(addr uint16) {
	if addr == null {
		cpu.SetMemory8((0x100 + uint16(cpu.Reg.S)), cpu.Reg.A)
		cpu.Reg.S--
	}
}

// PLA Pull A from stack
func (cpu *CPU) PLA(addr uint16) {
	if addr == null {
		value := cpu.FetchMemory8(0x0100 + uint16(cpu.Reg.S) + 1)
		cpu.Reg.A = value
		cpu.Reg.S++

		cpu.FlagN(value)
		cpu.FlagZ(value)
	}
}

// PHP Push P on stack
func (cpu *CPU) PHP(addr uint16) {
	if addr == null {
		cpu.SetMemory8((0x100 + uint16(cpu.Reg.S)), cpu.Reg.P)
		cpu.Reg.S--
	}
}

// PLP Pull P from stack
func (cpu *CPU) PLP(addr uint16) {
	if addr == null {
		value := cpu.FetchMemory8(0x0100 + uint16(cpu.Reg.S) + 1)
		cpu.Reg.P = value // pullがフラグのセットになっている
		cpu.Reg.S++
	}
}

// ============================================ NOP ============================================

// NOP No operation
func (cpu *CPU) NOP(addr uint16) {
}
