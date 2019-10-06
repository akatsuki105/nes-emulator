package emulator

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"

	"github.com/faiface/pixel"
)

// PPU Picture Processing Unit
type PPU struct {
	RAM          [0x4000]byte
	sRAM         [0x100]byte // Sprite RAM
	mirror       bool        // 0: 水平ミラー, 1:垂直ミラー
	ptr          uint16      // PPURAMのポインタ 0x2006に書き込まれたとき更新される
	scroll       [2]uint8    // (水平スクロールpixel, 垂直スクロールpixel)
	scrollFlag   bool        // trueなら2回目として書き込みする
	BGBuf        *pixel.PictureData
	BGPalleteOK  bool
	SPRBuf       *pixel.PictureData
	SPRPalleteOK bool
}

// CacheBG BGのデータをキャッシュする
func (ppu *PPU) CacheBG() {
	img := image.NewRGBA(image.Rect(0, 0, 2048, 32))

	for sprite := 0; sprite < 256; sprite++ {
		var spriteBytes [16]byte
		for i := 0; i < 16; i++ {
			spriteBytes[i] = ppu.RAM[uint(sprite)*16+uint(i)]
		}

		for pallete := 0; pallete < 4; pallete++ {
			var x, y uint
			for y = 0; y < 8; y++ {
				for x = 0; x < 8; x++ {
					color0 := (spriteBytes[y] & (0x01 << (7 - x))) >> (7 - x)
					color1 := ((spriteBytes[y+8] & (0x01 << (7 - x))) >> (7 - x)) << 1

					p := uint(pallete*4) + uint(color0+color1) // パレット番号 + パレット内番号
					if p%4 == 0 {
						p = 0
					}

					R, G, B := colors[ppu.RAM[0x3f00+p]][0], colors[ppu.RAM[0x3f00+p]][1], colors[ppu.RAM[0x3f00+p]][2]
					img.Set((int)(sprite*8+int(x)), (int)(pallete*8+int(y)), color.RGBA{R, G, B, 0})
				}
			}
		}
	}

	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, img, &jpeg.Options{Quality: 100}); err != nil {
		panic("error/jpeg")
	}

	tmp, _, _ := image.Decode(buf)
	pic := pixel.PictureDataFromImage(tmp)
	ppu.BGBuf = pic
	ppu.BGPalleteOK = true
}

// renderBlock 画面の(x,y)pixelのRGBAの出力を行う
func (ppu *PPU) outputBGRect(x, y uint) (rect pixel.Rect) {
	blockX, blockY, indexX, indexY := x/8, y/8, x%8, y%8

	scrollPixelX, scrollPixelY := uint(ppu.scroll[0]), uint(ppu.scroll[1])
	var spriteNum uint
	var attr byte
	if scrollPixelX+x > width && scrollPixelY+y > height {
		spriteNum = uint(ppu.RAM[0x2c00+(x-scrollPixelX/8)+(y-scrollPixelY/8)*0x20])
		attr = ppu.RAM[0x2fc0+(x-scrollPixelX/8)/4+(y-scrollPixelY/8)/4*0x08]
	} else if scrollPixelX+x > width && scrollPixelY+y <= height {
		relBlockX, relBlockY := blockX-(width-scrollPixelX)/8, blockY+scrollPixelY/8 // x, y がスクリーン原点を(0, 0)ブロックとしたときどのブロックにいるか
		spriteNum = uint(ppu.RAM[0x2400+relBlockX+relBlockY*0x20])
		attr = ppu.RAM[0x27c0+relBlockX/4+relBlockY/4*0x08]
	} else if scrollPixelX+x <= width && scrollPixelY+y > height {
		spriteNum = uint(ppu.RAM[0x2800+x+(y-scrollPixelY/8)*0x20])
		attr = ppu.RAM[0x2bc0+(x+scrollPixelX/8)/4+(y-scrollPixelY/8)/4*0x08]
	} else {
		spriteNum = uint(ppu.RAM[0x2000+(blockX+scrollPixelX/8)+(blockY+scrollPixelY/8)*0x20])
		attr = ppu.RAM[0x23c0+(blockX+scrollPixelX/8)/4+(blockY+scrollPixelY/8)/4*0x08]
	}

	var pallete byte
	if (x%4 < 2) && (y%4 < 2) {
		pallete = attr & 0x03
	} else if (x%4 > 2) && (y%4 < 2) {
		pallete = (attr & 0x0c) >> 2
	} else if (x%4 < 2) && (y%4 > 2) {
		pallete = (attr & 0x30) >> 4
	} else {
		pallete = (attr & 0xc0) >> 6
	}

	rect = pixel.R(float64(spriteNum*8+indexX), float64(32-uint(pallete)*8-indexY), float64(spriteNum*8+indexX+1), float64(32-uint(pallete)*8-indexY-1))
	return rect
}

