package emulator

// BCCRelative 0x90: if C is 1, jump in relative mode.
func (cpu *CPU) BCCRelative() {
	addr := cpu.RelativeAddressing()
	cpu.BCC(addr)
}

// STAIndirectIndexed 0x91: Store A into M in Indirect Indexed mode(Y)
func (cpu *CPU) STAIndirectIndexed() {
	addr := cpu.IndirectIndexedAddressing()
	cpu.STA(addr)
}

// STYZeroPageX 0x94: Store Y into M in ZeroPageX mode
func (cpu *CPU) STYZeroPageX() {
	addr := cpu.ZeroPageXAddressing()
	cpu.STY(addr)
}

// STAZeroPageX 0x95: Store A into M in ZeroPageX mode
func (cpu *CPU) STAZeroPageX() {
	addr := cpu.ZeroPageXAddressing()
	cpu.STA(addr)
}

// STXZeroPageY 0x96: Store X into M in ZeroPageY mode
func (cpu *CPU) STXZeroPageY() {
	addr := cpu.ZeroPageYAddressing()
	cpu.STX(addr)
}

// TYAImplied 0x98: Transfer Y into A in Implied mode
func (cpu *CPU) TYAImplied() {
	addr := cpu.ImpliedAddressing()
	cpu.TYA(addr)
}

// STAAbsoluteY 0x99: Store A into M in AbsoluteY mode
func (cpu *CPU) STAAbsoluteY() {
	addr := cpu.AbsoluteYAddressing()
	cpu.STA(addr)
}

// TXSImplied 0x9a: Transfer X into S in Implied mode
func (cpu *CPU) TXSImplied() {
	addr := cpu.ImpliedAddressing()
	cpu.TXS(addr)
}

// STAAbsoluteX 0x9d: Store A into M in AbsoluteX mode
func (cpu *CPU) STAAbsoluteX() {
	addr := cpu.AbsoluteXAddressing()
	cpu.STA(addr)
}
