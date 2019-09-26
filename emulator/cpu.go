package emulator

var (
	prgRomPageSize int = 16 * 1024 // プログラムROMのページサイズ
	chrRomPageSize int = 8 * 1024  // キャラクタROMのページサイズ
)

// CPURegister CPUのレジスタです。
type cpuRegister struct {
	A  byte   // アキュムレータ
	X  byte   // インデックスレジスタ
	Y  byte   // インデックスレジスタ
	S  byte   // スタックポインタ
	P  byte   // ステータスレジスタ
	PC uint16 // プログラムカウンタ
}

// CPU Central Processing Unit
type CPU struct {
	Reg    cpuRegister
	RAM    [0x10000]byte
	PPURAM [0x4000]byte
}

// CPURAM CPUからアクセスできるメモリマップ
type CPURAM struct {
	wRAM         [0x0800]byte // 0x0000-0x07FF
	wRAMMirror   [0x1800]byte // 0x0800-0x1FFF
	ppuReg       [0x0008]byte // 0x2000-0x2007
	ppuRegMirror [0x1ff8]byte // 0x2008-0x3FFF
	apu          [0x0020]byte // 0x4000-0x401F
	exROM        [0x1fe0]byte // 0x4020-0x5FFF
	exRAM        [0x2000]byte // 0x6000-0x7FFF
	prgROM0      [0x4000]byte // 0x8000-0xBFFF
	prgROM1      [0x4000]byte // 0xC000-0xFFFF
}

// InitIRQVector 割り込みベクタの初期化
func (cpu *CPU) InitIRQVector() {
	cpu.RAM[RESET0] = 0x00
	cpu.RAM[RESET1] = 0x80
}

// InitReg レジスタの初期化
func (cpu *CPU) InitReg() {
	cpu.Reg.PC = 0x8000
}

// LoadROM ROMのバイトデータからプログラムROMとページROMを取り出してRAMにロードする
func (cpu *CPU) LoadROM(rom []byte) {
	prgAddr := 0x0010
	prgPage := rom[4]

	chrAddr := prgAddr + int(prgPage)*prgRomPageSize
	chrPage := rom[5]

	prgBytes := rom[prgAddr : prgAddr+int(prgPage)*prgRomPageSize]
	chrBytes := rom[chrAddr : chrAddr+int(chrPage)*chrRomPageSize]

	// プログラムROMを0x8000~に配置
	for i := 0; i < len(prgBytes); i++ {
		cpu.RAM[0x8000+i] = prgBytes[i]
	}

	// キャラクタROMをPPUの0x0000~に配置
	for i := 0; i < len(chrBytes); i++ {
		cpu.PPURAM[i] = chrBytes[i]
	}
}

// FetchCode8 メモリから次の命令をフェッチする PCのインクリメントはしない
func (cpu *CPU) FetchCode8(index uint) byte {
	code := cpu.RAM[(uint)(cpu.Reg.PC)+index]
	return code
}

// FetchMemory8 引数で指定したアドレスから値を取得する
func (cpu *CPU) FetchMemory8(addr uint) byte {
	value := cpu.RAM[addr]
	return value
}

// SetMemory8 引数で指定したアドレスにvalueを書き込む
func (cpu *CPU) SetMemory8(addr uint, value byte) {
	cpu.RAM[addr] = value
}

// Reset エミュレータをリセットする
func (cpu *CPU) Reset() {
	// TODO メモリの初期化
	// TODO 画面の初期化
	cpu.InitReg()
}
