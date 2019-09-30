package emulator

// CPXImmediate 0xe0: Compare M and X (X - M)
func (cpu *CPU) CPXImmediate() {
	addr := cpu.ImmediateAddressing()
	cpu.CPX(addr)
}

// SBCIndexedIndirect 0xe1
func (cpu *CPU) SBCIndexedIndirect() {
	addr := cpu.IndexedIndirectAddressing()
	cpu.SBC(addr)
}

// CPXZeroPage 0xe4
func (cpu *CPU) CPXZeroPage() {
	addr := cpu.ZeroPageAddressing()
	cpu.CPX(addr)
}

// SBCZeroPage 0xe5
func (cpu *CPU) SBCZeroPage() {
	addr := cpu.ZeroPageAddressing()
	cpu.SBC(addr)
}

// INCZeroPage 0xe6: Increment M by one
func (cpu *CPU) INCZeroPage() {
	addr := cpu.ZeroPageAddressing()
	cpu.INC(addr)
}

// INXImplied 0xe8
func (cpu *CPU) INXImplied() {
	addr := cpu.ImpliedAddressing()
	cpu.INX(addr)
}

// SBCImmediate 0xe9
func (cpu *CPU) SBCImmediate() {
	addr := cpu.ImmediateAddressing()
	cpu.SBC(addr)
}

// NOPImplied 0xea
func (cpu *CPU) NOPImplied() {
	addr := cpu.ImpliedAddressing()
	cpu.NOP(addr)
}

// CPXAbsolute 0xec
func (cpu *CPU) CPXAbsolute() {
	addr := cpu.AbsoluteAddressing()
	cpu.CPX(addr)
}

// SBCAbsolute 0xed
func (cpu *CPU) SBCAbsolute() {
	addr := cpu.AbsoluteAddressing()
	cpu.SBC(addr)
}
