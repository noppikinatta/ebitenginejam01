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
