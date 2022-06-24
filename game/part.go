package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Part interface {
	Update([]Pole)
	GeoM() ebiten.GeoM
}

type BarPart struct {
	rootPole Pole
	tipPole  Pole
	loc      Location
	angle    float64
	width    float64
	height   float64
	rootV    Velocity
	tipV     Velocity
}

func (p *BarPart) Update(poles []Pole) {
	for _, pole := range poles {
		p.updateVelocity(pole)
	}
	p.loc = p.loc.Move(p.velocity())
	p.angle += p.angularVelocity()
	p.updatePoleLocs()
}

func (p *BarPart) updateVelocity(pole Pole) {
	ra := p.rootPole.Affected(pole)
	ta := p.tipPole.Affected(pole)
	p.rootV = p.rootV.Accelerate(ra)
	p.tipV = p.tipV.Accelerate(ta)
}

func (p *BarPart) velocity() Velocity {
	return p.rootV.Avarage(p.tipV)
}

func (p *BarPart) angularVelocity() float64 {
	v := p.velocity()

	sin := math.Sin(p.angle)
	cos := math.Cos(p.angle)

	avRoot := -sin*(p.rootV.X-v.X) + -cos*(p.rootV.Y-v.Y)
	avTip := sin*(p.tipV.X-v.X) + cos*(p.tipV.Y-v.Y)

	r := p.width / 2
	return (avRoot + avTip) / r
}

func (p *BarPart) updatePoleLocs() {
	r := p.width / 2

	x := r * math.Cos(p.angle)
	y := r * math.Sin(p.angle)

	p.rootPole.X = p.loc.X - x
	p.rootPole.Y = p.loc.Y - y

	p.tipPole.X = p.loc.X + x
	p.tipPole.Y = p.loc.Y + y
}

func (p *BarPart) GeoM() ebiten.GeoM {
	gm := ebiten.GeoM{}
	gm.Translate(-p.width/2, -p.height/2)
	gm.Rotate(p.angle)
	gm.Translate(p.width/2, p.height/2)
	gm.Translate(p.loc.X, p.loc.Y)
	return gm
}

type MonopolePart struct {
	pole     Pole
	loc      Location
	width    float64
	height   float64
	velocity Velocity
}
