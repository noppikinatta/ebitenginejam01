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

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/noppikinatta/ebitenginejam01/asset"
	"github.com/noppikinatta/ebitenginejam01/combine"
)

type completeState int

const (
	completeStateBack completeState = iota
	completeStateGo
)

type complete struct {
	img    *ebiten.Image
	drawer *combine.Drawer
	w, h   float64
	x, y   float64
	state  completeState
}

func newComplete(d *combine.Drawer, w, h float64) *complete {
	c := complete{
		img:    ebiten.NewImage(int(w), int(h)),
		drawer: d,
		w:      w,
		h:      h,
	}
	c.Reset()
	return &c
}

func (c *complete) SetLoc(x, y float64) {
	c.x = x
	c.y = y
}

func (c *complete) Update() {
	switch c.state {
	case completeStateBack:
		c.y += 5
		if c.y > (c.h + 100) {
			c.state = completeStateGo
			c.y = 400
			asset.PlaySound(asset.SEFly)
		}
	case completeStateGo:
		c.y -= 20
	}
}

func (c *complete) Draw(screen *ebiten.Image) {
	switch c.state {
	case completeStateBack:
		gm := ebiten.GeoM{}
		gm.Translate(c.x, c.y)
		c.drawer.Draw(screen, gm)
	case completeStateGo:
		c.img.Clear()
		c.drawer.Draw(c.img, ebiten.GeoM{})
		gm := ebiten.GeoM{}
		gm.Scale(2, 2)
		gm.Translate(0, c.y)
		opt := ebiten.DrawImageOptions{}
		opt.GeoM = gm
		screen.DrawImage(c.img, &opt)
	}
}

func (c *complete) End() bool {
	return c.state == completeStateGo && c.y < -500
}

func (c *complete) Reset() {
	c.state = completeStateBack
}
