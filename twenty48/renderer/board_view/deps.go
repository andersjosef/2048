package board_view

import (
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/eventhandler"
	"github.com/andersjosef/2048/twenty48/theme"
)

type BoardViewDeps struct {
	EventHandler
	ScreenControl
	Theme
	Layout
	Board interface {
		CurMatrixSnapshot() [co.BOARDSIZE][co.BOARDSIZE]int
		GetBoardDimentions() (x, y int)
	}
	IsGameOver func() bool
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

type Theme interface {
	Current() theme.Theme
	Fonts() *theme.FontSet
}

type Layout interface {
	Recalculate()
	BorderSize() float32
	GetStartPos() (x, y float32)
	TileSize() float32
	StartPos() (x, y float32)
}
