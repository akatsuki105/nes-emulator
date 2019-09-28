package emulator

// NOP No operation
func (cpu *CPU) NOP() {
}

// CMP Compare M and A (A - M)
func (cpu *CPU) CMP(immediate byte) {
	value := cpu.Reg.A - immediate
	value16 := uint16(cpu.Reg.A) - uint16(immediate)

	cpu.FlagN(value)
	cpu.FlagZ(value)
	cpu.FlagC(value16)
}

// CPX Compare M and X (X - M)
func (cpu *CPU) CPX(immediate byte) {
	value := cpu.Reg.X - immediate
	value16 := uint16(cpu.Reg.X) - uint16(immediate)

	cpu.FlagN(value)
	cpu.FlagZ(value)
	cpu.FlagC(value16)
}

// CPY Compare M and Y (Y - M)
func (cpu *CPU) CPY(immediate byte) {
	value := cpu.Reg.Y - immediate
	value16 := uint16(cpu.Reg.Y) - uint16(immediate)

	cpu.FlagN(value)
	cpu.FlagZ(value)
	cpu.FlagC(value16)
}

// DEC Decrement M by one (M - 1 -> M)
func (cpu *CPU) DEC(addr uint) {
	value := cpu.FetchMemory8(addr) - 1
	cpu.SetMemory8(addr, value)

	cpu.FlagN(value)
	cpu.FlagZ(value)
}

// DEX Decrement X by one (X - 1 -> X)
func (cpu *CPU) DEX() {
	cpu.Reg.X--
	cpu.FlagN(cpu.Reg.X)
	cpu.FlagZ(cpu.Reg.X)
}

// DEY Decrement Y by one (Y - 1 -> Y)
func (cpu *CPU) DEY() {
	cpu.Reg.Y--
	cpu.FlagN(cpu.Reg.Y)
	cpu.FlagZ(cpu.Reg.Y)
}

// LDA Load A from M (M -> A)
func (cpu *CPU) LDA(immediate byte) {
	cpu.Reg.A = immediate

	cpu.FlagN(immediate)
	cpu.FlagZ(immediate)
}

// LDX Load X from M (M -> X)
func (cpu *CPU) LDX(immediate byte) {
	cpu.Reg.X = immediate

	cpu.FlagN(immediate)
	cpu.FlagZ(immediate)
}

// LDY Load Y from M (M -> Y)
func (cpu *CPU) LDY(immediate byte) {
	cpu.Reg.Y = immediate

	cpu.FlagN(immediate)
	cpu.FlagZ(immediate)
}

// TAX Transfer A to X
func (cpu *CPU) TAX() {
	cpu.Reg.X = cpu.Reg.A
	cpu.FlagN(cpu.Reg.X)
	cpu.FlagZ(cpu.Reg.X)
}

// TAY Transfer A to Y
func (cpu *CPU) TAY() {
	cpu.Reg.Y = cpu.Reg.A
	cpu.FlagN(cpu.Reg.Y)
	cpu.FlagZ(cpu.Reg.Y)
}

// TXA Transfer X to A
func (cpu *CPU) TXA() {
	cpu.Reg.A = cpu.Reg.X
	cpu.FlagN(cpu.Reg.A)
	cpu.FlagZ(cpu.Reg.A)
}

// TYA Transfer Y to A
func (cpu *CPU) TYA() {
	cpu.Reg.A = cpu.Reg.Y
	cpu.FlagN(cpu.Reg.A)
	cpu.FlagZ(cpu.Reg.A)
}

// TSX Transfer S to X
func (cpu *CPU) TSX() {
	cpu.Reg.X = cpu.Reg.S
	cpu.FlagN(cpu.Reg.X)
	cpu.FlagZ(cpu.Reg.X)
}

// INY Increment Y by one (Y + 1 -> Y)
func (cpu *CPU) INY() {
	cpu.Reg.Y++
	cpu.FlagN(cpu.Reg.Y)
	cpu.FlagZ(cpu.Reg.Y)
}

// ROR Rotate right one bit
func (cpu *CPU) ROR(addr uint) {
	value := cpu.FetchMemory8(addr)
	cFlag := cpu.Reg.P & 0x01
	cpu.Reg.P = cpu.Reg.P | (value & 0x01) // valueのbit0をcにセット
	value = value >> 1
	value = value | (cFlag << 7) // valueのbit7にcをセット
	cpu.SetMemory8(addr, value)
	cpu.FlagN(value)
	cpu.FlagZ(value)
}

// ADC Add M to A with C (A + M + C -> A)
func (cpu *CPU) ADC(addr uint) {
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
