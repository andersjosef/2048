package twenty48

import (
	"github.com/andersjosef/2048/twenty48/theme"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type GameView interface {
	GameProvider
	ButtonManagerProvider
	ScreenControlProvider
	RendererProvider
}

type GameProvider interface {
	GetState() GameState
	GetCurrentTheme() theme.Theme
	GetFontSet() theme.FontSet
}

type ButtonManagerProvider interface {
	UpdatePosForButton(keyName string, posX, posY int)
}

type ScreenControlProvider interface {
	GetActualSize() (x, y int)
	GetIsFullScreen() bool
}

type RendererProvider interface {
	DrawDoubleText(screen *ebiten.Image, message string, xpos int, ypos int, offset float64, fontUsed *text.GoTextFace, isCentered bool)
}
