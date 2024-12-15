package opengl

import (
	"log"

	"github.com/go-gl/gl/v4.6-core/gl"
)

func createShaderProgram() uint32 {
	// Vertex shader source code
	vertexShaderSource := `#version 460 core
	layout(location = 0) in vec3 position;
	void main() {
		gl_Position = vec4(position, 1.0);
	}` + "\x00"

	// Fragment shader source code
	fragmentShaderSource := `#version 460 core
	out vec4 fragColor;
	void main() {
		fragColor = vec4(1.0, 0.0, 0.0, 1.0); // White color
	}` + "\x00"

	// Compile vertex shader
	vertexShader := gl.CreateShader(gl.VERTEX_SHADER)
	cVertexShaderSource, freeVertex := gl.Strs(vertexShaderSource)
	gl.ShaderSource(vertexShader, 1, cVertexShaderSource, nil)
	freeVertex()
	gl.CompileShader(vertexShader)
	checkShaderCompileStatus(vertexShader)

	// Compile fragment shader
	fragmentShader := gl.CreateShader(gl.FRAGMENT_SHADER)
	cFragmentShaderSource, freeFragment := gl.Strs(fragmentShaderSource)
	gl.ShaderSource(fragmentShader, 1, cFragmentShaderSource, nil)
	freeFragment()
	gl.CompileShader(fragmentShader)
	checkShaderCompileStatus(fragmentShader)

	// Link shaders into a program
	program := gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)
	checkProgramLinkStatus(program)

	// Cleanup
	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return program
}

func checkShaderCompileStatus(shader uint32) {
	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		logBuf := make([]byte, logLength+1)
		gl.GetShaderInfoLog(shader, logLength, nil, &logBuf[0])
		log.Fatalln("Shader compile error:", string(logBuf))
	}
}

func checkProgramLinkStatus(program uint32) {
	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		logBuf := make([]byte, logLength+1)
		gl.GetProgramInfoLog(program, logLength, nil, &logBuf[0])
		log.Fatalln("Program link error:", string(logBuf))
	}
}
