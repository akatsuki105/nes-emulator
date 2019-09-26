package emulator

// LDYImmediate 0xa0: Load into Y in Immediate mode
func (cpu *CPU) LDYImmediate() {
	immediate := cpu.FetchCode8(1)
	cpu.Reg.Y = immediate
	cpu.Reg.PC += 2
}

// LDAIndexedIndirectX 0xa1: Load into A in Index Indirect mode(X)
func (cpu *CPU) LDAIndexedIndirectX() {
	lower0 := uint16(cpu.FetchCode8(1)) + uint16(cpu.Reg.X)
	upper0 := uint16(0x00)
	addr0 := (upper0 << 8) | (lower0)

	lower1 := uint16(cpu.FetchMemory8(uint(addr0)))
	upper1 := uint16(cpu.FetchMemory8(uint(addr0 + 1)))
	addr1 := (upper1 << 8) | (lower1)

	value := cpu.FetchMemory8(uint(addr1))
	cpu.Reg.A = value
	cpu.Reg.PC += 2
}

// LDXImmediate 0xa2: Load into X in Immediate mode
func (cpu *CPU) LDXImmediate() {
	immediate := cpu.FetchCode8(1)
	cpu.Reg.X = immediate
	cpu.Reg.PC += 2
}

// LDYZeroPage 0xa4: Load into Y in ZeroPage mode
func (cpu *CPU) LDYZeroPage() {
	lower := uint16(cpu.FetchCode8(1))
	upper := uint16(0x00)
	addr := (upper << 8) | (lower)
	value := cpu.FetchMemory8(uint(addr))
	cpu.Reg.Y = value
	cpu.Reg.PC += 2
}

// LDAZeroPage 0xa5: Load into A in ZeroPage mode
func (cpu *CPU) LDAZeroPage() {
	lower := uint16(cpu.FetchCode8(1))
	upper := uint16(0x00)
	addr := (upper << 8) | (lower)
	value := cpu.FetchMemory8(uint(addr))
	cpu.Reg.A = value
	cpu.Reg.PC += 2
}

// LDXZeroPage 0xa6: Load into X in ZeroPage mode
func (cpu *CPU) LDXZeroPage() {
	lower := uint16(cpu.FetchCode8(1))
	upper := uint16(0x00)
	addr := (upper << 8) | (lower)
	value := cpu.FetchMemory8(uint(addr))
	cpu.Reg.X = value
	cpu.Reg.PC += 2
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
	lower := uint16(cpu.FetchCode8(1))
	upper := uint16(cpu.FetchCode8(2))
	addr := (upper << 8) | (lower)
	value := cpu.FetchMemory8(uint(addr))
	cpu.Reg.Y = value
	cpu.Reg.PC += 3
}

// LDAAbsolute 0xad: Load into A in Absolute mode
func (cpu *CPU) LDAAbsolute() {
	lower := uint16(cpu.FetchCode8(1))
	upper := uint16(cpu.FetchCode8(2))
	addr := (upper << 8) | (lower)
	value := cpu.FetchMemory8(uint(addr))
	cpu.Reg.A = value
	cpu.Reg.PC += 3
}

// LDXAbsolute 0xac: Load into X in Absolute mode
func (cpu *CPU) LDXAbsolute() {
	lower := uint16(cpu.FetchCode8(1))
	upper := uint16(cpu.FetchCode8(2))
	addr := (upper << 8) | (lower)
	value := cpu.FetchMemory8(uint(addr))
	cpu.Reg.X = value
	cpu.Reg.PC += 3
}
