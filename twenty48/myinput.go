package twenty48

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type MyInput struct {
	keys              []ebiten.Key
	keyIsBeingPressed bool
}

var m = &MyInput{
	keyIsBeingPressed: false,
}

// this is also the game logic I guess
func (m *MyInput) UpdateInput(b *Board) error {
	m.keys = inpututil.AppendPressedKeys(m.keys[:0])
	if len(m.keys) > 0 {
		if !m.keyIsBeingPressed {
			m.keyIsBeingPressed = true
			key_pressed := m.keys[len(m.keys)-1]
			var board_before_change [BOARDSIZE][BOARDSIZE]int = b.board
			// fmt.Println(key_pressed)
			switch b.game.state {
			case 1:
				switch fmt.Sprintf("%v", key_pressed) {
				case "D", "ArrowRight": // right@
					b.moveRight()
					// fmt.Println("right")
				case "A", "ArrowLeft": // left
					b.moveLeft()
					// fmt.Println("left")
				case "W", "ArrowUp":
					b.moveUp()
					// fmt.Println("up")
				case "S", "ArrowDown":
					b.moveDown()
					// fmt.Println("down")
				case "R": // reset button
					b.ResetGame()
				}
				b.addNewRandomPieceIfBoardChanged(board_before_change)
			case 2: // menu
				if fmt.Sprintf("%v", key_pressed) != "" {

					b.game.state = 1
				}
			}
		}
	} else {
		m.keyIsBeingPressed = false
	}
	return nil
}

func (b *Board) ResetGame() {
	b.board = [BOARDSIZE][BOARDSIZE]int{}
	b.game.score = 0
	b.randomNewPiece()
	b.game.state = 2 // swap to main menu
}
