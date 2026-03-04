package main

import (
	"fmt"
	"log"
	"os"
)

func (c *Chip8) LoadROM(filename string) {
	content, err := os.ReadFile(filename)
	romSize := len(content)
	if err != nil {
		log.Fatal(err)
	}
	for i, val := range content {
		c.Memory[0x200+i] = val
	}

	for i := 0x200; i < 0x200+romSize; i += 2 {
		fmt.Printf("%x ", uint16(c.Memory[i]<<8)|uint16(c.Memory[i+1]))
	}
	for true {
		c.cycle()
	}

}
