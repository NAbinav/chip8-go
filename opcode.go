package main

import (
	"fmt"
	"time"
)

func (c *Chip8) cycle() {
	// Fetch
	c.Opcode = uint16(c.Memory[c.PC])<<8 |
		uint16(c.Memory[c.PC+1])

	// Increment PC FIRST
	c.PC += 2

	// Decode & Execute
	c.table[(c.Opcode&0xF000)>>12]()
	c.PrintDisplay()
	time.Sleep(time.Second / 24)
}
func (c *Chip8) PrintDisplay() {
	fmt.Print("\033[H")
	for y := range 32 {
		for x := range 64 {
			if c.Display[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		if y < 31 {
			fmt.Print("\r\n")
		}
		fmt.Print("\033[K")
	}
}
