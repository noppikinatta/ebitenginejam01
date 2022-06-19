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
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/noppikinatta/ebitenginejam01/animation"
	"github.com/noppikinatta/ebitenginejam01/input"
)

type Scene struct {
	state   state
	fadeIn  *animation.FadeIn
	fadeOut *animation.FadeOut
	msg1    *messageScreen
	msg2    *messageScreen
}

// TODO: constructor

func (s *Scene) Update() error {
	s.updateState()

	switch s.state {
	case stateMsg1FadeIn, stateMsg2FadeIn:
		s.fadeIn.Update()
	case stateMsg1FadeOut, stateMsg2FadeOut:
		s.fadeOut.Update()
	}

	return nil
}

func (s *Scene) updateState() {
	switch s.state {
	case stateMsg1FadeIn:
		if s.fadeIn.End() {
			s.fadeIn.Reset()
			s.state = stateMsg1WaitClick
		}
	case stateMsg1WaitClick:
		if input.LeftMousedownOrTouched() {
			s.state = stateMsg1FadeOut
		}
	case stateMsg1FadeOut:
		if s.fadeOut.End() {
			s.fadeOut.Reset()
			s.state = stateMsg2FadeIn
		}
	case stateMsg2FadeIn:
		if s.fadeIn.End() {
			s.state = stateMsg2WaitClick
		}
	case stateMsg2WaitClick:
		if input.LeftMousedownOrTouched() {
			s.state = stateMsg2FadeOut
		}
	}
}

func (s *Scene) Draw(screen *ebiten.Image) {
	s.currentScreen().Draw(screen)
	switch s.state {
	case stateMsg1FadeIn, stateMsg2FadeIn:
		s.fadeIn.Draw(screen)
	case stateMsg1FadeOut, stateMsg2FadeOut:
		s.fadeOut.Draw(screen)
	}
}

func (s *Scene) currentScreen() *messageScreen {
	switch s.state {
	case stateMsg1FadeIn, stateMsg1WaitClick, stateMsg1FadeOut:
		return s.msg1
	}
	return s.msg2
}

func (s *Scene) End() bool {
	if s.state != stateMsg2FadeOut {
		return false
	}
	return s.fadeOut.End()
}

func (s *Scene) Reset() {
	s.state = stateMsg1FadeIn
	s.fadeIn.Reset()
	s.fadeOut.Reset()
	s.msg1.Reset()
	s.msg2.Reset()
}
