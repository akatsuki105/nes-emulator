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
	RAM    [0x4000]byte
	sRAM   [0x100]byte // Sprite RAM
	mirror bool        // 0: 水平ミラー, 1:垂直ミラー
	ptr    uint16      // PPURAMのポインタ 0x2006に書き込まれたとき更新される
	scroll [2]uint8    // (水平スクロールpixel, 垂直スクロールpixel)
}

// renderBlock 画面の(x,y)ブロックのRGBAの出力を行う
// CHR => 0x0000 BG => 0x1000なら0x0000-0x00ffはspr、0x100-0x1ffはbg　逆なら逆
func (ppu *PPU) outputBGBlockPicture(x, y uint) (pic *pixel.PictureData) {
	spriteNum := uint(ppu.RAM[0x2000+x+y*0x20])

	attr := ppu.RAM[0x23c0+(x/4)+(y/4)*0x08]
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

	var spriteBytes [16]byte
	for i := 0; i < 16; i++ {
		spriteBytes[i] = ppu.RAM[spriteNum*16+uint(i)]
	}

	img := image.NewRGBA(image.Rect(0, 0, 8, 8))

	var w, h uint
	for h = 0; h < 8; h++ {
		for w = 0; w < 8; w++ {
			color0 := (spriteBytes[h] & (0x01 << (7 - w))) >> (7 - w)
			color1 := ((spriteBytes[h+8] & (0x01 << (7 - w))) >> (7 - w)) << 1

			p := uint(pallete*4) + uint(color0+color1) // パレット番号 + パレット内番号
			R, G, B := colors[ppu.RAM[0x3f00+p]][0], colors[ppu.RAM[0x3f00+p]][1], colors[ppu.RAM[0x3f00+p]][2]
			img.Set((int)(w), (int)(h), color.RGBA{R, G, B, 0})
		}
	}

	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, img, &jpeg.Options{Quality: 100}); err != nil {
		fmt.Println("error:jpeg\n", err)
		return
	}

	tmp, _, _ := image.Decode(buf)
	pic = pixel.PictureDataFromImage(tmp)
	return pic
}

func (ppu *PPU) outputSpritePicture(spriteNum, attr byte) (pic *pixel.PictureData) {
	var spriteBytes [16]byte
	for i := 0; i < 16; i++ {
		spriteBytes[i] = ppu.RAM[0x1000+uint(spriteNum)*16+uint(i)]
	}

	img := image.NewRGBA(image.Rect(0, 0, 8, 8))

	pallete := attr & 0x03 // パレット番号
	lrTurn := attr & 0x40  // 左右反転
	udTurn := attr & 0x80  // 上下反転

	var w, h uint
	if lrTurn > 0 && udTurn > 0 {
		for h = 0; h < 8; h++ {
			for w = 0; w < 8; w++ {
				color0 := (spriteBytes[8-h] & (0x01 << w)) >> w
				color1 := (spriteBytes[16-h] & (0x01 << w) >> w) << 1

				p := uint(pallete*4) + uint(color0+color1) // パレット番号 + パレット内番号
				R, G, B := colors[ppu.RAM[0x3f10+p]][0], colors[ppu.RAM[0x3f10+p]][1], colors[ppu.RAM[0x3f10+p]][2]
				img.Set((int)(w), (int)(h), color.RGBA{R, G, B, 0})
			}
		}
	} else if lrTurn > 0 && udTurn == 0 {
		for h = 0; h < 8; h++ {
			for w = 0; w < 8; w++ {
				color0 := (spriteBytes[h] & (0x01 << w)) >> w
				color1 := (spriteBytes[h+8] & (0x01 << w) >> w) << 1

				p := uint(pallete*4) + uint(color0+color1) // パレット番号 + パレット内番号
				R, G, B := colors[ppu.RAM[0x3f10+p]][0], colors[ppu.RAM[0x3f10+p]][1], colors[ppu.RAM[0x3f10+p]][2]
				img.Set((int)(w), (int)(h), color.RGBA{R, G, B, 0})
			}
		}
	} else if lrTurn == 0 && udTurn > 0 {
		for h = 0; h < 8; h++ {
			for w = 0; w < 8; w++ {
				color0 := (spriteBytes[8-h] & (0x01 << (7 - w))) >> (7 - w)
				color1 := ((spriteBytes[16-h] & (0x01 << (7 - w))) >> (7 - w)) << 1

				p := uint(pallete*4) + uint(color0+color1) // パレット番号 + パレット内番号
				R, G, B := colors[ppu.RAM[0x3f10+p]][0], colors[ppu.RAM[0x3f10+p]][1], colors[ppu.RAM[0x3f10+p]][2]
				img.Set((int)(w), (int)(h), color.RGBA{R, G, B, 0})
			}
		}
	} else {
		for h = 0; h < 8; h++ {
			for w = 0; w < 8; w++ {
				color0 := (spriteBytes[h] & (0x01 << (7 - w))) >> (7 - w)
				color1 := ((spriteBytes[h+8] & (0x01 << (7 - w))) >> (7 - w)) << 1

				p := uint(pallete*4) + uint(color0+color1) // パレット番号 + パレット内番号
				R, G, B := colors[ppu.RAM[0x3f10+p]][0], colors[ppu.RAM[0x3f10+p]][1], colors[ppu.RAM[0x3f10+p]][2]
				img.Set((int)(w), (int)(h), color.RGBA{R, G, B, 0})
			}
		}
	}

	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, img, &jpeg.Options{Quality: 100}); err != nil {
		fmt.Println("error:jpeg\n", err)
		return
	}

	tmp, _, _ := image.Decode(buf)
	pic = pixel.PictureDataFromImage(tmp)
	return pic
}
