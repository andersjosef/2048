package screencontrol

import (
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/eventhandler"
	"github.com/hajimehoshi/ebiten/v2"
)

type ScreenControl struct {
	isFullscreen bool
	view         View
	actualWidth  int
	actualHeight int
	scale        float64
}

func InitScreenControl(g View) *ScreenControl {
	sc := &ScreenControl{
		isFullscreen: false,
		view:         g,
		scale:        1,
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
		sc.actualWidth = co.LOGICAL_WIDTH * int(sc.scale) * int(dpiScale)
		sc.actualHeight = co.LOGICAL_HEIGHT * int(sc.scale) * int(dpiScale)
	}
}

func (sc *ScreenControl) GetActualSize() (x, y int) {
	return sc.actualWidth, sc.actualHeight
}

func (sc *ScreenControl) ToggleFullScreen() {
	ebiten.SetFullscreen(!sc.isFullscreen)
	sc.SetFullScreen(!sc.isFullscreen)

	// Trigger screen changed event
	sc.view.Emit(eventhandler.Event{
		Type: eventhandler.EventScreenChanged,
	})
}

func (sc *ScreenControl) SetFullScreen(val bool) {
	sc.isFullscreen = val
	sc.UpdateActualDimentions()
}

func (sc *ScreenControl) IsFullScreen() bool {
	return sc.isFullscreen
}

func (sc *ScreenControl) GetScale() float64 {
	return sc.scale
}

func (sc *ScreenControl) IncrementScale() {
	sc.scale++
	sc.UpdateActualDimentions()
}

func (sc *ScreenControl) DecrementScale() bool {
	if sc.scale > 1 {
		sc.scale--
		sc.UpdateActualDimentions()
		return true
	}
	return false
}
