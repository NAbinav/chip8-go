package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"image/color"
)

var white = color.RGBA{255, 255, 255, 255}

const scale = 10

type Game struct {
	chip8  *Chip8
	player *audio.Player
}

func (g *Game) Update() error {
	handleInput(g.chip8)
	for i := 0; i < 10; i++ {
		g.chip8.cycle()
	}
	if g.chip8.DelayTimer > 0 {
		g.chip8.DelayTimer--
	}
	if g.chip8.SoundTimer > 0 {
		g.chip8.SoundTimer--
	}
	if g.chip8.SoundTimer > 0 && !g.player.IsPlaying() {
		g.player.Play()
	} else if g.chip8.SoundTimer == 0 && g.player.IsPlaying() {
		g.player.Pause()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	for y := 0; y < 32; y++ {
		for x := 0; x < 64; x++ {
			if g.chip8.Display[y][x] {
				for dy := 0; dy < scale; dy++ {
					for dx := 0; dx < scale; dx++ {
						screen.Set(x*scale+dx, y*scale+dy, white)
					}
				}
			}
		}
	}
}

func (g *Game) Layout(w, h int) (int, int) {
	return 64 * scale, 32 * scale
}

func main() {
	chip := &Chip8{}
	chip.Init()
	chip.LoadROM("./Brix.ch8")

	player, _ := audio.NewContext(sampleRate).NewPlayer(&beep{})
	game := &Game{chip8: chip, player: player}

	ebiten.SetWindowSize(64*scale, 32*scale)
	ebiten.SetWindowTitle("CHIP-8")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func handleInput(c *Chip8) {
	keys := map[ebiten.Key]byte{
		ebiten.Key1: 0x1,
		ebiten.Key2: 0x2,
		ebiten.Key3: 0x3,
		ebiten.Key4: 0xC,
		ebiten.KeyQ: 0x4,
		ebiten.KeyW: 0x5,
		ebiten.KeyE: 0x6,
		ebiten.KeyR: 0xD,
		ebiten.KeyA: 0x7,
		ebiten.KeyS: 0x8,
		ebiten.KeyD: 0x9,
		ebiten.KeyF: 0xE,
		ebiten.KeyZ: 0xA,
		ebiten.KeyX: 0x0,
		ebiten.KeyC: 0xB,
		ebiten.KeyV: 0xF,
	}

	for k, v := range keys {
		c.Key[v] = ebiten.IsKeyPressed(k)
	}
}
