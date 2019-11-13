package emulator

import (
	"sync"

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
func (cpu *CPU) handleJoypad() {
	var wait sync.WaitGroup
	wait.Add(8)
	for i, key := range keyList {
		go func(i int, key pixelgl.Button) {
			if cpu.win.Pressed(key) {
				cpu.joypad1.cmd[i] = 1
			}
			wait.Done()
		}(i, key)
	}
	wait.Wait()
}
