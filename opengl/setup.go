package opengl

import (
	"log"
	"opengl/fractals"
	"runtime"
	"strconv"
	"time"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var (
	shaderProgram uint32

	fps time.Duration = 60

	lastFrame time.Time = time.Now()
	nbFrames  int       = 0
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
	glfw.WindowHint(glfw.Samples, 4)

	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		log.Fatalln("Failed to create GLFW window:", err)
	}
	defer window.Destroy()

	window.MakeContextCurrent()

	glfw.SwapInterval(1)

	if err := gl.Init(); err != nil {
		log.Fatalln("Failed to initialize OpenGL:", err)
	}

	gl.Viewport(0, 0, int32(width), int32(height))

	gl.Enable(gl.MULTISAMPLE)

	shaderProgram = createShaderProgram()

	gl.UseProgram(shaderProgram)

	window.SwapBuffers()

	lastDraw := time.Now().Add(-time.Second)

	for !window.ShouldClose() {

		if lastDraw.Add(time.Second / fps).Before(time.Now()) {
			fpsCounter(window)

			gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

			DrawFractal()

			window.SwapBuffers()

			lastDraw = time.Now()
		}

		glfw.PollEvents()
	}
}

func DrawFractal() {
	fHandler := fractals.Handler{ShaderProgram: shaderProgram}
	// fHandler.DrawMountain(9)
	// fHandler.DrawDragonCurve(15)
	// fHandler.DrawSponge(4)
	fHandler.DrawTriangle(6)
}

func fpsCounter(window *glfw.Window) {
	current := time.Now()

	delta := current.Sub(lastFrame).Seconds()

	nbFrames++

	if delta >= 1.0 {
		fps := float64(nbFrames) / float64(delta)
		window.SetTitle("FPS: " + strconv.FormatFloat(fps, 'f', -1, 64))
		nbFrames = 0
		lastFrame = time.Now()
	}
}
