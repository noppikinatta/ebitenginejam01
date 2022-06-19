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
	fadein  *animation.Fadein
	fadeout *animation.Fadeout
	image   *image
}

func NewScene() *Scene {
	s := Scene{
		fadein:  animation.NewFadein(15),
		fadeout: animation.NewFadeout(15),
		image:   &image{},
	}
	s.Reset()
	return &s
}

func (s *Scene) Update() error {
	s.updateState()

	switch s.state {
	case stateFadein:
		s.fadein.Update()
	case stateFadeout:
		s.fadeout.Update()
	}

	s.image.Update()

	return nil
}

func (s *Scene) updateState() {
	switch s.state {
	case stateFadein:
		if s.fadein.End() {
			s.state = stateWaitClick
		}
	case stateWaitClick:
		if input.LeftMousedownOrTouched() {
			s.state = stateFadeout
		}
	}
}

func (s *Scene) Draw(screen *ebiten.Image) {
	s.image.Draw(screen)
	if s.state == stateFadein {
		s.fadein.Draw(screen)
	}
	if s.state == stateFadeout {
		s.fadeout.Draw(screen)
	}
}

func (s *Scene) End() bool {
	return s.fadeout.End()
}

func (s *Scene) Reset() {
	s.state = stateFadein
	s.fadein.Reset()
	s.fadeout.Reset()
	s.image.Reset()
}
