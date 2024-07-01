package p5go

import (
	"time"

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
	w.window.MakeContextCurrent()

	t1 := time.Now().UnixNano()

	space := 1000000000.0 / 60.0
	//  actually draw function is called 60 times per second
	for !w.window.ShouldClose() {

		t2 := time.Now().UnixNano()
		if (t2 - t1) > int64(space) {
			p.proc.Draw(w)

			go glfw.PollEvents()
			w.window.SwapBuffers()
			t1 = time.Now().UnixNano()
		}
	}

	return err
}
