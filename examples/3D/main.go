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
	red    int     = rand.IntN(255)
	green  int     = rand.IntN(255)
	blue   int     = rand.IntN(255)
	angleX float32 = 1.0
)

func main() {
	p := p5go.NewProgramm(&model{}, p5go.P3D)

	err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func (m *model) Setup() *p5go.Window {
	c := p5go.Canvas3D(1080, 1080, 1080)

	return c
}

func (m *model) Draw(window *p5go.Window) {
	window.Background(red, green, blue)
	window.Camera().RotateY(angleX)
	window.Shape().Box(200, 200, 200, 200)s

	angleX += 0.1
}
