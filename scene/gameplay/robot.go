package gameplay

import (
	"github.com/hajimehoshi/ebiten/v2"
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
