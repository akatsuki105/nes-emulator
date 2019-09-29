package emulator

// BRKImplied 0x00: Break
func (cpu *CPU) BRKImplied() {
	iFlag := cpu.Reg.P & 0x04
	cpu.Reg.PC++
	if iFlag == 0 {
		// allow BRK
		cpu.Reg.P = cpu.Reg.P | 0x10 // set B Flag
		cpu.Reg.PC++

		// push PC & SR
		upper0 := byte((cpu.Reg.PC) >> 8)
		lower0 := byte((cpu.Reg.PC))
		cpu.SetMemory8((0x100 + uint(cpu.Reg.S)), upper0)
		cpu.SetMemory8((0x100 + uint(cpu.Reg.S) + 1), lower0)
		cpu.SetMemory8((0x100 + uint(cpu.Reg.S) + 2), cpu.Reg.P)
		cpu.Reg.S += 3

		cpu.Reg.P = cpu.Reg.P | 0x04 // set I Flag

		upper1 := uint16(cpu.FetchMemory8(0xffff))
		lower1 := uint16(cpu.FetchMemory8(0xfffe))
		cpu.Reg.PC = (upper1 << 8) | lower1
	}
}

// ORAIndexedIndirect 0x01
func (cpu *CPU) ORAIndexedIndirect() {
	addr := cpu.IndexedIndirectAddressing()
	cpu.ORA(addr)
}

// ORAZeroPage 0x05
func (cpu *CPU) ORAZeroPage() {
	addr := cpu.ZeroPageAddressing()
	cpu.ORA(addr)
}

// ASLZeroPage 0x06
func (cpu *CPU) ASLZeroPage() {
	addr := cpu.ZeroPageAddressing()
	cpu.ASL(addr)
}

// PHPImplied 0x08
func (cpu *CPU) PHPImplied() {
	cpu.SetMemory8((0x100 + uint(cpu.Reg.S)), cpu.Reg.P)
	cpu.Reg.S++
	cpu.Reg.PC++
}

// ORAImmediate 0x09
func (cpu *CPU) ORAImmediate() {
	addr := cpu.ImmediateAddressing()
	cpu.ORA(addr)
}

// ASLAccumulator 0x0a
func (cpu *CPU) ASLAccumulator() {
	cpu.Reg.P = cpu.Reg.P | ((cpu.Reg.A & 0x80) >> 7) // Aのbit7をcにセット
	cpu.Reg.A = cpu.Reg.A << 1
	cpu.Reg.A = cpu.Reg.A | 0 // Aのbit0に0をセット
	cpu.FlagN(cpu.Reg.A)
	cpu.FlagZ(cpu.Reg.A)
}

// ORAAbsolute 0x0d
func (cpu *CPU) ORAAbsolute() {
	addr := cpu.AbsoluteAddressing()
	cpu.ORA(addr)
}

// ASLAbsolute 0x0e
func (cpu *CPU) ASLAbsolute() {
	addr := cpu.AbsoluteAddressing()
	cpu.ASL(addr)
}
