package p5go

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

type Renderer int

const (
	P2D Renderer = iota
	P3D
)

type Program struct {
	proc     Processing
	renderer Renderer
}

type Processing interface {
	Setup() *Window
	Draw(*Window)
}

func NewProgram(processing Processing, renderer Renderer) Program {
	return Program{proc: processing, renderer: renderer}
}

func (p Program) Run() error {
	var err error = nil

	err = glfw.Init()
	if err != nil {
		return err
	}

	defer glfw.Terminate()

	w := p.proc.Setup()
	w.window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		return err
	}
	program := gl.CreateProgram()

	fragmentShaderCompiled, err := compileShader(fragmentShader, gl.FRAGMENT_SHADER)
	if err != nil {
		log.Fatal(err)
	}

	err = attachShader(program, fragmentShaderCompiled)
	if err != nil {
		log.Fatal(err)
	}

	if p.renderer == P3D {
		cameraShaderCompiled, err := compileShader(cameraShader, gl.VERTEX_SHADER)
		if err != nil {
			log.Fatal(err)
		}

		err = attachShader(program, cameraShaderCompiled)
		if err != nil {
			log.Fatal(err)
		}
	}

	gl.UseProgram(program)

	t1 := time.Now().UnixNano()
	space := 1000000000.0 / 60.0
	//  actually draw function is called 60 times per second
	for !w.window.ShouldClose() {
		gl.UseProgram(program)

		t2 := time.Now().UnixNano()
		if (t2 - t1) > int64(space) {

			if p.renderer == P3D {
				renderMatrix(program, w.camera)
			}

			colorUniform := gl.GetUniformLocation(program, gl.Str("shapesColor\x00"))
			gl.Uniform3f(colorUniform, w.color[0], w.color[1], w.color[2])

			p.proc.Draw(w)
			w.window.SwapBuffers()
			glfw.PollEvents()
			w.mouseX, w.mouseY = glfw.GetCurrentContext().GetCursorPos()
			t1 = time.Now().UnixNano()
		}
	}

	return err
}

func attachShader(program uint32, shader uint32) error {
	gl.AttachShader(program, shader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(shader)
	return nil
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

func renderMatrix(program uint32, camera Camera) {
	projection := mgl32.Perspective(mgl32.DegToRad(camera.FOV), float32(camera.width/camera.height), camera.nearPlane, camera.farPlane)
	projectionUniform := gl.GetUniformLocation(program, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])

	view := mgl32.LookAt(
		camera.position.X(), camera.position.Y(), camera.position.Z(),
		camera.center.X(), camera.center.Y(), camera.center.Z(),
		camera.up.X(), camera.up.Y(), camera.up.Z(),
	)
	viewUniform := gl.GetUniformLocation(program, gl.Str("camera\x00"))
	gl.UniformMatrix4fv(viewUniform, 1, false, &view[0])

	uniformModel(camera.rotation, program)

	vertAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointerWithOffset(vertAttrib, 3, gl.FLOAT, false, 5*4, 0)
}

func uniformModel(vec mgl32.Vec3, program uint32) {
	var model mgl32.Mat4
	switch {
	case vec.X() > 0:
		model = mgl32.HomogRotate3DX(vec.X())
	case vec.Y() > 0:
		model = mgl32.HomogRotate3DY(vec.Y())
	case vec.Z() > 0:
		model = mgl32.HomogRotate3DZ(vec.Z())
	default:
		model = mgl32.Ident4()
	}

	modelUniform := gl.GetUniformLocation(program, gl.Str("model\x00"))
	gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])
}

var cameraShader = `
#version 330 core

uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;

in vec3 vert;

void main() {
    gl_Position = projection * camera * model * vec4(vert, 1);
}
` + "\x00"

var fragmentShader = `
#version 330 core

uniform vec3 shapesColor;

out vec4 outColor;

void main()
{
    outColor = vec4(shapesColor, 1.0);
}
` + "\x00"
