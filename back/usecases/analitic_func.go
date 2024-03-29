package usecases

import (
	"github.com/BohdanCh-w/DSO-back/entities"
)

type AnaliticCalculator struct {
	From     float64
	To       float64
	PointNum int
	Func     func(float64) float64
}

func (calc *AnaliticCalculator) Calculate() []entities.WavePoint {
	step := (calc.To - calc.From) / float64(calc.PointNum)

	res := make([]entities.WavePoint, 0, calc.PointNum+1)

	for i := calc.From; i <= calc.To; i += step {
		point := entities.WavePoint{X: i, Y: calc.Func(i)}
		res = append(res, point)
	}

	return res
}
