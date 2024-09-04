package p5go

import (
	"math"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

// Windows struct
type Window struct {
	window *glfw.Window
	width  int
	height int
	depth  int
	camera Camera
	shape  Shape
	mouseX float64
	mouseY float64
}

// Canvas2D returns a Window with width and height.
//
// Parameters should be in pixel.
func Canvas2D(width, height int) *Window {
	window, err := glfw.CreateWindow(width, height, "", nil, nil)
	if err != nil {
		panic(err)
	}

	return &Window{window: window, width: width, height: height}
}

// Canvas3D returns a Window with width, height and depth.
//
// Parameters should be in pixel.
func Canvas3D(width, height, depth int) *Window {
	window, err := glfw.CreateWindow(width, height, "", nil, nil)
	if err != nil {
		panic(err)
	}

	camera := centeredCamera(width, height)
	shape := Shape{
		width:  width,
		height: height,
		depth:  depth,
	}

	return &Window{window: window, width: width, height: height, depth: depth, camera: camera, shape: shape}
}

// Returns Camera object with default values:
//
// position: {0, 0, 2}
// center: {0, 0, 0}
// up: {0, 1, 0}
// FOV: 45.0
// nearPlane: 0.1
// farPlane: 100.0
func (w *Window) Camera() *Camera {
	return &w.camera
}

func (w *Window) Shape() Shape {
	return w.shape
}

// Returns height of window
func (w *Window) Height() float32 {
	return float32(w.height)
}

// Returns width of window
func (w *Window) Width() float32 {
	return float32(w.width)
}

// Returns current x-axis position of mouse
func (w *Window) MouseX() float32 {
	return float32(w.mouseX)
}

// Returns current y-axis position of mouse
func (w *Window) MouseY() float32 {
	return float32(w.mouseY)
}

// Background change color of window.
//
// Parameters accepted are RGB values: 0-255
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
	w.Ellipse(x1, y1, radius, radius)
}

func (w *Window) Ellipse(x1, y1, width, height float32) {
	triangles := 360
	angle := float32(360 / triangles)

	prevX := x1
	prevY := y1

	for i := 0; i <= triangles; i++ {
		newX := float64(x1) + (float64(width) * math.Cos(float64(toRadians(angle*float32(i)))))
		newY := float64(y1) + (float64(height) * math.Sin(float64(toRadians(angle*float32(i)))))

		w.Triangle(x1, y1, prevX, prevY, float32(newX), float32(newY))

		prevX = float32(newX)
		prevY = float32(newY)
	}
}

func (w *Window) Arc(x1, y1, width, height, start, stop float32) {
	triangles := 360
	angle := float32(360 / triangles)

	prevX := x1
	prevY := y1

	for i := start; i <= stop; i++ {
		newX := float64(x1) + (float64(width) * math.Cos(float64(toRadians(angle*float32(i)))))
		newY := float64(y1) + (float64(height) * math.Sin(float64(toRadians(angle*float32(i)))))

		w.Triangle(x1, y1, prevX, prevY, float32(newX), float32(newY))

		prevX = float32(newX)
		prevY = float32(newY)
	}
}

func toRadians(degrees float32) float32 {
	return degrees * math.Pi / 180
}

func (w *Window) Point(x1, y1 float32) {
	x1f := fromWorldToLocalSpace(x1, w.width)
	y1f := fromWorldToLocalSpace(y1, w.height)

	quad := []float32{
		x1f, y1f,
	}

	vbo := generatevbo(quad)
	vao := generate2Dvao(vbo)

	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.POINTS, 0, 1)
}

func fromWorldToLocalSpace(world float32, axis int) float32 {
	return ((world / float32(axis)) * 2) - 1.0
}

func generatevbo(array []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(array), gl.Ptr(array), gl.STATIC_DRAW)

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
