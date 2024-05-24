package p5go

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Programm struct {
	proc Processing
}

type Processing interface {
	Setup() *Window
	Draw(*Window)
}

type Window struct {
	window *glfw.Window
}

func NewProgramm(processing Processing) Programm {
	return Programm{proc: processing}
}

func (p Programm) Run() error {
	var err error = nil

	err = glfw.Init()
	if err != nil {
		return err
	}

	defer glfw.Terminate()

	if err := gl.Init(); err != nil {
		return err
	}

	w := p.proc.Setup()

	for !w.window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)

		p.proc.Draw(w)
		w.window.SwapBuffers()
		glfw.PollEvents()
	}

	return err
}
