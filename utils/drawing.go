package utils

import (
	"github.com/go-gl/gl/v4.6-core/gl"
)

func DrawLine(p1, p2 Point, color [3]float32, shaderProgram uint32) {
	vertices := []float32{
		p1.X, p1.Y, 0.0,
		p2.X, p2.Y, 0.0,
	}

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	// Create a Vertex Buffer Object (VBO)
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	// Enable vertex attribute array
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))

	// Use the shader program
	gl.UseProgram(shaderProgram)

	// Set the line color
	colorUniform := gl.GetUniformLocation(shaderProgram, gl.Str("lineColor\x00"))
	gl.Uniform3fv(colorUniform, 1, &color[0])

	// Draw the line
	gl.DrawArrays(gl.LINES, 0, 2)

	// Cleanup
	gl.DisableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)
	gl.DeleteBuffers(1, &vbo)
	gl.DeleteVertexArrays(1, &vao)
}
