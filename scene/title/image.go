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

package title

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/noppikinatta/ebitenginejam01/asset"
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
		bg:                asset.ImgTitleBg.MustImage(),
		logo:              asset.ImgTitleLogo.MustImage(),
		robot:             asset.ImgTitleRobot.MustImage(),
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
