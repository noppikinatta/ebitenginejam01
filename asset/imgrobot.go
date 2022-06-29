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

package asset

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

var robotPartImgCache map[RobotPart]*ebiten.Image

type RobotPart int

const (
	RobotPartBody RobotPart = iota
	RobotPartRightArm
	RobotPartLeftArm
	RobotPartTNT
	RobotPartEbitenN
	RobotPartEbitenS
	RobotPartLeftLeg
	RobotPartRightLeg
)

func (p RobotPart) SubImageRect() image.Rectangle {
	switch p {
	case RobotPartBody:
		return image.Rect(0, 0, 63, 79)
	case RobotPartRightArm:
		return image.Rect(64, 0, 127, 15)
	case RobotPartLeftArm:
		return image.Rect(128, 0, 191, 15)
	case RobotPartTNT:
		return image.Rect(64, 16, 127, 31)
	case RobotPartEbitenN:
		return image.Rect(64, 32, 127, 47)
	case RobotPartEbitenS:
		return image.Rect(128, 32, 191, 47)
	case RobotPartLeftLeg:
		return image.Rect(64, 48, 127, 63)
	case RobotPartRightLeg:
		return image.Rect(128, 48, 191, 63)
	}

	panic("wrong robot part.")
}

func robotPartFromCache(p RobotPart) (*ebiten.Image, bool) {
	if robotPartImgCache == nil {
		robotPartImgCache = make(map[RobotPart]*ebiten.Image)
	}

	img, ok := robotPartImgCache[p]
	return img, ok
}

func ImgRobotPart(p RobotPart) *ebiten.Image {
	img, ok := robotPartFromCache(p)
	if ok {
		return img
	}
	img = ImgGameplayRobot.MustImage()
	img = img.SubImage(p.SubImageRect()).(*ebiten.Image)
	robotPartImgCache[p] = img
	return img
}
