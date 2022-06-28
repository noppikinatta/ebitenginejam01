package result

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/noppikinatta/ebitenginejam01/asset"
)

type explosion struct {
	image *ebiten.Image
	opt   *ebiten.DrawImageOptions
	x, y  float64
	count int
	max   int
}

func newExplosion(x, y float64) *explosion {
	e := explosion{
		image: asset.ImgResultExplosion.MustImage(),
		opt:   &ebiten.DrawImageOptions{},
		x:     x,
		y:     y,
		max:   200,
	}
	return &e
}

func (e *explosion) Update() {
	if e.End() {
		return
	}
	e.count++
}

func (e *explosion) End() bool {
	return e.count >= e.max
}

func (e *explosion) Draw(screen *ebiten.Image) {
	a := float64(e.max-e.count) / float64(e.max)
	if a <= 0 {
		return
	}

	r := func() float64 { return 2*rndForResult.Float64() - 1 }
	gm := ebiten.GeoM{}
	gm.Scale(1.5, 1.5)
	gm.Translate(e.x+r()*2, e.y+r()*2)

	cm := ebiten.ColorM{}
	cm.Scale(1+a/4, 1+a/8, 1, a)

	e.opt.GeoM = gm
	e.opt.ColorM = cm

	screen.DrawImage(e.image, e.opt)
}

func (e *explosion) Reset() {
	e.count = 0
}
