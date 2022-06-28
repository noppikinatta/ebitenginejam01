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
	state         state
	fadeIn        *animation.FadeIn
	fadeOut       *animation.FadeOut
	ea            *enemyAppears
	msg           *messageScreen
	launch        *launching
	msgWaitFrames int
	msgWaitCount  int
}

func NewScene() *Scene {
	s := Scene{
		fadeIn:        animation.NewFadeIn(15),
		fadeOut:       animation.NewFadeOut(15),
		ea:            newEnemyAppears(),
		msg:           newMessageScreen(),
		launch:        newLaunching(),
		msgWaitFrames: 45,
	}
	s.Reset()
	return &s
}

func (s *Scene) Update() error {
	s.updateState()

	if s.state.FadingIn() {
		s.fadeIn.Update()
	}
	if s.state.FadingOut() {
		s.fadeOut.Update()
	}

	if s.state.Msg() {
		s.ea.Update()
		if s.msgWaitCount < s.msgWaitFrames {
			s.msgWaitCount++
		}
	}

	if s.state.Launching() {
		s.launch.Update()
	}

	return nil
}

func (s *Scene) updateState() {
	switch s.state {
	case stateMsgFadeIn:
		if s.fadeIn.End() {
			s.fadeIn.Reset()
			s.state = stateMsgWaitClick
		}
	case stateMsgWaitClick:
		if s.msgShown() && input.LeftMousedownOrTouched() {
			remaining := s.msg.Next()
			if !remaining {
				s.state = stateMsgFadeOut
			}
		}
	case stateMsgFadeOut:
		if s.fadeOut.End() {
			s.fadeOut.Reset()
			s.state = stateLaunchFadeIn
		}
	case stateLaunchFadeIn:
		if s.fadeIn.End() {
			s.state = stateLaunching
		}
	case stateLaunching:
		if s.launch.End() || input.LeftMousedownOrTouched() {
			s.state = stateLaunchFadeOut
		}
	}
}

func (s *Scene) msgShown() bool {
	return s.msgWaitCount >= s.msgWaitFrames
}

func (s *Scene) Draw(screen *ebiten.Image) {
	if s.state.Msg() {
		s.ea.Draw(screen)
		if s.msgShown() {
			s.msg.Draw(screen)
		}
	}

	if s.state.Launching() {
		s.launch.Draw(screen)
	}

	if s.state.FadingIn() {
		s.fadeIn.Draw(screen)
	}
	if s.state.FadingOut() {
		s.fadeOut.Draw(screen)
	}
}

func (s *Scene) End() bool {
	if s.state != stateLaunchFadeOut {
		return false
	}
	return s.fadeOut.End()
}

func (s *Scene) Reset() {
	s.state = stateMsgFadeIn
	s.fadeIn.Reset()
	s.fadeOut.Reset()
	s.msg.Reset()
	s.launch.Reset()
}
