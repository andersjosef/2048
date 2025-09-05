package board

import (
	co "github.com/andersjosef/2048/twenty48/constants"
)

func (b *Board) CurMatrixSnapshot() [co.BOARDSIZE][co.BOARDSIZE]int {
	return b.matrix
}

func (b *Board) PrevMatrixSnapshot() [co.BOARDSIZE][co.BOARDSIZE]int {
	return b.matrixBeforeChange
}

func (b *Board) GetBoardDimentions() (x, y int) {
	if len(b.matrix) == 0 {
		return 0, 0
	}
	return len(b.matrix[0]), len(b.matrix)
}
