package emulator

// BMIRelative 0x30
func (cpu *CPU) BMIRelative() {
	addr := cpu.RelativeAddressing()
	cpu.BMI(addr)
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
	addr := cpu.ImpliedAddressing()
	cpu.SEC(addr)
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
