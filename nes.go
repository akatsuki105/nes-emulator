package main

import (
	"flag"
	"io/ioutil"
	"path/filepath"

	"nes-emulator/emulator"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	flag.Parse()
	bytes := readNES(flag.Arg(0))

	cpu := &emulator.CPU{}
	cpu.LoadROM(bytes)
	cpu.InitReg()
	go cpu.Debug()

	pixelgl.Run(cpu.Render)
}

func readNES(path string) []byte {
	if path == "" {
		panic("please enter nes file path")
	}
	if filepath.Ext(path) != ".nes" {
		panic("please select .nes file")
	}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return bytes
}
