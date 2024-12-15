package fractals

import "opengl/utils"

func (h Handler) DrawKoch(iteration int) {
	start := utils.Point{X: -0.5, Y: -0.3}
	end := utils.Point{X: 0.5, Y: -0.3}

	mid := utils.Point{
		X: (end.X+start.X)/2 - (end.Y - start.Y),
		Y: (start.Y+end.Y)/2 + (end.X - start.X),
	}

	h.drawKochRec(iteration, end, start)
	h.drawKochRec(iteration, mid, end)
	h.drawKochRec(iteration, start, mid)
}

func (h Handler) drawKochRec(iteration int, start, end utils.Point) {
	if iteration == 0 {
		utils.DrawLine(start, end, [3]float32{1, 1, 1}, h.ShaderProgram)
	} else {
		p1 := utils.Point{
			X: start.X + (end.X-start.X)/3,
			Y: start.Y + (end.Y-start.Y)/3,
		}
		p2 := utils.Point{
			X: (start.X+end.X)/2 - (end.Y-start.Y)/3,
			Y: (start.Y+end.Y)/2 + (end.X-start.X)/3,
		}
		p3 := utils.Point{
			X: start.X + 2*(end.X-start.X)/3,
			Y: start.Y + 2*(end.Y-start.Y)/3,
		}

		h.drawKochRec(iteration-1, start, p1)
		h.drawKochRec(iteration-1, p1, p2)
		h.drawKochRec(iteration-1, p2, p3)
		h.drawKochRec(iteration-1, p3, end)
	}
}
