package fractals

import "opengl/utils"

func (h Handler) DrawVicsekCross(iteration int) {
	start := utils.Point{X: -0.75, Y: -0.75}

	h.drawVicsekCrossRec(iteration, start, 1.5)
}

func (h Handler) drawVicsekCrossRec(iteration int, start utils.Point, size float32) {
	if iteration == 0 {
		utils.FillRect(start, size, size, [3]float32{1, 1, 1}, h.ShaderProgram)
	} else {
		size = size / 3

		p := utils.Point{
			X: start.X + size,
			Y: start.Y,
		}
		h.drawVicsekCrossRec(iteration-1, p, size)

		p = utils.Point{
			X: start.X,
			Y: start.Y + size,
		}
		h.drawVicsekCrossRec(iteration-1, p, size)

		p = utils.Point{
			X: start.X + size,
			Y: start.Y + size,
		}
		h.drawVicsekCrossRec(iteration-1, p, size)

		p = utils.Point{
			X: start.X + (2 * size),
			Y: start.Y + size,
		}
		h.drawVicsekCrossRec(iteration-1, p, size)

		p = utils.Point{
			X: start.X + size,
			Y: start.Y + (size * 2),
		}
		h.drawVicsekCrossRec(iteration-1, p, size)
	}
}
