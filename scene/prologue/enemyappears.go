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
	"github.com/noppikinatta/ebitenginejam01/asset"
)

type enemyAppears struct {
	bg           *ebiten.Image
	enemy        *ebiten.Image
	optEnemy     *ebiten.DrawImageOptions
	enemyY       float64
	velocityY    float64
	jumpPower    float64
	gravityPower float64
}

func newEnemyAppears() *enemyAppears {
	ea := enemyAppears{
		bg:           asset.ImgResultBg.MustImage(),
		enemy:        asset.ImgResultEnemy.MustImage(),
		optEnemy:     &ebiten.DrawImageOptions{},
		jumpPower:    5,
		gravityPower: 1,
	}
	return &ea
}

func (ea *enemyAppears) Update() {
	ea.jump()
	gm := ebiten.GeoM{}
	gm.Translate(0, ea.enemyY+180)
	ea.optEnemy.GeoM = gm
}

func (ea *enemyAppears) jump() {
	if ea.enemyY < 0 {
		ea.velocityY += ea.gravityPower
	} else {
		ea.velocityY = -ea.jumpPower
	}

	ea.enemyY += ea.velocityY
	if ea.enemyY > 0 {
		ea.enemyY = 0
	}
}

func (ea *enemyAppears) Draw(screen *ebiten.Image) {
	screen.DrawImage(ea.bg, nil)
	screen.DrawImage(ea.enemy, ea.optEnemy)
}
