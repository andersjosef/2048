package twenty48

import (
	"testing"

	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/stretchr/testify/assert"
)

func TestReset(t *testing.T) {
	game, err := NewGame()
	assert.NoError(t, err)
	game.board.randomNewPiece()
	game.board.randomNewPiece()
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

func TestIsGameOver(t *testing.T) {
	var tests = []struct {
		name  string
		board [co.BOARDSIZE][co.BOARDSIZE]int
		want  bool
	}{
		{
			name: "Empty",
			board: [co.BOARDSIZE][co.BOARDSIZE]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			want: false,
		},
		{
			name: "Full",
			board: [co.BOARDSIZE][co.BOARDSIZE]int{
				{2, 4, 2, 4},
				{4, 2, 4, 2},
				{2, 4, 2, 4},
				{4, 2, 4, 2},
			},
			want: true,
		},
		{
			name: "Only UP/DOWN",
			board: [co.BOARDSIZE][co.BOARDSIZE]int{
				{2, 4, 2, 4},
				{4, 2, 4, 2},
				{8, 4, 2, 4},
				{8, 2, 4, 2},
			},
			want: false,
		},
		{
			name: "Only RIGHT/LEFT",
			board: [co.BOARDSIZE][co.BOARDSIZE]int{
				{2, 4, 2, 4},
				{4, 2, 4, 2},
				{2048, 2048, 2, 4},
				{4, 2, 4, 2},
			},
			want: false,
		},
		{
			name: "X",
			board: [co.BOARDSIZE][co.BOARDSIZE]int{
				{2, 0, 0, 4},
				{0, 2, 4, 0},
				{0, 4, 2, 0},
				{4, 0, 0, 2},
			},
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			game, err := NewGame()
			assert.NoError(t, err)

			game.board.matrix = tc.board
			assert.Equal(t, tc.want, game.board.isGameOver())
		})
	}
}
