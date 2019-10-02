package emulator

import (
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

	for !win.Closed() {
		cpu.setVBlank()

		// BG描画
		for y := 0; y < height/8; y++ {
			for x := 0; x < width/8; x++ {
				pic := cpu.PPU.outputBGBlockPicture(uint(x), uint(y))
				sprite := pixel.NewSprite(pic, pic.Bounds())
				matrix := pixel.IM.Moved(pixel.V(float64(x*8+4), float64(height-4-y*8)))
				sprite.Draw(win, matrix)
			}
		}

		// SPR描画
		for i := 0; i < 64; i++ {
			pixelX, pixelY := cpu.PPU.sRAM[i*4+3], (cpu.PPU.sRAM[i*4])
			spriteNum := cpu.PPU.sRAM[i*4+1]
			attr := cpu.PPU.sRAM[i*4+2]
			if attr&0x20 == 0 {
				pic := cpu.PPU.outputSpritePicture(spriteNum, attr)
				sprite := pixel.NewSprite(pic, pic.Bounds())
				matrix := pixel.IM.Moved(pixel.V(float64(pixelX+4), float64(height-4-pixelY)))
				sprite.Draw(win, matrix)
			}
		}

		cpu.clearVBlank()

		win.Update()
	}
}
