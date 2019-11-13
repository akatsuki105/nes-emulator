package emulator

import (
	"github.com/faiface/pixel/pixelgl"
)

// Joypad Joypadの状態を表す
type Joypad struct {
	ctr uint
	cmd [8]byte // A B Select Start Up Down Left Right
}

const (
	Player0 = 0
)

var Xbox360Controller map[string]int = map[string]int{"A": 1, "B": 0, "Start": 8, "Select": 7, "Horizontal": 0, "Vertical": 1}
var LogitechGamepadF310 map[string]int = map[string]int{"A": 1, "B": 0, "Start": 8, "Select": 7, "Horizontal": 0, "Vertical": 1}

// handleJoypad キー入力とジョイパッド入力の橋渡しを行う 今回は1Pのみ
func (cpu *CPU) handleJoypad() {
	// A
	if btnA(cpu.win) {
		cpu.joypad1.cmd[0] = 1
	}

	// B
	if btnB(cpu.win) {
		cpu.joypad1.cmd[1] = 1
	}

	// select
	if btnSelect(cpu.win) {
		cpu.joypad1.cmd[2] = 1
	}

	// start
	if btnStart(cpu.win) {
		cpu.joypad1.cmd[3] = 1
	}

	// 上
	if keyUp(cpu.win) {
		cpu.joypad1.cmd[4] = 1
	}

	// 下
	if keyDown(cpu.win) {
		cpu.joypad1.cmd[5] = 1
	}

	// 左
	if keyLeft(cpu.win) {
		cpu.joypad1.cmd[6] = 1
	}

	// 右
	if keyRight(cpu.win) {
		cpu.joypad1.cmd[7] = 1
	}
}

func btnA(win *pixelgl.Window) bool {
	switch win.JoystickName(Player0) {
	case "Xbox 360 Controller":
		return win.JoystickPressed(Player0, Xbox360Controller["A"])
	case "Logitech Gamepad F310":
		return win.JoystickPressed(Player0, LogitechGamepadF310["A"])
	default:
		return win.Pressed(pixelgl.KeyX)
	}
}

func btnB(win *pixelgl.Window) bool {
	switch win.JoystickName(Player0) {
	case "Xbox 360 Controller":
		return win.JoystickPressed(Player0, Xbox360Controller["B"])
	case "Logitech Gamepad F310":
		return win.JoystickPressed(Player0, LogitechGamepadF310["B"])
	default:
		return win.Pressed(pixelgl.KeyZ)
	}
}

func btnStart(win *pixelgl.Window) bool {
	switch win.JoystickName(Player0) {
	case "Xbox 360 Controller":
		return win.JoystickPressed(Player0, Xbox360Controller["Start"])
	case "Logitech Gamepad F310":
		return win.JoystickPressed(Player0, LogitechGamepadF310["Start"])
	default:
		return win.Pressed(pixelgl.KeyEnter)
	}
}

func btnSelect(win *pixelgl.Window) bool {
	switch win.JoystickName(Player0) {
	case "Xbox 360 Controller":
		return win.JoystickPressed(Player0, Xbox360Controller["Select"])
	case "Logitech Gamepad F310":
		return win.JoystickPressed(Player0, LogitechGamepadF310["Select"])
	default:
		return win.Pressed(pixelgl.KeyRightShift)
	}
}

func keyUp(win *pixelgl.Window) bool {
	switch win.JoystickName(Player0) {
	case "Xbox 360 Controller":
		return win.JoystickAxis(Player0, Xbox360Controller["Vertical"]) >= 1
	case "Logitech Gamepad F310":
		return win.JoystickAxis(Player0, LogitechGamepadF310["Vertical"]) <= -1
	default:
		return win.Pressed(pixelgl.KeyUp)
	}
}

func keyDown(win *pixelgl.Window) bool {
	switch win.JoystickName(Player0) {
	case "Xbox 360 Controller":
		return win.JoystickAxis(Player0, Xbox360Controller["Vertical"]) <= -1
	case "Logitech Gamepad F310":
		return win.JoystickAxis(Player0, LogitechGamepadF310["Vertical"]) >= 1
	default:
		return win.Pressed(pixelgl.KeyDown)
	}
}

func keyRight(win *pixelgl.Window) bool {
	switch win.JoystickName(Player0) {
	case "Xbox 360 Controller":
		return win.JoystickAxis(Player0, Xbox360Controller["Horizontal"]) >= 1
	case "Logitech Gamepad F310":
		return win.JoystickAxis(Player0, LogitechGamepadF310["Horizontal"]) >= 1
	default:
		return win.Pressed(pixelgl.KeyRight)
	}
}

func keyLeft(win *pixelgl.Window) bool {
	switch win.JoystickName(Player0) {
	case "Xbox 360 Controller":
		return win.JoystickAxis(Player0, Xbox360Controller["Horizontal"]) <= -1
	case "Logitech Gamepad F310":
		return win.JoystickAxis(Player0, LogitechGamepadF310["Horizontal"]) <= -1
	default:
		return win.Pressed(pixelgl.KeyLeft)
	}
}
