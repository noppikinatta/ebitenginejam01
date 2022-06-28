package result

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/noppikinatta/ebitenginejam01/asset"
	"github.com/noppikinatta/ebitenginejam01/combine"
	"github.com/noppikinatta/ebitenginejam01/part"
)

type randing struct {
	result *combine.CombinedResult
	x, y   float64
	end    bool
}

func newRanding(cr *combine.CombinedResult) *randing {
	r := randing{
		result: cr,
	}
	r.Reset()
	return &r
}

func (r *randing) Update() {

	r.y += 3

	switch r.result.Randing() {
	case combine.RandingTypeSuccess:
		if r.y > 180 {
			r.y = 180
			r.end = true
		}
	case combine.RandingTypeFailure:
		if r.y > 180 {
			if r.result.LeftLeg() != combine.CombinedTypeCorrectLeg {
				t := 180 - r.y
				gm := ebiten.GeoM{}
				gm.Rotate(-t / 40)
				gm.Translate(t*2, t*3)
				r.result.Drawer.SetOptGeoM(part.PartTypeLeftLeg, gm)
			}
			if r.result.RightLeg() != combine.CombinedTypeCorrectLeg {
				t := 180 - r.y
				gm := ebiten.GeoM{}
				gm.Rotate(t / 40)
				gm.Translate(-t*2, t*3)
				r.result.Drawer.SetOptGeoM(part.PartTypeRightLeg, gm)
			}
		}
		if r.y > 240 {
			r.y = 240
			r.end = true
		}
	}
}

func (r *randing) Draw(screen *ebiten.Image) {
	gm := ebiten.GeoM{}
	gm.Translate(r.x, r.y)
	r.result.Drawer.Draw(screen, gm)
}

func (r *randing) End() bool {
	return r.end
}

func (r *randing) Loc() (x, y float64) {
	return r.x, r.y
}

func (r *randing) Reset() {
	r.x = 200
	r.y = -400
	r.end = false
}

type attack struct {
	result      *combine.CombinedResult
	x, y        float64
	counter     int
	waitCounter int
	end         bool
}

func newAttack(cr *combine.CombinedResult) *attack {
	a := attack{
		result: cr,
	}
	a.Reset()
	return &a
}

func (a *attack) SetLoc(x, y float64) {
	a.x = x
	a.y = y
}

func (a *attack) Update() {
	if a.waitCounter < 60 {
		a.waitCounter++
		return
	}

	if a.counter > 60 {
		a.end = true
		return
	}

	a.counter++

	switch a.result.LeftArm() {
	case combine.CombinedTypeCorrectArm:
		gm := ebiten.GeoM{}
		gm.Translate(-float64(a.counter)*6, 0)
		a.result.Drawer.SetOptGeoM(part.PartTypeLeftArm, gm)
	default:
		gm := ebiten.GeoM{}
		ty := float64(a.counter) * float64(a.counter) * 0.1
		if ty > 110 {
			ty = 110
		}
		gm.Translate(-float64(a.counter)/2, ty)
		a.result.Drawer.SetOptGeoM(part.PartTypeLeftArm, gm)
	}

	switch a.result.RightArm() {
	case combine.CombinedTypeCorrectArm:
		bodyW, _ := asset.ImgRobotPart(asset.RobotPartBody).Size()
		// depends on all parts have same size
		_, partH := asset.ImgRobotPart(asset.RobotPartRightArm).Size()

		gm := ebiten.GeoM{}
		gm.Translate(-float64(bodyW), float64(partH)/2-armPoleYOffset)
		gm.Rotate(math.Pi)
		// I don't know why but adding half of partH/2 works fine
		gm.Translate(float64(bodyW), armPoleYOffset+float64(partH)/2)
		gm.Translate(-float64(a.counter)*6, 0)
		a.result.Drawer.SetOptGeoM(part.PartTypeRightArm, gm)
	default:
		gm := ebiten.GeoM{}
		ty := float64(a.counter) * float64(a.counter) * 0.1
		if ty > 110 {
			ty = 110
		}
		gm.Translate(float64(a.counter)/2, ty)
		a.result.Drawer.SetOptGeoM(part.PartTypeRightArm, gm)
	}
}

func (a *attack) Draw(screen *ebiten.Image) {
	gm := ebiten.GeoM{}
	gm.Translate(a.x, a.y)
	a.result.Drawer.Draw(screen, gm)
}

func (a *attack) End() bool {
	return a.end
}

func (a *attack) Reset() {
	a.waitCounter = 0
	a.counter = 0
	a.end = false
}

func (a *attack) MayDealDamage(d damager) {
	switch a.result.LeftArm() {
	case combine.CombinedTypeCorrectArm:
		if a.counter == 60 {
			d.Damage()
		}
	}

	switch a.result.RightArm() {
	case combine.CombinedTypeCorrectArm:
		if a.counter == 70 {
			d.Damage()
		}
	}
}
