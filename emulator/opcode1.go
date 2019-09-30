package emulator

// BPLRelative 0x10
func (cpu *CPU) BPLRelative() {
	addr := cpu.RelativeAddressing()

	nFlag := uint8(cpu.Reg.P & 0x80)
	if nFlag == 0 {
		cpu.Reg.PC = addr
	}
}

// ORAIndirectIndexed 0x11
func (cpu *CPU) ORAIndirectIndexed() {
	addr := cpu.IndirectIndexedAddressing()
	cpu.ORA(addr)
}

// ORAZeroPageX 0x15
func (cpu *CPU) ORAZeroPageX() {
	addr := cpu.ZeroPageXAddressing()
	cpu.ORA(addr)
}

// ASLZeroPageX 0x16
func (cpu *CPU) ASLZeroPageX() {
	addr := cpu.ZeroPageXAddressing()
	cpu.ASL(addr)
}

// CLCImplied 0x18
func (cpu *CPU) CLCImplied() {
	cpu.Reg.P = cpu.Reg.P & 0xfe
	cpu.Reg.PC++
}

// ORAAbsoluteY 0x19
func (cpu *CPU) ORAAbsoluteY() {
	addr := cpu.AbsoluteYAddressing()
	cpu.ORA(addr)
}

// ORAAbsoluteX 0x1d
func (cpu *CPU) ORAAbsoluteX() {
	addr := cpu.AbsoluteXAddressing()
	cpu.ORA(addr)
}

// ASLAbsoluteX 0x1e
func (cpu *CPU) ASLAbsoluteX() {
	addr := cpu.AbsoluteXAddressing()
	cpu.ASL(addr)
}
