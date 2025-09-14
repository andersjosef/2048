package screencontrol

import (
	"math"

	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/eventhandler"
	"github.com/hajimehoshi/ebiten/v2"
)

type ScreenControl struct {
	isFullscreen   bool
	d              Deps
	windowWidth    int
	windowHeight   int
	scale          float64
	fraction       float64
	scaleThreshold float64
}

func New(d Deps) *ScreenControl {
	sc := &ScreenControl{
		isFullscreen: false,
		d:            d,
		scale:        1,
		fraction:     1.5,
	}

	sc.scaleThreshold = sc.scale / (sc.fraction * 2)
	sc.updateActualDimentions()
	return sc
}

func (sc *ScreenControl) updateActualDimentions() {
	dpiScale := ebiten.Monitor().DeviceScaleFactor() // Accounting for high dpi monitors
	if sc.isFullscreen {
		sc.windowWidth, sc.windowHeight = ebiten.Monitor().Size()
		sc.windowWidth *= int(dpiScale)
		sc.windowHeight *= int(dpiScale)
	} else {
		sc.windowWidth = int(float64(co.LOGICAL_WIDTH) * sc.scale * dpiScale)
		sc.windowHeight = int(float64(co.LOGICAL_HEIGHT) * sc.scale * dpiScale)
	}
}

func (sc *ScreenControl) GetActualSize() (x, y int) {
	return sc.windowWidth, sc.windowHeight
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
	if sc.scale >= 1 {
		sc.scale++
	} else {
		sc.scale *= sc.fraction
	}
	sc.updateActualDimentions()
}

func (sc *ScreenControl) DecrementScale() bool {
	if sc.scale > 1 {
		sc.scale--
		sc.updateActualDimentions()
		return true
	} else if sc.scale > sc.scaleThreshold {
		sc.scale /= sc.fraction
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
