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
	if sc.fullscreen {
		sc.actualWidth, sc.actualHeight = ebiten.Monitor().Size()
	} else {
		sc.actualWidth = logicalWidth * int(sc.game.scale)
		sc.actualHeight = logicalHeight * int(sc.game.scale)
	}
}
