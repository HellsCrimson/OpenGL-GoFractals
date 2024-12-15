package opengl

import (
	"log"
	"opengl/fractals"
	"runtime"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var (
	shaderProgram uint32
)

func init() {
	runtime.LockOSThread()
}

func DoSetup(width, height int, title string) {
	if err := glfw.Init(); err != nil {
		log.Fatalln("Failed to initialize GLFW:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 6)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.Resizable, glfw.False)

	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		log.Fatalln("Failed to create GLFW window:", err)
	}
	defer window.Destroy()

	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		log.Fatalln("Failed to initialize OpenGL:", err)
	}

	gl.Viewport(0, 0, int32(width), int32(height))

	shaderProgram = createShaderProgram()

	gl.UseProgram(shaderProgram)

	window.SwapBuffers()

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	DrawFractal()

	window.SwapBuffers()

	for !window.ShouldClose() {
		glfw.PollEvents()
	}
}

func DrawFractal() {
	fHandler := fractals.Handler{ShaderProgram: shaderProgram}
	// fHandler.DrawMountain(9)
	// fHandler.DrawDragonCurve(20)
	fHandler.DrawSponge(0)
}
