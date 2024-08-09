package p5go

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	width    int
	height   int
	position mgl32.Vec3
	center   mgl32.Vec3
	up       mgl32.Vec3
	rotation mgl32.Vec3
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
		position: mgl32.Vec3{0.0, 0.0, 2.0},
		center:   mgl32.Vec3{0.0, 0.0, 0.0},
		up:       mgl32.Vec3{0.0, 1.0, 0.0},
		width:    width,
		height:   height,
	}
}

// Rotates a camera around the x-axis the amount specified by the angle parameter.
// Angles should be specified in degrees (values from 0 to 360).
func (c *Camera) RotateX(angle float32) {
	c.rotation = mgl32.Vec3{angle, 0, 0}
}

// Rotates a camera around the y-axis the amount specified by the angle parameter.
// Angles should be specified in degrees (values from 0 to 360).
func (c *Camera) RotateY(angle float32) {
	c.rotation = mgl32.Vec3{0, angle, 0}
}

// Rotates a camera around the z-axis the amount specified by the angle parameter.
// Angles should be specified in degrees (values from 0 to 360).
func (c *Camera) RotateZ(angle float32) {
	c.rotation = mgl32.Vec3{0, 0, angle}
}
