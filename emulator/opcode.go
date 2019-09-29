package emulator

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

// SBC Subtract M from A with C (A - M - not C -> A)
func (cpu *CPU) SBC(addr uint) {
	notCFlag := ^(cpu.Reg.P) & 0x01
	aFlag := cpu.Reg.A
	value := (cpu.Reg.A - cpu.FetchMemory8(addr) - notCFlag) & (0xff)                // キャリーオーバー対策のため
	value16 := uint16(cpu.Reg.A) - uint16(cpu.FetchMemory8(addr)) - uint16(notCFlag) // Cフラグ判定のため
	cpu.Reg.A = value

	cpu.FlagN(value)
	cpu.FlagV(aFlag, value, value16)
	cpu.FlagZ(value)
	cpu.FlagC(value16)
}

// AND "AND" M with A (A and M -> A)
func (cpu *CPU) AND(addr uint) {
	value := cpu.Reg.A & cpu.FetchMemory8(addr)
	cpu.Reg.A = value

	cpu.FlagN(value)
	cpu.FlagZ(value)
}

// ORA "OR" M with A (A or M -> A)
func (cpu *CPU) ORA(addr uint) {
	value := cpu.Reg.A | cpu.FetchMemory8(addr)
	cpu.Reg.A = value

	cpu.FlagN(value)
	cpu.FlagZ(value)
}

// EOR "Exclusive-OR" M with A (A eor M -> A)
func (cpu *CPU) EOR(addr uint) {
	value := cpu.Reg.A ^ cpu.FetchMemory8(addr)
	cpu.Reg.A = value

	cpu.FlagN(value)
	cpu.FlagZ(value)
}

// ASL Arithmetic shift left one bit
func (cpu *CPU) ASL(addr uint) {
	value := cpu.FetchMemory8(addr)
	cpu.Reg.P = cpu.Reg.P | ((value & 0x80) >> 7) // valueのbit7をcにセット
	value = value << 1
	value = value | 0 // valueのbit0に0をセット
	cpu.SetMemory8(addr, value)
	cpu.FlagN(value)
	cpu.FlagZ(value)
}

// LSR Logical shift right one bit
func (cpu *CPU) LSR(addr uint) {
	value := cpu.FetchMemory8(addr)
	cpu.Reg.P = cpu.Reg.P | (value & 0x01) // valueのbit0をcにセット
	value = value >> 1
	value = value | (0 << 7) // valueのbit7に0をセット
	cpu.SetMemory8(addr, value)
	cpu.FlagN(value)
	cpu.FlagZ(value)
}

// ROL Rotate left one bit
func (cpu *CPU) ROL(addr uint) {
	value := cpu.FetchMemory8(addr)
	cFlag := cpu.Reg.P & 0x01
	cpu.Reg.P = cpu.Reg.P | ((value & 0x80) >> 7) // valueのbit7をcにセット
	value = value << 1
	value = value | cFlag // valueのbit0にcをセット
	cpu.SetMemory8(addr, value)
	cpu.FlagN(value)
	cpu.FlagZ(value)
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

// NOP No operation
func (cpu *CPU) NOP() {
}

// CMP Compare M and A (A - M)
func (cpu *CPU) CMP(addr uint) {
	value := cpu.Reg.A - cpu.FetchMemory8(addr)
	value16 := uint16(cpu.Reg.A) - uint16(cpu.FetchMemory8(addr))

	cpu.FlagN(value)
	cpu.FlagZ(value)
	cpu.FlagC(value16)
}

// CPX Compare M and X (X - M)
func (cpu *CPU) CPX(addr uint) {
	value := cpu.Reg.X - cpu.FetchMemory8(addr)
	value16 := uint16(cpu.Reg.X) - uint16(cpu.FetchMemory8(addr))

	cpu.FlagN(value)
	cpu.FlagZ(value)
	cpu.FlagC(value16)
}

// CPY Compare M and Y (Y - M)
func (cpu *CPU) CPY(addr uint) {
	value := cpu.Reg.Y - cpu.FetchMemory8(addr)
	value16 := uint16(cpu.Reg.Y) - uint16(cpu.FetchMemory8(addr))

	cpu.FlagN(value)
	cpu.FlagZ(value)
	cpu.FlagC(value16)
}

// INC Increment M by one (M + 1 -> M)
func (cpu *CPU) INC(addr uint) {
	value := cpu.FetchMemory8(addr) + 1
	cpu.SetMemory8(addr, value)

	cpu.FlagN(value)
	cpu.FlagZ(value)
}

// DEC Decrement M by one (M - 1 -> M)
func (cpu *CPU) DEC(addr uint) {
	value := cpu.FetchMemory8(addr) - 1
	cpu.SetMemory8(addr, value)

	cpu.FlagN(value)
	cpu.FlagZ(value)
}

// INX Increment X by one (X + 1 -> X)
func (cpu *CPU) INX() {
	cpu.Reg.X++
	cpu.FlagN(cpu.Reg.X)
	cpu.FlagZ(cpu.Reg.X)
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
func (cpu *CPU) LDA(addr uint) {
	cpu.Reg.A = cpu.FetchMemory8(addr)

	cpu.FlagN(cpu.FetchMemory8(addr))
	cpu.FlagZ(cpu.FetchMemory8(addr))
}

// LDX Load X from M (M -> X)
func (cpu *CPU) LDX(addr uint) {
	cpu.Reg.X = cpu.FetchMemory8(addr)

	cpu.FlagN(cpu.FetchMemory8(addr))
	cpu.FlagZ(cpu.FetchMemory8(addr))
}

// LDY Load Y from M (M -> Y)
func (cpu *CPU) LDY(addr uint) {
	cpu.Reg.Y = cpu.FetchMemory8(addr)

	cpu.FlagN(cpu.FetchMemory8(addr))
	cpu.FlagZ(cpu.FetchMemory8(addr))
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

// BIT Test Bits in M with A
func (cpu *CPU) BIT(addr uint) {
	value := cpu.FetchMemory8(addr)

	// NZフラグを立てる
	cpu.FlagZ(value & cpu.Reg.A)
	cpu.FlagN(value)
	// bit6をVフラグに転送
	if (value & 0x40) != 0 {
		cpu.Reg.P = cpu.Reg.P | 0x40 // 0b0100_0000
	} else {
		cpu.Reg.P = cpu.Reg.P & 0xbf // 0b1011_1111
	}
}
