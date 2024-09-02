package p5go

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	width     int
	height    int
	position  mgl32.Vec3
	center    mgl32.Vec3
	up        mgl32.Vec3
	rotation  mgl32.Vec3
	FOV       float32
	nearPlane float32
	farPlane  float32
}

func centeredCamera(width, height int) Camera {
	return Camera{
		position:  mgl32.Vec3{0.0, 0.0, 2.0},
		center:    mgl32.Vec3{0.0, 0.0, 0.0},
		up:        mgl32.Vec3{0.0, 1.0, 0.0},
		width:     width,
		height:    height,
		FOV:       45.0,
		nearPlane: 0.1,
		farPlane:  100.0,
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

// Sets a perspective projection of camera. A perspective is a large frustum that defines the visible space.
// fovy should be specified in degrees (values from 0 to 360).
func (c *Camera) Perspective(fovy, near, far float32) {
	c.FOV = fovy
	c.nearPlane = near
	c.farPlane = far
}
