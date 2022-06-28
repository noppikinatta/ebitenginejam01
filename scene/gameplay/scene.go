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
	"github.com/noppikinatta/ebitenginejam01/magnet"
)

const (
	armPoleYOffset float64 = 24
	legPoleXOffset float64 = 16
)

type Scene struct {
	state    state
	fadeIn   *animation.FadeIn
	fadeOut  *animation.FadeOut
	bg       *ebiten.Image
	body     *body
	launcher *launcher
	result   *combine.CombinedResult
	comp     *complete
}

func NewScene(result *combine.CombinedResult) *Scene {
	bg := asset.ImgGameplayBg.MustImage()
	bgW, bgH := bg.Size()
	drawer := combine.NewDrawer(armPoleYOffset, legPoleXOffset)
	result.Drawer = drawer

	s := Scene{
		fadeIn:   animation.NewFadeIn(15),
		fadeOut:  animation.NewFadeOut(15),
		bg:       bg,
		body:     newBody(drawer),
		launcher: newLauncher(),
		result:   result,
		comp:     newComplete(drawer, float64(bgW), float64(bgH)),
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

	if s.state == stateCombine {
		s.body.Update()
		s.updateCombine()
	}

	if s.state == stateCombineComplete {
		s.comp.Update()
	}

	s.launcher.Update()
	s.updateParts()

	return nil
}

func (s *Scene) updateState() {
	switch s.state {
	case stateFadeIn:
		if s.fadeIn.End() {
			s.fadeIn.Reset()
			s.state = stateCombine
		}
	case stateCombine:
		if s.result.Complete() {
			s.state = stateCombineComplete
			s.comp.SetLoc(s.body.Loc())
		}
	case stateCombineComplete:
		if s.comp.End() {
			s.state = stateFadeOut
		}
	}
}

func (s *Scene) updateParts() {
	poles := make([]magnet.Pole, 0, len(s.body.Poles()))
	for k, v := range s.body.Poles() {
		if s.result.Combined(k) {
			continue
		}
		poles = append(poles, v)
	}

	for _, p := range s.launcher.Parts {
		p.Update(poles)
	}
}

func (s *Scene) updateCombine() {
	combinedParts := make([]robotPart, 0, len(s.launcher.Parts))
	for _, p := range s.launcher.Parts {
		for k, v := range s.body.Poles() {
			if s.result.Combined(k) {
				continue
			}

			r, ok := p.Combine(v)
			if !ok {
				continue
			}

			combinedParts = append(combinedParts, p)
			s.result.Set(k, r.CombinedType)
			s.result.Drawer.SetPart(k, r.Image, r.Inverse)
			break
		}
	}

	for _, p := range combinedParts {
		s.launcher.Remove(p)
	}
}

func (s *Scene) Draw(screen *ebiten.Image) {
	screen.DrawImage(s.bg, nil)

	if s.state == stateCombine {
		s.body.Draw(screen)
	}

	for _, p := range s.launcher.Parts {
		p.Draw(screen)
	}

	if s.state == stateCombineComplete {
		s.comp.Draw(screen)
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
	s.result.Reset()
	s.launcher.Reset()
	s.comp.Reset()
}
