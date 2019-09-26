package emulator

// LDYImmediate 0xa0: Load into Y in Immediate mode
func (cpu *CPU) LDYImmediate() {
	immediate := cpu.FetchCode8(1)
	cpu.Reg.Y = immediate
	cpu.Reg.PC += 2
}

// LDAIndexedIndirectX 0xa1: Load into A in IndexedIndirect mode(X)
func (cpu *CPU) LDAIndexedIndirectX() {
	addr := cpu.IndexedIndirectAddressing()
	value := cpu.FetchMemory8(addr)
	cpu.Reg.A = value
}

// LDXImmediate 0xa2: Load into X in Immediate mode
func (cpu *CPU) LDXImmediate() {
	immediate := cpu.FetchCode8(1)
	cpu.Reg.X = immediate
	cpu.Reg.PC += 2
}

// LDYZeroPage 0xa4: Load into Y in ZeroPage mode
func (cpu *CPU) LDYZeroPage() {
	addr := cpu.ZeroPageAddressing()
	value := cpu.FetchMemory8(uint(addr))
	cpu.Reg.Y = value
}

// LDAZeroPage 0xa5: Load into A in ZeroPage mode
func (cpu *CPU) LDAZeroPage() {
	addr := cpu.ZeroPageAddressing()
	value := cpu.FetchMemory8(uint(addr))
	cpu.Reg.A = value
}

// LDXZeroPage 0xa6: Load into X in ZeroPage mode
func (cpu *CPU) LDXZeroPage() {
	addr := cpu.ZeroPageAddressing()
	value := cpu.FetchMemory8(uint(addr))
	cpu.Reg.X = value
}

// TAYImplied 0xa8: Transfer A into Y in implied mode
func (cpu *CPU) TAYImplied() {
	cpu.Reg.Y = cpu.Reg.A
	cpu.Reg.PC++
}

// LDAImmediate 0xa9: Load into A in Immediate mode
func (cpu *CPU) LDAImmediate() {
	immediate := cpu.FetchCode8(1)
	cpu.Reg.A = immediate
	cpu.Reg.PC += 2
}

// TAXImplied 0xaa: Transfer A into X in implied mode
func (cpu *CPU) TAXImplied() {
	cpu.Reg.X = cpu.Reg.A
	cpu.Reg.PC++
}

// LDYAbsolute 0xac: Load into Y in Absolute mode
func (cpu *CPU) LDYAbsolute() {
	addr := cpu.AbsoluteAddressing()
	value := cpu.FetchMemory8(uint(addr))
	cpu.Reg.Y = value
}

// LDAAbsolute 0xad: Load into A in Absolute mode
func (cpu *CPU) LDAAbsolute() {
	addr := cpu.AbsoluteAddressing()
	value := cpu.FetchMemory8(uint(addr))
	cpu.Reg.A = value
}

// LDXAbsolute 0xac: Load into X in Absolute mode
func (cpu *CPU) LDXAbsolute() {
	addr := cpu.AbsoluteAddressing()
	value := cpu.FetchMemory8(uint(addr))
	cpu.Reg.X = value
}
