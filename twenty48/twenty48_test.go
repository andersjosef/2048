package twenty48

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyBoard(t *testing.T) {
	board := Board{}
	want := [BOARDSIZE][BOARDSIZE]int{}

	assert.Equal(t, want, board.board)
}

func TestInitializeGame(t *testing.T) {
	game, err := NewGame()
	count := 0

	// counts the number of pieces on the board
	for x := 0; x < len(game.board.board); x++ {
		for y := 0; y < len(game.board.board[0]); y++ {
			if game.board.board[x][y] != 0 {
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
	game, err := NewGame()

	assert.NoError(t, err)

	game.board.board = [BOARDSIZE][BOARDSIZE]int{
		{2, 2, 0, 0},
		{0, 2, 0, 0},
		{0, 0, 2, 0},
		{0, 0, 0, 2},
	}
	want := [BOARDSIZE][BOARDSIZE]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{2, 4, 2, 2},
	}

	game.board.moveDown()
	assert.Equal(t, want, game.board.board)

	game.board.board = [BOARDSIZE][BOARDSIZE]int{
		{2, 2, 2, 0},
		{0, 2, 0, 0},
		{0, 2, 2, 0},
		{0, 0, 0, 2},
	}
	want = [BOARDSIZE][BOARDSIZE]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 2, 0, 0},
		{2, 4, 4, 2},
	}
	game.board.moveDown()
	assert.Equal(t, want, game.board.board)

}

