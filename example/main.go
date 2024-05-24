package main

import (
	"log"

	p5go "github.com/stanipetrosyan/p5go"
)

type model struct {
}

func main() {
	p := p5go.NewProgramm(&model{})

	err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func (m *model) Setup() *p5go.Window {

}

func (m *model) Draw(window *p5go.Window) {
}
