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

package result

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/noppikinatta/ebitenginejam01/asset"
)

type showResult struct {
	victory *ebiten.Image
	failed  *ebiten.Image
	replay  *ebiten.Image
	ss      *ebiten.Image
	msg     *ebiten.Image
	optMsg  *ebiten.DrawImageOptions
	count   int
	max     int
}

func newShowResult() *showResult {
	sr := showResult{
		victory: asset.ImgResultMsgvictory.MustImage(),
		failed:  asset.ImgResultMsgfailed.MustImage(),
		replay:  asset.ImgResultToreplay.MustImage(),
		optMsg:  &ebiten.DrawImageOptions{},
		max:     100,
	}
	return &sr
}

func (sr *showResult) SetSS(img *ebiten.Image) {
	shadow := ebiten.NewImage(img.Size())
	shadow.Fill(color.RGBA{A: 64})
	img.DrawImage(shadow, nil)
	sr.ss = img
}

func (sr *showResult) Victory() {
	sr.msg = sr.victory
}

func (sr *showResult) Failed() {
	sr.msg = sr.failed
}

func (sr *showResult) Update() {
	if sr.count >= sr.max {
		return
	}
	sr.count++
}

func (sr *showResult) Draw(screen *ebiten.Image) {
	if sr.ss == nil || sr.msg == nil {
		return
	}
	screen.DrawImage(sr.ss, nil)

	a := float64(sr.count) / float64(sr.max)

	cm := ebiten.ColorM{}
	cm.Scale(1, 1, 1, a)
	sr.optMsg.ColorM = cm
	screen.DrawImage(sr.msg, sr.optMsg)

	if a > 0.5 {
		screen.DrawImage(sr.replay, nil)
	}
}

func (sr *showResult) Reset() {
	sr.ss = nil
	sr.msg = nil
	sr.count = 0
}
