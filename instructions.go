package main

import (
	"math/rand"
)

// 00E0 - CLS
func (c *Chip8) OP_00E0() {
	c.Display = [32][64]bool{}
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
	if c.V[Vx] == c.V[Vy] {
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
	} else {
		c.V[0xF] = 0
	}
	c.V[Vx] = byte(sum)
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
	if c.V[Vy] >= c.V[Vx] {
		c.V[0xF] = 1
	} else {
		c.V[0xF] = 0
	}
	c.V[Vx] = c.V[Vy] - c.V[Vx]
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
	if c.V[Vx] != c.V[Vy] {
		c.PC += 2
	}
}

// Annn - LD I, addr
func (c *Chip8) OP_ANNN() {
	c.I = c.Opcode & 0xFFF
}

// Bnnn - JP V0, addr
func (c *Chip8) OP_BNNN() {
	c.PC = (c.Opcode & 0x0FFF) + uint16(c.V[0])
}

// Cxkk - RND Vx, byte
func (c *Chip8) OP_Cxkk() {
	Vx := (c.Opcode & 0x0F00) >> 8
	kk := byte(c.Opcode & 0x00FF)
	c.V[Vx] = kk & uint8(rand.Intn(256))
}

// Dxyn - DRW Vx, Vy, nibble
func (c *Chip8) OP_Dxyn() {
	x := c.V[(c.Opcode&0x0F00)>>8] % 64
	y := c.V[(c.Opcode&0x00F0)>>4] % 32
	height := c.Opcode & 0x000F
	c.V[0xF] = 0
	for row := uint16(0); row < height; row++ {
		spriteByte := c.Memory[c.I+row]
		for col := uint16(0); col < 8; col++ {
			if (spriteByte & (0x80 >> col)) != 0 {
				px := (uint16(x) + col) % 64
				py := (uint16(y) + row) % 32
				if c.Display[py][px] {
					c.V[0xF] = 1
				}
				c.Display[py][px] = !c.Display[py][px]
			}
		}
	}
}

// Ex9E - SKP Vx
func (c *Chip8) OP_Ex9E() {
	Vx := (c.Opcode & 0x0F00) >> 8
	if c.Key[c.V[Vx]] {
		c.PC += 2
	}
}

// ExA1 - SKNP Vx
func (c *Chip8) OP_ExA1() {

	Vx := (c.Opcode & 0x0F00) >> 8
	if !c.Key[c.V[Vx]] {
		c.PC += 2
	}
}

// Fx07 - LD Vx, DT

func (c *Chip8) OP_Fx07() {
	Vx := (c.Opcode & 0x0F00) >> 8
	c.V[Vx] = c.DelayTimer
}

// Fx0A - LD Vx, K — wait for key press
func (c *Chip8) OP_Fx0A() {
	Vx := (c.Opcode & 0x0F00) >> 8
	for i := range 16 {
		if c.Key[i] {
			c.V[Vx] = byte(i)
			return
		}
	}
	c.PC -= 2 // no key pressed, retry
}

// Fx15 - LD DT, Vx

func (c *Chip8) OP_Fx15() {
	Vx := (c.Opcode & 0x0F00) >> 8
	c.DelayTimer = c.V[Vx]
}

// Fx18 - LD ST, Vx
func (c *Chip8) OP_Fx18() {
	Vx := (c.Opcode & 0x0F00) >> 8
	c.SoundTimer = c.V[Vx]
}

// Fx1E - ADD I, Vx
func (c *Chip8) OP_Fx1E() {
	Vx := (c.Opcode & 0x0F00) >> 8
	c.I += uint16(c.V[Vx])
}

// Fx29 - LD F, Vx
func (c *Chip8) OP_Fx29() {
	Vx := (c.Opcode & 0x0F00) >> 8
	c.I = 0x50 + 5*uint16(c.V[Vx])
}

// Fx33 - LD B, Vx
func (c *Chip8) OP_Fx33() {
	Vx := (c.Opcode & 0x0F00) >> 8
	val := c.V[Vx]
	c.Memory[c.I+2] = val % 10
	val /= 10
	c.Memory[c.I+1] = val % 10
	val /= 10
	c.Memory[c.I] = val % 10
}

// Fx55 - LD [I], Vx
func (c *Chip8) OP_Fx55() {
	x := (c.Opcode & 0x0F00) >> 8
	for i := uint16(0); i <= x; i++ {
		c.Memory[c.I+i] = c.V[i]
	}
}

// Fx65 - LD Vx, [I]
func (c *Chip8) OP_Fx65() {
	x := (c.Opcode & 0x0F00) >> 8
	for i := uint16(0); i <= x; i++ {
		c.V[i] = c.Memory[c.I+i]
	}
}
