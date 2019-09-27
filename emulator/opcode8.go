package emulator

// STAIndexedIndirect 0x81
func (cpu *CPU) STAIndexedIndirect() {
	addr := cpu.IndexedIndirectAddressing()
	cpu.SetMemory8(addr, cpu.Reg.A)
}

// STYZeroPage 0x84
func (cpu *CPU) STYZeroPage() {
	addr := cpu.ZeroPageAddressing()
	cpu.SetMemory8(addr, cpu.Reg.Y)
}

// STAZeroPage 0x85
func (cpu *CPU) STAZeroPage() {
	addr := cpu.ZeroPageAddressing()
	cpu.SetMemory8(addr, cpu.Reg.A)
}

// STXZeroPage 0x86
func (cpu *CPU) STXZeroPage() {
	addr := cpu.ZeroPageAddressing()
	cpu.SetMemory8(addr, cpu.Reg.X)
}

// DEYImplied 0x88
func (cpu *CPU) DEYImplied() {
	cpu.Reg.Y--
	cpu.Reg.PC++

	cpu.FlagN(cpu.Reg.Y)
	cpu.FlagZ(cpu.Reg.Y)
}

// TXAImplied 0x8a
func (cpu *CPU) TXAImplied() {
	cpu.Reg.A = cpu.Reg.X
	cpu.Reg.PC++

	cpu.FlagN(cpu.Reg.A)
	cpu.FlagZ(cpu.Reg.A)
}

// STYAbsolute 0xc
func (cpu *CPU) STYAbsolute() {
	addr := cpu.AbsoluteAddressing()
	cpu.SetMemory8(addr, cpu.Reg.Y)
}

// STAAbsolute 0xd
func (cpu *CPU) STAAbsolute() {
	addr := cpu.AbsoluteAddressing()
	cpu.SetMemory8(addr, cpu.Reg.A)
}

// STXAbsolute 0xe
func (cpu *CPU) STXAbsolute() {
	addr := cpu.AbsoluteAddressing()
	cpu.SetMemory8(addr, cpu.Reg.X)
}
