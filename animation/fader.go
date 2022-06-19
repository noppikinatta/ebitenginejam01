package animation

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var faderGlobalCacheImage *ebiten.Image

type FadeIn struct {
	fade
}

func NewFadeIn(frames int) *FadeIn {
	f := FadeIn{
		fade: fade{
			frames: frames,
		},
	}

	return &f
}

func (f *FadeIn) Draw(target *ebiten.Image) {
	f.fade.Draw(target, f.alpha())
}

func (f *FadeIn) alpha() float64 {
	a := 1.0 - float64(f.current)/float64(f.frames)
	if a < 0 {
		a = 0
	}
	return a
}

type FadeOut struct {
	fade
}

func NewFadeOut(frames int) *FadeOut {
	f := FadeOut{
		fade: fade{
			frames: frames,
		},
	}

	return &f
}

func (f *FadeOut) Draw(target *ebiten.Image) {
	f.fade.Draw(target, f.alpha())
}

func (f *FadeOut) alpha() float64 {
	a := float64(f.current) / float64(f.frames)
	if a > 1 {
		a = 1
	}
	return a
}

type fade struct {
	current int
	frames  int
	opt     *ebiten.DrawImageOptions
}

func (f *fade) Reset() {
	f.current = 0
}

func (f *fade) Update() {
	if !f.End() {
		f.current++
	}
}

func (f *fade) Draw(target *ebiten.Image, alpha float64) {
	if alpha == 0 {
		return
	}
	c := f.cache(target.Size())
	target.DrawImage(c, f.option(alpha))
}

func (f *fade) cache(w, h int) *ebiten.Image {
	if f.shouldUpdateCache(w, h) {
		f.updateCache(w, h)
	}

	return faderGlobalCacheImage
}

func (f *fade) shouldUpdateCache(w, h int) bool {
	if faderGlobalCacheImage == nil {
		return true
	}
	cw, ch := faderGlobalCacheImage.Size()
	if cw != w || ch != h {
		return true
	}
	return false
}

func (f *fade) updateCache(w, h int) {
	if faderGlobalCacheImage != nil {
		faderGlobalCacheImage.Dispose()
	}
	faderGlobalCacheImage = ebiten.NewImage(w, h)
	faderGlobalCacheImage.Fill(color.Black)
}

func (f *fade) option(alpha float64) *ebiten.DrawImageOptions {
	if f.opt == nil {
		f.opt = &ebiten.DrawImageOptions{}
	}
	cm := ebiten.ColorM{}
	cm.Scale(1, 1, 1, alpha)
	f.opt.ColorM = cm
	return f.opt
}

func (f *fade) End() bool {
	return f.current > f.frames
}
