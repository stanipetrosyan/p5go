package main

import (
	"log"
	"math/rand/v2"
	"time"

	p5go "github.com/stanipetrosyan/p5go"
)

type model struct {
}

var (
	red   int = rand.IntN(255)
	green int = rand.IntN(255)
	blue  int = rand.IntN(255)
)

func main() {
	p := p5go.NewProgramm(&model{})

	err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func (m *model) Setup() *p5go.Window {
	c := p5go.Canvas(640, 480)

	return c
}

func (m *model) Draw(window *p5go.Window) {
	time.Sleep(time.Second / 4)

	window.Background(red, green, blue)

	red = rand.IntN(255)
	green = rand.IntN(255)
	blue = rand.IntN(255)

	window.Rect(10, 200, 220, 220)
}