// CacheSPR スプライトをキャッシュする
func (ppu *PPU) CacheSPR() {
	img := image.NewRGBA(image.Rect(0, 0, 2048, 128))

	for sprite := 0; sprite < 256; sprite++ {
		var spriteBytes [16]byte
		for i := 0; i < 16; i++ {
			spriteBytes[i] = ppu.RAM[0x1000+uint(sprite)*16+uint(i)]
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
					if p == 0 || p == 4 || p == 8 || p == 12 {
						p -= 0x10
					}
					R, G, B := colors[ppu.RAM[0x3f10+p]][0], colors[ppu.RAM[0x3f10+p]][1], colors[ppu.RAM[0x3f10+p]][2]
					img.Set((int)(sprite*8+int(w)), (int)(pallete*8+int(h)), color.RGBA{R, G, B, 0})
				}
			}

			// 上下反転
			for h = 0; h < 8; h++ {
				for w = 0; w < 8; w++ {
					color0 := (spriteBytes[7-h] & (0x01 << (7 - w))) >> (7 - w)
					color1 := ((spriteBytes[15-h] & (0x01 << (7 - w))) >> (7 - w)) << 1

					p := int(pallete*4) + int(color0+color1) // パレット番号 + パレット内番号
					// パレットミラーリング
					if p == 0 || p == 4 || p == 8 || p == 12 {
						p -= 0x10
					}
					R, G, B := colors[ppu.RAM[0x3f10+p]][0], colors[ppu.RAM[0x3f10+p]][1], colors[ppu.RAM[0x3f10+p]][2]
					img.Set((int)(sprite*8+int(w)), (int)(pallete*8+int(h)+32), color.RGBA{R, G, B, 0})
				}
			}

			// 左右反転
			for h = 0; h < 8; h++ {
				for w = 0; w < 8; w++ {
					color0 := (spriteBytes[h] & (0x01 << w)) >> w
					color1 := (spriteBytes[h+8] & (0x01 << w) >> w) << 1

					p := int(pallete*4) + int(color0+color1) // パレット番号 + パレット内番号
					// パレットミラーリング
					if p == 0 || p == 4 || p == 8 || p == 12 {
						p -= 0x10
					}
					R, G, B := colors[ppu.RAM[0x3f10+p]][0], colors[ppu.RAM[0x3f10+p]][1], colors[ppu.RAM[0x3f10+p]][2]
					img.Set((int)(sprite*8+int(w)), (int)(pallete*8+int(h)+64), color.RGBA{R, G, B, 0})
				}
			}

			// 上下左右反転
			for h = 0; h < 8; h++ {
				for w = 0; w < 8; w++ {
					color0 := (spriteBytes[7-h] & (0x01 << w)) >> w
					color1 := (spriteBytes[15-h] & (0x01 << w) >> w) << 1

					p := int(pallete*4) + int(color0+color1) // パレット番号 + パレット内番号
					// パレットミラーリング
					if p == 0 || p == 4 || p == 8 || p == 12 {
						p -= 0x10
					}
					R, G, B := colors[ppu.RAM[0x3f10+p]][0], colors[ppu.RAM[0x3f10+p]][1], colors[ppu.RAM[0x3f10+p]][2]
					img.Set((int)(sprite*8+int(w)), (int)(pallete*8+int(h)+96), color.RGBA{R, G, B, 0})
				}
			}
		}
	}

	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, img, &jpeg.Options{Quality: 100}); err != nil {
		fmt.Println("error:jpeg\n", err)
		return
	}

	tmp, _, _ := image.Decode(buf)
	pic := pixel.PictureDataFromImage(tmp)
	ppu.SPRBuf = pic
	ppu.SPRPalleteOK = true
}

func (ppu *PPU) outputSpriteRect(spriteNum, attr byte) (rect pixel.Rect) {
	pallete := attr & 0x03 // パレット番号
	lrTurn := attr & 0x40  // 左右反転
	udTurn := attr & 0x80  // 上下反転

	if lrTurn == 0 && udTurn == 0 {
		rect = pixel.R(float64(spriteNum*8), float64(128-pallete*8), float64((spriteNum+1)*8), float64(128-(pallete+1)*8))
	} else if lrTurn == 0 && udTurn > 0 {
		rect = pixel.R(float64(spriteNum*8), float64(96-pallete*8), float64((spriteNum+1)*8), float64(96-(pallete+1)*8))
	} else if lrTurn > 0 && udTurn == 0 {
		rect = pixel.R(float64(spriteNum*8), float64(64-pallete*8), float64((spriteNum+1)*8), float64(64-(pallete+1)*8))
	} else {
		rect = pixel.R(float64(spriteNum*8), float64(32-pallete*8), float64((spriteNum+1)*8), float64(32-(pallete+1)*8))
	}
	return rect
}
