package fractals

import "opengl/utils"

func (h Handler) DrawVicsekTree(iteration int) {
	start := utils.Point{X: -0.75, Y: -0.75}

	h.drawVicsekTreeRec(iteration, start, 1.5)
}

func (h Handler) drawVicsekTreeRec(iteration int, start utils.Point, size float32) {
	if iteration == 0 {
		utils.FillRect(start, size, size, [3]float32{1, 1, 1}, h.ShaderProgram)
	} else {
		size = size / 3

		p := utils.Point{
			X: start.X,
			Y: start.Y,
		}
		h.drawVicsekTreeRec(iteration-1, p, size)

		p = utils.Point{
			X: start.X + size,
			Y: start.Y,
		}
		h.drawVicsekTreeRec(iteration-1, p, size)

		p = utils.Point{
			X: start.X + (2 * size),
			Y: start.Y,
		}
		h.drawVicsekTreeRec(iteration-1, p, size)

		p = utils.Point{
			X: start.X + size,
			Y: start.Y + size,
		}
		h.drawVicsekTreeRec(iteration-1, p, size)

		p = utils.Point{
			X: start.X + size,
			Y: start.Y + (size * 2),
		}
		h.drawVicsekTreeRec(iteration-1, p, size)
	}
}

func (h Handler) DrawVicsekTreeInv(iteration int) {
	start := utils.Point{X: -0.75, Y: -0.75}

	h.drawVicsekTreeRecInv(iteration, start, 1.5)
}

func (h Handler) drawVicsekTreeRecInv(iteration int, start utils.Point, size float32) {
	if iteration == 0 {
		utils.FillRect(start, size, size, [3]float32{1, 1, 1}, h.ShaderProgram)
	} else {
		size = size / 3

		p := utils.Point{
			X: start.X + size,
			Y: start.Y,
		}
		h.drawVicsekTreeRecInv(iteration-1, p, size)

		p = utils.Point{
			X: start.X + size,
			Y: start.Y + size,
		}
		h.drawVicsekTreeRecInv(iteration-1, p, size)

		p = utils.Point{
			X: start.X,
			Y: start.Y + (size * 2),
		}
		h.drawVicsekTreeRecInv(iteration-1, p, size)

		p = utils.Point{
			X: start.X + size,
			Y: start.Y + (size * 2),
		}
		h.drawVicsekTreeRecInv(iteration-1, p, size)

		p = utils.Point{
			X: start.X + (2 * size),
			Y: start.Y + (size * 2),
		}
		h.drawVicsekTreeRecInv(iteration-1, p, size)
	}
}
