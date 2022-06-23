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

package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/noppikinatta/ebitenginejam01/scene/gameplay"
	"github.com/noppikinatta/ebitenginejam01/scene/prologue"
	"github.com/noppikinatta/ebitenginejam01/scene/result"
	"github.com/noppikinatta/ebitenginejam01/scene/title"
)

type Scene interface {
	Update() error
	Draw(screen *ebiten.Image)
	End() bool
	Reset()
}

func AllScenes() *Container {
	t := title.NewScene()
	p := prologue.NewScene() // TODO: add constructors
	g := &gameplay.Scene{}
	r := &result.Scene{}

	c := NewContainer([]Scene{t, p, g, r})

	// TODO: return err
	_ = c.AddTransition(r, p)

	return c
}
