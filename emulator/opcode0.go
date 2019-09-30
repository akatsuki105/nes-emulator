package emulator

// BRKImplied 0x00: Break
func (cpu *CPU) BRKImplied() {
	addr := cpu.ImpliedAddressing()
	cpu.BRK(addr)
}

// ORAIndexedIndirect 0x01
func (cpu *CPU) ORAIndexedIndirect() {
	addr := cpu.IndexedIndirectAddressing()
	cpu.ORA(addr)
}

// ORAZeroPage 0x05
func (cpu *CPU) ORAZeroPage() {
	addr := cpu.ZeroPageAddressing()
	cpu.ORA(addr)
}

// ASLZeroPage 0x06
func (cpu *CPU) ASLZeroPage() {
	addr := cpu.ZeroPageAddressing()
	cpu.ASL(addr)
}

// PHPImplied 0x08
func (cpu *CPU) PHPImplied() {
	addr := cpu.ImpliedAddressing()
	cpu.PHP(addr)
}

// ORAImmediate 0x09
func (cpu *CPU) ORAImmediate() {
	addr := cpu.ImmediateAddressing()
	cpu.ORA(addr)
}

// ASLAccumulator 0x0a
func (cpu *CPU) ASLAccumulator() {
	addr := cpu.AccumulatorAddressing()
	cpu.ASL(addr)
}

// ORAAbsolute 0x0d
func (cpu *CPU) ORAAbsolute() {
	addr := cpu.AbsoluteAddressing()
	cpu.ORA(addr)
}

// ASLAbsolute 0x0e
func (cpu *CPU) ASLAbsolute() {
	addr := cpu.AbsoluteAddressing()
	cpu.ASL(addr)
}
