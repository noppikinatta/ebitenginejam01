package game

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

func (b *Body) UpdateLoc(x, y float64) {
	b.loc.X = x
	b.loc.Y = y

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
