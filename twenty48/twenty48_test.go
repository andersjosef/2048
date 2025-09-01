package twenty48

import (
	"testing"

	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/stretchr/testify/assert"
)

func TestReset(t *testing.T) {
	game, err := NewGame()

	game.score = 1000

	assert.NoError(t, err)
	ResetGame(game.input)

	assert.Equal(t, 0, game.score)
}

func TestScore(t *testing.T) {
	game, err := NewGame()
	assert.NoError(t, err)

	game.board.matrix = [co.BOARDSIZE][co.BOARDSIZE]int{
		{8, 16, 2, 0},
		{0, 0, 2, 0},
		{0, 0, 0, 0},
		{8, 2, 0, 0},
	}
	game.board.move(Up)
	assert.Equal(t, 20, game.score)
	game.board.move(Left)
	assert.Equal(t, 52, game.score)

}

func TestFullBoard(t *testing.T) {
	game, err := NewGame()
	assert.NoError(t, err)

	game.board.matrix = [co.BOARDSIZE][co.BOARDSIZE]int{
		{2, 8, 2, 8},
		{8, 2, 8, 2},
		{2, 8, 2, 8},
		{8, 2, 8, 2},
	}
	game.board.matrixBeforeChange = game.board.matrix

	want := game.board.matrix

	game.board.move(Down)
	game.board.addNewRandomPieceIfBoardChanged()

	assert.Equal(t, want, game.board.matrix)

	game.board.move(Up)
	game.board.addNewRandomPieceIfBoardChanged()

	assert.Equal(t, want, game.board.matrix)

	game.board.move(Left)
	game.board.addNewRandomPieceIfBoardChanged()

	assert.Equal(t, want, game.board.matrix)

	game.board.move(Right)
	game.board.addNewRandomPieceIfBoardChanged()

	assert.Equal(t, want, game.board.matrix)
}
