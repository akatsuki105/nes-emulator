package emulator

// FetchMemory8 引数で指定したアドレスから値を取得する
func (cpu *CPU) FetchMemory8(addr uint16) byte {
	value := cpu.RAM[addr]
	return value
}

// SetMemory8 引数で指定したアドレスにvalueを書き込む
func (cpu *CPU) SetMemory8(addr uint16, value byte) {
	cpu.RAM[addr] = value
}
