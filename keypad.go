package main

func (c *Chip8) KeyDown(k byte) {
	c.Key[k] = true
}

func (c *Chip8) KeyUp(k byte) {
	c.Key[k] = false
}
