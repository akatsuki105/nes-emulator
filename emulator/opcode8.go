package emulator

// STAIndexedIndirect 0x81
func (cpu *CPU) STAIndexedIndirect() {
	addr := cpu.IndexedIndirectAddressing()
	cpu.STA(addr)
}

// STYZeroPage 0x84
func (cpu *CPU) STYZeroPage() {
	addr := cpu.ZeroPageAddressing()
	cpu.STY(addr)
}

// STAZeroPage 0x85
func (cpu *CPU) STAZeroPage() {
	addr := cpu.ZeroPageAddressing()
	cpu.STA(addr)
}

// STXZeroPage 0x86
func (cpu *CPU) STXZeroPage() {
	addr := cpu.ZeroPageAddressing()
	cpu.STX(addr)
}

// DEYImplied 0x88
func (cpu *CPU) DEYImplied() {
	addr := cpu.ImpliedAddressing()
	cpu.DEY(addr)
}

// TXAImplied 0x8a
func (cpu *CPU) TXAImplied() {
	addr := cpu.ImpliedAddressing()
	cpu.TXA(addr)
}

// STYAbsolute 0xc
func (cpu *CPU) STYAbsolute() {
	addr := cpu.AbsoluteAddressing()
	cpu.STY(addr)
}

// STAAbsolute 0xd
func (cpu *CPU) STAAbsolute() {
	addr := cpu.AbsoluteAddressing()
	cpu.STA(addr)
}

// STXAbsolute 0xe
func (cpu *CPU) STXAbsolute() {
	addr := cpu.AbsoluteAddressing()
	cpu.STX(addr)
}
