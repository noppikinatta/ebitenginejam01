package prologue

import (
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/noppikinatta/ebitenginejam01/asset"
	"github.com/noppikinatta/ebitenginejam01/magnet"
)

var rndForLaunching *rand.Rand

func init() {
	rndForLaunching = rand.New(rand.NewSource(time.Now().UnixNano()))
}

type launching struct {
	parts       []*ebiten.Image
	locs        []magnet.Location
	vecs        []magnet.Velocity
	bg          *ebiten.Image
	lab         *ebiten.Image
	soundPlayed bool
}

func newLaunching() *launching {
	l := launching{
		parts: []*ebiten.Image{
			asset.ImgRobotPart(asset.RobotPartLeftArm),
			asset.ImgRobotPart(asset.RobotPartRightArm),
			asset.ImgRobotPart(asset.RobotPartLeftLeg),
			asset.ImgRobotPart(asset.RobotPartRightLeg),
			asset.ImgRobotPart(asset.RobotPartEbitenN),
			asset.ImgRobotPart(asset.RobotPartEbitenS),
			//asset.ImgRobotPart(asset.RobotPartTNT),
		},
		bg:  asset.ImgTitleBg.MustImage(),
		lab: asset.ImgPrologueBg.MustImage(),
	}
	l.Reset()
	return &l
}

func (l *launching) Update() {
	if !l.soundPlayed {
		asset.PlaySound(asset.SEFly)
		l.soundPlayed = true
	}
	for i := range l.locs {
		l.locs[i] = l.locs[i].Move(l.vecs[i])
	}
}

func (l *launching) Draw(screen *ebiten.Image) {
	screen.DrawImage(l.bg, nil)
	for i, p := range l.parts {
		l := l.locs[i]

		gm := ebiten.GeoM{}
		gm.Rotate(math.Pi / 2)
		gm.Translate(l.X, l.Y)

		opt := ebiten.DrawImageOptions{
			GeoM: gm,
		}

		screen.DrawImage(p, &opt)
	}
	screen.DrawImage(l.lab, nil)
}

func (l *launching) End() bool {
	for i := range l.locs {
		if l.locs[i].Y > 0 {
			return false
		}
	}
	return true
}

func (l *launching) Reset() {
	l.locs = make([]magnet.Location, len(l.parts))
	for i := range l.locs {
		l.locs[i] = magnet.Location{
			X: 140 + rndForLaunching.Float64()*20,
			Y: 500,
		}
	}
	l.vecs = make([]magnet.Velocity, len(l.parts))
	for i := range l.vecs {
		l.vecs[i] = magnet.Velocity{
			X: -0.5 + rndForLaunching.Float64(),
			Y: -5 - 5*rndForLaunching.Float64(),
		}
	}
	l.soundPlayed = false
}
