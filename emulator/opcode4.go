package emulator

// RTIImplied 0x40: Return from Interrupt
func (cpu *CPU) RTIImplied() {
	addr := cpu.ImpliedAddressing()
	cpu.RTI(addr)
}

// EORIndexedIndirect 0x41
func (cpu *CPU) EORIndexedIndirect() {
	addr := cpu.IndexedIndirectAddressing()
	cpu.EOR(addr)
}

// EORZeroPage 0x45
func (cpu *CPU) EORZeroPage() {
	addr := cpu.ZeroPageAddressing()
	cpu.EOR(addr)
}

// LSRZeroPage 0x46
func (cpu *CPU) LSRZeroPage() {
	addr := cpu.ZeroPageAddressing()
	cpu.LSR(addr)
}

// PHAImplied 0x48
func (cpu *CPU) PHAImplied() {
	addr := cpu.ImpliedAddressing()
	cpu.PHA(addr)
}

// EORImmediate 0x49
func (cpu *CPU) EORImmediate() {
	addr := cpu.ImmediateAddressing()
	cpu.EOR(addr)
}

// LSRAccumulator 0x4a
func (cpu *CPU) LSRAccumulator() {
	addr := cpu.AccumulatorAddressing()
	cpu.LSR(addr)
}

// JMPAbsolute 0x4c
func (cpu *CPU) JMPAbsolute() {
	addr := cpu.AbsoluteAddressing()
	cpu.JMP(addr)
}

// EORAbsolute 0x4d
func (cpu *CPU) EORAbsolute() {
	addr := cpu.AbsoluteAddressing()
	cpu.EOR(addr)
}

// LSRAbsolute 0x4e
func (cpu *CPU) LSRAbsolute() {
	addr := cpu.AbsoluteAddressing()
	cpu.LSR(addr)
}
