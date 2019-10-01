package emulator

// ImpliedAddressing Impliedのアドレスを返す
func (cpu *CPU) ImpliedAddressing() (addr uint16) {
	cpu.Reg.PC++
	return null
}

// AccumulatorAddressing Accumulatorのアドレスを返す
func (cpu *CPU) AccumulatorAddressing() (addr uint16) {
	cpu.Reg.PC++
	return null
}

// ImmediateAddressing Immediateのアドレスを返す
func (cpu *CPU) ImmediateAddressing() (addr uint16) {
	addr = cpu.Reg.PC + 1

	cpu.Reg.PC += 2
	return addr
}

// AbsoluteAddressing Absoluteのアドレスを返す
func (cpu *CPU) AbsoluteAddressing() (addr uint16) {
	lower := uint16(cpu.FetchCode8(1))
	upper := uint16(cpu.FetchCode8(2))
	addr = (upper << 8) | lower

	cpu.Reg.PC += 3
	return addr
}

// AbsoluteIndirectAddressing AbsoluteIndirectのアドレスを返す
func (cpu *CPU) AbsoluteIndirectAddressing() (addr uint16) {
	addr = cpu.AbsoluteAddressing()
	lower := uint16(cpu.FetchMemory8(addr))
	upper := uint16(cpu.FetchMemory8(addr + 1))
	addr = (upper << 8) | lower
	return addr
}

// AbsoluteXAddressing AbsoluteXのアドレスを返す
func (cpu *CPU) AbsoluteXAddressing() (addr uint16) {
	lower := uint16(cpu.FetchCode8(1))
	upper := uint16(cpu.FetchCode8(2))
	addr = (upper << 8) | lower
	addr += uint16(cpu.Reg.X)

	cpu.Reg.PC += 3
	return addr
}

// AbsoluteYAddressing AbsoluteYのアドレスを返す
func (cpu *CPU) AbsoluteYAddressing() (addr uint16) {
	lower := uint16(cpu.FetchCode8(1))
	upper := uint16(cpu.FetchCode8(2))
	addr = (upper << 8) | lower
	addr += uint16(cpu.Reg.Y)

	cpu.Reg.PC += 3
	return addr
}

// ZeroPageAddressing ZeroPageのアドレスを返す
func (cpu *CPU) ZeroPageAddressing() (addr uint16) {
	lower := uint16(cpu.FetchCode8(1))
	upper := uint16(0x00)
	addr = (upper << 8) | (lower)

	cpu.Reg.PC += 2
	return addr
}

// ZeroPageXAddressing ZeroPageYのアドレスを返す
func (cpu *CPU) ZeroPageXAddressing() (addr uint16) {
	lower := uint16(cpu.FetchCode8(1)) + uint16(cpu.Reg.X)
	upper := uint16(0x00)
	addr = (upper << 8) | (lower)

	cpu.Reg.PC += 2
	return addr
}

// ZeroPageYAddressing ZeroPageYのアドレスを返す
func (cpu *CPU) ZeroPageYAddressing() (addr uint16) {
	lower := uint16(cpu.FetchCode8(1)) + uint16(cpu.Reg.Y)
	upper := uint16(0x00)
	addr = (upper << 8) | (lower)

	cpu.Reg.PC += 2
	return addr
}

// IndexedIndirectAddressing IndexedIndirectのアドレスを返す
func (cpu *CPU) IndexedIndirectAddressing() (addr uint16) {
	lower0 := uint16(cpu.FetchCode8(1)) + uint16(cpu.Reg.X)
	upper0 := uint16(0x00)
	addr0 := (upper0 << 8) | (lower0)

	lower1 := uint16(cpu.FetchMemory8(addr0))
	upper1 := uint16(cpu.FetchMemory8(addr0 + 1))
	addr = (upper1 << 8) | (lower1)

	cpu.Reg.PC += 2
	return addr
}

// IndirectIndexedAddressing IndirectIndexedのアドレスを返す
func (cpu *CPU) IndirectIndexedAddressing() (addr uint16) {
	upper0 := uint16(0x00)
	lower0 := uint16(cpu.FetchCode8(1))
	addr0 := (upper0 << 8) | (lower0)

	upper1 := uint16(cpu.FetchMemory8(addr0))
	lower1 := uint16(cpu.FetchMemory8(addr0 + 1))
	addr = (upper1 << 8) | (lower1 + uint16(cpu.Reg.Y))

	cpu.Reg.PC += 2
	return addr
}

// RelativeAddressing Relativeのアドレスを返す
func (cpu *CPU) RelativeAddressing() (addr uint16) {
	delta := int8(cpu.FetchCode8(1))
	cpu.Reg.PC += 2
	addr = uint16(int32(cpu.Reg.PC) + int32(delta))
	return addr
}
