package emulator

// RTSImplied 0x60: Return from Subroutine
func (cpu *CPU) RTSImplied() {
	addr := cpu.ImpliedAddressing()
	cpu.RTS(addr)
}

// ADCIndexedIndirect 0x61
func (cpu *CPU) ADCIndexedIndirect() {
	addr := cpu.IndexedIndirectAddressing()
	cpu.ADC(addr)
}

// ADCZeroPage 0x65
func (cpu *CPU) ADCZeroPage() {
	addr := cpu.ZeroPageAddressing()
	cpu.ADC(addr)
}

// RORZeroPage 0x66
func (cpu *CPU) RORZeroPage() {
	addr := cpu.ZeroPageAddressing()
	cpu.ROR(addr)
}

// PLAImplied 0x68: Pull A from stack (stack -> A)
func (cpu *CPU) PLAImplied() {
	addr := cpu.ImpliedAddressing()
	cpu.PLA(addr)
}

// ADCImmediate 0x69
func (cpu *CPU) ADCImmediate() {
	addr := cpu.ImmediateAddressing()
	cpu.ADC(addr)
}

// RORAccumulator 0x6a
func (cpu *CPU) RORAccumulator() {
	addr := cpu.AccumulatorAddressing()
	cpu.ROR(addr)
}

// JMPAbsoluteIndirect 0x6c
func (cpu *CPU) JMPAbsoluteIndirect() {
	addr := cpu.AbsoluteIndirectAddressing()
	cpu.JMP(addr)
}

// ADCAbsolute 0x6d
func (cpu *CPU) ADCAbsolute() {
	addr := cpu.AbsoluteAddressing()
	cpu.ADC(addr)
}

// RORAbsolute 0x6d
func (cpu *CPU) RORAbsolute() {
	addr := cpu.AbsoluteAddressing()
	cpu.ROR(addr)
}
