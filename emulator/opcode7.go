package emulator

// BVSRelative 0x70
func (cpu *CPU) BVSRelative() {
	addr := cpu.RelativeAddressing()

	vFlag := uint8(cpu.Reg.P & 0x40)
	if vFlag > 0 {
		cpu.Reg.PC = uint16(addr)
	}
}

// ADCIndirectIndexed 0x71
func (cpu *CPU) ADCIndirectIndexed() {
	addr := cpu.IndirectIndexedAddressing()

	cFlag := cpu.Reg.P & 0x01
	aFlag := cpu.Reg.A
	value := (cpu.Reg.A + cpu.FetchMemory8(addr) + cFlag) & (0xff)                // キャリーオーバー対策のため
	value16 := uint16(cpu.Reg.A) + uint16(cpu.FetchMemory8(addr)) + uint16(cFlag) // Cフラグ判定のため
	cpu.Reg.A = value

	cpu.FlagN(value)
	cpu.FlagV(aFlag, value, value16)
	cpu.FlagZ(value)
	cpu.FlagC(value16)
}

// ADCZeroPageX 0x75
func (cpu *CPU) ADCZeroPageX() {
	addr := cpu.ZeroPageXAddressing()

	cFlag := cpu.Reg.P & 0x01
	aFlag := cpu.Reg.A
	value := (cpu.Reg.A + cpu.FetchMemory8(addr) + cFlag) & (0xff)                // キャリーオーバー対策のため
	value16 := uint16(cpu.Reg.A) + uint16(cpu.FetchMemory8(addr)) + uint16(cFlag) // Cフラグ判定のため
	cpu.Reg.A = value

	cpu.FlagN(value)
	cpu.FlagV(aFlag, value, value16)
	cpu.FlagZ(value)
	cpu.FlagC(value16)
}

// RORZeroPageX 0x76
func (cpu *CPU) RORZeroPageX() {
	addr := cpu.ZeroPageXAddressing()

	value := cpu.FetchMemory8(addr)
	cFlag := cpu.Reg.P & 0x01
	cpu.Reg.P = cpu.Reg.P | (value & 0x01) // valueのbit0をcにセット
	value = value >> 1
	value = value | (cFlag << 7) // valueのbit7にcをセット
	cpu.SetMemory8(addr, value)

	cpu.FlagN(value)
	cpu.FlagZ(value)
}

// SEIImplied 0x78
func (cpu *CPU) SEIImplied() {
	cpu.Reg.P = cpu.Reg.P | 0x04
}

// ADCAbsoluteY 0x79
func (cpu *CPU) ADCAbsoluteY() {
	addr := cpu.AbsoluteYAddressing()

	cFlag := cpu.Reg.P & 0x01
	aFlag := cpu.Reg.A
	value := (cpu.Reg.A + cpu.FetchMemory8(addr) + cFlag) & (0xff)                // キャリーオーバー対策のため
	value16 := uint16(cpu.Reg.A) + uint16(cpu.FetchMemory8(addr)) + uint16(cFlag) // Cフラグ判定のため
	cpu.Reg.A = value

	cpu.FlagN(value)
	cpu.FlagV(aFlag, value, value16)
	cpu.FlagZ(value)
	cpu.FlagC(value16)
}

// ADCAbsoluteX 0x7d
func (cpu *CPU) ADCAbsoluteX() {
	addr := cpu.AbsoluteXAddressing()

	cFlag := cpu.Reg.P & 0x01
	aFlag := cpu.Reg.A
	value := (cpu.Reg.A + cpu.FetchMemory8(addr) + cFlag) & (0xff)                // キャリーオーバー対策のため
	value16 := uint16(cpu.Reg.A) + uint16(cpu.FetchMemory8(addr)) + uint16(cFlag) // Cフラグ判定のため
	cpu.Reg.A = value

	cpu.FlagN(value)
	cpu.FlagV(aFlag, value, value16)
	cpu.FlagZ(value)
	cpu.FlagC(value16)
}

// RORAbsoluteX 0x7e
func (cpu *CPU) RORAbsoluteX() {
	addr := cpu.AbsoluteXAddressing()

	value := cpu.FetchMemory8(addr)
	cFlag := cpu.Reg.P & 0x01
	cpu.Reg.P = cpu.Reg.P | (value & 0x01) // valueのbit0をcにセット
	value = value >> 1
	value = value | (cFlag << 7) // valueのbit7にcをセット
	cpu.SetMemory8(addr, value)

	cpu.FlagN(value)
	cpu.FlagZ(value)
}
