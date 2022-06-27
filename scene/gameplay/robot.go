package gameplay

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/noppikinatta/ebitenginejam01/asset"
	"github.com/noppikinatta/ebitenginejam01/combine"
	"github.com/noppikinatta/ebitenginejam01/input"
	"github.com/noppikinatta/ebitenginejam01/magnet"
	"github.com/noppikinatta/ebitenginejam01/part"
)

type body struct {
	Drawer *combine.Drawer
	magnet *magnet.Body
}

func newBody(drawer *combine.Drawer) *body {
	img := asset.ImgRobotPart(asset.RobotPartBody)
	w, h := img.Size()

	b := body{
		Drawer: drawer,
		magnet: magnet.NewBody(float64(w), float64(h), armPoleYOffset, legPoleXOffset),
	}

	return &b
}

func (b *body) Poles() map[part.PartType]magnet.Pole {
	return b.magnet.Poles()
}

func (b *body) Update() {
	x, y := input.CursorPosition()

	b.magnet.UpdateLoc(float64(x), float64(y))
}

func (b *body) Combine(cc []combiner) []Result {
	rr := make([]Result, 0)

	for _, c := range cc {
		for pt, p := range b.magnet.Poles() {
			r, ok := c.Combine(p)
			if ok {
				r.PartType = pt
				rr = append(rr, r)
				goto next
			}
		}
	next:
	}

	return rr
}

func (b *body) Draw(screen *ebiten.Image) {
	gm := b.magnet.GeoM()
	b.Drawer.Draw(screen, gm)
}

type Result struct {
	PartType     part.PartType
	CombinedType combine.CombinedType
	Image        *ebiten.Image
	Inverse      bool
}

type updater interface {
	Update(poles []magnet.Pole)
}

type combiner interface {
	Combine(pole magnet.Pole) (result Result, ok bool)
}

type updateCombiner interface {
	updater
	combiner
}

type leftArm struct {
	image  *ebiten.Image
	magnet *magnet.BarMagnet
	opt    *ebiten.DrawImageOptions
}

func newLeftArm() *leftArm {
	img := asset.ImgRobotPart(asset.RobotPartLeftArm)
	w, h := img.Size()

	la := leftArm{
		image:  img,
		magnet: magnet.NewBarMagnet(float64(w), float64(h), magnet.PoleTypeN, magnet.PoleTypeS),
		opt:    &ebiten.DrawImageOptions{},
	}

	return &la
}

func (la *leftArm) Update(poles []magnet.Pole) {
	la.magnet.Update(poles)
}

func (la *leftArm) Draw(screen *ebiten.Image) {
	gm := la.magnet.GeoM()
	la.opt.GeoM = gm
	screen.DrawImage(la.image, la.opt)

	x1, y1, x2, y2, c := la.magnet.RootVDebug()
	ebitenutil.DrawLine(screen, x1, y1, x2, y2, c)
	x1, y1, x2, y2, c = la.magnet.TipVDebug()
	ebitenutil.DrawLine(screen, x1, y1, x2, y2, c)
}

func (la *leftArm) Combine(pole magnet.Pole) (result Result, ok bool) {
	if la.magnet.StickRoot(pole) {
		r := Result{
			CombinedType: combine.CombinedTypeCorrectArm,
			Image:        la.image,
			Inverse:      false,
		}
		return r, true
	}
	if la.magnet.StickTip(pole) {
		r := Result{
			CombinedType: combine.CombinedTypeInverseArm,
			Image:        la.image,
			Inverse:      true,
		}
		return r, true
	}

	return Result{}, false
}
