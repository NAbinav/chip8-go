package main

import (
	"math/rand"
)

// 00E0 - CLS
func (c *Chip8) OP_00E0() {
	c.Display = [64][32]bool{}
}

// 00EE - RET
func (c *Chip8) OP_00EE() {
	c.SP--
	c.PC = c.Stack[c.SP]
}

// 0nnn - SYS addr -> No need
// 1nnn - JP addr
func (c *Chip8) OP_1NNN() {
	c.PC = c.Opcode & 0x0FFF //take nnn
}

// 2nnn - CALL addr
// calls subroutine at nnn
func (c *Chip8) OP_2NNN() {
	c.Stack[c.SP] = c.PC
	c.SP++
	c.PC = c.Opcode & 0x0FFF //take nnn
}

// 3xkk - SE Vx, byte
func (c *Chip8) OP_3xkk() {
	Vx := (c.Opcode & 0x0F00) >> 8
	kk := c.Opcode & 0x00FF
	if uint16(c.V[Vx]) == kk {
		c.PC += 2
	}
}

// 4xkk - SNE Vx, byte
func (c *Chip8) OP_4xkk() {
	Vx := (c.Opcode & 0x0F00) >> 8
	kk := c.Opcode & 0x00FF
	if uint16(c.V[Vx]) != kk {
		c.PC += 2
	}
}

// 5xy0 - SE Vx, Vy
func (c *Chip8) OP_5xy0() {
	Vx := (c.Opcode & 0x0F00) >> 8
	Vy := (c.Opcode & 0x00F0) >> 4
	if uint16(c.V[Vx]) == uint16(c.V[Vy]) {
		c.PC += 2
	}
}

//	6xkk - LD Vx, byte
//
// set Vx=kk
func (c *Chip8) OP_6xkk() {
	Vx := (c.Opcode & 0x0F00) >> 8
	kk := c.Opcode & 0x00FF
	c.V[Vx] = byte(kk)
}

// 7xkk - ADD Vx, byte
func (c *Chip8) OP_7xkk() {
	Vx := (c.Opcode & 0x0F00) >> 8
	kk := c.Opcode & 0x00FF
	c.V[Vx] += byte(kk)
}

// 8xy0 - LD Vx, Vy
func (c *Chip8) OP_8xy0() {
	Vx := (c.Opcode & 0x0F00) >> 8
	Vy := (c.Opcode & 0x00F0) >> 4
	c.V[Vx] = c.V[Vy]
}

// 8xy1 - OR Vx, Vy
func (c *Chip8) OP_8xy1() {
	Vx := (c.Opcode & 0x0F00) >> 8
	Vy := (c.Opcode & 0x00F0) >> 4
	c.V[Vx] |= c.V[Vy]
}

// 8xy2 - AND Vx, Vy
func (c *Chip8) OP_8xy2() {
	Vx := (c.Opcode & 0x0F00) >> 8
	Vy := (c.Opcode & 0x00F0) >> 4
	c.V[Vx] &= c.V[Vy]
}

// 8xy3 - XOR Vx, Vy
func (c *Chip8) OP_8xy3() {
	Vx := (c.Opcode & 0x0F00) >> 8
	Vy := (c.Opcode & 0x00F0) >> 4
	c.V[Vx] ^= c.V[Vy]
}

// 8xy4 - ADD Vx, Vy
func (c *Chip8) OP_8xy4() {
	Vx := (c.Opcode & 0x0F00) >> 8
	Vy := (c.Opcode & 0x00F0) >> 4
	sum := uint16(c.V[Vx]) + uint16(c.V[Vy])
	if sum > 255 {
		c.V[0xF] = 1
	}
	c.V[Vx] = byte(sum) & 0xFF
}

// 8xy5 - SUB Vx, Vy
func (c *Chip8) OP_8xy5() {
	Vx := (c.Opcode & 0x0F00) >> 8
	Vy := (c.Opcode & 0x00F0) >> 4
	if c.V[Vx] >= c.V[Vy] {
		c.V[0xF] = 1
	} else {
		c.V[0xF] = 0
	}
	c.V[Vx] -= c.V[Vy]
}

// 8xy6 - SHR Vx {, Vy}
func (c *Chip8) OP_8xy6() {
	Vx := (c.Opcode & 0x0F00) >> 8
	c.V[0xF] = c.V[Vx] & 0x1
	c.V[Vx] >>= 1
}

// 8xy7 - SUBN Vx, Vy
func (c *Chip8) OP_8xy7() {
	Vx := (c.Opcode & 0x0F00) >> 8
	Vy := (c.Opcode & 0x00F0) >> 4
	if c.V[Vx] < c.V[Vy] {
		c.V[0xF] = 1
	} else {
		c.V[0xF] = 0
	}
	c.V[Vy] -= c.V[Vx]
}

// 8xyE - SHL Vx {, Vy}
func (c *Chip8) OP_8xyE() {
	Vx := (c.Opcode & 0x0F00) >> 8
	c.V[0xF] = (c.V[Vx] & 0x80) >> 7
	c.V[Vx] <<= 1
}

// 9xy0 - SNE Vx, Vy
func (c *Chip8) OP_9xy0() {
	Vx := (c.Opcode & 0x0F00) >> 8
	Vy := (c.Opcode & 0x00F0) >> 4
	if uint16(c.V[Vx]) != uint16(c.V[Vy]) {
		c.PC += 2
	}

}

// Annn - LD I, addr
func (c *Chip8) OP_ANNN() {
	c.I = c.Opcode & 0xFFF
}

// Bnnn - JP V0, addr
func (c *Chip8) OP_BNNN() {
	address := c.Opcode & 0xFFF
	c.PC = address + uint16(c.V[0])
}

// Cxkk - RND Vx, byte
func random() uint8 {
	return uint8(rand.Intn(256))
}

func (c *Chip8) OP_Cxkk() {
	vx := c.Opcode & 0x0F00 >> 8
	val := uint8(c.Opcode & 0xFF)
	c.V[vx] = val & uint8(random())
}

// Dxyn - DRW Vx, Vy, nibble
func (c *Chip8) OP_Dxyn() {
	vx := (c.Opcode & 0x0F00) >> 8
	vy := (c.Opcode & 0x00F0) >> 4
	h := c.Opcode & 0x000F
	x := c.V[vx] % 64
	y := c.V[vy] % 32
	c.V[0xF] = 0
	for row := uint16(0); row < h; row++ {
		pix := c.Memory[c.I+row]

		for col := uint16(0); col < 8; col++ {
			if (pix & (0x80 >> col)) != 0 {
				x := (x + uint8(col)) % 64
				y := (y + uint8(row)) % 32

				if c.Display[x][y] {
					c.V[0xF] = 1
				}

				c.Display[x][y] = !c.Display[x][y]
			}

		}
	}
}

// Ex9E - SKP Vx
// Skip next instruction if key with the value of Vx is pressed.
func (c *Chip8) OP_Ex9E() {

	vx := (c.Opcode & 0x0F00) >> 8
	if c.Key[c.V[vx]] {

	}

}

//            ExA1 - SKNP Vx
//            Fx07 - LD Vx, DT
//            Fx0A - LD Vx, K
//            Fx15 - LD DT, Vx
//            Fx18 - LD ST, Vx
//            Fx1E - ADD I, Vx
//            Fx29 - LD F, Vx
//            Fx33 - LD B, Vx
//            Fx55 - LD [I], Vx
//            Fx65 - LD Vx, [I]
