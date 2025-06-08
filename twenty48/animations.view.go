package twenty48

import "github.com/hajimehoshi/ebiten/v2"

type View interface {
	BoardProvider
}

type BoardProvider interface {
	DrawBackgoundBoard(screen *ebiten.Image)
	GetBoardDimentions() (x, y int)
	DrawMovingMatrix(screen *ebiten.Image, x, y int, movDistX, movDistY float32)
}
