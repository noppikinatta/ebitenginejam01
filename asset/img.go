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
	"embed"
	"fmt"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed img/*.png
var imgDir embed.FS

var imgCache map[Img]*ebiten.Image

func fromCache(i Img) (*ebiten.Image, bool) {
	// concurrent unsafe but can implement easily

	if imgCache == nil {
		imgCache = make(map[Img]*ebiten.Image)
	}

	img, ok := imgCache[i]
	return img, ok
}

type Img string

func (i Img) MustImage() *ebiten.Image {
	img, ok := fromCache(i)
	if ok {
		return img
	}

	// MustImage can panic, because Img type values are defined by const and fixed
	img, err := i.createImage()
	if err != nil {
		panic(err)
	}
	imgCache[i] = img

	return img
}

func (i Img) createImage() (*ebiten.Image, error) {
	imgFile, err := imgDir.Open(i.path())
	if err != nil {
		return nil, err
	}

	baseImg, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, err
	}

	return ebiten.NewImageFromImage(baseImg), nil
}

func (i Img) path() string {
	return fmt.Sprintf("img/%s", i)
}
