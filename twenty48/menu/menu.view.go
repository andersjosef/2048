package menu

import (
	"github.com/andersjosef/2048/twenty48/eventhandler"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Deps struct {
	Buttons
	Renderer
	EventHandler

	GetSnapShot func() Snapshot
}

type EventHandler interface {
	Register(eventType eventhandler.EventType, handler func(eventhandler.Event))
}

type Buttons interface {
	UpdatePosForButton(keyName string, posX, posY int) (exists bool)
	UpdateTextForButton(keyName, newText string) (exists bool)
	ButtonExists(keyname string) (exists bool)
}

type Renderer interface {
	DrawDoubleText(screen *ebiten.Image, message string, xpos int, ypos int, offset float64, fontUsed *text.GoTextFace, isCentered bool)
}
