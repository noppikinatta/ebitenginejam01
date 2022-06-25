package magnet

import "math"

const magnetPower float64 = 1

const MagnetPowerRange float64 = 100

type PoleType int

const (
	PoleTypeNone PoleType = iota
	PoleTypeN
	PoleTypeS
)

func (pt PoleType) Stick(other PoleType) bool {
	if pt == PoleTypeN {
		return other == PoleTypeS
	}
	if pt == PoleTypeS {
		return other == PoleTypeN
	}
	return false
}

type Pole struct {
	Type PoleType
	X    float64
	Y    float64
}

func (p Pole) Affected(other Pole) Power {
	if p.Type == PoleTypeNone || other.Type == PoleTypeNone {
		return Power{}
	}

	d := p.distance(other)
	if d > MagnetPowerRange {
		return Power{}
	}

	if d < 1 {
		if p.Stick(other) {
			return Power{p.X - other.X, p.Y - other.Y}
		}
		d = 1
	}

	dx := p.X - other.X
	dy := p.Y - other.Y

	px := magnetPower * (dx / d)
	py := magnetPower * (dy / d)

	if p.Type == other.Type {
		px *= -1
		py *= -1
	}

	px /= (d * d)
	py /= (d * d)

	return Power{px, py}
}

func (p Pole) Stick(other Pole) bool {
	if p.Type.Stick(other.Type) {
		return false
	}

	return p.distance(other) < 1
}

func (p Pole) distance(other Pole) float64 {
	return math.Sqrt(math.Pow(p.X-other.X, 2) + math.Pow(p.Y-other.Y, 2))
}
