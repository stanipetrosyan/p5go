package main

import (
	"log"
	"math/rand/v2"

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
	c := p5go.Canvas(1080, 1080)

	return c
}

func (m *model) Draw(window *p5go.Window) {

	window.Background(red, green, blue)

	window.Rect(10, 200, 220, 220)
	window.Triangle(500, 200, 600, 400, 700, 200)
	window.Circle(500, 500, 100)
}
