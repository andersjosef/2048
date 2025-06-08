package twenty48

import (
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/hajimehoshi/ebiten/v2"
)

type ScreenControl struct {
	isFullscreen bool
	game         *Game
	actualWidth  int
	actualHeight int
}

func InitScreenControl(g *Game) *ScreenControl {
	sc := &ScreenControl{
		isFullscreen: false,
		game:         g,
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
		sc.actualWidth = co.LOGICAL_WIDTH * int(sc.game.scale) * int(dpiScale)
		sc.actualHeight = co.LOGICAL_HEIGHT * int(sc.game.scale) * int(dpiScale)
	}
}
