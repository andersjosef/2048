package animations

import (
	"github.com/andersjosef/2048/twenty48/eventhandler"
	"github.com/hajimehoshi/ebiten/v2"
)

type View interface {
	BoardProvider
}

type BoardProvider interface {
	DrawBackgoundBoard(screen *ebiten.Image)
	GetBoardDimentions() (x, y int)
	DrawMovingMatrix(screen *ebiten.Image, x, y int, movDistX, movDistY float32)
	GetBusHandler() *eventhandler.EventBus
}
