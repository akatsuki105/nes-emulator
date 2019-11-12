package emulator

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"sync"

	"github.com/faiface/pixel"
)

// PPU Picture Processing Unit
type PPU struct {
	RAM                [0x4000]byte
	sRAM               [0x100]byte // Sprite RAM
	mirror             bool        // 0: 水平ミラー, 1:垂直ミラー
	ptr                uint16      // PPURAMのポインタ 0x2006に書き込まれたとき更新される
	ppudataBuf         byte        // PPUDATAからreadしたときのbuffer
	scroll             [2]uint8    // (水平スクロールpixel, 垂直スクロールpixel)
	scrollFlag         bool        // trueなら2回目として書き込みする
	raster             uint16
	BGBuf              *pixel.PictureData
	newBGBuf           *pixel.PictureData
	BGPalleteModified  bool
	BGBufModified      bool
	SPRBuf             *pixel.PictureData
	newSPRBuf          *pixel.PictureData
	SPRPalleteModified bool
	SPRBufModified     bool
}

// CacheBG BGのデータをキャッシュする
func (cpu *CPU) CacheBG() {
	img := image.NewRGBA(image.Rect(0, 0, 2048, 32))

	baseAddr := cpu.getBaseAddr("BG")

	var wait sync.WaitGroup
	wait.Add(256)
	for sprite := 0; sprite < 256; sprite++ {
		go func(sprite int) {
			var spriteBytes [16]byte
			for i := 0; i < 16; i++ {
				spriteBytes[i] = cpu.PPU.RAM[baseAddr+uint(sprite)*16+uint(i)]
			}

			for pallete := 0; pallete < 4; pallete++ {
				var x, y uint
				for y = 0; y < 8; y++ {
					for x = 0; x < 8; x++ {
						color0 := (spriteBytes[y] & (0x01 << (7 - x))) >> (7 - x)
						color1 := ((spriteBytes[y+8] & (0x01 << (7 - x))) >> (7 - x)) << 1

						p := uint(pallete*4) + uint(color0+color1) // パレット番号 + パレット内番号
						if p%4 == 0 {
							p = 0x10
						}

						R, G, B := colors[cpu.PPU.RAM[0x3f00+p]][0], colors[cpu.PPU.RAM[0x3f00+p]][1], colors[cpu.PPU.RAM[0x3f00+p]][2]
						img.Set((int)(sprite*8+int(x)), (int)(pallete*8+int(y)), color.RGBA{R, G, B, 0xff})
					}
				}
			}

			wait.Done()
		}(sprite)
	}
	wait.Wait()

	buf := new(bytes.Buffer)
	if err := png.Encode(buf, img); err != nil {
		panic("error/png")
	}

	tmp, _, _ := image.Decode(buf)
	pic := pixel.PictureDataFromImage(tmp)
	cpu.PPU.newBGBuf = pic
	cpu.PPU.BGBufModified = true
	cpu.PPU.BGPalleteModified = false
}

