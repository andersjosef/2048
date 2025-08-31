package board

import (
	"testing"

	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/eventhandler"
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
	board, err := New(d)
	count := 0

	// counts the number of pieces on the board
	for x := range len(board.matrix) {
		for y := range len(board.matrix[0]) {
			if board.matrix[x][y] != 0 {
				count++
			}
		}
	}
	assert.NoError(t, err)

	// assert.Equal(t, 2, int(game.state))
	// assert.Equal(t, 0, game.score)
	assert.Equal(t, 2, count)
}
