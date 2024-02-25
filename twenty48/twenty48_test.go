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
	assert.Equal(t, 2, count)
}

func TestMoveLeft(t *testing.T) {
	game, err := NewGame()
	assert.NoError(t, err)

	game.board.board = [BOARDSIZE][BOARDSIZE]int{
		{2, 2, 0, 0},
		{0, 2, 0, 0},
		{0, 0, 2, 0},
		{0, 0, 0, 2},
	}
	want := [4][4]int{
		{4, 0, 0, 0},
		{2, 0, 0, 0},
		{2, 0, 0, 0},
		{2, 0, 0, 0},
	}

	game.board.moveLeft()
	assert.Equal(t, want, game.board.board)

	game.board.board = [BOARDSIZE][BOARDSIZE]int{
		{2, 2, 2, 0},
		{0, 2, 0, 0},
		{0, 0, 2, 0},
		{0, 0, 0, 2},
	}
	want = [BOARDSIZE][BOARDSIZE]int{
		{4, 2, 0, 0},
		{2, 0, 0, 0},
		{2, 0, 0, 0},
		{2, 0, 0, 0},
	}
	game.board.moveLeft()
	assert.Equal(t, want, game.board.board)

}

func TestMoveUp(t *testing.T) {
	game, err := NewGame()

	assert.NoError(t, err)

	game.board.board = [BOARDSIZE][BOARDSIZE]int{
		{2, 2, 0, 0},
		{0, 2, 0, 0},
		{0, 0, 2, 0},
		{0, 0, 0, 2},
	}
	want := [4][4]int{
		{2, 4, 2, 2},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	game.board.moveUp()
	assert.Equal(t, want, game.board.board)

	game.board.board = [BOARDSIZE][BOARDSIZE]int{
		{2, 2, 2, 0},
		{0, 2, 0, 0},
		{0, 2, 2, 0},
		{0, 0, 0, 2},
	}
	want = [BOARDSIZE][BOARDSIZE]int{
		{2, 4, 4, 2},
		{0, 2, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	game.board.moveUp()
	assert.Equal(t, want, game.board.board)

}

func TestMoveRight(t *testing.T) {
	game, err := NewGame()

	assert.NoError(t, err)

	game.board.board = [BOARDSIZE][BOARDSIZE]int{
		{2, 2, 0, 0},
		{0, 2, 0, 0},
		{0, 0, 2, 0},
		{0, 0, 0, 2},
	}
	want := [BOARDSIZE][BOARDSIZE]int{
		{0, 0, 0, 4},
		{0, 0, 0, 2},
		{0, 0, 0, 2},
		{0, 0, 0, 2},
	}

	game.board.moveRight()
	assert.Equal(t, want, game.board.board)

	game.board.board = [BOARDSIZE][BOARDSIZE]int{
		{2, 2, 2, 0},
		{0, 2, 0, 0},
		{0, 2, 2, 0},
		{0, 0, 0, 2},
	}
	want = [BOARDSIZE][BOARDSIZE]int{
		{0, 0, 2, 4},
		{0, 0, 0, 2},
		{0, 0, 0, 4},
		{0, 0, 0, 2},
	}
	game.board.moveRight()
	assert.Equal(t, want, game.board.board)

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
	board_before_change := game.board.board
	game.board.moveDown()
	game.board.addNewRandomPieceIfBoardChanged(board_before_change)
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
	board_before_change = game.board.board
	game.board.moveUp()
	game.board.addNewRandomPieceIfBoardChanged(board_before_change)
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
	early := game.board.board

	want := game.board.board

	game.board.moveDown()
	game.board.addNewRandomPieceIfBoardChanged(early)

	assert.Equal(t, want, game.board.board)

	game.board.moveUp()
	game.board.addNewRandomPieceIfBoardChanged(early)

	assert.Equal(t, want, game.board.board)

	game.board.moveLeft()
	game.board.addNewRandomPieceIfBoardChanged(early)

	assert.Equal(t, want, game.board.board)

	game.board.moveRight()
	game.board.addNewRandomPieceIfBoardChanged(early)

	assert.Equal(t, want, game.board.board)
}
