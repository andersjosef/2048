package twenty48

import (
	"github.com/andersjosef/2048/twenty48/animations"
)

type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
)

func (d *Direction) String() string {
	str := ""
	switch *d {
	case Left:
		str = "LEFT"
	case Up:
		str = "UP"
	case Right:
		str = "RIGHT"
	case Down:
		str = "DOWN"
	}
	return str
}

func (b *Board) updateBoardBeforeChange() {
	b.matrixBeforeChange = b.matrix
}

func (b *Board) move(dir Direction) {
	b.updateBoardBeforeChange()

	var snap [4][4]int
	copy(snap[:], b.matrix[:])
	apply := getTransform(dir)
	apply.pre(&snap) // Do pre matrix manipulation

	var newMat [4][4]int
	var allDeltas []animations.MoveDelta
	for rowIndex, row := range snap {
		newRow, d1, scoreGain := processRow(rowIndex, row)
		b.game.score += scoreGain
		allDeltas = append(allDeltas, d1...)
		newMat[rowIndex] = newRow
	}

	apply.post(&newMat)                            // Do post matrix manipulations
	b.game.animation.Play(allDeltas, dir.String()) // Trigger animation
	b.matrix = newMat                              // Set matrix to what has been manipulated
	b.addNewRandomPieceIfBoardChanged()
	b.game.gameOver = b.isGameOver()

}

// func (b *Board) moveLeft() {
// 	b.move(Left)
// }

// func (b *Board) moveUp() {
// 	b.move(Up)
// }

// func (b *Board) moveRight() {
// 	b.move(Right)
// }

// func (b *Board) moveDown() {
// 	b.move(Down)
// }

type transform struct {
	pre  func(*[4][4]int)
	post func(*[4][4]int)
}

func getTransform(dir Direction) transform {
	switch dir {
	case Left:
		return transform{noop, noop}
	case Up:
		return transform{transpose, transpose}
	case Right:
		return transform{reverseAllRows, reverseAllRows}
	case Down:
		return transform{
			func(m *[4][4]int) { transpose(m); reverseAllRows(m) },
			func(m *[4][4]int) { reverseAllRows(m); transpose(m) },
		}
	default:
		return transform{noop, noop}
	}
}

func noop(_ *[4][4]int) {}

func reverseAllRows(m *[4][4]int) {
	for i := range m {
		reverseRow(&m[i])
	}
}
