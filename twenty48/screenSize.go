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

func (s *ScreenControl) ChangeBoardPosition() {
	var newWidth, newHeight int = s.GetRealWidthHeight()
	startPosX = float32((newWidth - (BOARDSIZE * int(TILESIZE))) / 2)
	startPosY = float32((newHeight - (BOARDSIZE * int(TILESIZE))) / 2)
	s.game.board.createBoardImage()
	s.game.screenSizeChanged = false
}

func (s *ScreenControl) GetRealWidthHeight() (int, int) {
	var newWidth, newHeight int
	if s.fullscreen { // changing to full screen
		newWidth, newHeight = ebiten.Monitor().Size()
	} else { // changing to small
		newWidth, newHeight = SCREENWIDTH, SCREENHEIGHT
	}
	return newWidth, newHeight
}
