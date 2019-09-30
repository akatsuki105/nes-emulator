package emulator

// BNERelative 0xd0
func (cpu *CPU) BNERelative() {
	addr := cpu.RelativeAddressing()
	cpu.BNE(addr)
}

// CMPIndirectIndexed 0xd1
func (cpu *CPU) CMPIndirectIndexed() {
	addr := cpu.IndirectIndexedAddressing()
	cpu.CMP(addr)
}

// CMPZeroPageX 0xd5
func (cpu *CPU) CMPZeroPageX() {
	addr := cpu.ZeroPageXAddressing()
	cpu.CMP(addr)
}

// DECZeroPageX 0xd6
func (cpu *CPU) DECZeroPageX() {
	addr := cpu.ZeroPageXAddressing()
	cpu.DEC(addr)
}

// CLDImplied 0xc8
func (cpu *CPU) CLDImplied() {
	addr := cpu.ImpliedAddressing()
	cpu.CLD(addr)
}

// CMPAbsoluteY 0xc9
func (cpu *CPU) CMPAbsoluteY() {
	addr := cpu.AbsoluteYAddressing()
	cpu.CMP(addr)
}
