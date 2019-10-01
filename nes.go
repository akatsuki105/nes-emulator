package main

import (
	"flag"
	"io/ioutil"

	"./emulator"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	flag.Parse()
	bytes := readFile(flag.Arg(0))

	cpu := &emulator.CPU{}
	cpu.InitReg()
	cpu.LoadROM(bytes)

}

func render() {
	cfg := pixelgl.WindowConfig{
		Title:  "nes-emulator",
		Bounds: pixel.R(0, 0, 256, 240),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	for !win.Closed() {
		win.Update()
	}
}

func readFile(path string) []byte {
	if path == "" {
		panic("please enter nes file path")
	}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return bytes
}
