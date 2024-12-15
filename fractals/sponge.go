package fractals

import "opengl/utils"

func (h Handler) DrawSponge(iteration int) {
	start := utils.Point{X: -0.75, Y: -0.75}
	utils.FillRect(start, 1.5, 1.5, [3]float32{1, 1, 1}, h.ShaderProgram)
	h.drawSpongeRec(iteration, start, 0.5)
}

func (h Handler) drawSpongeRec(iteration int, start utils.Point, size float32) {
	if iteration == 0 {
		utils.FillRect(start, size, size, [3]float32{1, 1, 1}, h.ShaderProgram)
	} else {
		newStart := start
		h.drawSpongeRec(iteration-1, newStart, size/3)

		newStart = utils.Point{X: start.X + size, Y: start.Y}
		h.drawSpongeRec(iteration-1, newStart, size/3)

		newStart = utils.Point{X: start.X + (2 * size), Y: start.Y}
		h.drawSpongeRec(iteration-1, newStart, size/3)

		newStart = utils.Point{X: start.X, Y: start.Y + size}
		h.drawSpongeRec(iteration-1, newStart, size/3)

		// Center
		newStart = utils.Point{X: start.X + size, Y: start.Y + size}
		utils.FillRect(newStart, size, size, [3]float32{0, 0, 0}, h.ShaderProgram)

		newStart = utils.Point{X: start.X + (2 * size), Y: start.Y + size}
		h.drawSpongeRec(iteration-1, newStart, size/3)

		newStart = utils.Point{X: start.X, Y: start.Y + (2 * size)}
		h.drawSpongeRec(iteration-1, newStart, size/3)

		newStart = utils.Point{X: start.X + size, Y: start.Y + (2 * size)}
		h.drawSpongeRec(iteration-1, newStart, size/3)

		newStart = utils.Point{X: start.X + (2 * size), Y: start.Y + (2 * size)}
		h.drawSpongeRec(iteration-1, newStart, size/3)
	}
}
