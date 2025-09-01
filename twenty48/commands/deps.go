package commands

import (
	"github.com/andersjosef/2048/twenty48/board"
	"github.com/andersjosef/2048/twenty48/eventhandler"
)

type Deps struct {
	Board
	EventHandler
	ScreenControl
	Menu

	SetCloseGame          func(bool)
	IncrementCurrentTheme func()
	ToggleInfo            func()
}

type Board interface {
	Move(board.Direction)
	CreateBoardImage()
}

type EventHandler interface {
	Register(eventType eventhandler.EventType, handler func(eventhandler.Event))
	Dispatch()
	Emit(event eventhandler.Event)
}

type ScreenControl interface {
	ToggleFullScreen()
}

type Menu interface {
	UpdateDynamicText()
}
