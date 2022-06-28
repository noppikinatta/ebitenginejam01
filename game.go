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

package ebitenginejam01

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/noppikinatta/ebitenginejam01/asset"
	"github.com/noppikinatta/ebitenginejam01/scene"
)

type Game struct {
	scenes      *scene.Container
	soundLoaded bool
}

func NewGame(scenes *scene.Container) *Game {
	g := Game{
		scenes: scenes,
	}
	return &g
}

func (g *Game) Update() error {
	if !g.soundLoaded {
		err := asset.LoadSounds()
		if err != nil {
			return err
		}
		g.soundLoaded = true
	}
	return g.scenes.Current().Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.scenes.Current().Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth / 2, outsideHeight / 2
}
