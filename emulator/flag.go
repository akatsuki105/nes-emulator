package emulator

// setVBlank VBlankにする
func (cpu *CPU) setVBlank() {
	cpu.RAM[0x2002] = cpu.RAM[0x2002] | 0x80
	cpu.NMI(null)
}

// clearVBlank VBlankを解除
func (cpu *CPU) clearVBlank() {
	cpu.RAM[0x2002] = cpu.RAM[0x2002] & 0x7f
}

// FlagN Nフラグを立てるか判定する
func (cpu *CPU) FlagN(b byte) {
	if (b & 0x80) != 0 {
		cpu.Reg.P = cpu.Reg.P | 0x80 // 0b1000_0000
	} else {
		cpu.Reg.P = cpu.Reg.P & 0x7f // 0b0111_1111
	}
}

// FlagV Vフラグを立てるか判定する
func (cpu *CPU) FlagV(b0, b1 byte, u16 uint16) {
	if ((b0>>7)^(b1>>7) != 0) && (uint16(b1) != u16) {
		cpu.Reg.P = cpu.Reg.P | 0x40 // 0b0100_0000
	} else {
		cpu.Reg.P = cpu.Reg.P & 0xbf // 0b1011_1111
	}
}

// FlagZ Zフラグを立てるか判定する
func (cpu *CPU) FlagZ(b byte) {
	if b == 0 {
		cpu.Reg.P = cpu.Reg.P | 0x02 // 0b0000_0010
	} else {
		cpu.Reg.P = cpu.Reg.P & 0xfd // 0b1111_1101
	}
}

// FlagC Cフラグを立てるか判定する
func (cpu *CPU) FlagC(instruction string, u16 uint16) {
	if (u16 >> 8) != 0 {
		if instruction == "ADC" {
			cpu.setCFlag()
		} else {
			cpu.clearCFlag()
		}
	} else {
		if instruction == "ADC" {
			cpu.clearCFlag()
		} else {
			cpu.setCFlag()
		}
	}
}

func (cpu *CPU) setCFlag() {
	cpu.Reg.P = cpu.Reg.P | 0x01 // 0b0000_0001
}

func (cpu *CPU) clearCFlag() {
	cpu.Reg.P = cpu.Reg.P & 0xfe // 0b1111_1110
}
