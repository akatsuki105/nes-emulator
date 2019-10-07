package emulator

import (
	"sync"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	width  = 256
	height = 240
)

var (
	lineWait  sync.WaitGroup
	lineMutex sync.Mutex
)

// Render 画面描画を行う
func (cpu *CPU) Render() {
	cfg := pixelgl.WindowConfig{
		Title:  "nes-emulator",
		Bounds: pixel.R(0, 0, width, height),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	go cpu.handleJoypad(win)

	cpu.PPU.CacheBG()
	BGBatch := pixel.NewBatch(&pixel.TrianglesData{}, cpu.PPU.BGBuf)

	cpu.PPU.CacheSPR()
	SPRBatch := pixel.NewBatch(&pixel.TrianglesData{}, cpu.PPU.SPRBuf)

	for !win.Closed() {
		// SPR探索
		var pixel2sprite map[uint16]([2]byte)
		pixel2sprite = map[uint16]([2]byte){}
		for i := 0; i < 64; i++ {
			pixelX, pixelY := cpu.PPU.sRAM[i*4+3], (cpu.PPU.sRAM[i*4])+1
			spriteNum := cpu.PPU.sRAM[i*4+1]
			attr := cpu.PPU.sRAM[i*4+2]
			pixel2sprite[(uint16(pixelY)<<8)|uint16(pixelX)] = [2]byte{spriteNum, attr}
		}

		// BG・SPR描画
		BGBatch.Clear()
		SPRBatch.Clear()
		for y := 0; y < height; y++ {
			cpu.RAM[0x2002] &= 0xbf
			if y < height {
				lineWait.Add(width / 8)
				for x := 0; x < width/8; x++ {
					go func(x int) {
						var spriteWait sync.WaitGroup
						// sprite 描画
						spriteWait.Add(8)
						for i := 0; i < 8; i++ {
							go func(i int) {
								sprite, ok := pixel2sprite[(uint16(y)<<8)|uint16(x*8+i)]
								if ok {
									spriteNum, attr := sprite[0], sprite[1]
									if spriteNum == 0 {
										cpu.RAM[0x2002] |= 0x40
									}
									if attr&0x20 == 0 {
										rect := cpu.PPU.outputSpriteRect(spriteNum, attr)
										SPRSprite := pixel.NewSprite(cpu.PPU.SPRBuf, rect)
										matrix := pixel.IM.Moved(pixel.V(float64(x*8+i), float64(height-y)))
										lineMutex.Lock()
										SPRSprite.Draw(SPRBatch, matrix)
										lineMutex.Unlock()
									}
								}
								spriteWait.Done()
							}(i)
						}
						spriteWait.Wait()

						// BG描画
						scrollPixelX, scrollPixelY := cpu.PPU.scroll[0], cpu.PPU.scroll[1]
						rect := cpu.PPU.outputBGRect(uint(x), uint(y), uint(scrollPixelX), uint(scrollPixelY))
						BGSprite := pixel.NewSprite(cpu.PPU.BGBuf, rect)
						matrix := pixel.IM.Moved(pixel.V(float64(uint8(x*8)-(scrollPixelX%8)), float64(height-y)))

						lineMutex.Lock()
						BGSprite.Draw(BGBatch, matrix)
						lineMutex.Unlock()

						lineWait.Done()
					}(x)
				}
				lineWait.Wait()
			}
		}

		cpu.mutex.Lock()
		cpu.setVBlank()
		cpu.mutex.Unlock()

		BGBatch.Draw(win)
		SPRBatch.Draw(win)

		win.Update()

		if !cpu.PPU.BGPalleteOK {
			cpu.PPU.CacheBG()
			BGBatch = pixel.NewBatch(&pixel.TrianglesData{}, cpu.PPU.BGBuf)
		}
		if !cpu.PPU.SPRPalleteOK {
			cpu.PPU.CacheSPR()
			SPRBatch = pixel.NewBatch(&pixel.TrianglesData{}, cpu.PPU.SPRBuf)
		}

	}
}
