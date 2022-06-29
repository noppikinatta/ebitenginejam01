// Copyright 2022 noppikinatta
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
