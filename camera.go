package p5go

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	width    int
	height   int
	position mgl32.Vec3
	center   mgl32.Vec3
	up       mgl32.Vec3
}

func NewCamera(width, height int, position mgl32.Vec3, center mgl32.Vec3) Camera {
	return Camera{
		width:    width,
		height:   height,
		position: position,
		center:   center,
		up:       mgl32.Vec3{0.0, 1.0, 0.0},
	}
}

func CenteredCamera(width, height int) Camera {
	return Camera{
		position: mgl32.Vec3{0.0, 0.0, 1.0},
		center:   mgl32.Vec3{0.0, 0.0, 0.0},
		up:       mgl32.Vec3{0.0, 0.0, 0.0},
		width:    width,
		height:   height,
	}
}

func NewMatrix(program uint32, camera Camera, FOVdeg, nearPlane, farPlane float32) {
	projection := mgl32.Perspective(mgl32.DegToRad(FOVdeg), float32(camera.width/camera.height), nearPlane, farPlane)
	projectionUniform := gl.GetUniformLocation(program, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])

	view := mgl32.LookAt(
		camera.position.X(), camera.position.Y(), camera.position.Z(),
		camera.center.X(), camera.center.Y(), camera.center.Z(),
		camera.up.X(), camera.up.Y(), camera.up.Z(),
	)
	viewUniform := gl.GetUniformLocation(program, gl.Str("camera\x00"))
	gl.UniformMatrix4fv(viewUniform, 1, false, &view[0])

	// model := mgl32.HomogRotate3D(float32(0.0), mgl32.Vec3{0, 1, 0})
	model := mgl32.Ident4()
	modelUniform := gl.GetUniformLocation(program, gl.Str("model\x00"))
	gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])

	vertAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointerWithOffset(vertAttrib, 3, gl.FLOAT, false, 5*4, 0)
}
