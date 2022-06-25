package gameplay

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/noppikinatta/ebitenginejam01/asset"
	"github.com/noppikinatta/ebitenginejam01/input"
	"github.com/noppikinatta/ebitenginejam01/magnet"
)

type body struct {
	image  *ebiten.Image
	magnet *magnet.Body
	opt    *ebiten.DrawImageOptions
	parts  []*combinedPart
}

func newBody() *body {
	img := asset.ImgRobotPart(asset.RobotPartBody)
	w, h := img.Size()

	b := body{
		image:  img,
		magnet: magnet.NewBody(float64(w), float64(h), 24, 16),
		opt:    &ebiten.DrawImageOptions{},
	}

	return &b
}

func (b *body) Poles() []magnet.Pole {
	return b.magnet.Poles()
}

func (b *body) Update() {
	x, y := input.CursorPosition()

	b.magnet.UpdateLoc(float64(x), float64(y))
}

func (b *body) Draw(screen *ebiten.Image) {
	gm := b.magnet.GeoM()
	b.opt.GeoM = gm
	screen.DrawImage(b.image, b.opt)

	for _, p := range b.parts {
		p.Draw(screen, gm)
	}
}

type combinedPart struct {
	image *ebiten.Image
	gm    ebiten.GeoM
	opt   *ebiten.DrawImageOptions
}

func (p *combinedPart) Draw(screen *ebiten.Image, bodyGM ebiten.GeoM) {
	gm := p.gm
	gm.Concat(bodyGM)

	p.opt.GeoM = gm

	screen.DrawImage(p.image, p.opt)
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
