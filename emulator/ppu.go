package emulator

import (
	"image"
	"image/color"
)

// PPU Picture Processing Unit
type PPU struct {
	RAM    [0x4000]byte
	mirror bool     // 0: 水平ミラー, 1:垂直ミラー
	ptr    uint16   // PPURAMのポインタ 0x2006に書き込まれたとき更新される
	scroll [2]uint8 // (水平スクロールpixel, 垂直スクロールpixel)
}

// renderBlock 画面の(x,y)ブロックのRGBAの出力を行う
// CHR => 0x0000 BG => 0x1000なら0x0000-0x00ffはspr、0x100-0x1ffはbg　逆なら逆
func (ppu *PPU) outputBlock(x, y uint) (img *image.RGBA) {
	spriteNum := uint(ppu.RAM[0x2000+x+y*0x20])
	// if spriteNum > 0 {
	// 	fmt.Printf("%x: (%d, %d) sprite: %d\n", 0x2000+x+y*0x20, x, y, spriteNum)
	// }

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

	img = ppu.outputImage(spriteBytes, pallete)
	return img
}

func (ppu *PPU) outputImage(bytes [16]byte, pallete byte) (img *image.RGBA) {
	img = image.NewRGBA(image.Rect(0, 0, 8, 8))

	var x, y uint
	for y = 0; y < 8; y++ {
		for x = 0; x < 8; x++ {
			color0 := (bytes[y] & (0x01 << (7 - x))) >> (7 - x)
			color1 := ((bytes[y+8] & (0x01 << (7 - x))) >> (7 - x)) << 1

			p := uint(pallete*4) + uint(color0+color1) // パレット番号 + パレット内番号
			R, G, B := colors[ppu.RAM[0x3f00+p]][0], colors[ppu.RAM[0x3f00+p]][1], colors[ppu.RAM[0x3f00+p]][2]
			img.Set((int)(x), (int)(y), color.RGBA{R, G, B, 0})
		}
	}
	return img
}
