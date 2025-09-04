package board

import (
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/eventhandler"
	"github.com/andersjosef/2048/twenty48/theme"
)

type Deps struct {
	EventHandler
	ScreenControl
	Core

	SetGameOver       func(isGameOver bool)
	SetGameState      func(co.GameState)
	IsGameOver        func() bool
	GetCurrentTheme   func() theme.Theme
	GetCurrentFontSet func() theme.FontSet
}

type EventHandler interface {
	Register(eventType eventhandler.EventType, handler func(eventhandler.Event))
	Dispatch()
	Emit(event eventhandler.Event)
}

type ScreenControl interface {
	GetScale() float64
	GetActualSize() (x, y int)
}

type Core interface {
	AddScore(int)
}
