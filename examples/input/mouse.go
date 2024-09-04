package main

import (
	"log"
	"math/rand/v2"
	"runtime"

	p5go "github.com/stanipetrosyan/p5go"
)

func init() { runtime.LockOSThread() }

type model struct {
}

var (
	red   int = rand.IntN(255)
	green int = rand.IntN(255)
	blue  int = rand.IntN(255)
)

func main() {
	p := p5go.NewProgram(&model{}, p5go.P2D)

	err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func (m *model) Setup() *p5go.Window {
	c := p5go.Canvas2D(1080, 1080)

	return c
}

func (m *model) Draw(window *p5go.Window) {
	window.Background(red, green, blue)
	window.Fill(16, 200, 13)

	window.Rect(window.MouseX(), window.Height()/2, window.MouseY()/2+10, window.MouseY()/2+10)
	inverseX := window.Width() - window.MouseX()
	inverceY := window.Height() - window.MouseY()
	window.Rect(inverseX, window.Height()/2, (inverceY/2)+10, (inverceY/2)+10)
}
