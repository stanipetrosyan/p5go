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

var (
	circleX float32 = 100
	circleY float32 = 100

	xspeed float32 = 10.0
	yspeed float32 = 1.8

	xdirection float32 = 1.0
	ydirection float32 = 1.0

	radius float32 = 50
)

func (m *model) Setup() *p5go.Window {
	c := p5go.Canvas(1080, 1080)

	return c
}

func (m *model) Draw(window *p5go.Window) {
	window.Background(red, green, blue)

	circleX = circleX + (xspeed * xdirection)
	circleY = circleY + (yspeed * ydirection)

	if circleX > 1080-radius || circleX < radius {
		xdirection *= -1
	}
	if circleY > 1080-radius || circleY < radius {
		ydirection *= -1
	}
	window.Circle(circleX, circleY, radius)
}
