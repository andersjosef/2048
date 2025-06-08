package screencontrol

import (
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/hajimehoshi/ebiten/v2"
)

type ScreenControl struct {
	isFullscreen bool
	view         GameView
	actualWidth  int
	actualHeight int
}

func InitScreenControl(g GameView) *ScreenControl {
	sc := &ScreenControl{
		isFullscreen: false,
		view:         g,
	}

	sc.UpdateActualDimentions()
	return sc
}

func (sc *ScreenControl) UpdateActualDimentions() {
	dpiScale := ebiten.Monitor().DeviceScaleFactor() // Accounting for high dpi monitors
	if sc.isFullscreen {
		sc.actualWidth, sc.actualHeight = ebiten.Monitor().Size()
		sc.actualWidth *= int(dpiScale)
		sc.actualHeight *= int(dpiScale)
	} else {
		sc.actualWidth = co.LOGICAL_WIDTH * int(sc.view.GetScale()) * int(dpiScale)
		sc.actualHeight = co.LOGICAL_HEIGHT * int(sc.view.GetScale()) * int(dpiScale)
	}
}

func (sc *ScreenControl) GetActualSize() (x, y int) {
	return sc.actualWidth, sc.actualHeight
}

func (sc *ScreenControl) IsFullScreen() bool {
	return sc.isFullscreen
}

func (sc *ScreenControl) SetFullScreen(val bool) {
	sc.isFullscreen = val
}
