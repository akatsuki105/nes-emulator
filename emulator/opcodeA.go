package emulator

// LDYImmediate 0xa0: Load into Y in Immediate mode
func (cpu *CPU) LDYImmediate() {
	addr := cpu.ImmediateAddressing()
	cpu.LDY(addr)
}

// LDAIndexedIndirectX 0xa1: Load into A in IndexedIndirect mode(X)
func (cpu *CPU) LDAIndexedIndirectX() {
	addr := cpu.IndexedIndirectAddressing()
	cpu.LDA(addr)
}

// LDXImmediate 0xa2: Load into X in Immediate mode
func (cpu *CPU) LDXImmediate() {
	addr := cpu.ImmediateAddressing()
	cpu.LDX(addr)
}

// LDYZeroPage 0xa4: Load into Y in ZeroPage mode
func (cpu *CPU) LDYZeroPage() {
	addr := cpu.ZeroPageAddressing()
	cpu.LDY(addr)
}

// LDAZeroPage 0xa5: Load into A in ZeroPage mode
func (cpu *CPU) LDAZeroPage() {
	addr := cpu.ZeroPageAddressing()
	cpu.LDA(addr)
}

// LDXZeroPage 0xa6: Load into X in ZeroPage mode
func (cpu *CPU) LDXZeroPage() {
	addr := cpu.ZeroPageAddressing()
	cpu.LDX(addr)
}

// TAYImplied 0xa8: Transfer A into Y in implied mode
func (cpu *CPU) TAYImplied() {
	cpu.Reg.PC++
	cpu.TAY()
}

// LDAImmediate 0xa9: Load into A in Immediate mode
func (cpu *CPU) LDAImmediate() {
	addr := cpu.ImmediateAddressing()
	cpu.LDA(addr)
}

// TAXImplied 0xaa: Transfer A into X in implied mode
func (cpu *CPU) TAXImplied() {
	cpu.Reg.PC++
	cpu.TAX()
}

// LDYAbsolute 0xac: Load into Y in Absolute mode
func (cpu *CPU) LDYAbsolute() {
	addr := cpu.AbsoluteAddressing()
	cpu.LDY(addr)
}

// LDAAbsolute 0xad: Load into A in Absolute mode
func (cpu *CPU) LDAAbsolute() {
	addr := cpu.AbsoluteAddressing()
	cpu.LDA(addr)
}

// LDXAbsolute 0xac: Load into X in Absolute mode
func (cpu *CPU) LDXAbsolute() {
	addr := cpu.AbsoluteAddressing()
	cpu.LDX(addr)
}
