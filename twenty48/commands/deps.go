package commands

import (
	"github.com/andersjosef/2048/twenty48/board"
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/eventhandler"
)

type Deps struct {
	Board
	BoardView
	EventHandler
	ScreenControl
	FSM

	IncrementCurrentTheme func()
	ToggleInfo            func()
	ScaleWindow           func()
}

type Board interface {
	Move(board.Direction)
}

type BoardView interface {
	CreateBoardImage()
}

type EventHandler interface {
	Register(eventType eventhandler.EventType, handler func(eventhandler.Event))
	Dispatch()
	Emit(event eventhandler.Event)
}

type ScreenControl interface {
	ToggleFullScreen()
	IncrementScale()
	DecrementScale() bool
}

type FSM interface {
	Switch(co.GameState)
}
