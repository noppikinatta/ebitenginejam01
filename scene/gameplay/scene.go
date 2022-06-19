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

import "github.com/hajimehoshi/ebiten/v2"

type Scene struct {
	canClick bool
}

func (s *Scene) Update() error {
	return nil // TODO: implement
}

func (s *Scene) Draw(screen *ebiten.Image) {

}

func (s *Scene) End() bool {
	return false // TODO: implement
}

func (s *Scene) Reset() {

}

// 1. Title Scene Animation
// 2. Can click After N frames
// 3. FadeOut Screen
// 4. Next can return next scene after faded out
