package emulator

// BCSRelative 0xb0: if C is 1, jump in relative mode.
func (cpu *CPU) BCSRelative() {
	addr := cpu.RelativeAddressing()

	cFlag := uint8(cpu.Reg.P & 0x01)
	if cFlag > 0 {
		cpu.Reg.PC = uint16(addr) // jump
	}
}

// LDAIndirectIndexed 0xb1: Load into A in Indirect Indexed mode(Y)
func (cpu *CPU) LDAIndirectIndexed() {
	addr := cpu.IndirectIndexedAddressing()
	value := cpu.FetchMemory8(uint(addr))
	cpu.Reg.A = value

	cpu.FlagN(value)
	cpu.FlagZ(value)
}

// LDYZeroPageX 0xb4: Load into Y in ZeroPageX mode
func (cpu *CPU) LDYZeroPageX() {
	addr := cpu.ZeroPageXAddressing()
	value := cpu.FetchMemory8(uint(addr))
	cpu.Reg.Y = value

	cpu.FlagN(value)
	cpu.FlagZ(value)
}

// LDAZeroPageX 0xb5: Load into A in ZeroPageX mode
func (cpu *CPU) LDAZeroPageX() {
	addr := cpu.ZeroPageXAddressing()
	value := cpu.FetchMemory8(uint(addr))
	cpu.Reg.A = value

	cpu.FlagN(value)
	cpu.FlagZ(value)
}

// LDXZeroPageY 0xb6: Load into X in ZeroPageY mode
func (cpu *CPU) LDXZeroPageY() {
	addr := cpu.ZeroPageYAddressing()
	value := cpu.FetchMemory8(uint(addr))
	cpu.Reg.X = value

	cpu.FlagN(value)
	cpu.FlagZ(value)
}

// CLVImplied 0xb8: Clear V flag
func (cpu *CPU) CLVImplied() {
	cpu.Reg.P = cpu.Reg.P & 0xbf // 0b1011_1111
	cpu.Reg.PC++
}

// LDAAbsoluteY 0xb9: Load into A in AbsoluteY mode
func (cpu *CPU) LDAAbsoluteY() {
	addr := cpu.AbsoluteYAddressing()
	value := cpu.FetchMemory8(uint(addr))
	cpu.Reg.A = value

	cpu.FlagN(value)
	cpu.FlagZ(value)
}

// TSXImplied 0xba: Transfer S into X
func (cpu *CPU) TSXImplied() {
	cpu.Reg.X = cpu.Reg.S
	cpu.Reg.PC++

	cpu.FlagN(cpu.Reg.X)
	cpu.FlagZ(cpu.Reg.X)
}

// LDYAbsoluteX 0xbc: Load into Y in AbsoluteX mode
func (cpu *CPU) LDYAbsoluteX() {
	addr := cpu.AbsoluteXAddressing()
	value := cpu.FetchMemory8(uint(addr))
	cpu.Reg.Y = value

	cpu.FlagN(value)
	cpu.FlagZ(value)
}

// LDAAbsoluteX 0xbd: Load into A in AbsoluteX mode
func (cpu *CPU) LDAAbsoluteX() {
	addr := cpu.AbsoluteXAddressing()
	value := cpu.FetchMemory8(uint(addr))
	cpu.Reg.A = value

	cpu.FlagN(value)
	cpu.FlagZ(value)
}

// LDXAbsoluteY 0xbe: Load into Y in AbsoluteY mode
func (cpu *CPU) LDXAbsoluteY() {
	addr := cpu.AbsoluteYAddressing()
	value := cpu.FetchMemory8(uint(addr))
	cpu.Reg.X = value

	cpu.FlagN(value)
	cpu.FlagZ(value)
}
