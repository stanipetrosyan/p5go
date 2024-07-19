package p5go

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Matrix struct {
	FOVdeg    float32
	nearPlane float32
	farPlane  float32
}

type Camera struct {
	Position    mgl32.Vec3
	Orientation mgl32.Vec3
	Up          mgl32.Vec3

	width  int
	height int

	speed       float32
	sensitivity float32
}

func NewCamera(width, height int, position mgl32.Vec3) Camera {
	return Camera{
		Position:    position,
		Orientation: mgl32.Vec3{0.0, 0.0, -1.0},
		Up:          mgl32.Vec3{0.0, 1.0, 0.0},
		width:       width,
		height:      height,
		speed:       0.1,
		sensitivity: 100.0,
	}
}

func NewMatrix(program uint32, camera Camera, FOVdeg, nearPlane, farPlane float32) {
	projection := mgl32.Perspective(mgl32.DegToRad(45.0), (800 / 640), 0.1, 10.0)
	projectionUniform := gl.GetUniformLocation(program, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])

	view := mgl32.LookAtV(mgl32.Vec3{4, 5, 5}, mgl32.Vec3{2, 2, 2}, mgl32.Vec3{0, 1, 0})
	viewUniform := gl.GetUniformLocation(program, gl.Str("camera\x00"))
	gl.UniformMatrix4fv(viewUniform, 1, false, &view[0])

	model := mgl32.Ident4()
	modelUniform := gl.GetUniformLocation(program, gl.Str("model\x00"))
	gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])

	vertAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointerWithOffset(vertAttrib, 3, gl.FLOAT, false, 5*4, 0)
}
