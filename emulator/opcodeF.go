package emulator

// BEQRelative 0xf0
func (cpu *CPU) BEQRelative() {
	addr := cpu.RelativeAddressing()

	zFlag := uint8(cpu.Reg.P & 0x02)
	if zFlag > 0 {
		cpu.Reg.PC = addr
	}
}

// SBCIndirectIndexed 0xf1
func (cpu *CPU) SBCIndirectIndexed() {
	addr := cpu.IndirectIndexedAddressing()
	cpu.SBC(addr)
}

// SBCZeroPageX 0xf5
func (cpu *CPU) SBCZeroPageX() {
	addr := cpu.ZeroPageXAddressing()
	cpu.SBC(addr)
}

// INCZeroPageX 0xf6
func (cpu *CPU) INCZeroPageX() {
	addr := cpu.ZeroPageXAddressing()
	cpu.INC(addr)
}

// SEDImplied 0xf8: Set Decimal mode
func (cpu *CPU) SEDImplied() {
	cpu.Reg.PC++
	cpu.Reg.P = cpu.Reg.P | 0x08
}

// SBCAbsoluteY 0xf9
func (cpu *CPU) SBCAbsoluteY() {
	addr := cpu.AbsoluteYAddressing()
	cpu.SBC(addr)
}

// SBCAbsoluteX 0xfd
func (cpu *CPU) SBCAbsoluteX() {
	addr := cpu.AbsoluteXAddressing()
	cpu.SBC(addr)
}

// INCAbsoluteX 0xfe
func (cpu *CPU) INCAbsoluteX() {
	addr := cpu.AbsoluteXAddressing()
	cpu.INC(addr)
}