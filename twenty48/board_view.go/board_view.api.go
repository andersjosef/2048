package board_view

import "github.com/hajimehoshi/ebiten/v2"

func (b *BoardView) DrawBackgoundBoard(screen *ebiten.Image) {
	screen.DrawImage(b.boardImage, b.boardImageOptions)
}

func (b *BoardView) DrawMovingMatrix(
	screen *ebiten.Image,
	x,
	y int,
	movDistX,
	movDistY float32,
	value int,
) {
	// matrix := b.d.Board.CurMatrixSnapshot()
	b.DrawTile(
		screen,
		b.sizes.startPosX,
		b.sizes.startPosY,
		x,
		y,
		value,
		movDistX,
		movDistY,
	)

}
