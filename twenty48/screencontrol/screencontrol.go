package screencontrol

import (
	"math"

	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/eventhandler"
	"github.com/hajimehoshi/ebiten/v2"
)

type ScreenControl struct {
	isFullscreen bool
	d            Deps
	actualWidth  int
	actualHeight int
	scale        float64
}

func New(d Deps) *ScreenControl {
	sc := &ScreenControl{
		isFullscreen: false,
		d:            d,
		scale:        1,
	}

	sc.updateActualDimentions()
	return sc
}

func (sc *ScreenControl) updateActualDimentions() {
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
	sc.setFullScreen(!sc.isFullscreen)

	// Trigger screen changed event
	sc.d.Emit(eventhandler.Event{
		Type: eventhandler.EventScreenChanged,
	})
}

func (sc *ScreenControl) setFullScreen(val bool) {
	sc.isFullscreen = val
	sc.updateActualDimentions()
}

func (sc *ScreenControl) IsFullScreen() bool {
	return sc.isFullscreen
}

func (sc *ScreenControl) GetScale() float64 {
	return sc.scale
}

func (sc *ScreenControl) IncrementScale() {
	sc.scale++
	sc.updateActualDimentions()
}

func (sc *ScreenControl) DecrementScale() bool {
	if sc.scale > 1 {
		sc.scale--
		sc.updateActualDimentions()
		return true
	}
	return false
}

func (g *ScreenControl) LayoutF(logicWinWidth, logicWinHeight float64) (float64, float64) {
	scale := ebiten.Monitor().DeviceScaleFactor()
	canvasWidth := math.Ceil(logicWinWidth * scale)
	canvasHeight := math.Ceil(logicWinHeight * scale)
	return canvasWidth, canvasHeight
}
