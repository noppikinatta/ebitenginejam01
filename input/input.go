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

package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var touchIDs []ebiten.TouchID
var touchIDsJust []ebiten.TouchID

func LeftMousedownOrTouched() bool {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return true
	}

	return len(justPressedTouchIDs()) > 0
}

func CursorPosition() (x, y int) {
	x, y = ebiten.CursorPosition()

	if x == 0 && y == 0 {
		tt := pressedTouchIDs()
		if len(tt) > 0 {
			x, y = ebiten.TouchPosition(tt[0])
		}
	}

	return x, y
}

func pressedTouchIDs() []ebiten.TouchID {
	if len(touchIDs) > 0 {
		touchIDs = touchIDs[:0]
	}
	touchIDs = ebiten.AppendTouchIDs(touchIDs)
	return touchIDs
}

func justPressedTouchIDs() []ebiten.TouchID {
	if len(touchIDsJust) > 0 {
		touchIDsJust = touchIDsJust[:0]
	}
	touchIDsJust = inpututil.AppendJustPressedTouchIDs(touchIDsJust)
	return touchIDsJust
}
