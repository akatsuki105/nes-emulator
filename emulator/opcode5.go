package emulator

// BVCRelative 0x50
func (cpu *CPU) BVCRelative() {
	addr := cpu.RelativeAddressing()
	cpu.BVC(addr)
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
	addr := cpu.ImpliedAddressing()
	cpu.CLI(addr)
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
