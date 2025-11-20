package board

/*
	This file contains functions intended purely for ease of debugging
*/

func (b *Board) SetBoard(newMatrix [4][4]int) {
	b.matrix = newMatrix
}
