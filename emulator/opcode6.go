package emulator

// RTSImplied 0x60: Return from Subroutine
func (cpu *CPU) RTSImplied() {
	// TODO: 実装
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
	// TODO: 実装
}

// ADCImmediate 0x69
func (cpu *CPU) ADCImmediate() {
	addr := uint(cpu.Reg.PC + 1)
	cpu.Reg.PC += 2
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
