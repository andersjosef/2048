package animations

import (
	"github.com/andersjosef/2048/twenty48/eventhandler"
	"github.com/hajimehoshi/ebiten/v2"
)

type Deps struct {
	Board
	BoardView
	EventHandler
	Layout
	State
}

type EventHandler interface {
	Register(eventType eventhandler.EventType, handler func(eventhandler.Event))
	Emit(eventhandler.Event)
}

type Board interface {
	GetBoardDimentions() (x, y int)
}
type BoardView interface {
	DrawBackgoundBoard(screen *ebiten.Image)
	GetTile(v int) (img *ebiten.Image, ok bool)
}

type Layout interface {
	GetStartPos() (x, y float32)
	TileSize() float32
	BorderSize() float32
}
type State interface {
	IsGameOver() bool
}
