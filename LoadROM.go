package main

import (
	"log"
	"os"
)

func (c *Chip8) LoadROM(filename string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	for i, val := range content {
		c.Memory[0x200+i] = val
	}

}

// TODO: Embed the ROM file so that it is easy to load in browser when covnerted to wasm
