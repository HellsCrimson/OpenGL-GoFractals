package fractals

import "opengl/utils"

func (h Handler) DrawVicsekTreeComb(iteration int) {
	start := utils.Point{X: -0.75, Y: -0.75}

	h.drawVicsekTreeComb(iteration, start, 1.5)
}

func (h Handler) drawVicsekTreeComb(iteration int, start utils.Point, size float32) {
	if iteration == 0 {
		utils.FillRect(start, size, size, [3]float32{1, 1, 1}, h.ShaderProgram)
	} else {
		size = size / 3

		p := utils.Point{
			X: start.X,
			Y: start.Y,
		}
		h.drawVicsekTreeComb(iteration-1, p, size)

		p = utils.Point{
			X: start.X + size,
			Y: start.Y,
		}
		h.drawVicsekTreeComb(iteration-1, p, size)

		p = utils.Point{
			X: start.X + (2 * size),
			Y: start.Y,
		}
		h.drawVicsekTreeComb(iteration-1, p, size)

		p = utils.Point{
			X: start.X + size,
			Y: start.Y + size,
		}
		h.drawVicsekTreeComb(iteration-1, p, size)

		p = utils.Point{
			X: start.X,
			Y: start.Y + (size * 2),
		}
		h.drawVicsekTreeComb(iteration-1, p, size)

		p = utils.Point{
			X: start.X + size,
			Y: start.Y + (size * 2),
		}
		h.drawVicsekTreeComb(iteration-1, p, size)

		p = utils.Point{
			X: start.X + (2 * size),
			Y: start.Y + (size * 2),
		}
		h.drawVicsekTreeComb(iteration-1, p, size)
	}
}

func (h Handler) DrawVicsekTreeCombInv(iteration int) {
	start := utils.Point{X: -0.75, Y: -0.75}

	h.drawVicsekTreeCombInv(iteration, start, 1.5)
}

func (h Handler) drawVicsekTreeCombInv(iteration int, start utils.Point, size float32) {
	if iteration == 0 {
		utils.FillRect(start, size, size, [3]float32{1, 1, 1}, h.ShaderProgram)
	} else {
		size = size / 3

		p := utils.Point{
			X: start.X,
			Y: start.Y,
		}
		h.drawVicsekTreeCombInv(iteration-1, p, size)

		p = utils.Point{
			X: start.X,
			Y: start.Y + size,
		}
		h.drawVicsekTreeCombInv(iteration-1, p, size)

		p = utils.Point{
			X: start.X,
			Y: start.Y + (2 * size),
		}
		h.drawVicsekTreeCombInv(iteration-1, p, size)

		p = utils.Point{
			X: start.X + size,
			Y: start.Y + size,
		}
		h.drawVicsekTreeCombInv(iteration-1, p, size)

		p = utils.Point{
			X: start.X + (2 * size),
			Y: start.Y,
		}
		h.drawVicsekTreeCombInv(iteration-1, p, size)

		p = utils.Point{
			X: start.X + (2 * size),
			Y: start.Y + size,
		}
		h.drawVicsekTreeCombInv(iteration-1, p, size)

		p = utils.Point{
			X: start.X + (2 * size),
			Y: start.Y + (2 * size),
		}
		h.drawVicsekTreeCombInv(iteration-1, p, size)
	}
}
