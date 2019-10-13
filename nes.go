package main

import (
	"flag"
	"io/ioutil"

	"./emulator"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	flag.Parse()
	bytes := readFile(flag.Arg(0))

	cpu := &emulator.CPU{}
	cpu.LoadROM(bytes)
	cpu.InitReg()

	pixelgl.Run(cpu.Render)
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
