package title

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/noppikinatta/ebitenginejam01/assets"
)

type image struct {
	bg                *ebiten.Image
	logo              *ebiten.Image
	robot             *ebiten.Image
	optBG             *ebiten.DrawImageOptions
	optLogo           *ebiten.DrawImageOptions
	optRobot          *ebiten.DrawImageOptions
	logoWhiteInFrames int
	logoWhiteInCount  int
	robotMoveFrames   int
	robotMoveCount    int
}

func newImage() *image {
	img := image{
		bg:                assets.ImgTitleBg.MustImage(),
		logo:              assets.ImgTitleLogo.MustImage(),
		robot:             assets.ImgTitleRobot.MustImage(),
		optBG:             &ebiten.DrawImageOptions{},
		optLogo:           &ebiten.DrawImageOptions{},
		optRobot:          &ebiten.DrawImageOptions{},
		logoWhiteInFrames: 30,
		robotMoveFrames:   60,
	}
	img.Reset()
	return &img
}

func (img *image) Update() {
	if img.logoWhiteInCount < img.logoWhiteInFrames {
		img.logoWhiteInCount++
	}

	if img.robotMoveCount < img.robotMoveFrames {
		img.robotMoveCount++
	}

	img.optLogo.ColorM = img.logoColorM()
	img.optRobot.GeoM = img.robotGeoM()
}

func (img *image) logoColorM() ebiten.ColorM {
	cm := ebiten.ColorM{}
	t := float64(img.logoWhiteInFrames-img.logoWhiteInCount) / float64(img.logoWhiteInFrames)
	cm.Translate(t, t, t, 0)

	return cm
}

func (img *image) robotGeoM() ebiten.GeoM {
	gm := ebiten.GeoM{}
	rate := float64(img.robotMoveFrames-img.robotMoveCount) / float64(img.robotMoveFrames)
	dx := rate * float64(img.robot.Bounds().Dx())
	gm.Translate(dx, 0)

	return gm
}

func (img *image) Draw(screen *ebiten.Image) {
	screen.DrawImage(img.bg, img.optBG)
	screen.DrawImage(img.robot, img.optRobot)
	screen.DrawImage(img.logo, img.optLogo)
}

func (img *image) Reset() {
	img.logoWhiteInCount = 0
	img.robotMoveCount = 0
}
