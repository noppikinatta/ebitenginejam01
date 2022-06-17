package title

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type image struct {
	img *ebiten.Image
}

func (img *image) Update() {
	if img.img == nil {
		return
	}
	// TODO: titile image animation
	img.img.Fill(color.RGBA{G: 200, B: 200, A: 255})
	ebitenutil.DebugPrint(img.img, "title")
}

func (img *image) Draw(screen *ebiten.Image) {
	if img.img == nil {
		img.img = ebiten.NewImage(screen.Size())
		img.img.Fill(color.Black)
	}
	screen.DrawImage(img.img, nil)
}

func (img *image) Reset() {
	// TODO: reset animation progress
}
