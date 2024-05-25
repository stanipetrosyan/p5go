package p5go

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Window struct {
	window *glfw.Window
}

func Canvas(width, height int) *Window {
	window, err := glfw.CreateWindow(width, height, "", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	return &Window{window: window}
}

func (w *Window) Background(red, green, blue int) {
	r_float := (1.0 / 255) * float32(red)
	g_float := (1.0 / 255) * float32(green)
	b_float := (1.0 / 255) * float32(blue)
	gl.ClearColor(r_float, g_float, b_float, 1.0)
}
