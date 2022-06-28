package magnet

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/noppikinatta/ebitenginejam01/part"
)

type Body struct {
	parts          map[part.PartType]Pole
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
		parts: map[part.PartType]Pole{
			part.PartTypeLeftArm:  {Type: PoleTypeS},
			part.PartTypeRightArm: {Type: PoleTypeN},
			part.PartTypeLeftLeg:  {Type: PoleTypeN},
			part.PartTypeRightLeg: {Type: PoleTypeS},
		},
	}
	b.updatePoleLocs()
	return &b
}

func (b *Body) Loc() (x, y float64) {
	gm := b.GeoM()
	tx := gm.Element(0, 2)
	ty := gm.Element(1, 2)
	return tx, ty
}

func (b *Body) Pole(pt part.PartType) Pole {
	return b.parts[pt]
}

func (b *Body) Poles() map[part.PartType]Pole {
	return b.parts
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

	p := b.parts[part.PartTypeLeftArm]
	p.X = left
	p.Y = top + b.armPoleYOffset
	b.parts[part.PartTypeLeftArm] = p

	p = b.parts[part.PartTypeRightArm]
	p.X = right
	p.Y = top + b.armPoleYOffset
	b.parts[part.PartTypeRightArm] = p

	p = b.parts[part.PartTypeLeftLeg]
	p.X = left + b.legPoleXOffset
	p.Y = bottom
	b.parts[part.PartTypeLeftLeg] = p

	p = b.parts[part.PartTypeRightLeg]
	p.X = right - b.legPoleXOffset
	p.Y = bottom
	b.parts[part.PartTypeRightLeg] = p
}

func (b *Body) GeoM() ebiten.GeoM {
	gm := ebiten.GeoM{}
	gm.Translate(b.loc.X-b.width/2, b.loc.Y-b.height/2)
	return gm
}
