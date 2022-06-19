package prologue

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type messageScreen struct {
	bg           *ebiten.Image
	chara        *ebiten.Image
	msgJa        []*ebiten.Image
	msgEn        []*ebiten.Image
	msgIdx       int
	img          *ebiten.Image
	shouldUpdate bool
}

func newMessageScreen(bg, chara *ebiten.Image, msgJa, msgEn []*ebiten.Image) *messageScreen {
	s := messageScreen{
		bg:           bg,
		chara:        chara,
		msgJa:        msgJa,
		msgEn:        msgEn,
		msgIdx:       0,
		img:          ebiten.NewImageFromImage(bg),
		shouldUpdate: true,
	}

	return &s
}

func (s *messageScreen) Next() bool {
	newIdx := s.msgIdx + 1
	if newIdx >= s.length() {
		return false
	}
	s.msgIdx = newIdx
	s.shouldUpdate = true
	return true
}

func (s *messageScreen) length() int {
	lenJa := len(s.msgJa)
	lenEn := len(s.msgEn)
	if lenJa > lenEn {
		return lenJa
	}
	return lenEn
}

func (s *messageScreen) Draw(screen *ebiten.Image) {
	if s.shouldUpdate {
		s.update()
	}
	screen.DrawImage(s.img, nil)
}

func (s *messageScreen) update() {
	s.img.DrawImage(s.bg, nil)

	msgJa, msgEn := s.messages()

	s.img.DrawImage(msgJa, nil)
	s.img.DrawImage(s.bgForEn())
	s.img.DrawImage(msgEn, nil)
}

func (s *messageScreen) messages() (jp, en *ebiten.Image) {
	iJa := s.msgIdx
	iEn := s.msgIdx

	if iJa >= len(s.msgJa) {
		iJa = len(s.msgJa) - 1
	}
	if iEn >= len(s.msgEn) {
		iEn = len(s.msgEn) - 1
	}

	return s.msgJa[iJa], s.msgEn[iEn]
}

func (s *messageScreen) bgForEn() (*ebiten.Image, *ebiten.DrawImageOptions) {
	// TODO: may cached

	bgW, bgH := s.bg.Size()
	img := ebiten.NewImage(bgW, bgH/4)

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
