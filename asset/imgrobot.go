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
		return image.Rect(64, 16, 79, 31)
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
