package board

import (
	"math/rand"
	"testing"

	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/eventhandler"
	"github.com/andersjosef/2048/twenty48/shared"
	"github.com/andersjosef/2048/twenty48/theme"
	"github.com/stretchr/testify/assert"
)

type MockEventHandler struct{}

func (MockEventHandler) Register(eventType eventhandler.EventType, handler func(eventhandler.Event)) {
}
func (MockEventHandler) Dispatch()                         {}
func (MockEventHandler) Emit(eventType eventhandler.Event) {}

type MockScreenControl struct{}

func (MockScreenControl) GetActualSize() (x, y int) { return 50, 50 }
func (MockScreenControl) ToggleFullScreen()         {}
func (MockScreenControl) IsFullScreen() bool        { return false }
func (MockScreenControl) IncrementScale()           {}
func (MockScreenControl) DecrementScale() bool      { return false }
func (MockScreenControl) GetScale() float64         { return 1 }

func TestEmptyBoard(t *testing.T) {
	board := Board{}
	want := [co.BOARDSIZE][co.BOARDSIZE]int{}

	assert.Equal(t, want, board.matrix)
}

func TestInitializeBoard(t *testing.T) {
	d := Deps{
		EventHandler:    MockEventHandler{},
		GetCurrentTheme: func() theme.Theme { return theme.Theme{} },
		ScreenControl:   MockScreenControl{},
	}
	count := 0
	sum := 0

	board, err := New(d)
	// counts the number of pieces on the board
	for x := range len(board.matrix) {
		for y := range len(board.matrix[0]) {
			val := board.matrix[x][y]
			if val != 0 {
				count++
				sum += val
			}
		}
	}

	assert.NoError(t, err)
	assert.GreaterOrEqual(t, sum, 0)
	assert.LessOrEqual(t, sum, 8)
	assert.Equal(t, 2, count)
}

// func TestAddNewRandomPieceIfBoardChanged(t *testing.T) {
// 	game, err := NewGame()

// 	assert.NoError(t, err)
// 	game.board.matrix = [co.BOARDSIZE][co.BOARDSIZE]int{
// 		{2, 2, 2, 0},
// 		{0, 2, 0, 0},
// 		{0, 2, 2, 0},
// 		{0, 0, 0, 2},
// 	}
// 	game.board.move(Down)
// 	count := 0
// 	for x := 0; x < len(game.board.matrix); x++ {
// 		for y := 0; y < len(game.board.matrix[0]); y++ {
// 			if game.board.matrix[x][y] != 0 {
// 				count++
// 			}
// 		}
// 	}
// 	assert.Equal(t, 6, count)

// 	game.board.matrix = [co.BOARDSIZE][co.BOARDSIZE]int{
// 		{2, 2, 2, 2},
// 		{0, 0, 0, 0},
// 		{0, 0, 0, 0},
// 		{0, 0, 0, 0},
// 	}
// 	game.board.matrixBeforeChange = game.board.matrix
// 	game.board.move(Up)
// 	game.board.addNewRandomPieceIfBoardChanged()
// 	count = 0
// 	for x := 0; x < len(game.board.matrix); x++ {
// 		for y := 0; y < len(game.board.matrix[0]); y++ {
// 			if game.board.matrix[x][y] != 0 {
// 				count++
// 			}
// 		}
// 	}
// 	assert.Equal(t, 4, count)
// }

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
			moveFunc:    func(b *Board) { b.Move(Left) },
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
			moveFunc:    func(b *Board) { b.Move(Left) },
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
			moveFunc:    func(b *Board) { b.Move(Right) },
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
			moveFunc:    func(b *Board) { b.Move(Right) },
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
			moveFunc:    func(b *Board) { b.Move(Up) },
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
			moveFunc:    func(b *Board) { b.Move(Up) },
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
			moveFunc:    func(b *Board) { b.Move(Down) },
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
			moveFunc:    func(b *Board) { b.Move(Down) },
			wantedScore: 8,
		},
	}

	// creates a test for every entry in the list above
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			eventHandler := eventhandler.NewEventBus()

			// Listen to score given
			gotScore := 0
			eventHandler.Register(eventhandler.EventMoveMade, func(e eventhandler.Event) {
				data, ok := e.Data.(shared.MoveData)
				if !ok {
					return
				}

				gotScore += data.ScoreGain
			})

			d := Deps{
				EventHandler:    eventHandler,
				GetCurrentTheme: func() theme.Theme { return theme.Theme{} },
				ScreenControl:   MockScreenControl{},
				SetGameOver:     func(_ bool) {},
			}
			board, err := New(d)
			board.matrix = tc.initialBoard
			tc.moveFunc(board)

			assert.NoError(t, err)
			assert.Equal(t, tc.want, board.matrix)
			assert.Equal(t, tc.wantedScore, gotScore)
		})
	}
}
