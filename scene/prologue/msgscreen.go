package prologue

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/noppikinatta/ebitenginejam01/asset"
)

type messageScreen struct {
	chara        *ebiten.Image
	msg          []*ebiten.Image
	optMsgShadow *ebiten.DrawImageOptions
	msgIdx       int
	img          *ebiten.Image
	shouldUpdate bool
}

func newMessageScreen() *messageScreen {
	s := messageScreen{
		chara: asset.ImgPrologueDoctor.MustImage(),
		msg: []*ebiten.Image{
			asset.ImgPrologueMsg1.MustImage(),
			asset.ImgPrologueMsg2.MustImage(),
		},
		optMsgShadow: &ebiten.DrawImageOptions{},
		img:          ebiten.NewImage(asset.ImgResultBg.MustImage().Size()),
	}
	s.optMsgShadow.GeoM.Translate(2, 2)
	s.optMsgShadow.ColorM.Scale(0, 0, 0, 1)

	s.Reset()

	return &s
}

func (s *messageScreen) Next() bool {
	newIdx := s.msgIdx + 1
	if newIdx >= len(s.msg) {
		return false
	}
	s.msgIdx = newIdx
	s.shouldUpdate = true
	return true
}

func (s *messageScreen) Draw(screen *ebiten.Image) {
	if s.shouldUpdate {
		s.update()
	}
	screen.DrawImage(s.img, nil)
}

func (s *messageScreen) update() {
	s.img.Clear()
	s.img.DrawImage(s.chara, nil)
	s.img.DrawImage(s.bgForMsg())
	s.img.DrawImage(s.msg[s.msgIdx], s.optMsgShadow)
	s.img.DrawImage(s.msg[s.msgIdx], nil)
}

func (s *messageScreen) bgForMsg() (*ebiten.Image, *ebiten.DrawImageOptions) {
	bgW, bgH := s.img.Size()
	img := ebiten.NewImage(bgW, bgH/3)

	img.Fill(color.RGBA{A: 64})

	geom := ebiten.GeoM{}
	geom.Translate(0, float64(bgH-img.Bounds().Dy()))

	opt := ebiten.DrawImageOptions{GeoM: geom}

	return img, &opt
}

func (s *messageScreen) Reset() {
	s.msgIdx = 0
	s.shouldUpdate = true
}
