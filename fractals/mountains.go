package fractals

import (
	"math"
	"opengl/utils"

	"math/rand/v2"
)

func (h Handler) DrawMountain(iteration int) {
	start := utils.Point{X: -0.9, Y: -0.8}
	end := utils.Point{X: 0.9, Y: -0.8}

	h.drawMountainRec(iteration, start, end)
}

func (h Handler) drawMountainRec(iteration int, start, end utils.Point) {
	if iteration == 0 {
		utils.DrawLine(start, end, [3]float32{1, 0, 0}, h.ShaderProgram)
	} else {
		newEnd := utils.Point{
			X: (start.X + end.X) / 2,
			Y: (start.Y+end.Y)/2 + rand.Float32()*(float32(math.Abs(float64(end.X-start.X)))/5+0.2),
		}
		h.drawMountainRec(iteration-1, start, newEnd)
		h.drawMountainRec(iteration-1, newEnd, end)
	}
}
