package gameplay

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/noppikinatta/ebitenginejam01/asset"
	"github.com/noppikinatta/ebitenginejam01/combine"
	"github.com/noppikinatta/ebitenginejam01/input"
	"github.com/noppikinatta/ebitenginejam01/magnet"
	"github.com/noppikinatta/ebitenginejam01/part"
)

type body struct {
	Drawer *combine.Drawer
	magnet *magnet.Body
}

func newBody(drawer *combine.Drawer) *body {
	img := asset.ImgRobotPart(asset.RobotPartBody)
	w, h := img.Size()

	b := body{
		Drawer: drawer,
		magnet: magnet.NewBody(float64(w), float64(h), armPoleYOffset, legPoleXOffset),
	}

	return &b
}

func (b *body) Poles() map[part.PartType]magnet.Pole {
	return b.magnet.Poles()
}

func (b *body) Update() {
	x, y := input.CursorPosition()

	b.magnet.UpdateLoc(float64(x), float64(y))
}

func (b *body) Draw(screen *ebiten.Image) {
	gm := b.magnet.GeoM()
	b.Drawer.Draw(screen, gm)
}

func (b *body) Loc() (x, y float64) {
	return b.magnet.Loc()
}

type Result struct {
	CombinedType combine.CombinedType
	Image        *ebiten.Image
	Inverse      bool
}

var rndForParts *rand.Rand

