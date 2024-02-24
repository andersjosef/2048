package twenty48

import (
	"testing"
)

func TestEmptyBoard(t *testing.T) {
	board := Board{}
	want := [BOARDSIZE][BOARDSIZE]int{}

	if board.board != want {
		t.Fatalf(`board.board = %v, want %v error`, board.board, want)
	}

}

func TestAddTwoRandom(t *testing.T) {
	board := Board{}
	count := 0
	for i := 0; i < 2; i++ {
		board.randomNewPiece()
	}
	for x := 0; x < len(board.board); x++ {
		for y := 0; y < len(board.board[0]); y++ {
			if board.board[x][y] != 0 {
				count++
			}
		}
	}
	if count < 2 {
		t.Fatalf(`less than two pieces are changed, count = %v, want 2. board = %v error`, count, board.board)
	} else if count > 2 {
		t.Fatalf(`less than two pieces are changed, count = %v, want 2. board = %v error`, count, board.board)
	}
}

func TestMoveLeft(t *testing.T) {
	board := Board{}

	board.board = [BOARDSIZE][BOARDSIZE]int{
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

	board.moveLeft()
	if board.board != want {
		t.Fatalf(`board.board = %v, want %v error`, board.board, want)
	}

	board.board = [BOARDSIZE][BOARDSIZE]int{
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
	board.moveLeft()
	if board.board != want {
		t.Fatalf(`board.board = %v, want %v error`, board.board, want)
	}

}

func TestMoveUp(t *testing.T) {
	board := Board{}

	board.board = [BOARDSIZE][BOARDSIZE]int{
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

	board.moveUp()
	if board.board != want {
		t.Fatalf(`board.board = %v, want %v error`, board.board, want)
	}
	board.board = [BOARDSIZE][BOARDSIZE]int{
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
	board.moveUp()
	if board.board != want {
		t.Fatalf(`board.board = %v, want %v error`, board.board, want)
	}

}

func TestMoveRight(t *testing.T) {
	board := Board{}

	board.board = [BOARDSIZE][BOARDSIZE]int{
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

	board.moveRight()
	if board.board != want {
		t.Fatalf(`board.board = %v, want %v error`, board.board, want)
	}
	board.board = [BOARDSIZE][BOARDSIZE]int{
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
	board.moveRight()
	if board.board != want {
		t.Fatalf(`board.board = %v, want %v error`, board.board, want)
	}

}

func TestMoveDown(t *testing.T) {
	board := Board{}

	board.board = [BOARDSIZE][BOARDSIZE]int{
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

	board.moveDown()
	if board.board != want {
		t.Fatalf(`board.board = %v, want %v error`, board.board, want)
	}

	board.board = [BOARDSIZE][BOARDSIZE]int{
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
	board.moveDown()
	if board.board != want {
		t.Fatalf(`board.board = %v, want %v error`, board.board, want)
	}

}

func TestAddNewRandomPieceIfBoardChanged(t *testing.T) {
	board := Board{}
	board.board = [BOARDSIZE][BOARDSIZE]int{
		{2, 2, 2, 0},
		{0, 2, 0, 0},
		{0, 2, 2, 0},
		{0, 0, 0, 2},
	}
	board_before_change := board.board
	board.moveDown()
	board.addNewRandomPieceIfBoardChanged(board_before_change)
	count := 0
	for x := 0; x < len(board.board); x++ {
		for y := 0; y < len(board.board[0]); y++ {
			if board.board[x][y] != 0 {
				count++
			}
		}
	}

	if count < 6 {
		t.Fatalf(`less than 6 pieces are changed, count = %v, want 2. board = %v error`, count, board.board)
	} else if count > 6 {
		t.Fatalf(`less than 6 pieces are changed, count = %v, want 2. board = %v error`, count, board.board)
	}
	board.board = [BOARDSIZE][BOARDSIZE]int{
		{2, 2, 2, 2},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	board_before_change = board.board
	board.moveUp()
	board.addNewRandomPieceIfBoardChanged(board_before_change)
	count = 0
	for x := 0; x < len(board.board); x++ {
		for y := 0; y < len(board.board[0]); y++ {
			if board.board[x][y] != 0 {
				count++
			}
		}
	}
	if count < 4 {
		t.Fatalf(`less than 4 pieces are changed, count = %v, want 2. board = %v error`, count, board.board)
	} else if count > 4 {
		t.Fatalf(`less than 4 pieces are changed, count = %v, want 2. board = %v error`, count, board.board)
	}
}

func TestReset(t *testing.T) {
	game, _ := NewGame()
	game.board.randomNewPiece()
	game.board.randomNewPiece()
	game.board.ResetGame()

	if game.score != 0 {
		t.Fatalf(`score is not zero, score = %v, want 0. board = %v error`, game.score, game.board.board)
	}
}
