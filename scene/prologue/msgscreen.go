// Copyright 2022 noppikinatta
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package prologue

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/noppikinatta/ebitenginejam01/asset"
)

type messageScreen struct {
	chara        *ebiten.Image
	optChara     *ebiten.DrawImageOptions
	msg          []*ebiten.Image
	optMsgShadow *ebiten.DrawImageOptions
	msgIdx       int
	img          *ebiten.Image
	shouldUpdate bool
}

func newMessageScreen() *messageScreen {
	s := messageScreen{
		chara:    asset.ImgPrologueDoctor.MustImage(),
		optChara: &ebiten.DrawImageOptions{},
		msg: []*ebiten.Image{
			asset.ImgPrologueMsg1.MustImage(),
			asset.ImgPrologueMsg2.MustImage(),
		},
		optMsgShadow: &ebiten.DrawImageOptions{},
		img:          ebiten.NewImage(asset.ImgResultBg.MustImage().Size()),
	}
	w, h := s.chara.Size()
	s.optChara.GeoM.Scale(0.75, 0.75)
	s.optChara.GeoM.Translate(float64(w)*0.25, float64(h)*0.25)
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
	s.img.DrawImage(s.chara, s.optChara)
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
