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
