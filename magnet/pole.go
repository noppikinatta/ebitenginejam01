package magnet

import "math"

const magnetPower float64 = 5

const MagnetPowerRange float64 = 1000

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

	dx := other.X - p.X
	dy := other.Y - p.Y

	px := magnetPower * (dx / d)
	py := magnetPower * (dy / d)

	if p.Type == other.Type {
		px *= -0.5
		py *= -0.5
	}

	px /= d
	py /= d

	return Power{px, py}
}

func (p Pole) Stick(other Pole) bool {
	if !p.Type.Stick(other.Type) {
		return false
	}

	return p.distance(other) < 5
}

func (p Pole) distance(other Pole) float64 {
	return math.Sqrt(math.Pow(p.X-other.X, 2) + math.Pow(p.Y-other.Y, 2))
}
