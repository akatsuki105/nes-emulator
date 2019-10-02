package emulator

const (
	// 何もないことを表す
	null    uint16 = 0x2008
	joypad1 uint16 = 0x4016
	joypad2 uint16 = 0x4017
)

var (
	// colors {R, G, B}
	colors [64][3]uint8 = [64][3]uint8{
		{0x80, 0x80, 0x80}, {0x00, 0x3D, 0xA6}, {0x00, 0x12, 0xB0}, {0x44, 0x00, 0x96},
		{0xA1, 0x00, 0x5E}, {0xC7, 0x00, 0x28}, {0xBA, 0x06, 0x00}, {0x8C, 0x17, 0x00},
		{0x5C, 0x2F, 0x00}, {0x10, 0x45, 0x00}, {0x05, 0x4A, 0x00}, {0x00, 0x47, 0x2E},
		{0x00, 0x41, 0x66}, {0x00, 0x00, 0x00}, {0x05, 0x05, 0x05}, {0x05, 0x05, 0x05},
		{0xC7, 0xC7, 0xC7}, {0x00, 0x77, 0xFF}, {0x21, 0x55, 0xFF}, {0x82, 0x37, 0xFA},
		{0xEB, 0x2F, 0xB5}, {0xFF, 0x29, 0x50}, {0xFF, 0x22, 0x00}, {0xD6, 0x32, 0x00},
		{0xC4, 0x62, 0x00}, {0x35, 0x80, 0x00}, {0x05, 0x8F, 0x00}, {0x00, 0x8A, 0x55},
		{0x00, 0x99, 0xCC}, {0x21, 0x21, 0x21}, {0x09, 0x09, 0x09}, {0x09, 0x09, 0x09},
		{0xFF, 0xFF, 0xFF}, {0x0F, 0xD7, 0xFF}, {0x69, 0xA2, 0xFF}, {0xD4, 0x80, 0xFF},
		{0xFF, 0x45, 0xF3}, {0xFF, 0x61, 0x8B}, {0xFF, 0x88, 0x33}, {0xFF, 0x9C, 0x12},
		{0xFA, 0xBC, 0x20}, {0x9F, 0xE3, 0x0E}, {0x2B, 0xF0, 0x35}, {0x0C, 0xF0, 0xA4},
		{0x05, 0xFB, 0xFF}, {0x5E, 0x5E, 0x5E}, {0x0D, 0x0D, 0x0D}, {0x0D, 0x0D, 0x0D},
		{0xFF, 0xFF, 0xFF}, {0xA6, 0xFC, 0xFF}, {0xB3, 0xEC, 0xFF}, {0xDA, 0xAB, 0xEB},
		{0xFF, 0xA8, 0xF9}, {0xFF, 0xAB, 0xB3}, {0xFF, 0xD2, 0xB0}, {0xFF, 0xEF, 0xA6},
		{0xFF, 0xF7, 0x9C}, {0xD7, 0xE8, 0x95}, {0xA6, 0xED, 0xAF}, {0xA2, 0xF2, 0xDA},
		{0x99, 0xFF, 0xFC}, {0xDD, 0xDD, 0xDD}, {0x11, 0x11, 0x11}, {0x11, 0x11, 0x11},
	}

	// instructions {"Opcode", "Addressing"}
	instructions [256][2]string = [256][2]string{
		{"BRK", "impl"}, {"ORA", "X,ind"}, {"*", "*"}, {"*", "*"}, {"*", "*"}, {"ORA", "zpg"}, {"ASL", "zpg"}, {"*", "*"}, {"PHP", "impl"}, {"ORA", "#"}, {"ASL", "A"}, {"*", "*"}, {"*", "*"}, {"ORA", "abs"}, {"ASL", "abs"}, {"*", "*"},
		{"BPL", "rel"}, {"ORA", "ind,Y"}, {"*", "*"}, {"*", "*"}, {"*", "*"}, {"ORA", "zpg,X"}, {"ASL", "zpg,X"}, {"*", "*"}, {"CLC", "impl"}, {"ORA", "abs,Y"}, {"*", "*"}, {"*", "*"}, {"*", "*"}, {"ORA", "abs,X"}, {"ASL", "abs,X"}, {"*", "*"},
		{"JSR", "abs"}, {"AND", "X,ind"}, {"*", "*"}, {"*", "*"}, {"BIT", "zpg"}, {"AND", "zpg"}, {"ROL", "zpg"}, {"*", "*"}, {"PLP", "impl"}, {"AND", "#"}, {"ROL", "A"}, {"*", "*"}, {"BIT", "abs"}, {"AND", "abs"}, {"ROL", "abs"}, {"*", "*"},
		{"BMI", "rel"}, {"AND", "ind,Y"}, {"*", "*"}, {"*", "*"}, {"*", "*"}, {"AND", "zpg,X"}, {"ROL", "zpg,X"}, {"*", "*"}, {"SEC", "impl"}, {"AND", "abs,Y"}, {"*", "*"}, {"*", "*"}, {"*", "*"}, {"AND", "abs,X"}, {"ROL", "abs,X"}, {"*", "*"},
		{"RTI", "impl"}, {"EOR", "X,ind"}, {"*", "*"}, {"*", "*"}, {"*", "*"}, {"EOR", "zpg"}, {"LSR", "zpg"}, {"*", "*"}, {"PHP", "impl"}, {"EOR", "#"}, {"LSR", "A"}, {"*", "*"}, {"JMP", "abs"}, {"EOR", "abs"}, {"LSR", "abs"}, {"*", "*"},
		{"BVC", "rel"}, {"EOR", "ind,Y"}, {"*", "*"}, {"*", "*"}, {"*", "*"}, {"EOR", "zpg,X"}, {"LSR", "zpg,X"}, {"*", "*"}, {"CLI", "impl"}, {"EOR", "abs,Y"}, {"*", "*"}, {"*", "*"}, {"*", "*"}, {"EOR", "abs,X"}, {"LSR", "abs,X"}, {"*", "*"},
		{"RTS", "impl"}, {"ADC", "X,ind"}, {"*", "*"}, {"*", "*"}, {"*", "*"}, {"ADC", "zpg"}, {"ROR", "zpg"}, {"*", "*"}, {"PLA", "impl"}, {"ADC", "#"}, {"ROR", "A"}, {"*", "*"}, {"JMP", "ind"}, {"ADC", "abs"}, {"ROR", "abs"}, {"*", "*"},
		{"BVS", "rel"}, {"ADC", "ind,Y"}, {"*", "*"}, {"*", "*"}, {"*", "*"}, {"ADC", "zpg,X"}, {"ROR", "zpg,X"}, {"*", "*"}, {"SEI", "impl"}, {"ADC", "abs,Y"}, {"*", "*"}, {"*", "*"}, {"*", "*"}, {"ADC", "abs,X"}, {"ROR", "abs,X"}, {"*", "*"},
		{"*", "*"}, {"STA", "ind"}, {"*", "*"}, {"*", "*"}, {"STY", "zpg"}, {"STA", "zpg"}, {"STX", "zpg"}, {"*", "*"}, {"DEY", "impl"}, {"*", "*"}, {"TAX", "impl"}, {"*", "*"}, {"STY", "abs"}, {"STA", "abs"}, {"STX", "abs"}, {"*", "*"},
		{"BCC", "rel"}, {"STA", "ind,Y"}, {"*", "*"}, {"*", "*"}, {"STY", "zpg,X"}, {"STA", "zpg,X"}, {"STX", "zpg,Y"}, {"*", "*"}, {"TYA", "impl"}, {"STA", "abs,Y"}, {"TXS", "impl"}, {"*", "*"}, {"*", "*"}, {"STA", "abs,X"}, {"*", "*"}, {"*", "*"}, // 0x90
		{"LDY", "#"}, {"LDA", "X,ind"}, {"LDX", "#"}, {"*", "*"}, {"LDY", "zpg"}, {"LDA", "zpg"}, {"LDX", "zpg"}, {"*", "*"}, {"TAY", "impl"}, {"LDA", "#"}, {"TAX", "impl"}, {"*", "*"}, {"LDY", "abs"}, {"LDA", "abs"}, {"LDX", "abs"}, {"*", "*"},
		{"BCS", "rel"}, {"LDA", "ind,Y"}, {"*", "*"}, {"*", "*"}, {"LDY", "zpg,X"}, {"LDA", "zpg,X"}, {"LDX", "zpg,Y"}, {"*", "*"}, {"CLV", "impl"}, {"LDA", "abs,Y"}, {"TSX", "impl"}, {"*", "*"}, {"LDY", "abs,X"}, {"LDA", "abs,X"}, {"LDX", "abs,Y"}, {"*", "*"},
		{"CPY", "#"}, {"CMP", "X,ind"}, {"*", "*"}, {"*", "*"}, {"CPY", "zpg"}, {"CMP", "zpg"}, {"DEC", "zpg"}, {"*", "*"}, {"INY", "impl"}, {"CMP", "#"}, {"DEX", "impl"}, {"*", "*"}, {"CPY", "abs"}, {"CMP", "abs"}, {"DEC", "abs"}, {"*", "*"},
		{"BNE", "rel"}, {"CMP", "ind,Y"}, {"*", "*"}, {"*", "*"}, {"*", "*"}, {"CMP", "zpg,X"}, {"DEC", "zpg,X"}, {"*", "*"}, {"CLD", "impl"}, {"CMP", "abs,Y"}, {"*", "*"}, {"*", "*"}, {"*", "*"}, {"CMP", "abs,X"}, {"DEC", "abs,X"}, {"*", "*"},
		{"CPX", "#"}, {"SBC", "X,ind"}, {"*", "*"}, {"*", "*"}, {"CPX", "zpg"}, {"SBC", "zpg"}, {"INC", "zpg"}, {"*", "*"}, {"INX", "impl"}, {"SBC", "#"}, {"NOP", "impl"}, {"*", "*"}, {"CPX", "abs"}, {"SBC", "abs"}, {"INC", "abs"}, {"*", "*"},
		{"BEQ", "rel"}, {"SBC", "ind,Y"}, {"*", "*"}, {"*", "*"}, {"*", "*"}, {"SBC", "zpg,X"}, {"INC", "zpg,X"}, {"*", "*"}, {"SED", "impl"}, {"SBC", "abs,Y"}, {"*", "*"}, {"*", "*"}, {"*", "*"}, {"SBC", "abs,X"}, {"INC", "abs,X"}, {"*", "*"},
	}
)
