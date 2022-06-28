package result

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/noppikinatta/ebitenginejam01/asset"
)

type damager interface {
	Damage()
}

type enemy struct {
	enemy       *ebiten.Image
	optEnemy    *ebiten.DrawImageOptions
	flame       *ebiten.Image
	optFrame    *ebiten.DrawImageOptions
	attackCount int
	attackEnd   bool
	damage      float64
}

func newEnemy() *enemy {
	e := enemy{
		enemy:    asset.ImgResultEnemy.MustImage(),
		optEnemy: &ebiten.DrawImageOptions{},
		flame:    asset.ImgResultFlame.MustImage(),
		optFrame: &ebiten.DrawImageOptions{},
	}
	e.Reset()
	return &e
}

func (e *enemy) Update() {
	if e.attackCount > 0 {
		e.attackCount--
		if e.attackCount == 0 {
			e.attackEnd = true
		}
	}
	if e.damage > 0 {
		e.damage--
	}
}

func (e *enemy) Attack() {
	asset.PlaySound(asset.SEFire)
	e.attackCount = 150
	e.attackEnd = false
}

func (e *enemy) AttackEnd() bool {
	return e.attackEnd
}

func (e *enemy) Damage() {
	e.damage = 30
}

func (e *enemy) Draw(screen *ebiten.Image) {
	e.drawEnemy(screen)
	if e.attackCount > 0 {
		e.drawFlame(screen)
	}
}

func (e *enemy) drawEnemy(screen *ebiten.Image) {
	gm := ebiten.GeoM{}
	gm.Translate(-e.damage, 180)
	e.optEnemy.GeoM = gm
	screen.DrawImage(e.enemy, e.optEnemy)
}

func (e *enemy) drawFlame(screen *ebiten.Image) {
	r := func() float64 { return 2*rndForResult.Float64() - 1 }
	gm := ebiten.GeoM{}
	gm.Translate(2*r(), 180+2*r())
	cm := ebiten.ColorM{}
	cm.Scale(1, 1, 1, float64(e.attackCount)/180)
	e.optFrame.GeoM = gm
	e.optFrame.ColorM = cm
	screen.DrawImage(e.flame, e.optFrame)
}

func (e *enemy) Reset() {
	e.attackCount = 0
	e.attackEnd = false
	e.damage = 0
}
