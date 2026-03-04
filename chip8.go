package main

//memory:
//0x000-0x1FF (contains font set in emu)
//0x050-0x0A0 (built in 4x5 pixel font set (0-F))
//0x200-0xFFF - Program ROM and work RAM

type Chip8 struct {
	Memory     [4096]byte
	V          [16]byte //CPU Registers V0-VE
	Opcode     uint16
	I          uint16 //Index Registers
	PC         uint16
	Display    [64][32]bool
	DelayTimer byte
	SoundTimer byte
	Stack      [16]uint16
	SP         uint8
	Key        [16]bool
	table      [16]func()
	table0     [16]func()
	table8     [16]func()
	tableE     [16]func()
	tableF     [256]func()
}
