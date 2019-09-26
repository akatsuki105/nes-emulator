package emulator

// NOP No operation
func (cpu *CPU) NOP() {
	cpu.Reg.PC++
}
