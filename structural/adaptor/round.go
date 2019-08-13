package adaptor

type RoundPeg struct {
	radius float32
}

func (peg *RoundPeg) getRadius() float32 {
	return peg.radius
}

type RoundHole struct {
	radius float32
}

func (hole *RoundHole) getRadius() float32 {
	return hole.radius
}

func (hole *RoundHole) isFit(peg RoundPegIntf) bool {
	return hole.getRadius() >= peg.getRadius()
}
