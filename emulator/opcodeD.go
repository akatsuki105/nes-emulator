package emulator

// BNERelative 0xd0
func (cpu *CPU) BNERelative() {
	addr := cpu.RelativeAddressing()

	zFlag := uint8(cpu.Reg.P & 0x02)
	if zFlag == 0 {
		cpu.Reg.PC = addr
	}
}

// CMPIndirectIndexed 0xd1
func (cpu *CPU) CMPIndirectIndexed() {
	addr := cpu.IndirectIndexedAddressing()
	cpu.CMP(addr)
}

// CMPZeroPageX 0xd5
func (cpu *CPU) CMPZeroPageX() {
	addr := cpu.ZeroPageXAddressing()
	cpu.CMP(addr)
}

// DECZeroPageX 0xd6
func (cpu *CPU) DECZeroPageX() {
	addr := cpu.ZeroPageXAddressing()
	cpu.DEC(addr)
}

// CLDImplied 0xc8: Clear D Flag
func (cpu *CPU) CLDImplied() {
	cpu.Reg.PC++
	cpu.Reg.P = cpu.Reg.P & 0xfb // 0b1111_1011
}

// CMPAbsoluteY 0xc9
func (cpu *CPU) CMPAbsoluteY() {
	addr := cpu.AbsoluteYAddressing()
	cpu.CMP(addr)
}