package emulator

// BMIRelative 0x30
func (cpu *CPU) BMIRelative() {
	addr := cpu.RelativeAddressing()

	nFlag := uint8(cpu.Reg.P & 0x80)
	if nFlag > 0 {
		cpu.Reg.PC = uint16(addr)
	}
}

// ANDIndirectIndexed 0x31
func (cpu *CPU) ANDIndirectIndexed() {
	addr := cpu.IndirectIndexedAddressing()
	cpu.AND(addr)
}

// ANDZeroPageX 0x35
func (cpu *CPU) ANDZeroPageX() {
	addr := cpu.ZeroPageAddressing()
	cpu.AND(addr)
}

// ROLZeroPageX 0x36
func (cpu *CPU) ROLZeroPageX() {
	addr := cpu.ZeroPageXAddressing()
	cpu.ROL(addr)
}

// SECImplied 0x38: Set C flag
func (cpu *CPU) SECImplied() {
	cpu.Reg.P = cpu.Reg.P | 0x01
	cpu.Reg.PC++
}

// ANDAbsoluteY 0x39
func (cpu *CPU) ANDAbsoluteY() {
	addr := cpu.AbsoluteYAddressing()
	cpu.AND(addr)
}

// ANDAbsoluteX 0x3d
func (cpu *CPU) ANDAbsoluteX() {
	addr := cpu.AbsoluteXAddressing()
	cpu.AND(addr)
}

// ROLAbsoluteX 0x3e
func (cpu *CPU) ROLAbsoluteX() {
	addr := cpu.AbsoluteXAddressing()
	cpu.ROL(addr)
}
