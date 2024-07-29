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

	w := p.proc.Setup()
	w.window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		return err
	}

	cameraShaderCompiled, err := compileShader(cameraShader, gl.VERTEX_SHADER)
	if err != nil {
		log.Fatal(err)
	}

	program := gl.CreateProgram()

	gl.AttachShader(program, cameraShaderCompiled)
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

	gl.DeleteShader(cameraShaderCompiled)
	gl.UseProgram(program)

	projection := mgl32.Perspective(mgl32.DegToRad(45.0), float32(w.width/w.height), 0.1, 10.0)
	projectionUniform := gl.GetUniformLocation(program, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])

	camera := mgl32.LookAtV(mgl32.Vec3{3, 3, 3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
	cameraUniform := gl.GetUniformLocation(program, gl.Str("camera\x00"))
	gl.UniformMatrix4fv(cameraUniform, 1, false, &camera[0])

	model := mgl32.Ident4()
	modelUniform := gl.GetUniformLocation(program, gl.Str("model\x00"))
	gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])

	vertAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointerWithOffset(vertAttrib, 3, gl.FLOAT, false, 5*4, 0)

	t1 := time.Now().UnixNano()
	space := 1000000000.0 / 60.0
	//  actually draw function is called 60 times per second
	for !w.window.ShouldClose() {

		t2 := time.Now().UnixNano()
		if (t2 - t1) > int64(space) {
			// NewMatrix(program, w.camera, 45.0, 0.1, 10.0)

			model := mgl32.HomogRotate3D(float32(45), mgl32.Vec3{0, 1, 0})
			gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])

			p.proc.Draw(w)
			w.window.SwapBuffers()
			glfw.PollEvents()
			t1 = time.Now().UnixNano()
		}
	}

	return err
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
