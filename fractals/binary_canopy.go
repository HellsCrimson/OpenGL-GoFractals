package fractals

import (
	"math"
	"opengl/utils"
)

func (h Handler) DrawBinaryCanopy(iteration int, initialAngle float64) {
	start := utils.Point{X: 0.0, Y: -0.8}
	end := utils.Point{X: 0.0, Y: -0.2}

	h.drawBinaryCanopyRec(iteration, start, end, initialAngle, 0.5, [3]float32{0, 0, 1})
}

func (h Handler) drawBinaryCanopyRec(iteration int, start, end utils.Point, angle, length float64, color [3]float32) {
	if iteration == 0 {
		utils.DrawLine(start, end, color, h.ShaderProgram)
	} else {

		utils.DrawLine(start, end, color, h.ShaderProgram)

		mid := end
		angleOffset := 30.0
		length *= 0.7

		leftAngle := angle + angleOffset
		leftEnd := utils.Point{
			X: mid.X + float32(length*math.Cos(leftAngle*math.Pi/180)),
			Y: mid.Y + float32(length*math.Sin(leftAngle*math.Pi/180)),
		}
		h.drawBinaryCanopyRec(iteration-1, mid, leftEnd, leftAngle, length, [3]float32{1, 0, 0})

		rightAngle := angle - angleOffset
		rightEnd := utils.Point{
			X: mid.X + float32(length*math.Cos(rightAngle*math.Pi/180)),
			Y: mid.Y + float32(length*math.Sin(rightAngle*math.Pi/180)),
		}
		h.drawBinaryCanopyRec(iteration-1, mid, rightEnd, rightAngle, length, [3]float32{0, 1, 0})
	}
}
