package emulator

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	width  = 256
	height = 240
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

	go cpu.VBlank()

	for !win.Closed() {
		// SPR探索
		var pixel2sprite map[uint16]([2]byte)
		pixel2sprite = map[uint16]([2]byte){}
		for i := 0; i < 64; i++ {
			pixelX, pixelY := cpu.PPU.sRAM[i*4+3], (cpu.PPU.sRAM[i*4])
			spriteNum := cpu.PPU.sRAM[i*4+1]
			attr := cpu.PPU.sRAM[i*4+2]
			pixel2sprite[(uint16(pixelY)<<8)|uint16(pixelX)] = [2]byte{spriteNum, attr}
		}

		// BG・SPR描画
		BGBatch.Clear()
		SPRBatch.Clear()
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				sprite, ok := pixel2sprite[(uint16(y)<<8)|uint16(x)]
				if ok {
					spriteNum, attr := sprite[0], sprite[1]
					if attr&0x20 == 0 {
						rect := cpu.PPU.outputSpriteRect(spriteNum, attr)
						SPRSprite := pixel.NewSprite(cpu.PPU.SPRBuf, rect)
						matrix := pixel.IM.Moved(pixel.V(float64(x+4), float64(height-4-y)))
						SPRSprite.Draw(SPRBatch, matrix)
					}
				}

				rect := cpu.PPU.outputBGRect(uint(x), uint(y))
				BGSprite := pixel.NewSprite(cpu.PPU.BGBuf, rect)
				matrix := pixel.IM.Moved(pixel.V(float64(x), float64(height-y)))
				BGSprite.Draw(BGBatch, matrix)
			}
		}
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

// VBlank VBlankを起こす
func (cpu *CPU) VBlank() {
	for range time.Tick(5 * time.Millisecond) {
		cpu.mutex.Lock()
		cpu.setVBlank()
		cpu.mutex.Unlock()
	}
}
