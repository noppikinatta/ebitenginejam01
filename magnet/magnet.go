package magnet

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type BarMagnet struct {
	rootPole Pole
	tipPole  Pole
	loc      Location
	angle    float64
	width    float64
	height   float64
	rootV    Velocity
	tipV     Velocity
}

func (m *BarMagnet) Update(poles []Pole) {
	for _, pole := range poles {
		m.updateVelocity(pole)
	}
	m.loc = m.loc.Move(m.velocity())
	m.angle += m.angularVelocity()
	m.updatePoleLocs()
}

func (m *BarMagnet) updateVelocity(pole Pole) {
	ra := m.rootPole.Affected(pole)
	ta := m.tipPole.Affected(pole)
	m.rootV = m.rootV.Accelerate(ra)
	m.tipV = m.tipV.Accelerate(ta)
}

func (m *BarMagnet) velocity() Velocity {
	return m.rootV.Avarage(m.tipV)
}

func (m *BarMagnet) angularVelocity() float64 {
	v := m.velocity()

	sin := math.Sin(m.angle)
	cos := math.Cos(m.angle)

	avRoot := -sin*(m.rootV.X-v.X) + -cos*(m.rootV.Y-v.Y)
	avTip := sin*(m.tipV.X-v.X) + cos*(m.tipV.Y-v.Y)

	r := m.width / 2
	return (avRoot + avTip) / r
}

func (m *BarMagnet) updatePoleLocs() {
	r := m.width / 2

	x := r * math.Cos(m.angle)
	y := r * math.Sin(m.angle)

	m.rootPole.X = m.loc.X - x
	m.rootPole.Y = m.loc.Y - y

	m.tipPole.X = m.loc.X + x
	m.tipPole.Y = m.loc.Y + y
}

func (m *BarMagnet) StickRoot(pole Pole) bool {
	return m.rootPole.Stick(pole)
}

func (m *BarMagnet) StickTip(pole Pole) bool {
	return m.tipPole.Stick(pole)
}

func (m *BarMagnet) GeoM() ebiten.GeoM {
	gm := ebiten.GeoM{}
	gm.Translate(-m.width/2, -m.height/2)
	gm.Rotate(m.angle)
	gm.Translate(m.width/2, m.height/2)
	gm.Translate(m.loc.X, m.loc.Y)
	return gm
}

type MonopoleMagnet struct {
	pole     Pole
	loc      Location
	width    float64
	height   float64
	velocity Velocity
}

func (m *MonopoleMagnet) Update(poles []Pole) {
	for _, pole := range poles {
		m.updateVelocity(pole)
	}
	m.loc = m.loc.Move(m.velocity)
	m.updatePoleLoc()
}

func (m *MonopoleMagnet) updateVelocity(pole Pole) {
	a := m.pole.Affected(pole)
	m.velocity = m.velocity.Accelerate(a)
}

func (m *MonopoleMagnet) updatePoleLoc() {
	m.pole.X = m.loc.X
	m.pole.Y = m.loc.Y
}

func (m *MonopoleMagnet) Stick(pole Pole) bool {
	return m.pole.Stick(pole)
}

func (m *MonopoleMagnet) GeoM() ebiten.GeoM {
	gm := ebiten.GeoM{}
	gm.Translate(m.loc.X, m.loc.Y)
	return gm
}
