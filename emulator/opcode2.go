package emulator

// JSRAbsolute 0x20: Jump to new location saving return address
func (cpu *CPU) JSRAbsolute() {
	addr := cpu.AbsoluteAddressing()
	cpu.JSR(addr)
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
	addr := cpu.ImpliedAddressing()
	cpu.PLP(addr)
}

// ANDImmediate 0x29
func (cpu *CPU) ANDImmediate() {
	addr := cpu.ImmediateAddressing()
	cpu.AND(addr)
}

// ROLAccumulator 0x2a
func (cpu *CPU) ROLAccumulator() {
	addr := cpu.AccumulatorAddressing()
	cpu.ROL(addr)
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
