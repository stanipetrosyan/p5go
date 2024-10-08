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
	// background(51);
	//  fill(255, 204);
	//  rect(mouseX, height/2, mouseY/2+10, mouseY/2+10);
	//  fill(255, 204);
	//  int inverseX = width-mouseX;
	//  int inverseY = height-mouseY;
	//  rect(inverseX, height/2, (inverseY/2)+10, (inverseY/2)+10);

	window.Background(red, green, blue)

	window.Rect(window.MouseX(), window.Height()/2, window.MouseY()/2+10, window.MouseY()/2+10)
	inverseX := window.Width() - window.MouseX()
	inverceY := window.Height() - window.MouseY()
	window.Rect(inverseX, window.Height()/2, (inverceY/2)+10, (inverceY/2)+10)

	// window.Rect(700, 800, 220, 220)
	// window.Triangle(500, 200, 600, 400, 700, 200)
	// window.Circle(200, 400, 100)
	// window.Ellipse(500, 500, 100, 200)
	// window.Point(20, 20)
}
