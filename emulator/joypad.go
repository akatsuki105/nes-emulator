package emulator

import (
	"sync"
	"time"

	"github.com/faiface/pixel/pixelgl"
)

// Joypad Joypadの状態を表す
type Joypad struct {
	ctr uint
	cmd [8]byte
}

var keyList = [8]pixelgl.Button{
	pixelgl.KeyX,
	pixelgl.KeyZ,
	pixelgl.KeyRightShift,
	pixelgl.KeyEnter,
	pixelgl.KeyUp,
	pixelgl.KeyDown,
	pixelgl.KeyLeft,
	pixelgl.KeyRight,
}

// handleJoypad キー入力とジョイパッド入力の橋渡しを行う 今回は1Pのみ
func (cpu *CPU) handleJoypad(win *pixelgl.Window) {
	var wait sync.WaitGroup
	for range time.Tick(time.Millisecond) {
		wait.Add(8)
		for i, key := range keyList {
			go func(i int, key pixelgl.Button) {
				if win.Pressed(key) {
					cpu.joypad1.cmd[i] = 1
				}
				wait.Done()
			}(i, key)
		}
		wait.Wait()
	}
}
