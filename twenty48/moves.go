package twenty48

import co "github.com/andersjosef/2048/twenty48/constants"

func (b *Board) updateBoardBeforeChange() {
	b.boardBeforeChange = b.board
}

func (b *Board) moveLeft() {
	b.updateBoardBeforeChange()
	for i := range b.board {
		// Shift tiles to the left
		compactTiles(i, b, true)
		// Merge tiles and shift again if needed
		mergeTiles(&b.board[i], b)
		compactTiles(i, b, false)
	}
	b.game.animation.ActivateAnimation("LEFT")

	b.addNewRandomPieceIfBoardChanged()

	b.game.gameOver = b.isGameOver()
}

func (b *Board) moveUp() {
	b.updateBoardBeforeChange()
	transpose(&b.board)
	for i := range b.board {
		// Shift tiles "left" (actually up, due to transposition)
		compactTiles(i, b, true)
		// Merge tiles and shift again if needed
		mergeTiles(&b.board[i], b)
		compactTiles(i, b, false)
	}
	transpose(&b.board) // Transpose back to the original orientation
	transpose(&b.game.animation.arrayOfChange)
	b.game.animation.ActivateAnimation("UP")

	b.addNewRandomPieceIfBoardChanged()

	b.game.gameOver = b.isGameOver()
}

func (b *Board) moveRight() {
	b.updateBoardBeforeChange()
	for i := range b.board {
		// Reverse the row to treat the right end as the left
		reverseRow(&b.board[i])
		// Shift tiles "left" (actually right, due to reversal)
		compactTiles(i, b, true)
		// Merge tiles and shift again if needed
		mergeTiles(&b.board[i], b)
		compactTiles(i, b, false)
		// Reverse back to original orientation
		reverseRow(&b.board[i])
		reverseRow(&b.game.animation.arrayOfChange[i])

		b.game.animation.ActivateAnimation("RIGHT")
	}

	b.addNewRandomPieceIfBoardChanged()

	b.game.gameOver = b.isGameOver()
}
func (b *Board) moveDown() {
	b.updateBoardBeforeChange()
	transpose(&b.board)
	for i := range b.board {
		// Reverse the row (which is actually a column due to transposition)
		reverseRow(&b.board[i])
		// Shift tiles "left" (actually down, due to reversal and transposition)
		compactTiles(i, b, true)
		// Merge tiles and shift again if needed
		mergeTiles(&b.board[i], b)
		compactTiles(i, b, false)
		// Reverse back to treat the bottom as the top
		reverseRow(&b.board[i])
		reverseRow(&b.game.animation.arrayOfChange[i])
	}
	transpose(&b.board) // Transpose back to the original orientation
	transpose(&b.game.animation.arrayOfChange)
	b.game.animation.ActivateAnimation("DOWN")

	b.addNewRandomPieceIfBoardChanged()

	b.game.gameOver = b.isGameOver()
}

func reverseRow(row *[co.BOARDSIZE]int) {
	for i, j := 0, len(*row)-1; i < j; i, j = i+1, j-1 {
		(*row)[i], (*row)[j] = (*row)[j], (*row)[i]
	}
}

// Moves all tiles to the left
func compactTiles(rowIndex int, b *Board, beforeMerge bool) {
	insertPos := 0

	// these two are for adding an extra move to the animation to make it pretty
	lastVal := -1
	extraMov := 0
	for i, val := range b.board[rowIndex] {
		if val != 0 {
			if val == lastVal {
				extraMov++
			}
			if beforeMerge {
				b.game.animation.arrayOfChange[rowIndex][i] = (i - insertPos) + extraMov // delta movement to the left
			}
			(b.board[rowIndex])[insertPos] = val
			insertPos++
			lastVal = val

		}
	}
	// Fill the rest with 0s
	for i := insertPos; i < len(b.board[rowIndex]); i++ {
		b.board[rowIndex][i] = 0
	}
}

func mergeTiles(row *[co.BOARDSIZE]int, b *Board) {
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
func transpose(board *[co.BOARDSIZE][co.BOARDSIZE]int) {
	for i := 0; i < len(*board); i++ {
		for j := i; j < len((*board)[0]); j++ {
			(*board)[i][j], (*board)[j][i] = (*board)[j][i], (*board)[i][j]
		}
	}
}
