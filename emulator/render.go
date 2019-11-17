package emulator

import (
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	width    = 256
	height   = 240
	overload = 12
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

	cpu.win = win
	go cpu.handleJoypad()

	var (
		frames   = 0
		second   = time.Tick(time.Second)
		expanded = uint(1) // 画面拡大率
	)

	for !win.Closed() {
		backgroundRender := false
		spriteRender := false
		if cpu.RAM[0x2001]&0x08 != 0 {
			backgroundRender = true
		}
		if cpu.RAM[0x2001]&0x10 != 0 {
			spriteRender = true
		}

		// SPR探索
		var spriteList [][4]byte
		for i := 0; i < 64; i++ {
			pixelX, pixelY := cpu.PPU.sRAM[i*4+3], (cpu.PPU.sRAM[i*4])+1
			spriteNum := cpu.PPU.sRAM[i*4+1]
			attr := cpu.PPU.sRAM[i*4+2]
			// ラスタースクロールのフラグを立てる
			if i == 0 {
				cpu.PPU.raster = uint16(pixelY)
			}
			spriteList = append(spriteList, [4]byte{pixelX, pixelY, spriteNum, attr})
		}

		// BG・SPR描画
		for y := 0; y < height; y++ {

			// ラスタースクロール
			if cpu.PPU.raster != 0 && (uint16(y) == cpu.PPU.raster) {
				cpu.spriteZeroHit(cpu.PPU.raster)
				cpu.PPU.raster = 0
			}

			for j := 0; j < int(math.Ceil(341/overload)); j++ {
				cpu.exec()
			}

			if y%8 == 0 {
				lineWait.Add(width / 8)

				for x := 0; x < width/8; x++ {
					go func(x, y int) {
						// BG描画
						scrollPixelX, scrollPixelY := cpu.PPU.scroll[0], cpu.PPU.scroll[1]
						mainScreen := cpu.RAM[0x2000] & 0x03

						if backgroundRender {
							cpu.setBGTile(uint(x), uint(y/8), uint(scrollPixelX), uint(scrollPixelY), mainScreen)
						}

						lineWait.Done()
					}(x, y)
				}
				lineWait.Wait()
			}
		}

		// sprite 描画
		for _, sprite := range spriteList {
			pixelX, pixelY, spriteNum, attr := sprite[0], sprite[1], sprite[2], sprite[3]
			if attr&0x20 == 0 {
				if spriteRender {
					cpu.setSPRTile(pixelX, pixelY, spriteNum, attr)
				}
			}
		}

		cpu.mutex.Lock()
		cpu.setVBlank()
		cpu.mutex.Unlock()

		var wait sync.WaitGroup
		wait.Add(1)
		go func() {
			for i := 0; i < 22; i++ {
				for j := 0; j < int(math.Ceil(341/overload)); j++ {
					cpu.exec()
				}
			}
			wait.Done()
		}()

		pic := pixel.PictureDataFromImage(cpu.PPU.displayImage)
		matrix := pixel.IM.Moved(win.Bounds().Center())
		matrix = matrix.ScaledXY(win.Bounds().Center(), pixel.V(float64(expanded), float64(expanded)))
		sprite := pixel.NewSprite(pic, pic.Bounds())
		sprite.Draw(win, matrix)

		win.Update()

		frames++
		select {
		case <-second:
			fmt.Printf("%s | FPS: %d\n", cfg.Title, frames)
			frames = 0
		default:
		}

		// coredump
		if win.Pressed(pixelgl.KeyD) && win.Pressed(pixelgl.KeyS) {
			cpu.dump()
		}
		if win.Pressed(pixelgl.KeyL) {
			cpu.load()
		}
		// resize window
		if win.Pressed(pixelgl.KeyE) {
			expanded *= 2
			time.Sleep(time.Millisecond * 400)
			win.SetBounds(pixel.R(0, 0, float64(width*expanded), float64(height*expanded)))
		}
		if win.Pressed(pixelgl.KeyR) {
			if expanded >= 2 {
				expanded /= 2
				time.Sleep(time.Millisecond * 400)
				win.SetBounds(pixel.R(0, 0, float64(width*expanded), float64(height*expanded)))
			}
		}

		wait.Wait()
		cpu.RAM[0x2002] &= 0xbf // clear Raster
		cpu.clearVBlank()
	}
}
