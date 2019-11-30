package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"nes-emulator/emulator"

	"github.com/akatsuki-py/tfd"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	path, err := tfd.CreateSelectDialog([]string{"nes"}, false)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bytes := readNES(path)

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
