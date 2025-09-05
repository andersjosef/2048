package animations

import (
	"github.com/andersjosef/2048/twenty48/eventhandler"
	"github.com/hajimehoshi/ebiten/v2"
)

type Deps struct {
	Board
	BoardView
	EventHandler
}

type EventHandler interface {
	Register(eventType eventhandler.EventType, handler func(eventhandler.Event))
}

type Board interface {
	GetBoardDimentions() (x, y int)
}
type BoardView interface {
	DrawBackgoundBoard(screen *ebiten.Image)
	DrawMovingMatrix(screen *ebiten.Image, x, y int, movDistX, movDistY float32)
}
