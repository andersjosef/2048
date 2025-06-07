package twenty48

import "github.com/hajimehoshi/ebiten/v2"

type ScreenControl struct {
	fullscreen   bool
	game         *Game
	actualWidth  int
	actualHeight int
}

func InitScreenControl(g *Game) *ScreenControl {
	sc := &ScreenControl{
		fullscreen: false,
		game:       g,
	}

	sc.UpdateActualDimentions()
	return sc
}

func (sc *ScreenControl) UpdateActualDimentions() {
	dpiScale := ebiten.Monitor().DeviceScaleFactor() // Accounting for high dpi monitors
	if sc.fullscreen {
		sc.actualWidth, sc.actualHeight = ebiten.Monitor().Size()
		sc.actualWidth *= int(dpiScale)
		sc.actualHeight *= int(dpiScale)
	} else {
		sc.actualWidth = LOGICAL_WIDTH * int(sc.game.scale) * int(dpiScale)
		sc.actualHeight = LOGICAL_HEIGHT * int(sc.game.scale) * int(dpiScale)
	}
}
