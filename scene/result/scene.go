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
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/noppikinatta/ebitenginejam01/animation"
	"github.com/noppikinatta/ebitenginejam01/asset"
	"github.com/noppikinatta/ebitenginejam01/combine"
	"github.com/noppikinatta/ebitenginejam01/input"
)

const (
	armPoleYOffset float64 = 24 // FIXME: duplicate from gameplay scene
)

var rndForResult *rand.Rand

func init() {
	rndForResult = rand.New(rand.NewSource(time.Now().UnixNano()))
}

type Scene struct {
	state       state
	fadeIn      *animation.FadeIn
	fadeOut     *animation.FadeOut
	bg          *ebiten.Image
	result      *combine.CombinedResult
	randing     *randing
	attack      *attack
	enemy       *enemy
	explRobot   *explosion
	explEnemy   *explosion
	showR       *showResult
	bgm2Stopped bool
}

func NewScene(result *combine.CombinedResult) *Scene {
	s := Scene{
		fadeIn:    animation.NewFadeIn(15),
		fadeOut:   animation.NewFadeOut(15),
		bg:        asset.ImgResultBg.MustImage(),
		result:    result,
		randing:   newRanding(result),
		attack:    newAttack(result),
		enemy:     newEnemy(),
		explRobot: newExplosion(160, 100),
		explEnemy: newExplosion(-64, 120),
		showR:     newShowResult(),
	}
	s.Reset()
	return &s
}

func (s *Scene) Update() error {
	if !s.bgm2Stopped {
		asset.StopSound(asset.BGM2)
		s.bgm2Stopped = true
	}

	s.updateState()

	if s.state == stateFadeIn {
		s.fadeIn.Update()
	}
	if s.state == stateFadeOut {
		s.fadeOut.Update()
	}

	if s.state == stateRanding {
		s.randing.Update()
	}

	if s.state == stateRobotAttack {
		s.attack.Update()
		s.attack.MayDealDamage(s.enemy)
	}

	if s.state == stateRobotExplode {
		s.explRobot.Update()
	}

	if s.state == stateEnemyExplode {
		s.explEnemy.Update()
	}

	if s.state == stateShowResult {
		s.showR.Update()
	}

	s.enemy.Update()

	return nil
}

func (s *Scene) updateState() {
	switch s.state {
	case stateFadeIn:
		if s.fadeIn.End() {
			s.fadeIn.Reset()
			s.state = stateRanding
		}
	case stateRanding:
		if s.randing.End() {
			if s.result.Randing() == combine.RandingTypeSuccess {
				s.state = stateRobotAttack
				s.attack.SetLoc(s.randing.Loc())
			} else {
				s.state = stateRobotExplode
			}
		}
	case stateRobotAttack:
		if s.attack.End() {
			if s.result.Attack() == combine.AttackTypeSuccess {
				s.state = stateEnemyExplode
			} else {
				s.state = stateEnemyAttack
				s.enemy.Attack()
			}
		}
	case stateEnemyAttack:
		if s.enemy.AttackEnd() {
			s.state = stateRobotExplode
		}
	case stateEnemyExplode:
		if s.explEnemy.End() || input.LeftMousedownOrTouched() {
			img := ebiten.NewImage(s.bg.Size())
			s.Draw(img)
			s.showR.SetSS(img)
			s.showR.Victory()
			s.state = stateShowResult
			asset.PlaySound(asset.BGM3)
		}
	case stateRobotExplode:
		if s.explRobot.End() || input.LeftMousedownOrTouched() {
			img := ebiten.NewImage(s.bg.Size())
			s.Draw(img)
			s.showR.SetSS(img)
			s.showR.Failed()
			s.state = stateShowResult
			asset.PlaySound(asset.BGM3)
		}
	case stateShowResult:
		if input.LeftMousedownOrTouched() {
			s.state = stateFadeOut
		}
	}
}

func (s *Scene) Draw(screen *ebiten.Image) {
	screen.DrawImage(s.bg, nil)

	if s.state == stateRanding {
		s.randing.Draw(screen)
	}

	if s.state.DrawAttack() {
		s.attack.Draw(screen)
	}

	if s.state.DrawEnemy() {
		s.enemy.Draw(screen)
	}

	if s.state == stateEnemyExplode {
		s.explEnemy.Draw(screen)
	}

	if s.state == stateRobotExplode {
		s.explRobot.Draw(screen)
	}

	if s.state == stateShowResult || s.state == stateFadeOut {
		s.showR.Draw(screen)
	}

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
	s.randing.Reset()
	s.attack.Reset()
	s.enemy.Reset()
	s.explRobot.Reset()
	s.explEnemy.Reset()
	s.showR.Reset()
	s.bgm2Stopped = false
}
