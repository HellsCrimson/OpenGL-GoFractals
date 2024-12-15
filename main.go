package main

import "opengl/opengl"

const (
	width  = 1000
	height = 800
	title  = "OpenGL Window"
)

func main() {
	opengl.DoSetup(width, height, title)
}
