package emulator

// LDYImmediate 0xa0: Load into Y in Immediate mode
func (cpu *CPU) LDYImmediate() {
	immediate := cpu.FetchCode8(1)
	cpu.Reg.PC += 2
	cpu.LDY(immediate)
}

// LDAIndexedIndirectX 0xa1: Load into A in IndexedIndirect mode(X)
func (cpu *CPU) LDAIndexedIndirectX() {
	addr := cpu.IndexedIndirectAddressing()
	cpu.LDA(cpu.FetchMemory8(addr))
}

// LDXImmediate 0xa2: Load into X in Immediate mode
func (cpu *CPU) LDXImmediate() {
	immediate := cpu.FetchCode8(1)
	cpu.Reg.PC += 2
	cpu.LDX(immediate)
}

// LDYZeroPage 0xa4: Load into Y in ZeroPage mode
func (cpu *CPU) LDYZeroPage() {
	addr := cpu.ZeroPageAddressing()
	cpu.LDY(cpu.FetchMemory8(addr))
}

// LDAZeroPage 0xa5: Load into A in ZeroPage mode
func (cpu *CPU) LDAZeroPage() {
	addr := cpu.ZeroPageAddressing()
	cpu.LDA(cpu.FetchMemory8(addr))
}

// LDXZeroPage 0xa6: Load into X in ZeroPage mode
func (cpu *CPU) LDXZeroPage() {
	addr := cpu.ZeroPageAddressing()
	cpu.LDX(cpu.FetchMemory8(addr))
}

// TAYImplied 0xa8: Transfer A into Y in implied mode
func (cpu *CPU) TAYImplied() {
	cpu.Reg.PC++
	cpu.TAY()
}

// LDAImmediate 0xa9: Load into A in Immediate mode
func (cpu *CPU) LDAImmediate() {
	immediate := cpu.FetchCode8(1)
	cpu.Reg.PC += 2
	cpu.LDA(immediate)
}

// TAXImplied 0xaa: Transfer A into X in implied mode
func (cpu *CPU) TAXImplied() {
	cpu.Reg.PC++
	cpu.TAX()
}

// LDYAbsolute 0xac: Load into Y in Absolute mode
func (cpu *CPU) LDYAbsolute() {
	addr := cpu.AbsoluteAddressing()
	cpu.LDY(cpu.FetchMemory8(addr))
}

// LDAAbsolute 0xad: Load into A in Absolute mode
func (cpu *CPU) LDAAbsolute() {
	addr := cpu.AbsoluteAddressing()
	cpu.LDA(cpu.FetchMemory8(addr))
}

// LDXAbsolute 0xac: Load into X in Absolute mode
func (cpu *CPU) LDXAbsolute() {
	addr := cpu.AbsoluteAddressing()
	cpu.LDX(cpu.FetchMemory8(addr))
}