// outputBGRect 画面の(x,y)ブロックのRGBAの出力を行う
func (ppu *PPU) outputBGRect(x, y, scrollPixelX, scrollPixelY uint, mainScreen byte) (rect pixel.Rect) {
	var spriteNum uint
	var attr byte

	switch mainScreen {
	case 1:
		scrollPixelX += width
	case 2:
		scrollPixelY += height
	case 3:
		scrollPixelX += width
		scrollPixelY += height
	}

	var blockX, blockY uint
	if (scrollPixelX+x*8 >= width && scrollPixelX+x*8 < width*2) && (scrollPixelY+y*8 >= height && scrollPixelY+y*8 < height*2) {
		blockX, blockY = (x*8-width+scrollPixelX)/8, (y*8-height+scrollPixelY)/8
		spriteNum = uint(ppu.RAM[0x2c00+blockX+blockY*0x20])
		attr = ppu.RAM[0x2fc0+blockX/4+blockY/4*0x08]
	} else if (scrollPixelX+x*8 >= width && scrollPixelX+x*8 < width*2) && scrollPixelY+y*8 < height {
		blockX, blockY = (x*8-width+scrollPixelX)/8, y+scrollPixelY/8
		spriteNum = uint(ppu.RAM[0x2400+blockX+blockY*0x20])
		attr = ppu.RAM[0x27c0+blockX/4+blockY/4*0x08]
	} else if scrollPixelX+x*8 < width && (scrollPixelY+y*8 >= height && scrollPixelY+y*8 < height*2) {
		blockX, blockY = (x*8+scrollPixelX)/8, (y*8-height+scrollPixelY)/8
		spriteNum = uint(ppu.RAM[0x2800+blockX+blockY*0x20])
		attr = ppu.RAM[0x2bc0+blockX/4+blockY/4*0x08]
	} else {
		blockX, blockY = ((x*8+scrollPixelX)/8)%(width/8*2), ((y*8+scrollPixelY)/8)%(height/8*2)
		spriteNum = uint(ppu.RAM[0x2000+blockX+blockY*0x20])
		attr = ppu.RAM[0x23c0+blockX/4+blockY/4*0x08]
	}

	var pallete byte
	if (blockX%4 < 2) && (blockY%4 < 2) {
		pallete = attr & 0x03
	} else if (blockX%4 >= 2) && (blockY%4 < 2) {
		pallete = (attr & 0x0c) >> 2
	} else if (blockX%4 < 2) && (blockY%4 >= 2) {
		pallete = (attr & 0x30) >> 4
	} else {
		pallete = (attr & 0xc0) >> 6
	}

	rect = pixel.R(float64(spriteNum*8), float64(32-uint(pallete+1)*8), float64((spriteNum+1)*8), float64(32-uint(pallete)*8))
	return rect
}

