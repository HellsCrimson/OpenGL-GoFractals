package main

import "opengl/opengl"

const (
	width  = 800
	height = 600
	title  = "OpenGL Window"
)

func main() {
	opengl.DoSetup(width, height, title)
}