func TestAddNewRandomPieceIfBoardChanged(t *testing.T) {
	game, err := NewGame()

	assert.NoError(t, err)
	game.board.board = [BOARDSIZE][BOARDSIZE]int{
		{2, 2, 2, 0},
		{0, 2, 0, 0},
		{0, 2, 2, 0},
		{0, 0, 0, 2},
	}
	game.board.boardBeforeChange = game.board.board
	game.board.moveDown()
	game.board.addNewRandomPieceIfBoardChanged()
	count := 0
	for x := 0; x < len(game.board.board); x++ {
		for y := 0; y < len(game.board.board[0]); y++ {
			if game.board.board[x][y] != 0 {
				count++
			}
		}
	}
	assert.Equal(t, 6, count)

	game.board.board = [BOARDSIZE][BOARDSIZE]int{
		{2, 2, 2, 2},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	game.board.boardBeforeChange = game.board.board
	game.board.moveUp()
	game.board.addNewRandomPieceIfBoardChanged()
	count = 0
	for x := 0; x < len(game.board.board); x++ {
		for y := 0; y < len(game.board.board[0]); y++ {
			if game.board.board[x][y] != 0 {
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
	game.board.ResetGame()

	assert.Equal(t, 0, game.score)
}

func TestScore(t *testing.T) {
	game, err := NewGame()
	assert.NoError(t, err)

	game.board.board = [BOARDSIZE][BOARDSIZE]int{
		{8, 16, 2, 0},
		{0, 0, 2, 0},
		{0, 0, 0, 0},
		{8, 2, 0, 0},
	}
	game.board.moveUp()
	assert.Equal(t, 20, game.score)
	game.board.moveLeft()
	assert.Equal(t, 52, game.score)

}

func TestFullBoard(t *testing.T) {
	game, err := NewGame()
	assert.NoError(t, err)

	game.board.board = [BOARDSIZE][BOARDSIZE]int{
		{2, 8, 2, 8},
		{8, 2, 8, 2},
		{2, 8, 2, 8},
		{8, 2, 8, 2},
	}
	game.board.boardBeforeChange = game.board.board

	want := game.board.board

	game.board.moveDown()
	game.board.addNewRandomPieceIfBoardChanged()

	assert.Equal(t, want, game.board.board)

	game.board.moveUp()
	game.board.addNewRandomPieceIfBoardChanged()

	assert.Equal(t, want, game.board.board)

	game.board.moveLeft()
	game.board.addNewRandomPieceIfBoardChanged()

	assert.Equal(t, want, game.board.board)

	game.board.moveRight()
	game.board.addNewRandomPieceIfBoardChanged()

	assert.Equal(t, want, game.board.board)
}

func TestMoves(t *testing.T) {
	var tests = []struct {
		name         string
		initialBoard [BOARDSIZE][BOARDSIZE]int
		want         [BOARDSIZE][BOARDSIZE]int
		moveFunc     func(*Board)
		wantedScore  int
	}{
		{
			name: "Move Left 1",
			initialBoard: [BOARDSIZE][BOARDSIZE]int{
				{2, 2, 0, 0},
				{0, 2, 0, 0},
				{0, 0, 2, 0},
				{0, 0, 0, 2},
			},
			want: [BOARDSIZE][BOARDSIZE]int{
				{4, 0, 0, 0},
				{2, 0, 0, 0},
				{2, 0, 0, 0},
				{2, 0, 0, 0},
			},
			moveFunc:    func(b *Board) { b.moveLeft() },
			wantedScore: 4,
		},
		{
			name: "Move Left 2",
			initialBoard: [BOARDSIZE][BOARDSIZE]int{
				{2, 2, 2, 0},
				{0, 2, 0, 0},
				{0, 0, 2, 0},
				{0, 0, 0, 2},
			},
			want: [BOARDSIZE][BOARDSIZE]int{
				{4, 2, 0, 0},
				{2, 0, 0, 0},
				{2, 0, 0, 0},
				{2, 0, 0, 0},
			},
			moveFunc:    func(b *Board) { b.moveLeft() },
			wantedScore: 4,
		},
		{
			name: "Move right 1",
			initialBoard: [BOARDSIZE][BOARDSIZE]int{
				{2, 2, 0, 0},
				{0, 2, 0, 0},
				{0, 0, 2, 0},
				{0, 0, 0, 2},
			},
			want: [BOARDSIZE][BOARDSIZE]int{
				{0, 0, 0, 4},
				{0, 0, 0, 2},
				{0, 0, 0, 2},
				{0, 0, 0, 2},
			},
			moveFunc:    func(b *Board) { b.moveRight() },
			wantedScore: 4,
		},
		{
			name: "Move Right 2",
			initialBoard: [BOARDSIZE][BOARDSIZE]int{
				{2, 2, 2, 0},
				{0, 2, 0, 0},
				{0, 2, 2, 0},
				{0, 0, 0, 2},
			},
			want: [BOARDSIZE][BOARDSIZE]int{
				{0, 0, 2, 4},
				{0, 0, 0, 2},
				{0, 0, 0, 4},
				{0, 0, 0, 2},
			},
			moveFunc:    func(b *Board) { b.moveRight() },
			wantedScore: 8,
		},
		{
			name: "Move Up 1",
			initialBoard: [BOARDSIZE][BOARDSIZE]int{
				{2, 2, 0, 0},
				{0, 2, 0, 0},
				{0, 0, 2, 0},
				{0, 0, 0, 2},
			},
			want: [BOARDSIZE][BOARDSIZE]int{
				{2, 4, 2, 2},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			moveFunc:    func(b *Board) { b.moveUp() },
			wantedScore: 4,
		},
		{
			name: "Move Up 2",
			initialBoard: [BOARDSIZE][BOARDSIZE]int{
				{2, 2, 2, 0},
				{0, 2, 0, 0},
				{0, 2, 2, 0},
				{0, 0, 0, 2},
			},
			want: [BOARDSIZE][BOARDSIZE]int{
				{2, 4, 4, 2},
				{0, 2, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			moveFunc:    func(b *Board) { b.moveUp() },
			wantedScore: 8,
		},
		{
			name: "Move Down 1",
			initialBoard: [BOARDSIZE][BOARDSIZE]int{
				{2, 2, 0, 0},
				{0, 2, 0, 0},
				{0, 0, 2, 0},
				{0, 0, 0, 2},
			},
			want: [BOARDSIZE][BOARDSIZE]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{2, 4, 2, 2},
			},
			moveFunc:    func(b *Board) { b.moveDown() },
			wantedScore: 4,
		},
		{
			name: "Move Down 2",
			initialBoard: [BOARDSIZE][BOARDSIZE]int{
				{2, 2, 2, 0},
				{0, 2, 0, 0},
				{0, 2, 2, 0},
				{0, 0, 0, 2},
			},
			want: [BOARDSIZE][BOARDSIZE]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 2, 0, 0},
				{2, 4, 4, 2},
			},
			moveFunc:    func(b *Board) { b.moveDown() },
			wantedScore: 8,
		},
	}

	// creates a test for every entry in the list above
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			game, err := NewGame()
			assert.NoError(t, err)
			game.board.board = tc.initialBoard
			tc.moveFunc(game.board)
			assert.Equal(t, tc.want, game.board.board)
			assert.Equal(t, tc.wantedScore, game.score)
		})
	}
}

func TestInitAnimation(t *testing.T) {
	g := &Game{}
	a := InitAnimation(g)

	wantArrayOfChange := [BOARDSIZE][BOARDSIZE]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	assert.Equal(t, false, a.isAnimating)
	assert.Equal(t, wantArrayOfChange, a.arrayOfChange)

}

func TestResetAnimation(t *testing.T) {
	g := &Game{}
	a := InitAnimation(g)

	a.arrayOfChange = [BOARDSIZE][BOARDSIZE]int{
		{0, 1, 0, 0},
		{0, 0, 0, 3},
		{0, 0, 2, 0},
		{0, 0, 0, 0},
	}

	wantArrayOfChange := [BOARDSIZE][BOARDSIZE]int{}

	a.ResetArray()

	assert.Equal(t, false, a.isAnimating)
	assert.Equal(t, wantArrayOfChange, a.arrayOfChange)

}
