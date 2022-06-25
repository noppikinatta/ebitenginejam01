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

package gameplay

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/noppikinatta/ebitenginejam01/animation"
	"github.com/noppikinatta/ebitenginejam01/asset"
	"github.com/noppikinatta/ebitenginejam01/combine"
)

type Scene struct {
	state   state
	fadeIn  *animation.FadeIn
	fadeOut *animation.FadeOut
	bg      *ebiten.Image
	body    *body
	leftarm *leftArm
	result  *combine.CombinedResult
}

func NewScene() *Scene {
	s := Scene{
		fadeIn:  animation.NewFadeIn(15),
		fadeOut: animation.NewFadeOut(15),
		bg:      asset.ImgGameplayBg.MustImage(),
		body:    newBody(),
		leftarm: newLeftArm(),
	}
	s.Reset()
	return &s
}

func (s *Scene) Update() error {
	s.updateState()

	if s.state == stateFadeIn {
		s.fadeIn.Update()
	}
	if s.state == stateFadeOut {
		s.fadeOut.Update()
	}

	s.body.Update()

	poles := s.body.Poles()

	s.leftarm.Update(poles)
	return nil // TODO: implement
}

func (s *Scene) updateState() {
	switch s.state {
	case stateFadeIn:
		if s.fadeIn.End() {
			s.fadeIn.Reset()
			s.state = stateCombine
		}
	}
}

func (s *Scene) Draw(screen *ebiten.Image) {
	screen.DrawImage(s.bg, nil)
	s.body.Draw(screen)
	s.leftarm.Draw(screen)

	if s.state == stateFadeIn {
		s.fadeIn.Draw(screen)
	}
	if s.state == stateFadeOut {
		s.fadeOut.Draw(screen)
	}
}

func (s *Scene) End() bool {
	return false // TODO: implement
}

func (s *Scene) Reset() {
	s.state = stateFadeIn
	s.fadeIn.Reset()
	s.fadeOut.Reset()

}
