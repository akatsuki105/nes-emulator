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

	cpu.CacheBG()
	cpu.PPU.BGBuf = cpu.PPU.newBGBuf
	cpu.PPU.BGBufModified = false
	BGBatch := pixel.NewBatch(&pixel.TrianglesData{}, cpu.PPU.BGBuf)

	cpu.CacheSPR()
	cpu.PPU.SPRBuf = cpu.PPU.newSPRBuf
	cpu.PPU.SPRBufModified = false
	SPRBatch := pixel.NewBatch(&pixel.TrianglesData{}, cpu.PPU.SPRBuf)

	go func() {
		for range time.Tick(time.Millisecond * 80) {
			if cpu.PPU.BGPalleteModified {
				cpu.CacheBG()
			}
			if cpu.PPU.SPRPalleteModified {
				cpu.CacheSPR()
			}
		}
	}()

	go cpu.handleJoypad(win)

	var (
		frames = 0
		second = time.Tick(time.Second)
	)

	for !win.Closed() {
		// SPR探索
		var pixel2sprite map[uint16]([3]byte)
		pixel2sprite = map[uint16]([3]byte){}
		for i := 0; i < 64; i++ {
			pixelX, pixelY := cpu.PPU.sRAM[i*4+3], (cpu.PPU.sRAM[i*4])+1
			spriteNum := cpu.PPU.sRAM[i*4+1]
			attr := cpu.PPU.sRAM[i*4+2]
			pixel2sprite[(uint16(pixelY)<<8)|uint16(pixelX)] = [3]byte{spriteNum, attr, byte(i)}
		}

		// BG・SPR描画
		BGBatch.Clear()
		SPRBatch.Clear()
		for y := 0; y < height/8; y++ {

			// ラスタースクロール
			if cpu.PPU.raster != 0 {
				cpu.spriteZeroHit(cpu.PPU.raster)
				cpu.PPU.raster = 0
			}

			go func() {
				for i := 0; i < 8; i++ {
					for j := 0; j < int(math.Ceil(341/overload)); j++ {
						cpu.exec()
					}
				}
			}()

			lineWait.Add(width / 8)

			for x := 0; x < width/8; x++ {
				go func(x, y int) {
					// sprite 描画
					var spriteWait sync.WaitGroup
					spriteWait.Add(64)
					for i := 0; i < 64; i++ {
						go func(i int) {
							indexX, indexY := i%8, i/8
							sprite, ok := pixel2sprite[(uint16(y*8+indexY)<<8)|uint16(x*8+indexX)]
							if ok {
								spriteNum, attr, index := sprite[0], sprite[1], sprite[2]

								// ラスタースクロールのフラグを立てる
								if index == 0 {
									if x == 0 {
										cpu.spriteZeroHit(uint16(y*8 + indexY))
										cpu.PPU.raster = 0
									} else {
										cpu.PPU.raster = uint16(y*8 + indexY)
									}
								}

								if attr&0x20 == 0 {
									rect := cpu.PPU.outputSpriteRect(uint(spriteNum), attr)
									SPRSprite := pixel.NewSprite(cpu.PPU.SPRBuf, rect)
									matrix := pixel.IM.Moved(pixel.V(float64(x*8+indexX+4), float64(height-y*8-indexY-4)))
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
					mainScreen := cpu.RAM[0x2000] & 0x03
					rect := cpu.PPU.outputBGRect(uint(x), uint(y), uint(scrollPixelX), uint(scrollPixelY), mainScreen)
					BGSprite := pixel.NewSprite(cpu.PPU.BGBuf, rect)
					matrix := pixel.IM.Moved(pixel.V(float64(uint8(x*8)-(scrollPixelX%8)+4), float64(uint8(height-y*8)+(scrollPixelY%8)-4)))

					lineMutex.Lock()
					BGSprite.Draw(BGBatch, matrix)
					lineMutex.Unlock()

					lineWait.Done()
				}(x, y)
			}
			lineWait.Wait()
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

		BGBatch.Draw(win)
		SPRBatch.Draw(win)

		if cpu.PPU.BGBufModified {
			cpu.PPU.BGBuf = cpu.PPU.newBGBuf
			cpu.PPU.BGBufModified = false
			BGBatch = pixel.NewBatch(&pixel.TrianglesData{}, cpu.PPU.BGBuf)
		}
		if cpu.PPU.SPRBufModified {
			cpu.PPU.SPRBuf = cpu.PPU.newSPRBuf
			cpu.PPU.SPRBufModified = false
			SPRBatch = pixel.NewBatch(&pixel.TrianglesData{}, cpu.PPU.SPRBuf)
		}

		win.Update()

		frames++
		select {
		case <-second:
			fmt.Printf("%s | FPS: %d\n", cfg.Title, frames)
			frames = 0
		default:
		}

		// coredump
		if win.Pressed(pixelgl.KeyQ) {
			cpu.dump()
		}
		if win.Pressed(pixelgl.KeyW) {
			cpu.load()
		}

		wait.Wait()
		cpu.RAM[0x2002] &= 0xbf // clear Raster
		cpu.clearVBlank()
	}
}
