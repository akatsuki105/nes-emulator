package emulator

// CPYImmediate 0xc0: Compare (Y - M)
func (cpu *CPU) CPYImmediate() {
	immediate := cpu.FetchCode8(1)
	value := cpu.Reg.Y - immediate // Y - M
	value16 := uint16(cpu.Reg.Y) - uint16(immediate)
	cpu.Reg.PC += 2

	cpu.FlagN(value)
	cpu.FlagZ(value)
	cpu.FlagC(value16)
}

// CMPIndexedIndirect 0xc1: Compare (A - M)
func (cpu *CPU) CMPIndexedIndirect() {
	addr := cpu.IndexedIndirectAddressing()

	value := cpu.Reg.A - cpu.FetchMemory8(addr)
	value16 := uint16(cpu.Reg.A) - uint16(cpu.FetchMemory8(addr))

	cpu.FlagN(value)
	cpu.FlagZ(value)
	cpu.FlagC(value16)
}

// CPYZeroPage 0xc4: Compare (Y - M)
func (cpu *CPU) CPYZeroPage() {
	addr := cpu.ZeroPageAddressing()

	value := cpu.Reg.Y - cpu.FetchMemory8(addr)
	value16 := uint16(cpu.Reg.Y) - uint16(cpu.FetchMemory8(addr))

	cpu.FlagN(value)
	cpu.FlagZ(value)
	cpu.FlagC(value16)
}

// CMPZeroPage 0xc5: Compare (A - M)
func (cpu *CPU) CMPZeroPage() {
	addr := cpu.ZeroPageAddressing()

	value := cpu.Reg.A - cpu.FetchMemory8(addr)
	value16 := uint16(cpu.Reg.A) - uint16(cpu.FetchMemory8(addr))

	cpu.FlagN(value)
	cpu.FlagZ(value)
	cpu.FlagC(value16)
}

// DECZeroPage 0xc6
func (cpu *CPU) DECZeroPage() {
	addr := cpu.ZeroPageAddressing()

	value := cpu.FetchMemory8(addr) - 1
	cpu.SetMemory8(addr, value)

	cpu.FlagN(value)
	cpu.FlagZ(value)
}

// INYImplied 0xc8
func (cpu *CPU) INYImplied() {
	cpu.Reg.Y++
	cpu.Reg.PC++

	cpu.FlagN(cpu.Reg.Y)
	cpu.FlagZ(cpu.Reg.Y)
}

// CMPImmediate 0xc9
func (cpu *CPU) CMPImmediate() {
	immediate := cpu.FetchCode8(1)
	value := cpu.Reg.A - immediate // A - M
	value16 := uint16(cpu.Reg.A) - uint16(immediate)
	cpu.Reg.PC += 2

	cpu.FlagN(value)
	cpu.FlagZ(value)
	cpu.FlagC(value16)
}

// DEXImplied 0xca
func (cpu *CPU) DEXImplied() {
	cpu.Reg.X--
	cpu.Reg.PC++

	cpu.FlagN(cpu.Reg.X)
	cpu.FlagZ(cpu.Reg.X)
}

// CPYAbsolute 0xcc
func (cpu *CPU) CPYAbsolute() {
	addr := cpu.AbsoluteAddressing()

	value := cpu.Reg.Y - cpu.FetchMemory8(addr)
	value16 := uint16(cpu.Reg.Y) - uint16(cpu.FetchMemory8(addr))

	cpu.FlagN(value)
	cpu.FlagZ(value)
	cpu.FlagC(value16)
}

// CMPAbsolute 0xcd
func (cpu *CPU) CMPAbsolute() {
	addr := cpu.AbsoluteAddressing()

	value := cpu.Reg.A - cpu.FetchMemory8(addr)
	value16 := uint16(cpu.Reg.A) - uint16(cpu.FetchMemory8(addr))

	cpu.FlagN(value)
	cpu.FlagZ(value)
	cpu.FlagC(value16)
}

// DECAbsolute 0xce
func (cpu *CPU) DECAbsolute() {
	addr := cpu.AbsoluteAddressing()

	value := cpu.FetchMemory8(addr) - 1
	cpu.SetMemory8(addr, value)

	cpu.FlagN(value)
	cpu.FlagZ(value)
}
