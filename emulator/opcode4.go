package emulator

// RTIImplied 0x40: Return from Interrupt
func (cpu *CPU) RTIImplied() {
	// ステータスレジスタをpop
	SR := cpu.FetchMemory8((0x100 + uint(cpu.Reg.S) - 1))
	cpu.Reg.S--
	cpu.Reg.P = SR
	// PCをpop
	lower := uint16(cpu.FetchMemory8((0x100 + uint(cpu.Reg.S) - 1)))
	cpu.Reg.S--
	upper := uint16(cpu.FetchMemory8((0x100 + uint(cpu.Reg.S) - 1)))
	cpu.Reg.S--
	cpu.Reg.PC = (upper << 8) | lower // ここでPCにリターンしているかつ割り込みなのでインクリメントは必要ない
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
	cpu.SetMemory8((0x100 + uint(cpu.Reg.S)), cpu.Reg.A)
	cpu.Reg.S++
	cpu.Reg.PC++
}

// EORImmediate 0x49
func (cpu *CPU) EORImmediate() {
	addr := cpu.ImmediateAddressing()
	cpu.EOR(addr)
}

// LSRAccumulator 0x4a
func (cpu *CPU) LSRAccumulator() {
	cpu.Reg.P = cpu.Reg.P | (cpu.Reg.A & 0x01) // Aのbit0をcにセット
	cpu.Reg.A = cpu.Reg.A >> 1
	cpu.Reg.A = cpu.Reg.A | (0 << 7) // Aのbit7に0をセット
	cpu.FlagN(cpu.Reg.A)
	cpu.FlagZ(cpu.Reg.A)

	cpu.Reg.PC++
}

// JMPAbsolute 0x4c
func (cpu *CPU) JMPAbsolute() {
	addr := cpu.AbsoluteAddressing()
	cpu.Reg.PC = uint16(addr)
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
