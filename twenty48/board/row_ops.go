package board

import (
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/shared"
)

func processRow(rowIndex int, row [4]int) (outRow [4]int, deltas []shared.MoveDelta, scoreGain int) {
	compacted, deltas := compactRow(rowIndex, row, true)
	merged, scoreGain := mergeRow(compacted)
	outRow, _ = compactRow(rowIndex, merged, false)
	return
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

func compactRow(rowIndex int, row [4]int, applyExtra bool) (newRow [4]int, deltas []shared.MoveDelta) {
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
			delta := shared.MoveDelta{
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
