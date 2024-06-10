package p5go

import (
	"math"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Window struct {
	window *glfw.Window
	width  int
	height int
}

func Canvas(width, height int) *Window {
	window, err := glfw.CreateWindow(width, height, "", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	return &Window{window: window, width: width, height: height}
}

func (w *Window) Background(red, green, blue int) {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	r_float := (1.0 / 255) * float32(red)
	g_float := (1.0 / 255) * float32(green)
	b_float := (1.0 / 255) * float32(blue)
	gl.ClearColor(r_float, g_float, b_float, 1.0)
}

func (w *Window) Line(x1, y1, x2, y2 float32) {
	x1f := fromWorldToLocalSpace(x1, w.width)
	y1f := fromWorldToLocalSpace(y1, w.height)
	x2f := fromWorldToLocalSpace(x2, w.width)
	y2f := fromWorldToLocalSpace(y2, w.height)

	line := []float32{
		x1f, y1f,
		x2f, y2f,
	}

	vbo := generatevbo(line)
	vao := generate2Dvao(vbo)

	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.LINES, 0, 2)
}

func (w *Window) Triangle(x1, y1, x2, y2, x3, y3 float32) {
	x1f := fromWorldToLocalSpace(x1, w.width)
	y1f := fromWorldToLocalSpace(y1, w.height)
	x2f := fromWorldToLocalSpace(x2, w.width)
	y2f := fromWorldToLocalSpace(y2, w.height)
	x3f := fromWorldToLocalSpace(x3, w.width)
	y3f := fromWorldToLocalSpace(y3, w.height)

	triangle := []float32{
		x1f, y1f,
		x2f, y2f,
		x3f, y3f,
	}

	vbo := generatevbo(triangle)
	vao := generate2Dvao(vbo)

	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
}

func (w *Window) Quad(x1, y1, x2, y2, x3, y3, x4, y4 float32) {

	x1f := fromWorldToLocalSpace(x1, w.width)
	y1f := fromWorldToLocalSpace(y1, w.height)
	x2f := fromWorldToLocalSpace(x2, w.width)
	y2f := fromWorldToLocalSpace(y2, w.height)
	x3f := fromWorldToLocalSpace(x3, w.width)
	y3f := fromWorldToLocalSpace(y3, w.height)
	x4f := fromWorldToLocalSpace(x4, w.width)
	y4f := fromWorldToLocalSpace(y4, w.height)

	quad := []float32{
		x1f, y1f,
		x2f, y2f,
		x3f, y3f,
		x4f, y4f,
	}

	vbo := generatevbo(quad)
	vao := generate2Dvao(vbo)

	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.QUADS, 0, 4)
}

func (w *Window) Rect(x1, y1, width, height float32) {
	w.Triangle(x1, y1, x1+width, y1, x1, y1+height)
	w.Triangle(x1+width, y1, x1, y1+height, x1+width, y1+height)
}

func (w *Window) Square(x1, y1, size float32) {
	w.Rect(x1, y1, size, size)
}

func (w *Window) Circle(x1, y1, radius float32) {
	triangles := 128
	twicePi := math.Pi * 2.0

	prevX := x1
	prevY := y1 - radius

	for i := 0; i <= triangles; i++ {
		delta := float64(i) * twicePi
		newX := float64(x1) + (float64(radius) * math.Cos(delta/float64(triangles)))
		newY := float64(y1) + (float64(radius) * math.Sin(delta/float64(triangles)))

		w.Triangle(x1, y1, prevX, prevY, float32(newX), float32(newY))

		prevX = float32(newX)
		prevY = float32(newY)
	}

}

func fromWorldToLocalSpace(world float32, axis int) float32 {
	return ((world / float32(axis)) * 2) - 1.0
}

func generatevbo(array []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 9*len(array), gl.Ptr(array), gl.STATIC_DRAW)

	return vbo
}

func generate2Dvao(vbo uint32) uint32 {
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 0, nil)

	return vao
}
