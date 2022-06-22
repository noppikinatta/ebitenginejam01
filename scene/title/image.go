package title

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type image struct {
	img *ebiten.Image
}

func (img *image) Update() {
	// TODO: titile image animation
}

func (img *image) Draw(screen *ebiten.Image) {
	screen.DrawImage(img.img, nil)
}

func (img *image) Reset() {
	// TODO: reset animation progress
}
