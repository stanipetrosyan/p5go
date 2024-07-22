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

var eyeZ float32 = 0.0
var distance float32 = 0.1

func (m *model) Draw(window *p5go.Window) {
	window.Background(red, green, blue)

	window.Camera(0.0, 0.0, eyeZ)

	window.Square(200, 200, 200)

	eyeZ += distance
}
