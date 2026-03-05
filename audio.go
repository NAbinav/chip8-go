package main

import (
	"math"
)

const sampleRate = 48000

type beep struct{ pos int }

func (b *beep) Read(buf []byte) (int, error) {
	for i := 0; i < len(buf); i += 4 {
		v := int16(math.Sin(float64(b.pos)*440*2*math.Pi/sampleRate) * 3000)
		buf[i], buf[i+1], buf[i+2], buf[i+3] = byte(v), byte(v>>8), byte(v), byte(v>>8)
		b.pos++
	}
	return len(buf), nil
}
