package fractals

import (
	"opengl/utils"
)

func drawDragonCurveRec(iteration int, start, end utils.Point) {
	if iteration == 0 {
		utils.DrawLine(start, end)
	} else {
		newPoint := utils.Point{
			X: (start.X + end.X + end.Y - start.Y) / 2,
			Y: (start.Y + end.Y - end.X + start.X) / 2,
		}
		drawDragonCurveRec(iteration-1, start, newPoint)
		drawDragonCurveRecRev(iteration-1, newPoint, end)
	}
}

func drawDragonCurveRecRev(iteration int, start, end utils.Point) {
	if iteration == 0 {
		utils.DrawLine(start, end)
	} else {
		newPoint := utils.Point{
			X: (start.X + end.X - end.Y + start.Y) / 2,
			Y: (start.Y + end.Y + end.X - start.X) / 2,
		}
		drawDragonCurveRec(iteration-1, start, newPoint)
		drawDragonCurveRecRev(iteration-1, newPoint, end)
	}
}

func DrawDragonCurve(iterations int) {
	start := utils.Point{X: -0.5, Y: 0.0}
	end := utils.Point{X: 0.5, Y: 0.0}

	drawDragonCurveRec(iterations, start, end)
}