// CacheSPR スプライトをキャッシュする
func (cpu *CPU) CacheSPR() {
	img := image.NewRGBA(image.Rect(0, 0, 2048, 128))

	baseAddr := cpu.getBaseAddr("SPR")

	var wait sync.WaitGroup
	wait.Add(256)
	for sprite := 0; sprite < 256; sprite++ {
		go func(sprite int) {
			var spriteBytes [16]byte
			for i := 0; i < 16; i++ {
				spriteBytes[i] = cpu.PPU.RAM[baseAddr+uint(sprite)*16+uint(i)]
			}

			for pallete := 0; pallete < 4; pallete++ {
				var w, h uint

				// 反転無し
				for h = 0; h < 8; h++ {
					for w = 0; w < 8; w++ {
						color0 := (spriteBytes[h] & (0x01 << (7 - w))) >> (7 - w)
						color1 := ((spriteBytes[h+8] & (0x01 << (7 - w))) >> (7 - w)) << 1

						p := int(pallete*4) + int(color0+color1) // パレット番号 + パレット内番号
						// パレットミラーリング
						var transparent uint8
						if p == 0 || p == 4 || p == 8 || p == 12 {
							transparent = 0
						} else {
							transparent = 255
						}
						R, G, B := colors[cpu.PPU.RAM[0x3f10+p]][0], colors[cpu.PPU.RAM[0x3f10+p]][1], colors[cpu.PPU.RAM[0x3f10+p]][2]
						img.Set((int)(sprite*8+int(w)), (int)(pallete*8+int(h)), color.RGBA{R, G, B, transparent})
					}
				}

				// 上下反転
				for h = 0; h < 8; h++ {
					for w = 0; w < 8; w++ {
						color0 := (spriteBytes[7-h] & (0x01 << (7 - w))) >> (7 - w)
						color1 := ((spriteBytes[15-h] & (0x01 << (7 - w))) >> (7 - w)) << 1

						p := int(pallete*4) + int(color0+color1) // パレット番号 + パレット内番号
						// パレットミラーリング
						var transparent uint8
						if p == 0 || p == 4 || p == 8 || p == 12 {
							transparent = 0
						} else {
							transparent = 255
						}
						R, G, B := colors[cpu.PPU.RAM[0x3f10+p]][0], colors[cpu.PPU.RAM[0x3f10+p]][1], colors[cpu.PPU.RAM[0x3f10+p]][2]
						img.Set((int)(sprite*8+int(w)), (int)(pallete*8+int(h)+32), color.RGBA{R, G, B, transparent})
					}
				}

				// 左右反転
				for h = 0; h < 8; h++ {
					for w = 0; w < 8; w++ {
						color0 := (spriteBytes[h] & (0x01 << w)) >> w
						color1 := (spriteBytes[h+8] & (0x01 << w) >> w) << 1

						p := int(pallete*4) + int(color0+color1) // パレット番号 + パレット内番号
						// パレットミラーリング
						var transparent uint8
						if p == 0 || p == 4 || p == 8 || p == 12 {
							transparent = 0
						} else {
							transparent = 255
						}
						R, G, B := colors[cpu.PPU.RAM[0x3f10+p]][0], colors[cpu.PPU.RAM[0x3f10+p]][1], colors[cpu.PPU.RAM[0x3f10+p]][2]
						img.Set((int)(sprite*8+int(w)), (int)(pallete*8+int(h)+64), color.RGBA{R, G, B, transparent})
					}
				}

				// 上下左右反転
				for h = 0; h < 8; h++ {
					for w = 0; w < 8; w++ {
						color0 := (spriteBytes[7-h] & (0x01 << w)) >> w
						color1 := (spriteBytes[15-h] & (0x01 << w) >> w) << 1

						p := int(pallete*4) + int(color0+color1) // パレット番号 + パレット内番号

						// パレットミラーリング
						var transparent uint8
						if p == 0 || p == 4 || p == 8 || p == 12 {
							transparent = 0
						} else {
							transparent = 255
						}
						R, G, B := colors[cpu.PPU.RAM[0x3f10+p]][0], colors[cpu.PPU.RAM[0x3f10+p]][1], colors[cpu.PPU.RAM[0x3f10+p]][2]
						img.Set((int)(sprite*8+int(w)), (int)(pallete*8+int(h)+96), color.RGBA{R, G, B, transparent})
					}
				}
			}
			wait.Done()
		}(sprite)
	}
	wait.Wait()

	buf := new(bytes.Buffer)
	if err := png.Encode(buf, img); err != nil {
		fmt.Println("error:png", err)
		return
	}

	tmp, _, _ := image.Decode(buf)
	pic := pixel.PictureDataFromImage(tmp)
	cpu.PPU.newSPRBuf = pic
	cpu.PPU.SPRBufModified = true
	cpu.PPU.SPRPalleteModified = false
}

func (ppu *PPU) outputSpriteRect(spriteNum uint, attr byte) (rect pixel.Rect) {
	pallete := attr & 0x03 // パレット番号
	lrTurn := attr & 0x40  // 左右反転
	udTurn := attr & 0x80  // 上下反転

	if lrTurn == 0 && udTurn == 0 {
		// そのまま
		rect = pixel.R(float64(spriteNum*8), float64(128-pallete*8), float64((spriteNum+1)*8), float64(128-(pallete+1)*8))
	} else if lrTurn == 0 && udTurn != 0 {
		// 上下反転
		rect = pixel.R(float64(spriteNum*8), float64(96-pallete*8), float64((spriteNum+1)*8), float64(96-(pallete+1)*8))
	} else if lrTurn != 0 && udTurn == 0 {
		// 左右反転
		rect = pixel.R(float64(spriteNum*8), float64(64-pallete*8), float64((spriteNum+1)*8), float64(64-(pallete+1)*8))
	} else {
		// 上下左右反転
		rect = pixel.R(float64(spriteNum*8), float64(32-pallete*8), float64((spriteNum+1)*8), float64(32-(pallete+1)*8))
	}
	return rect
}

func (cpu *CPU) getBaseAddr(name string) uint {
	if name == "BG" {
		if cpu.RAM[0x2000]&0x10 > 0 {
			return 0x1000
		}
	} else if name == "SPR" {
		if cpu.RAM[0x2000]&0x08 > 0 {
			return 0x1000
		}
	}
	return 0x0000
}
