package twenty48

import (
	"testing"

	"math/rand"

	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/stretchr/testify/assert"
)

func TestEmptyBoard(t *testing.T) {
	board := Board{}
	want := [co.BOARDSIZE][co.BOARDSIZE]int{}

	assert.Equal(t, want, board.matrix)
}

func TestInitializeGame(t *testing.T) {
	game, err := NewGame()
	count := 0

	// counts the number of pieces on the board
	for x := 0; x < len(game.board.matrix); x++ {
		for y := 0; y < len(game.board.matrix[0]); y++ {
			if game.board.matrix[x][y] != 0 {
				count++
			}
		}
	}
	assert.NoError(t, err)

	assert.Equal(t, game, game.board.game)

	assert.Equal(t, 2, int(game.state))
	assert.Equal(t, 0, game.score)
	assert.Equal(t, 2, count)
}

func TestMoveDown(t *testing.T) {
	rand.Seed(42)
	game, err := NewGame()

	assert.NoError(t, err)

	game.board.matrix = [co.BOARDSIZE][co.BOARDSIZE]int{
		{2, 2, 0, 0},
		{0, 2, 0, 0},
		{0, 0, 2, 0},
		{0, 0, 0, 2},
	}
	want := [co.BOARDSIZE][co.BOARDSIZE]int{
		{2, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{2, 4, 2, 2},
	}

	game.board.move(Down)
	assert.Equal(t, want, game.board.matrix)

	game.board.matrix = [co.BOARDSIZE][co.BOARDSIZE]int{
		{2, 2, 2, 0},
		{0, 2, 0, 0},
		{0, 2, 2, 0},
		{0, 0, 0, 2},
	}
	want = [co.BOARDSIZE][co.BOARDSIZE]int{
		{0, 0, 0, 0},
		{0, 2, 0, 0},
		{0, 2, 0, 0},
		{2, 4, 4, 2},
	}
	game.board.move(Down)
	assert.Equal(t, want, game.board.matrix)

}

func TestAddNewRandomPieceIfBoardChanged(t *testing.T) {
	game, err := NewGame()

	assert.NoError(t, err)
	game.board.matrix = [co.BOARDSIZE][co.BOARDSIZE]int{
		{2, 2, 2, 0},
		{0, 2, 0, 0},
		{0, 2, 2, 0},
		{0, 0, 0, 2},
	}
	game.board.move(Down)
	count := 0
	for x := 0; x < len(game.board.matrix); x++ {
		for y := 0; y < len(game.board.matrix[0]); y++ {
			if game.board.matrix[x][y] != 0 {
				count++
			}
		}
	}
	assert.Equal(t, 6, count)

	game.board.matrix = [co.BOARDSIZE][co.BOARDSIZE]int{
		{2, 2, 2, 2},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	game.board.matrixBeforeChange = game.board.matrix
	game.board.move(Up)
	game.board.addNewRandomPieceIfBoardChanged()
	count = 0
	for x := 0; x < len(game.board.matrix); x++ {
		for y := 0; y < len(game.board.matrix[0]); y++ {
			if game.board.matrix[x][y] != 0 {
				count++
			}
		}
	}
	assert.Equal(t, 4, count)
}

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

func TestMoves(t *testing.T) {
	rand.Seed(42)
	var tests = []struct {
		name         string
		initialBoard [co.BOARDSIZE][co.BOARDSIZE]int
		want         [co.BOARDSIZE][co.BOARDSIZE]int
		moveFunc     func(*Board)
		wantedScore  int
	}{
		{
			name: "Move Left 1",
			initialBoard: [co.BOARDSIZE][co.BOARDSIZE]int{
				{2, 2, 0, 0},
				{0, 2, 0, 0},
				{0, 0, 2, 0},
				{0, 0, 0, 2},
			},
			want: [co.BOARDSIZE][co.BOARDSIZE]int{
				{4, 0, 0, 0},
				{2, 0, 0, 0},
				{2, 0, 0, 0},
				{2, 0, 0, 2},
			},
			moveFunc:    func(b *Board) { b.move(Left) },
			wantedScore: 4,
		},
		{
			name: "Move Left 2",
			initialBoard: [co.BOARDSIZE][co.BOARDSIZE]int{
				{2, 2, 2, 0},
				{0, 2, 0, 0},
				{0, 0, 2, 0},
				{0, 0, 0, 2},
			},
			want: [co.BOARDSIZE][co.BOARDSIZE]int{
				{4, 2, 0, 0},
				{2, 0, 2, 0},
				{2, 0, 0, 0},
				{2, 0, 0, 0},
			},
			moveFunc:    func(b *Board) { b.move(Left) },
			wantedScore: 4,
		},
		{
			name: "Move right 1",
			initialBoard: [co.BOARDSIZE][co.BOARDSIZE]int{
				{2, 2, 0, 0},
				{0, 2, 0, 0},
				{0, 0, 2, 0},
				{0, 0, 0, 2},
			},
			want: [co.BOARDSIZE][co.BOARDSIZE]int{
				{0, 2, 0, 4},
				{0, 0, 0, 2},
				{0, 0, 0, 2},
				{0, 0, 0, 2},
			},
			moveFunc:    func(b *Board) { b.move(Right) },
			wantedScore: 4,
		},
		{
			name: "Move Right 2",
			initialBoard: [co.BOARDSIZE][co.BOARDSIZE]int{
				{2, 2, 2, 0},
				{0, 2, 0, 0},
				{0, 2, 2, 0},
				{0, 0, 0, 2},
			},
			want: [co.BOARDSIZE][co.BOARDSIZE]int{
				{0, 2, 2, 4},
				{0, 0, 0, 2},
				{0, 0, 0, 4},
				{0, 0, 0, 2},
			},
			moveFunc:    func(b *Board) { b.move(Right) },
			wantedScore: 8,
		},
		{
			name: "Move Up 1",
			initialBoard: [co.BOARDSIZE][co.BOARDSIZE]int{
				{2, 2, 0, 0},
				{0, 2, 0, 0},
				{0, 0, 2, 0},
				{0, 0, 0, 2},
			},
			want: [co.BOARDSIZE][co.BOARDSIZE]int{
				{2, 4, 2, 2},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 2, 0},
			},
			moveFunc:    func(b *Board) { b.move(Up) },
			wantedScore: 4,
		},
		{
			name: "Move Up 2",
			initialBoard: [co.BOARDSIZE][co.BOARDSIZE]int{
				{2, 2, 2, 0},
				{0, 2, 0, 0},
				{0, 2, 2, 0},
				{0, 0, 0, 2},
			},
			want: [co.BOARDSIZE][co.BOARDSIZE]int{
				{2, 4, 4, 2},
				{0, 2, 0, 2},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			moveFunc:    func(b *Board) { b.move(Up) },
			wantedScore: 8,
		},
		{
			name: "Move Down 1",
			initialBoard: [co.BOARDSIZE][co.BOARDSIZE]int{
				{2, 2, 0, 0},
				{0, 2, 0, 0},
				{0, 0, 2, 0},
				{0, 0, 0, 2},
			},
			want: [co.BOARDSIZE][co.BOARDSIZE]int{
				{0, 0, 0, 2},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{2, 4, 2, 2},
			},
			moveFunc:    func(b *Board) { b.move(Down) },
			wantedScore: 4,
		},
		{
			name: "Move Down 2",
			initialBoard: [co.BOARDSIZE][co.BOARDSIZE]int{
				{2, 2, 2, 0},
				{0, 2, 0, 0},
				{0, 2, 2, 0},
				{0, 0, 0, 2},
			},
			want: [co.BOARDSIZE][co.BOARDSIZE]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{2, 2, 0, 0},
				{2, 4, 4, 2},
			},
			moveFunc:    func(b *Board) { b.move(Down) },
			wantedScore: 8,
		},
	}

	// creates a test for every entry in the list above
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			game, err := NewGame()
			assert.NoError(t, err)
			game.board.matrix = tc.initialBoard
			tc.moveFunc(game.board)
			assert.Equal(t, tc.want, game.board.matrix)
			assert.Equal(t, tc.wantedScore, game.score)
		})
	}
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
