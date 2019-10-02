package emulator

import (
	"time"

	"github.com/faiface/pixel/pixelgl"
)

// Joypad Joypadの状態を表す
type Joypad struct {
	ctr uint
	cmd [8]byte
}

// handleJoypad キー入力とジョイパッド入力の橋渡しを行う 今回は1Pのみ
func (cpu *CPU) handleJoypad(win *pixelgl.Window) {
	for range time.Tick(5 * time.Nanosecond) {
		// ジョイパッド入力処理
		if win.Pressed(pixelgl.KeyX) {
			cpu.joypad1.cmd[0] = 1 // A
		}
		if win.Pressed(pixelgl.KeyZ) {
			cpu.joypad1.cmd[1] = 1 // B
		}
		if win.Pressed(pixelgl.KeyEnter) {
			cpu.joypad1.cmd[2] = 1 // START
		}
		if win.Pressed(pixelgl.KeyRightShift) {
			cpu.joypad1.cmd[3] = 1 // SELECT
		}
		if win.Pressed(pixelgl.KeyUp) {
			cpu.joypad1.cmd[4] = 1 // 上
		}
		if win.Pressed(pixelgl.KeyDown) {
			cpu.joypad1.cmd[5] = 1 // 下
		}
		if win.Pressed(pixelgl.KeyLeft) {
			cpu.joypad1.cmd[6] = 1 // 左
		}
		if win.Pressed(pixelgl.KeyRight) {
			cpu.joypad1.cmd[7] = 1 // 右
		}
	}
}
