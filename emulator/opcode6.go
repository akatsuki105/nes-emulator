package emulator

// RTSImplied 0x60: Return from Subroutine
func (cpu *CPU) RTSImplied() {
	lower := uint16(cpu.FetchMemory8((0x100 + uint(cpu.Reg.S) - 1)))
	cpu.Reg.S--
	upper := uint16(cpu.FetchMemory8((0x100 + uint(cpu.Reg.S) - 1)))
	cpu.Reg.S--
	cpu.Reg.PC = (upper << 8) | lower
	cpu.Reg.PC++
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
	value := cpu.FetchMemory8(0x0100 + uint(cpu.Reg.S) - 1)
	cpu.Reg.A = value
	cpu.Reg.S--
	cpu.Reg.PC++

	cpu.FlagN(value)
	cpu.FlagZ(value)
}

// ADCImmediate 0x69
func (cpu *CPU) ADCImmediate() {
	addr := cpu.ImmediateAddressing()
	cpu.ADC(addr)
}

// RORAccumulator 0x6a
func (cpu *CPU) RORAccumulator() {
	cFlag := cpu.Reg.P & 0x01
	cpu.Reg.P = cpu.Reg.P | (cpu.Reg.A & 0x01) // valueのbit0をcにセット
	cpu.Reg.A = cpu.Reg.A >> 1
	cpu.Reg.A = cpu.Reg.A | (cFlag << 7) // valueのbit7にcをセット
	cpu.FlagN(cpu.Reg.A)
	cpu.FlagZ(cpu.Reg.A)

	cpu.Reg.PC++
}

// JMPAbsoluteIndirect 0x6c
func (cpu *CPU) JMPAbsoluteIndirect() {
	addr := cpu.AbsoluteIndirectAddressing()
	cpu.Reg.PC = uint16(addr)
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
