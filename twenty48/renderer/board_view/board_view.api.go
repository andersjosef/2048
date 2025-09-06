package board_view

import "github.com/hajimehoshi/ebiten/v2"

func (b *BoardView) DrawBackgoundBoard(screen *ebiten.Image) {
	screen.DrawImage(b.emptyBoard, b.opts)
}

func (b *BoardView) DrawMovingMatrix(
	screen *ebiten.Image,
	x,
	y int,
	movDistX,
	movDistY float32,
	value int,
) {
	b.drawTile(
		screen,
		x,
		y,
		value,
		movDistX,
		movDistY,
	)

}
