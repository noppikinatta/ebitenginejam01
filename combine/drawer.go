package combine

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/noppikinatta/ebitenginejam01/asset"
	"github.com/noppikinatta/ebitenginejam01/part"
)

type Drawer struct {
	image      *ebiten.Image
	opt        *ebiten.DrawImageOptions
	parts      map[part.PartType]*DrawerPart
	armYOffset float64
	legXOffset float64
}

func NewDrawer(armYOffset, legXOffset float64) *Drawer {
	img := asset.ImgRobotPart(asset.RobotPartBody)

	d := Drawer{
		image:      img,
		opt:        &ebiten.DrawImageOptions{},
		parts:      make(map[part.PartType]*DrawerPart),
		armYOffset: armYOffset,
		legXOffset: legXOffset,
	}
	return &d
}

func (d *Drawer) Draw(screen *ebiten.Image, gm ebiten.GeoM) {
	d.opt.GeoM = gm

	screen.DrawImage(d.image, d.opt)

	for _, p := range d.parts {
		p.Draw(screen, gm)
	}
}

func (d *Drawer) SetPart(p part.PartType, image *ebiten.Image, inverse bool) {
	bodyW, bodyH := d.image.Size()

	gm := ebiten.GeoM{}
	w, h := image.Size()
	if inverse {
		gm.Translate(-float64(w)/2, -float64(h)/2)
		gm.Rotate(math.Pi)
		gm.Translate(float64(w)/2, float64(h)/2)
	}

	switch p {
	case part.PartTypeLeftArm:
		// lotate PI
		gm.Translate(-float64(w)/2, -float64(h)/2)
		gm.Rotate(math.Pi)
		gm.Translate(float64(w)/2, float64(h)/2)

		gm.Translate(-float64(w), d.armYOffset-float64(h)/2)
	case part.PartTypeRightArm:
		gm.Translate(float64(bodyW), d.armYOffset-float64(h)/2)
	case part.PartTypeLeftLeg:
		// lotate PI/2
		gm.Translate(-float64(w)/2, -float64(h)/2)
		gm.Rotate(math.Pi / 2)
		gm.Translate(float64(w)/2, float64(h)/2)

		gm.Translate(-float64(bodyW)/2+d.legXOffset, float64(bodyH)+float64(w-h)/2)
	case part.PartTypeRightLeg:
		// lotate PI/2
		gm.Translate(-float64(w)/2, -float64(h)/2)
		gm.Rotate(math.Pi / 2)
		gm.Translate(float64(w)/2, float64(h)/2)

		// I don't know why but adding 2 to X offset is good
		gm.Translate(+float64(bodyW)/2-d.legXOffset+2, float64(bodyH)+float64(w-h)/2)
	}

	d.parts[p] = NewDrawPart(image, gm)
}

type DrawerPart struct {
	image *ebiten.Image
	gm    ebiten.GeoM
	opt   *ebiten.DrawImageOptions
}

func NewDrawPart(image *ebiten.Image, gm ebiten.GeoM) *DrawerPart {
	p := DrawerPart{
		image: image,
		gm:    gm,
		opt:   &ebiten.DrawImageOptions{},
	}
	return &p
}

func (p *DrawerPart) Draw(screen *ebiten.Image, bodyGM ebiten.GeoM) {
	gm := p.gm
	gm.Concat(bodyGM)
	p.opt.GeoM = gm

	screen.DrawImage(p.image, p.opt)
}