func init() {
	rndForParts = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func randLoc() magnet.Location {
	x := 30 + rndForParts.Intn(260)
	return magnet.Location{
		X: float64(x),
		Y: 400,
	}
}

func randVelocity() magnet.Velocity {
	vy := -2 * rndForParts.Float64()
	return magnet.Velocity{
		X: 0,
		Y: float64(vy),
	}
}

type updater interface {
	Update(poles []magnet.Pole)
}

type combiner interface {
	Combine(pole magnet.Pole) (result Result, ok bool)
}

type drawer interface {
	Draw(screen *ebiten.Image)
}

type robotPart interface {
	updater
	combiner
	drawer
}

type limb struct {
	image  *ebiten.Image
	magnet *magnet.BarMagnet
	opt    *ebiten.DrawImageOptions
	arm    bool
}

func newLeftArm() robotPart {
	img := asset.ImgRobotPart(asset.RobotPartLeftArm)
	w, h := img.Size()

	l := limb{
		image: img,
		magnet: magnet.NewBarMagnet(
			float64(w),
			float64(h),
			magnet.PoleTypeN,
			magnet.PoleTypeS,
			randLoc(),
			randVelocity(),
			randVelocity(),
		),
		opt: &ebiten.DrawImageOptions{},
		arm: true,
	}

	return &l
}

func newRightArm() robotPart {
	img := asset.ImgRobotPart(asset.RobotPartRightArm)
	w, h := img.Size()

	l := limb{
		image: img,
		magnet: magnet.NewBarMagnet(
			float64(w),
			float64(h),
			magnet.PoleTypeS,
			magnet.PoleTypeN,
			randLoc(),
			randVelocity(),
			randVelocity(),
		),
		opt: &ebiten.DrawImageOptions{},
		arm: true,
	}

	return &l
}

func newLeftLeg() robotPart {
	img := asset.ImgRobotPart(asset.RobotPartLeftLeg)
	w, h := img.Size()

	l := limb{
		image: img,
		magnet: magnet.NewBarMagnet(
			float64(w),
			float64(h),
			magnet.PoleTypeS,
			magnet.PoleTypeN,
			randLoc(),
			randVelocity(),
			randVelocity(),
		),
		opt: &ebiten.DrawImageOptions{},
	}

	return &l
}

func newRightLeg() robotPart {
	img := asset.ImgRobotPart(asset.RobotPartRightLeg)
	w, h := img.Size()

	l := limb{
		image: img,
		magnet: magnet.NewBarMagnet(
			float64(w),
			float64(h),
			magnet.PoleTypeN,
			magnet.PoleTypeS,
			randLoc(),
			randVelocity(),
			randVelocity(),
		),
		opt: &ebiten.DrawImageOptions{},
	}

	return &l
}

func (l *limb) Update(poles []magnet.Pole) {
	l.magnet.Update(poles)
}

func (l *limb) Draw(screen *ebiten.Image) {
	gm := l.magnet.GeoM()
	l.opt.GeoM = gm
	screen.DrawImage(l.image, l.opt)
}

func (l *limb) Combine(pole magnet.Pole) (result Result, ok bool) {
	if l.magnet.StickRoot(pole) {
		r := Result{
			Image:   l.image,
			Inverse: false,
		}
		if l.arm {
			r.CombinedType = combine.CombinedTypeCorrectArm
		} else {
			r.CombinedType = combine.CombinedTypeCorrectLeg
		}
		return r, true
	}
	if l.magnet.StickTip(pole) {
		r := Result{
			Image:   l.image,
			Inverse: true,
		}
		if l.arm {
			r.CombinedType = combine.CombinedTypeInverseArm
		} else {
			r.CombinedType = combine.CombinedTypeInverseLeg
		}
		return r, true
	}

	return Result{}, false
}

type ebi struct {
	image        *ebiten.Image
	magnet       *magnet.BarMagnet
	opt          *ebiten.DrawImageOptions
	combinedType combine.CombinedType
}

func newEbitenN() robotPart {
	img := asset.ImgRobotPart(asset.RobotPartEbitenN)
	w, h := img.Size()

	l := ebi{
		image: img,
		magnet: magnet.NewBarMagnet(
			float64(w),
			float64(h),
			magnet.PoleTypeN,
			magnet.PoleTypeNone,
			randLoc(),
			randVelocity(),
			randVelocity(),
		),
		opt:          &ebiten.DrawImageOptions{},
		combinedType: combine.CombinedTypeEbitenN,
	}

	return &l
}

func newEbitenS() robotPart {
	img := asset.ImgRobotPart(asset.RobotPartEbitenS)
	w, h := img.Size()

	e := ebi{
		image: img,
		magnet: magnet.NewBarMagnet(
			float64(w),
			float64(h),
			magnet.PoleTypeS,
			magnet.PoleTypeNone,
			randLoc(),
			randVelocity(),
			randVelocity(),
		),
		opt:          &ebiten.DrawImageOptions{},
		combinedType: combine.CombinedTypeEbitenS,
	}

	return &e
}

func (e *ebi) Update(poles []magnet.Pole) {
	e.magnet.Update(poles)
}

func (e *ebi) Draw(screen *ebiten.Image) {
	gm := e.magnet.GeoM()
	e.opt.GeoM = gm
	screen.DrawImage(e.image, e.opt)
}

func (e *ebi) Combine(pole magnet.Pole) (result Result, ok bool) {
	if e.magnet.StickRoot(pole) {
		r := Result{
			CombinedType: e.combinedType,
			Image:        e.image,
			Inverse:      true, // I drew shrimp pictures backwards...
		}
		return r, true
	}

	return Result{}, false
}

type tnt struct {
	image  *ebiten.Image
	magnet *magnet.BarMagnet
	opt    *ebiten.DrawImageOptions
}

func newTNT() robotPart {
	img := asset.ImgRobotPart(asset.RobotPartTNT)
	w, h := img.Size()

	t := tnt{
		image: img,
		magnet: magnet.NewBarMagnet(
			float64(w),
			float64(h),
			magnet.PoleTypeN,
			magnet.PoleTypeNone,
			randLoc(),
			randVelocity(),
			randVelocity(),
		),
		opt: &ebiten.DrawImageOptions{},
	}

	return &t
}

func (t *tnt) Update(poles []magnet.Pole) {
	t.magnet.Update(poles)
}

func (t *tnt) Draw(screen *ebiten.Image) {
	gm := t.magnet.GeoM()
	t.opt.GeoM = gm
	screen.DrawImage(t.image, t.opt)
}

func (t *tnt) Combine(pole magnet.Pole) (result Result, ok bool) {
	if t.magnet.StickRoot(pole) {
		r := Result{
			CombinedType: combine.CombinedTypeTNT,
			Image:        t.image,
			Inverse:      false,
		}
		return r, true
	}

	return Result{}, false
}
