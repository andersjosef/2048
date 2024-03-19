package twenty48

func (b *Board) moveLeft() {
	b.game.animation.ResetArray()
	for i := range b.board {
		// Shift tiles to the left
		compactTiles(i, b)
		// Merge tiles and shift again if needed
		mergeTiles(&b.board[i], b)
		compactTiles(i, b)
	}
	// fmt.Println(b.game.animation.arrayOfChange)
	b.game.animation.ActivateAnimation("LEFT")

}

func (b *Board) moveUp() {
	transpose(&b.board)
	for i := range b.board {
		// Shift tiles "left" (actually up, due to transposition)
		compactTiles(i, b)
		// Merge tiles and shift again if needed
		mergeTiles(&b.board[i], b)
		compactTiles(i, b)
	}
	transpose(&b.board)                        // Transpose back to the original orientation
	transpose(&b.game.animation.arrayOfChange) // Transpose back to the original orientation
	b.game.animation.ActivateAnimation("UP")

}

func (b *Board) moveRight() {
	for i := range b.board {
		// Reverse the row to treat the right end as the left
		reverseRow(&b.board[i])
		// Shift tiles "left" (actually right, due to reversal)
		compactTiles(i, b)
		// Merge tiles and shift again if needed
		mergeTiles(&b.board[i], b)
		compactTiles(i, b)
		// Reverse back to original orientation
		reverseRow(&b.board[i])
		reverseRow(&b.game.animation.arrayOfChange[i])

		b.game.animation.ActivateAnimation("RIGHT")
	}
}
func (b *Board) moveDown() {
	transpose(&b.board)
	for i := range b.board {
		// Reverse the row (which is actually a column due to transposition)
		reverseRow(&b.board[i])
		// Shift tiles "left" (actually down, due to reversal and transposition)
		compactTiles(i, b)
		// Merge tiles and shift again if needed
		mergeTiles(&b.board[i], b)
		compactTiles(i, b)
		// Reverse back to treat the bottom as the top
		reverseRow(&b.board[i])
		reverseRow(&b.game.animation.arrayOfChange[i])
	}
	transpose(&b.board)                        // Transpose back to the original orientation
	transpose(&b.game.animation.arrayOfChange) // Transpose back to the original orientation
	b.game.animation.ActivateAnimation("DOWN")
}

func reverseRow(row *[BOARDSIZE]int) {
	for i, j := 0, len(*row)-1; i < j; i, j = i+1, j-1 {
		(*row)[i], (*row)[j] = (*row)[j], (*row)[i]
	}
}

// Moves all tiles to the left
func compactTiles(rowIndex int, b *Board) {
	insertPos := 0
	for i, val := range b.board[rowIndex] {
		if val != 0 {
			b.game.animation.arrayOfChange[rowIndex][i] = (i - insertPos) // delta movement to the left
			(b.board[rowIndex])[insertPos] = val
			insertPos++
		}
	}
	// Fill the rest with 0s
	for i := insertPos; i < len(b.board[rowIndex]); i++ {
		b.board[rowIndex][i] = 0
	}
}

func mergeTiles(row *[BOARDSIZE]int, b *Board) {
	for i := 0; i < len(*row)-1; i++ {
		if (*row)[i] == (*row)[i+1] && (*row)[i] != 0 {
			(*row)[i] *= 2
			b.game.score += (*row)[i]
			(*row)[i+1] = 0
			i++ // Skip next tile as it has been merged and set to 0
		}
	}
}

// Swap cols and rows
func transpose(board *[BOARDSIZE][BOARDSIZE]int) {
	for i := 0; i < len(*board); i++ {
		for j := i; j < len((*board)[0]); j++ {
			(*board)[i][j], (*board)[j][i] = (*board)[j][i], (*board)[i][j]
		}
	}
}
