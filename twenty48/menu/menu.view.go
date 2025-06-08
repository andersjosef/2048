package menu

import (
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/eventhandler"
	"github.com/andersjosef/2048/twenty48/theme"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type View interface {
	GameProvider
	ButtonManagerProvider
	ScreenControlProvider
	RendererProvider
}

type GameProvider interface {
	GetState() co.GameState
	GetPreviousState() co.GameState
	GetCurrentTheme() theme.Theme
	GetFontSet() theme.FontSet
	GetScore() int
	GetBusHandler() *eventhandler.EventBus
}

type ButtonManagerProvider interface {
	UpdatePosForButton(keyName string, posX, posY int) (exists bool)
	UpdateTextForButton(keyName, newText string) (exists bool)
	ButtonExists(keyname string) (exists bool)
}

type ScreenControlProvider interface {
	GetActualSize() (x, y int)
	IsFullScreen() bool
}

type RendererProvider interface {
	DrawDoubleText(screen *ebiten.Image, message string, xpos int, ypos int, offset float64, fontUsed *text.GoTextFace, isCentered bool)
}
