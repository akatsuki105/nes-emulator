package emulator

// CPYImmediate 0xc0: Compare (Y - M)
func (cpu *CPU) CPYImmediate() {
	immediate := cpu.FetchCode8(1)
	cpu.Reg.PC += 2
	cpu.CPY(immediate)
}

// CMPIndexedIndirect 0xc1: Compare (A - M)
func (cpu *CPU) CMPIndexedIndirect() {
	addr := cpu.IndexedIndirectAddressing()
	cpu.CMP(cpu.FetchMemory8(addr))
}

// CPYZeroPage 0xc4: Compare (Y - M)
func (cpu *CPU) CPYZeroPage() {
	addr := cpu.ZeroPageAddressing()
	cpu.CPY(cpu.FetchMemory8(addr))
}

// CMPZeroPage 0xc5: Compare (A - M)
func (cpu *CPU) CMPZeroPage() {
	addr := cpu.ZeroPageAddressing()
	cpu.CMP(cpu.FetchMemory8(addr))
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
	immediate := cpu.FetchCode8(1)
	cpu.Reg.PC += 2
	cpu.CMP(immediate)
}

// DEXImplied 0xca
func (cpu *CPU) DEXImplied() {
	cpu.Reg.PC++
	cpu.DEX()
}

// CPYAbsolute 0xcc
func (cpu *CPU) CPYAbsolute() {
	addr := cpu.AbsoluteAddressing()
	cpu.CPY(cpu.FetchMemory8(addr))
}

// CMPAbsolute 0xcd
func (cpu *CPU) CMPAbsolute() {
	addr := cpu.AbsoluteAddressing()
	cpu.CMP(cpu.FetchMemory8(addr))
}

// DECAbsolute 0xce
func (cpu *CPU) DECAbsolute() {
	addr := cpu.AbsoluteAddressing()
	cpu.DEC(addr)
}
