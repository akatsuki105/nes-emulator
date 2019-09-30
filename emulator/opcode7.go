package emulator

// BVSRelative 0x70
func (cpu *CPU) BVSRelative() {
	addr := cpu.RelativeAddressing()
	cpu.BVS(addr)
}

// ADCIndirectIndexed 0x71
func (cpu *CPU) ADCIndirectIndexed() {
	addr := cpu.IndirectIndexedAddressing()
	cpu.ADC(addr)
}

// ADCZeroPageX 0x75
func (cpu *CPU) ADCZeroPageX() {
	addr := cpu.ZeroPageXAddressing()
	cpu.ADC(addr)
}

// RORZeroPageX 0x76
func (cpu *CPU) RORZeroPageX() {
	addr := cpu.ZeroPageXAddressing()
	cpu.ROR(addr)
}

// SEIImplied 0x78
func (cpu *CPU) SEIImplied() {
	addr := cpu.ImpliedAddressing()
	cpu.SEI(addr)
}

// ADCAbsoluteY 0x79
func (cpu *CPU) ADCAbsoluteY() {
	addr := cpu.AbsoluteYAddressing()
	cpu.ADC(addr)
}

// ADCAbsoluteX 0x7d
func (cpu *CPU) ADCAbsoluteX() {
	addr := cpu.AbsoluteXAddressing()
	cpu.ADC(addr)
}

// RORAbsoluteX 0x7e
func (cpu *CPU) RORAbsoluteX() {
	addr := cpu.AbsoluteXAddressing()
	cpu.ROR(addr)
}
