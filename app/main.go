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

package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/noppikinatta/ebitenginejam01"
	"github.com/noppikinatta/ebitenginejam01/scene"
)

func main() {
	ebiten.SetWindowSize(640, 640)
	ebiten.SetWindowTitle("Ebiten Game Jam")

	s := scene.AllScenes()
	g := ebitenginejam01.NewGame(s)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
