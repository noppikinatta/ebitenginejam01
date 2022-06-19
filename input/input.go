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
