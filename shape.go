package p5go

import "github.com/go-gl/gl/v4.1-core/gl"

type Shape struct {
	width  int
	height int
	depth  int
}

func (s *Shape) triangle(x1, y1, z1, x2, y2, z2, x3, y3, z3 float32) {
	x1f := fromWorldToLocalSpace(x1, s.width)
	y1f := fromWorldToLocalSpace(y1, s.height)
	z1f := fromWorldToLocalSpace(z1, s.depth)
	x2f := fromWorldToLocalSpace(x2, s.width)
	y2f := fromWorldToLocalSpace(y2, s.height)
	z2f := fromWorldToLocalSpace(z2, s.depth)
	x3f := fromWorldToLocalSpace(x3, s.width)
	y3f := fromWorldToLocalSpace(y3, s.height)
	z3f := fromWorldToLocalSpace(z3, s.depth)

	triangle := []float32{
		x1f, y1f, z1f,
		x2f, y2f, z2f,
		x3f, y3f, z3f,
	}
	vbo := generatevbo(triangle)
	vao := generate3Dvao(vbo)

	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
}

func (w Shape) Box(x1, y1, z1, size float32) {
	// face front
	w.triangle(x1, y1, z1, x1+size, y1+size, z1, x1, y1+size, z1)
	w.triangle(x1, y1, z1, x1+size, y1, z1, x1+size, y1+size, z1)

	// // face left
	w.triangle(x1, y1, z1, x1, y1+size, z1+size, x1, y1+size, z1)
	w.triangle(x1, y1, z1, x1, y1+size, z1+size, x1, y1, z1+size)

	// // face back
	w.triangle(x1, y1+size, z1+size, x1, y1, z1+size, x1+size, y1+size, z1+size)
	w.triangle(x1+size, y1, z1+size, x1, y1, z1+size, x1+size, y1+size, z1+size)

	// // face right
	w.triangle(x1+size, y1+size, z1, x1+size, y1+size, z1+size, x1+size, y1, z1+size)
	w.triangle(x1+size, y1, z1, x1+size, y1+size, z1, x1+size, y1, z1+size)

	// face up
	w.triangle(x1, y1+size, z1, x1, y1+size, z1+size, x1+size, y1+size, z1)
	w.triangle(x1+size, y1+size, z1, x1, y1+size, z1+size, x1+size, y1+size, z1+size)

	// face bottom
	w.triangle(x1, y1, z1, x1, y1, z1+size, x1+size, y1, z1)
	w.triangle(x1+size, y1, z1, x1, y1, z1+size, x1+size, y1, z1+size)
}

func generate3Dvao(vbo uint32) uint32 {
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}
