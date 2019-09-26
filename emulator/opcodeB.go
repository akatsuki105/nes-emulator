package emulator

// BCSRelative 0xb0: if C is 1, jump in relative mode.
func (cpu *CPU) BCSRelative() {
	cFlag := uint8(cpu.Reg.P & 0x01)
	addr := int8(cpu.Reg.PC) + 1 + int8(cpu.FetchCode8(1))
	if cFlag > 0 {
		cpu.Reg.PC = uint16(addr)
	} else {
		cpu.Reg.PC += 2
	}
}

// LDAIndirectIndexed 0xb1: Load into A in Indirect Index mode(Y)
func (cpu *CPU) LDAIndirectIndexed() {
	upper0 := uint16(0x00)
	lower0 := uint16(cpu.FetchCode8(1))
	addr0 := (upper0 << 8) | (lower0)

	upper1 := uint16(cpu.FetchMemory8(uint(addr0)))
	lower1 := uint16(cpu.FetchMemory8(uint(addr0 + 1)))
	addr1 := (upper1 << 8) | (lower1) + uint16(cpu.Reg.Y)

	value := cpu.FetchMemory8(uint(addr1))
	cpu.Reg.A = value
	cpu.Reg.PC += 2
}

// LDYZeroPageX 0xb4: Load into Y in ZeroPageX mode
func (cpu *CPU) LDYZeroPageX() {
	lower := uint16(cpu.FetchCode8(1)) + uint16(cpu.Reg.X)
	upper := uint16(0x00)
	addr := (upper << 8) | (lower)

	value := cpu.FetchMemory8(uint(addr))
	cpu.Reg.Y = value
	cpu.Reg.PC += 2
}

// LDAZeroPageX 0xb5: Load into A in ZeroPageX mode
func (cpu *CPU) LDAZeroPageX() {
	lower := uint16(cpu.FetchCode8(1)) + uint16(cpu.Reg.X)
	upper := uint16(0x00)
	addr := (upper << 8) | (lower)

	value := cpu.FetchMemory8(uint(addr))
	cpu.Reg.A = value
	cpu.Reg.PC += 2
}

// LDXZeroPageY 0xb6: Load into X in ZeroPageY mode
func (cpu *CPU) LDXZeroPageY() {
	lower := uint16(cpu.FetchCode8(1)) + uint16(cpu.Reg.Y)
	upper := uint16(0x00)
	addr := (upper << 8) | (lower)

	value := cpu.FetchMemory8(uint(addr))
	cpu.Reg.X = value
	cpu.Reg.PC += 2
}

// CLVImplied 0xb8: Clear V flag
func (cpu *CPU) CLVImplied() {
	cpu.Reg.P = cpu.Reg.P & 0xbf
	cpu.Reg.PC++
}

// LDAAbsoluteY 0xb9: Load into A in AbsoluteY mode
func (cpu *CPU) LDAAbsoluteY() {
	lower := uint16(cpu.FetchCode8(1))
	upper := uint16(cpu.FetchCode8(2))
	addr := (upper << 8) | (lower) + uint16(cpu.Reg.Y)
	value := cpu.FetchMemory8(uint(addr))
	cpu.Reg.A = value
	cpu.Reg.PC += 3
}

// TSXImplied 0xba: Transfer S into X
func (cpu *CPU) TSXImplied() {
	cpu.Reg.X = cpu.Reg.S
	cpu.Reg.PC++
}

// LDYAbsoluteX 0xbc: Load into Y in AbsoluteX mode
func (cpu *CPU) LDYAbsoluteX() {
	lower := uint16(cpu.FetchCode8(1))
	upper := uint16(cpu.FetchCode8(2))
	addr := (upper << 8) | (lower) + uint16(cpu.Reg.X)
	value := cpu.FetchMemory8(uint(addr))
	cpu.Reg.Y = value
	cpu.Reg.PC += 3
}

// LDAAbsoluteX 0xbd: Load into A in AbsoluteX mode
func (cpu *CPU) LDAAbsoluteX() {
	lower := uint16(cpu.FetchCode8(1))
	upper := uint16(cpu.FetchCode8(2))
	addr := (upper << 8) | (lower) + uint16(cpu.Reg.X)
	value := cpu.FetchMemory8(uint(addr))
	cpu.Reg.A = value
	cpu.Reg.PC += 3
}

// LDXAbsoluteY 0xbe: Load into Y in AbsoluteY mode
func (cpu *CPU) LDXAbsoluteY() {
	lower := uint16(cpu.FetchCode8(1))
	upper := uint16(cpu.FetchCode8(2))
	addr := (upper << 8) | (lower) + uint16(cpu.Reg.Y)
	value := cpu.FetchMemory8(uint(addr))
	cpu.Reg.X = value
	cpu.Reg.PC += 3
}
