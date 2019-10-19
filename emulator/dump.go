package emulator

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func (cpu *CPU) dump() {
	fmt.Println("dump...")
	time.Sleep(time.Millisecond * 200)

	dumpfile, err := os.Create("./dumpfile")
	if err != nil {
		fmt.Println("dump failed.")
	}
	defer dumpfile.Close()

	buf := bytes.NewBuffer(nil)
	_ = gob.NewEncoder(buf).Encode(&cpu)
	data := buf.Bytes()

	_, err = dumpfile.Write(data)
	if err != nil {
		fmt.Println("dump failed.")
	} else {
		fmt.Println("dump success!")
	}
}

func (cpu *CPU) load() {
	fmt.Println("loading dumpfile...")
	time.Sleep(time.Millisecond * 200)

	data, err := ioutil.ReadFile("./dumpfile")
	if err != nil {
		fmt.Println("loading dumpfile failed.")
		return
	}

	buf := bytes.NewBuffer(data)
	_ = gob.NewDecoder(buf).Decode(cpu)
	fmt.Println("loading dumpfile success!")
}
