package emulator

// JSRAbsolute 0x20: Jump to new location saving return address
func (cpu *CPU) JSRAbsolute() {
	addr := cpu.AbsoluteAddressing()
	upper := byte((cpu.Reg.PC - 1) >> 8)
	lower := byte((cpu.Reg.PC - 1))
	cpu.SetMemory8((0x100 + uint(cpu.Reg.S)), upper)
	cpu.SetMemory8((0x100 + uint(cpu.Reg.S) + 1), lower)
	cpu.Reg.S += 2
	cpu.Reg.PC = uint16(addr)
}

// ANDIndexedIndirect 0x21
func (cpu *CPU) ANDIndexedIndirect() {
	addr := cpu.IndexedIndirectAddressing()
	cpu.AND(addr)
}

// BITZeroPage 0x24
func (cpu *CPU) BITZeroPage() {
	addr := cpu.ZeroPageAddressing()
	cpu.BIT(addr)
}

// ANDZeroPage 0x25
func (cpu *CPU) ANDZeroPage() {
	addr := cpu.ZeroPageAddressing()
	cpu.AND(addr)
}

// ROLZeroPage 0x26
func (cpu *CPU) ROLZeroPage() {
	addr := cpu.ZeroPageAddressing()
	cpu.ROL(addr)
}

// PLPImplied 0x28: Pull P from stack (stack -> P)
func (cpu *CPU) PLPImplied() {
	value := cpu.FetchMemory8(0x0100 + uint(cpu.Reg.S) - 1)
	cpu.Reg.P = value // pullがフラグのセットになっている
	cpu.Reg.S--
	cpu.Reg.PC++
}

// ANDImmediate 0x29
func (cpu *CPU) ANDImmediate() {
	addr := cpu.ImmediateAddressing()
	cpu.AND(addr)
}

// ROLAccumulator 0x2a
func (cpu *CPU) ROLAccumulator() {
	cFlag := cpu.Reg.P & 0x01
	cpu.Reg.P = cpu.Reg.P | ((cpu.Reg.A & 0x80) >> 7) // Aのbit7をcにセット
	cpu.Reg.A = cpu.Reg.A << 1
	cpu.Reg.A = cpu.Reg.A | cFlag // Aのbit0にcをセット
	cpu.FlagN(cpu.Reg.A)
	cpu.FlagZ(cpu.Reg.A)

	cpu.Reg.PC++
}

// BITAbsolute 0x2c
func (cpu *CPU) BITAbsolute() {
	addr := cpu.AbsoluteAddressing()
	cpu.BIT(addr)
}

// ANDAbsolute 0x2d
func (cpu *CPU) ANDAbsolute() {
	addr := cpu.AbsoluteAddressing()
	cpu.AND(addr)
}

// ROLAbsolute 0x2e
func (cpu *CPU) ROLAbsolute() {
	addr := cpu.AbsoluteAddressing()
	cpu.ROL(addr)
}
