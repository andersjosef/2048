package twenty48

import "github.com/hajimehoshi/ebiten/v2"

func (b *Board) DrawBackgoundBoard(screen *ebiten.Image) {
	screen.DrawImage(b.boardImage, b.boardImageOptions)
}

func (b *Board) GetBoardDimentions() (x, y int) {
	if len(b.matrix) == 0 {
		return 0, 0
	}
	return len(b.matrix[0]), len(b.matrix)
}

func (b *Board) DrawMovingMatrix(
	screen *ebiten.Image,
	x,
	y int,
	movDistX,
	movDistY float32,
) {
	b.DrawTile(
		screen,
		b.sizes.startPosX,
		b.sizes.startPosY,
		x,
		y,
		b.matrixBeforeChange[y][x],
		movDistX,
		movDistY,
	)

}
