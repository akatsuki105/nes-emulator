package emulator

// BCSRelative 0xb0: if C is 1, jump in relative mode.
func (cpu *CPU) BCSRelative() {
	addr := cpu.RelativeAddressing()
	cpu.BCS(addr)
}

// LDAIndirectIndexed 0xb1: Load into A in Indirect Indexed mode(Y)
func (cpu *CPU) LDAIndirectIndexed() {
	addr := cpu.IndirectIndexedAddressing()
	cpu.LDA(addr)
}

// LDYZeroPageX 0xb4: Load into Y in ZeroPageX mode
func (cpu *CPU) LDYZeroPageX() {
	addr := cpu.ZeroPageXAddressing()
	cpu.LDY(addr)
}

// LDAZeroPageX 0xb5: Load into A in ZeroPageX mode
func (cpu *CPU) LDAZeroPageX() {
	addr := cpu.ZeroPageXAddressing()
	cpu.LDA(addr)
}

// LDXZeroPageY 0xb6: Load into X in ZeroPageY mode
func (cpu *CPU) LDXZeroPageY() {
	addr := cpu.ZeroPageYAddressing()
	cpu.LDX(addr)
}

// CLVImplied 0xb8: Clear V flag
func (cpu *CPU) CLVImplied() {
	addr := cpu.ImpliedAddressing()
	cpu.CLV(addr)
}

// LDAAbsoluteY 0xb9: Load into A in AbsoluteY mode
func (cpu *CPU) LDAAbsoluteY() {
	addr := cpu.AbsoluteYAddressing()
	cpu.LDA(addr)
}

// TSXImplied 0xba: Transfer S into X
func (cpu *CPU) TSXImplied() {
	addr := cpu.ImpliedAddressing()
	cpu.TSX(addr)
}

// LDYAbsoluteX 0xbc: Load into Y in AbsoluteX mode
func (cpu *CPU) LDYAbsoluteX() {
	addr := cpu.AbsoluteXAddressing()
	cpu.LDY(addr)
}

// LDAAbsoluteX 0xbd: Load into A in AbsoluteX mode
func (cpu *CPU) LDAAbsoluteX() {
	addr := cpu.AbsoluteXAddressing()
	cpu.LDA(addr)
}

// LDXAbsoluteY 0xbe: Load into Y in AbsoluteY mode
func (cpu *CPU) LDXAbsoluteY() {
	addr := cpu.AbsoluteYAddressing()
	cpu.LDX(addr)
}
