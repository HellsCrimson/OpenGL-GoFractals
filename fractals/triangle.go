package fractals

import "opengl/utils"

func (h Handler) DrawTriangle(iteration int) {
	p1 := utils.Point{X: -0.75, Y: -0.75}
	p2 := utils.Point{X: 0.75, Y: -0.75}
	p3 := utils.Point{X: 0, Y: 0.75}

	h.drawTriangleRec(iteration, p1, p2, p3)
}

func (h Handler) drawTriangleRec(iteration int, p1, p2, p3 utils.Point) {
	if iteration == 0 {
		utils.FillTriangle(p1, p2, p3, [3]float32{1, 1, 1}, h.ShaderProgram)
	} else {
		h.drawTriangleRec(iteration-1, p1, middle(p1, p2), middle(p1, p3))
		h.drawTriangleRec(iteration-1, p2, middle(p2, p1), middle(p2, p3))
		h.drawTriangleRec(iteration-1, p3, middle(p3, p1), middle(p3, p2))
	}
}

func middle(p1, p2 utils.Point) utils.Point {
	return utils.Point{
		X: (p1.X + p2.X) / 2,
		Y: (p1.Y + p2.Y) / 2,
	}
}
