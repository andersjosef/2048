package twenty48

import "github.com/hajimehoshi/ebiten/v2"

type ScreenControl struct {
	fullscreen bool
	game       *Game
}

func InitScreenControl(g *Game) *ScreenControl {
	sc := &ScreenControl{
		fullscreen: false,
		game:       g,
	}

	return sc
}

func (g *Game) ChangeBoardPosition() {
	var newWidth, newHeight int = g.GetRealWidthHeight()
	startPosX = float32((newWidth - (BOARDSIZE * int(TILESIZE))) / 2)
	startPosY = float32((newHeight - (BOARDSIZE * int(TILESIZE))) / 2)
	g.board.createBoardImage()
	g.screenSizeChanged = false
}

func (g *Game) GetRealWidthHeight() (int, int) {
	var newWidth, newHeight int
	if g.screenControl.fullscreen { // changing to full screen
		newWidth, newHeight = ebiten.ScreenSizeInFullscreen()
	} else { // changing to small
		newWidth, newHeight = SCREENWIDTH, SCREENHEIGHT
	}
	return newWidth, newHeight
}
