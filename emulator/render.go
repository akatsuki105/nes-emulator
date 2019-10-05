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
		// BG描画
		BGBatch.Clear()
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				rect := cpu.PPU.outputBGRect(uint(x), uint(y))
				sprite := pixel.NewSprite(cpu.PPU.BGBuf, rect)
				matrix := pixel.IM.Moved(pixel.V(float64(x), float64(height-y)))
				sprite.Draw(BGBatch, matrix)
			}
		}
		BGBatch.Draw(win)

		// SPR描画
		SPRBatch.Clear()
		for i := 0; i < 64; i++ {
			pixelX, pixelY := cpu.PPU.sRAM[i*4+3], (cpu.PPU.sRAM[i*4])
			spriteNum := cpu.PPU.sRAM[i*4+1]
			attr := cpu.PPU.sRAM[i*4+2]
			if attr&0x20 == 0 {
				rect := cpu.PPU.outputSpriteRect(spriteNum, attr)
				sprite := pixel.NewSprite(cpu.PPU.SPRBuf, rect)
				matrix := pixel.IM.Moved(pixel.V(float64(pixelX+4), float64(height-4-pixelY)))
				sprite.Draw(SPRBatch, matrix)
			}
		}
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
