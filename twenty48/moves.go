package twenty48

import (
	"github.com/andersjosef/2048/twenty48/animations"
	co "github.com/andersjosef/2048/twenty48/constants"
)

func (b *Board) updateBoardBeforeChange() {
	b.matrixBeforeChange = b.matrix
}

func (b *Board) moveLeft() {
	b.updateBoardBeforeChange()

	var allDeltas []animations.MoveDelta
	var newMatrix [4][4]int

	for rowIndex, row := range b.matrix {
		slideRow, d1 := compactRow(rowIndex, row, true)

		mergedRow, scoreGain := mergeRow(slideRow)
		b.game.score += scoreGain

		finalRow, _ := compactRow(rowIndex, mergedRow, false)

		// Write to new matrix
		newMatrix[rowIndex] = finalRow

		// Collect all deltas
		allDeltas = append(allDeltas, d1...)

	}
	b.game.animation.Play(allDeltas, "LEFT")

	b.matrix = newMatrix
	b.addNewRandomPieceIfBoardChanged()
	b.game.gameOver = b.isGameOver()
}

func (b *Board) moveUp() {
	b.updateBoardBeforeChange()
	var snapShot [4][4]int
	copy(snapShot[:], b.matrix[:])

	var allDeltas []animations.MoveDelta
	var newMatrix [4][4]int

	transpose(&snapShot)
	for rowIndex, row := range snapShot {
		slideRow, d1 := compactRow(rowIndex, row, true)

		mergedRow, scoreGain := mergeRow(slideRow)
		b.game.score += scoreGain

		finalRow, _ := compactRow(rowIndex, mergedRow, false)

		// Write to new matrix
		newMatrix[rowIndex] = finalRow

		// Collect all deltas
		allDeltas = append(allDeltas, d1...)

	}
	transpose(&newMatrix)
	b.game.animation.Play(allDeltas, "UP")

	b.matrix = newMatrix
	b.addNewRandomPieceIfBoardChanged()
	b.game.gameOver = b.isGameOver()
}

func (b *Board) moveRight() {
	b.updateBoardBeforeChange()
	var snapShot [4][4]int
	copy(snapShot[:], b.matrix[:])

	var allDeltas []animations.MoveDelta
	var newMatrix [4][4]int

	for rowIndex, row := range snapShot {
		reverseRow(&row)
		slideRow, d1 := compactRow(rowIndex, row, true)

		mergedRow, scoreGain := mergeRow(slideRow)
		b.game.score += scoreGain

		finalRow, _ := compactRow(rowIndex, mergedRow, false)

		reverseRow(&finalRow)

		// Write to new matrix
		newMatrix[rowIndex] = finalRow

		// Collect all deltas
		allDeltas = append(allDeltas, d1...)

	}
	b.game.animation.Play(allDeltas, "RIGHT")

	b.matrix = newMatrix
	b.addNewRandomPieceIfBoardChanged()
	b.game.gameOver = b.isGameOver()
}

func (b *Board) moveDown() {
	b.updateBoardBeforeChange()
	var snapShot [4][4]int
	copy(snapShot[:], b.matrix[:])

	var allDeltas []animations.MoveDelta
	var newMatrix [4][4]int

	transpose(&snapShot)
	for rowIndex, row := range snapShot {
		reverseRow(&row)
		slideRow, d1 := compactRow(rowIndex, row, true)

		mergedRow, scoreGain := mergeRow(slideRow)
		b.game.score += scoreGain

		finalRow, _ := compactRow(rowIndex, mergedRow, false)

		reverseRow(&finalRow)

		// Write to new matrix
		newMatrix[rowIndex] = finalRow

		// Collect all deltas
		allDeltas = append(allDeltas, d1...)

	}
	transpose(&newMatrix)
	b.matrix = newMatrix
	b.game.animation.Play(allDeltas, "DOWN")

	b.addNewRandomPieceIfBoardChanged()
	b.game.gameOver = b.isGameOver()
}

func reverseRow(row *[co.BOARDSIZE]int) {
	for i, j := 0, len(*row)-1; i < j; i, j = i+1, j-1 {
		(*row)[i], (*row)[j] = (*row)[j], (*row)[i]
	}
}

// Swap cols and rows
func transpose(board *[co.BOARDSIZE][co.BOARDSIZE]int) {
	for i := range len(*board) {
		for j := i; j < len((*board)[0]); j++ {
			(*board)[i][j], (*board)[j][i] = (*board)[j][i], (*board)[i][j]
		}
	}
}

func compactRow(rowIndex int, row [4]int, applyExtra bool) (newRow [4]int, deltas []animations.MoveDelta) {
	insertPos, lastVal, extraMov := 0, -1, 0

	for col, val := range row {
		if val == 0 {
			continue
		}
		if val == lastVal {
			extraMov++
		}

		// Record how far it will slide
		if applyExtra {
			delta := animations.MoveDelta{
				FromRow:    rowIndex,
				FromCol:    col,
				ToRow:      rowIndex,
				ToCol:      insertPos,
				ValueMoved: val,
			}

			if extraMov > 0 {
				delta.ToCol -= extraMov
			}
			// Append the delta to deltas for animation
			deltas = append(deltas, delta)
		}

		newRow[insertPos] = val
		insertPos++
		lastVal = val
	}

	return newRow, deltas
}

func mergeRow(row [4]int) (newRow [4]int, scoreGain int) {
	copy(newRow[:], row[:]) // Copy row info over in new row

	for i := range 3 {
		if newRow[i] != 0 && newRow[i] == newRow[i+1] {
			newRow[i] *= 2         // Update number
			scoreGain += newRow[i] // Update score
			newRow[i+1] = 0        // Remove value that was merged into current val
			i++                    // Skip to next possible val spot
		}
	}
	return newRow, scoreGain
}
