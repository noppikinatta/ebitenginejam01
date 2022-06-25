package magnet

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Body struct {
	leftArmPole    Pole
	rightArmPole   Pole
	leftLegPole    Pole
	rightLegPole   Pole
	armPoleYOffset float64
	legPoleXOffset float64
	loc            Location
	width          float64
	height         float64
}

func NewBody(width, height, armPoleYOffset, legPoleXOffset float64) *Body {
	b := Body{
		width:          width,
		height:         height,
		armPoleYOffset: armPoleYOffset,
		legPoleXOffset: legPoleXOffset,
		leftArmPole:    Pole{Type: PoleTypeS},
		rightArmPole:   Pole{Type: PoleTypeN},
		leftLegPole:    Pole{Type: PoleTypeN},
		rightLegPole:   Pole{Type: PoleTypeS},
	}
	b.updatePoleLocs()
	return &b
}

func (b *Body) LeftArmPole() Pole {
	return b.leftArmPole
}

func (b *Body) RightArmPole() Pole {
	return b.rightArmPole
}

func (b *Body) LeftLegPole() Pole {
	return b.leftLegPole
}

func (b *Body) RightLegPole() Pole {
	return b.rightLegPole
}

func (b *Body) Poles() []Pole {
	return []Pole{
		b.leftArmPole,
		b.rightArmPole,
		b.leftLegPole,
		b.rightLegPole,
	}
}

func (b *Body) UpdateLoc(x, y float64) {
	b.loc.X = x
	b.loc.Y = y
	b.updatePoleLocs()
}

func (b *Body) updatePoleLocs() {
	x := b.loc.X
	y := b.loc.Y

	left := x - b.width/2
	top := y - b.height/2
	right := x + b.width/2
	bottom := y + b.height/2

	b.leftArmPole.X = left
	b.leftArmPole.Y = top + b.armPoleYOffset

	b.rightArmPole.X = right
	b.rightArmPole.Y = b.leftArmPole.Y

	b.leftLegPole.X = left + b.legPoleXOffset
	b.leftLegPole.Y = bottom

	b.rightLegPole.X = right - b.legPoleXOffset
	b.rightLegPole.Y = bottom
}

func (b *Body) GeoM() ebiten.GeoM {
	gm := ebiten.GeoM{}
	gm.Translate(b.loc.X-b.width/2, b.loc.Y-b.height/2)
	return gm
}
