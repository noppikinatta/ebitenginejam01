package magnet

import (
	"image/color"
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

func NewBarMagnet(width, height float64, root, tip PoleType, loc Location, rootV, tipV Velocity) *BarMagnet {
	m := BarMagnet{
		rootPole: Pole{Type: root},
		tipPole:  Pole{Type: tip},
		width:    width,
		height:   height,
		loc:      loc,
		rootV:    rootV,
		tipV:     tipV,
	}
	m.updatePoleLocs()
	return &m
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
	decay := 0.99
	m.rootV.X *= decay
	m.rootV.Y *= decay
	m.tipV.X *= decay
	m.tipV.Y *= decay
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
	gm.Translate(m.loc.X-m.width/2, m.loc.Y-m.height/2)
	return gm
}

func (m *BarMagnet) RootVDebug() (x1, y1, x2, y2 float64, c color.Color) {
	x1 = m.rootPole.X
	y1 = m.rootPole.Y
	x2 = x1 + m.rootV.X
	y2 = y1 + m.rootV.Y
	c = color.Black
	if m.rootPole.Type == PoleTypeN {
		c = color.RGBA{200, 0, 0, 255}
	}
	if m.rootPole.Type == PoleTypeS {
		c = color.RGBA{0, 0, 200, 255}
	}
	return
}

func (m *BarMagnet) TipVDebug() (x1, y1, x2, y2 float64, c color.Color) {
	x1 = m.tipPole.X
	y1 = m.tipPole.Y
	x2 = x1 + m.tipV.X
	y2 = y1 + m.tipV.Y
	c = color.Black
	if m.tipPole.Type == PoleTypeN {
		c = color.RGBA{200, 0, 0, 255}
	}
	if m.tipPole.Type == PoleTypeS {
		c = color.RGBA{0, 0, 200, 255}
	}
	return
}

type MonopoleMagnet struct {
	pole     Pole
	loc      Location
	width    float64
	height   float64
	velocity Velocity
}

func NewMonopoleMagnet(width, height float64, pt PoleType, loc Location, v Velocity) *MonopoleMagnet {
	m := MonopoleMagnet{
		pole:     Pole{Type: pt},
		loc:      loc,
		width:    width,
		height:   height,
		velocity: v,
	}
	m.updatePoleLoc()
	return &m
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
	gm.Translate(m.loc.X-m.width/2, m.loc.Y-m.height/2)
	return gm
}
