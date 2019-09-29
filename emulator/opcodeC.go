package emulator

// CPYImmediate 0xc0: Compare (Y - M)
func (cpu *CPU) CPYImmediate() {
	addr := cpu.ImmediateAddressing()
	cpu.CPY(addr)
}

// CMPIndexedIndirect 0xc1: Compare (A - M)
func (cpu *CPU) CMPIndexedIndirect() {
	addr := cpu.IndexedIndirectAddressing()
	cpu.CMP(addr)
}

// CPYZeroPage 0xc4: Compare (Y - M)
func (cpu *CPU) CPYZeroPage() {
	addr := cpu.ZeroPageAddressing()
	cpu.CPY(addr)
}

// CMPZeroPage 0xc5: Compare (A - M)
func (cpu *CPU) CMPZeroPage() {
	addr := cpu.ZeroPageAddressing()
	cpu.CMP(addr)
}

// DECZeroPage 0xc6
func (cpu *CPU) DECZeroPage() {
	addr := cpu.ZeroPageAddressing()
	cpu.DEC(addr)
}

// INYImplied 0xc8
func (cpu *CPU) INYImplied() {
	cpu.Reg.PC++
	cpu.INY()
}

// CMPImmediate 0xc9
func (cpu *CPU) CMPImmediate() {
	addr := cpu.ImmediateAddressing()
	cpu.CMP(addr)
}

// DEXImplied 0xca
func (cpu *CPU) DEXImplied() {
	cpu.Reg.PC++
	cpu.DEX()
}

// CPYAbsolute 0xcc
func (cpu *CPU) CPYAbsolute() {
	addr := cpu.AbsoluteAddressing()
	cpu.CPY(addr)
}

// CMPAbsolute 0xcd
func (cpu *CPU) CMPAbsolute() {
	addr := cpu.AbsoluteAddressing()
	cpu.CMP(addr)
}

// DECAbsolute 0xce
func (cpu *CPU) DECAbsolute() {
	addr := cpu.AbsoluteAddressing()
	cpu.DEC(addr)
}
