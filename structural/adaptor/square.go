package adaptor

import (
	"math"
)

type RoundPegIntf interface {
	getRadius() float32
}

type SquarePeg struct {
	width float32
}

func (peg *SquarePeg) getWidth() float32 {
	return peg.width
}

func (peg *SquarePeg) getSquare() float32 {
	return float32(math.Pow(float64(peg.width), float64(2)))
}

type SquarePegAdaptor struct {
	peg SquarePeg
}

func (adaptor *SquarePegAdaptor) getRadius() float32 {
	return float32(math.Sqrt(math.Pow(float64(adaptor.peg.width/2), 2) * 2))
}
