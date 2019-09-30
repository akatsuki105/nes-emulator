package emulator

// BVCRelative 0x50
func (cpu *CPU) BVCRelative() {
	addr := cpu.RelativeAddressing()

	vFlag := uint8(cpu.Reg.P & 0x40)
	if vFlag == 0 {
		cpu.Reg.PC = addr
	}
}

// EORIndirectIndexed 0x51
func (cpu *CPU) EORIndirectIndexed() {
	addr := cpu.IndirectIndexedAddressing()
	cpu.EOR(addr)
}

// EORZeroPageX 0x55
func (cpu *CPU) EORZeroPageX() {
	addr := cpu.ZeroPageXAddressing()
	cpu.EOR(addr)
}

// LSRZeroPageX 0x56
func (cpu *CPU) LSRZeroPageX() {
	addr := cpu.ZeroPageXAddressing()
	cpu.LSR(addr)
}

// CLIImplied 0x58
func (cpu *CPU) CLIImplied() {
	cpu.Reg.PC++
	cpu.Reg.P = cpu.Reg.P & 0xfb // 0b11111011
}

// EORAbsoluteY 0x59
func (cpu *CPU) EORAbsoluteY() {
	addr := cpu.AbsoluteYAddressing()
	cpu.EOR(addr)
}

// EORAbsoluteX 0x5d
func (cpu *CPU) EORAbsoluteX() {
	addr := cpu.AbsoluteXAddressing()
	cpu.EOR(addr)
}

// LSRAbsoluteX 0x5e
func (cpu *CPU) LSRAbsoluteX() {
	addr := cpu.AbsoluteXAddressing()
	cpu.LSR(addr)
}
