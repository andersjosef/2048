package board

import (
	"math/rand"

	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/eventhandler"
	"github.com/andersjosef/2048/twenty48/shared"
)

type Board struct {
	matrix [co.BOARDSIZE][co.BOARDSIZE]int
	d      Deps
}

func New(d Deps) (*Board, error) {

	b := &Board{
		d: d,
	}

	// Add the two start pieces
	for range 2 {
		b.randomNewPiece()
	}

	b.registerEvents()

	return b, nil
}

func (b *Board) registerEvents() {
	b.d.Register(
		eventhandler.EventResetGame,
		func(_ eventhandler.Event) {
			b.matrix = [co.BOARDSIZE][co.BOARDSIZE]int{}
			b.randomNewPiece()
			b.randomNewPiece()

		},
	)
	b.d.Register(
		eventhandler.EventMoveMade,
		func(e eventhandler.Event) {
			data, ok := e.Data.(shared.MoveData)
			if !ok {
				return
			}
			b.d.Core.AddScore(data.ScoreGain)
			b.updateMatrix(data.NewBoard)
			b.d.SetGameOver(b.isGameOver())
		},
	)
}

func (b *Board) randomNewPiece() {
	x, y := b.GetBoardDimentions()

	// TODO: fix flawed logic, there is a much larger chance for a tile after a big block
	// Will start at a random position, then check every available spot after
	// until all tiles are checked
	for count := rand.Intn(x * y); count < count+x*y-1; count++ {
		posX := count % x
		posY := (count / y) % y
		if b.matrix[posX][posY] == 0 {
			if rand.Float32() > 0.16 {
				b.matrix[posX][posY] = 2 // 84%
			} else {
				b.matrix[posX][posY] = 4 // 16% chance of 4 spawning
			}
			break
		}
	}
}

// Set new matrix and add a random new piece
func (b *Board) updateMatrix(newBoard [co.BOARDSIZE][co.BOARDSIZE]int) {
	if b.matrix != newBoard {
		b.matrix = newBoard
		b.randomNewPiece()
	}
}

// Check if its gameOver
func (b *Board) isGameOver() bool {
	// Check if there are any empty spaces left, meaning its possible to play
	length, height := b.GetBoardDimentions()

	// Check for empty tiles
	for i := range height {
		for j := range length {
			if b.matrix[i][j] == 0 {
				return false
			}
		}
	}

	// Check for vertical merges
	for i := range height - 1 {
		for j := range length {
			if b.matrix[i][j] == b.matrix[i+1][j] {
				return false
			}
		}
	}

	// Check for horisontal merges
	for i := range height {
		for j := range length - 1 {
			if b.matrix[i][j] == b.matrix[i][j+1] {
				return false
			}
		}
	}
	return true

}
