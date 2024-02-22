package twenty48

import (
	"testing"
)

func TestEmptyBoard(t *testing.T) {
	board := Board{}
	want := [4][4]int{}

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

	board.board = [4][4]int{
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

}

func TestMoveUp(t *testing.T) {
	board := Board{}

	board.board = [4][4]int{
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

}
