package input

import (
	"github.com/andersjosef/2048/twenty48/commands"
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/eventhandler"
)

type Deps struct {
	EventHandler
	Buttons
	ScreenControl

	Cmds       commands.Commands
	GetState   func() co.GameState
	SetState   func(co.GameState)
	IsGameOver func() bool
}

type EventHandler interface {
	Register(eventType eventhandler.EventType, handler func(eventhandler.Event))
	// Dispatch()
	Emit(event eventhandler.Event)
}

type Buttons interface {
	CheckButtons() bool
	UpdatePosForButton(string, int, int) bool
}

type ScreenControl interface {
	GetActualSize() (x, y int)
}
