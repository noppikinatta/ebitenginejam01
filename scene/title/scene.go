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
	"github.com/noppikinatta/ebitenginejam01/animation"
	"github.com/noppikinatta/ebitenginejam01/input"
)

type Scene struct {
	state   state
	fadeIn  *animation.FadeIn
	fadeOut *animation.FadeOut
	image   *image
}

func NewScene() *Scene {
	s := Scene{
		fadeIn:  animation.NewFadeIn(15),
		fadeOut: animation.NewFadeOut(15),
		image:   &image{},
	}
	s.Reset()
	return &s
}

func (s *Scene) Update() error {
	s.updateState()

	switch s.state {
	case stateFadeIn:
		s.fadeIn.Update()
	case stateFadeOut:
		s.fadeOut.Update()
	}

	s.image.Update()

	return nil
}

func (s *Scene) updateState() {
	switch s.state {
	case stateFadeIn:
		if s.fadeIn.End() {
			s.state = stateWaitClick
		}
	case stateWaitClick:
		if input.LeftMousedownOrTouched() {
			s.state = stateFadeOut
		}
	}
}

func (s *Scene) Draw(screen *ebiten.Image) {
	s.image.Draw(screen)
	if s.state == stateFadeIn {
		s.fadeIn.Draw(screen)
	}
	if s.state == stateFadeOut {
		s.fadeOut.Draw(screen)
	}
}

func (s *Scene) End() bool {
	return s.fadeOut.End()
}

func (s *Scene) Reset() {
	s.state = stateFadeIn
	s.fadeIn.Reset()
	s.fadeOut.Reset()
	s.image.Reset()
}
