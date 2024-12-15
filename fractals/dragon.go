package fractals

import (
	"opengl/utils"
)

func (h Handler) drawDragonCurveRec(iteration int, start, end utils.Point) {
	if iteration == 0 {
		utils.DrawLine(start, end, [3]float32{1, 0, 0}, h.ShaderProgram)
	} else {
		newPoint := utils.Point{
			X: (start.X + end.X + end.Y - start.Y) / 2,
			Y: (start.Y + end.Y - end.X + start.X) / 2,
		}
		h.drawDragonCurveRec(iteration-1, start, newPoint)
		h.drawDragonCurveRecRev(iteration-1, newPoint, end)
	}
}

func (h Handler) drawDragonCurveRecRev(iteration int, start, end utils.Point) {
	if iteration == 0 {
		utils.DrawLine(start, end, [3]float32{0, 0, 1}, h.ShaderProgram)
	} else {
		newPoint := utils.Point{
			X: (start.X + end.X - end.Y + start.Y) / 2,
			Y: (start.Y + end.Y + end.X - start.X) / 2,
		}
		h.drawDragonCurveRec(iteration-1, start, newPoint)
		h.drawDragonCurveRecRev(iteration-1, newPoint, end)
	}
}

func (h Handler) DrawDragonCurve(iterations int) {
	start := utils.Point{X: -0.5, Y: 0.0}
	end := utils.Point{X: 0.5, Y: 0.0}

	h.drawDragonCurveRec(iterations, start, end)
}
