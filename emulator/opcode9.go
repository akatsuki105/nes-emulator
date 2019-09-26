package emulator

// BCCRelative 0x90: if C is 1, jump in relative mode.
func (cpu *CPU) BCCRelative() {
	cFlag := uint8(cpu.Reg.P & 0x01)
	addr := int8(cpu.Reg.PC) + 1 + int8(cpu.FetchCode8(1))
	if cFlag == 0 {
		cpu.Reg.PC = uint16(addr)
	} else {
		cpu.Reg.PC += 2
	}
}

// STAIndirectIndexed 0x91: Store A into M in Indirect Indexed mode(Y)
func (cpu *CPU) STAIndirectIndexed() {
	addr := cpu.IndirectIndexedAddressing()
	cpu.SetMemory8(addr, cpu.Reg.A)
}

// STYZeroPageX 0x94: Store Y into M in ZeroPageX mode
func (cpu *CPU) STYZeroPageX() {
	addr := cpu.ZeroPageXAddressing()
	cpu.SetMemory8(addr, cpu.Reg.Y)
}

// STAZeroPageX 0x95: Store A into M in ZeroPageX mode
func (cpu *CPU) STAZeroPageX() {
	addr := cpu.ZeroPageXAddressing()
	cpu.SetMemory8(addr, cpu.Reg.A)
}

// STXZeroPageY 0x96: Store X into M in ZeroPageY mode
func (cpu *CPU) STXZeroPageY() {
	addr := cpu.ZeroPageYAddressing()
	cpu.SetMemory8(addr, cpu.Reg.X)
}

// TYAImplied 0x98: Transfer Y into A in Implied mode
func (cpu *CPU) TYAImplied() {
	cpu.Reg.A = cpu.Reg.Y
	cpu.Reg.PC++
}

// STAAbsoluteY 0x99: Store A into M in AbsoluteY mode
func (cpu *CPU) STAAbsoluteY() {
	addr := cpu.AbsoluteYAddressing()
	cpu.SetMemory8(addr, cpu.Reg.A)
}

// TXSImplied 0x9a: Transfer X into S in Implied mode
func (cpu *CPU) TXSImplied() {
	cpu.Reg.S = cpu.Reg.X
	cpu.Reg.PC++
}

// STAAbsoluteX 0x9d: Store A into M in AbsoluteX mode
func (cpu *CPU) STAAbsoluteX() {
	addr := cpu.AbsoluteXAddressing()
	cpu.SetMemory8(addr, cpu.Reg.A)
}
