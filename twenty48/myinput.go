package twenty48

import (
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

type ActionFunc func(*Board)

// buttons while main game loop
var keyActionsMainGameLoop = map[ebiten.Key]ActionFunc{
	ebiten.KeyArrowRight: (*Board).moveRight,
	ebiten.KeyD:          (*Board).moveRight,
	ebiten.KeyArrowLeft:  (*Board).moveLeft,
	ebiten.KeyA:          (*Board).moveLeft,
	ebiten.KeyArrowUp:    (*Board).moveUp,
	ebiten.KeyW:          (*Board).moveUp,
	ebiten.KeyArrowDown:  (*Board).moveDown,
	ebiten.KeyS:          (*Board).moveDown,
	ebiten.KeyR:          (*Board).ResetGame,
	ebiten.KeyF:          (*Board).ToggleFullScreen,
	ebiten.KeyEscape:     (*Board).CloseGame,
	ebiten.KeyQ:          (*Board).SwitchDefaultDarkMode,
}

// buttons while main menu
var keyActionsMainMenu = map[ebiten.Key]ActionFunc{
	ebiten.KeyEscape: (*Board).CloseGame,
	ebiten.KeyF:      (*Board).ToggleFullScreen,
	ebiten.KeyQ:      (*Board).SwitchDefaultDarkMode,
}

// this is also the game logic I guess
func (m *MyInput) UpdateInput(b *Board) error {
	m.keys = inpututil.AppendPressedKeys(m.keys[:0])

	if len(m.keys) > 0 && !m.keyIsBeingPressed {
		m.keyIsBeingPressed = true
		key_pressed := m.keys[len(m.keys)-1]

		// fmt.Println(key_pressed)
		if action, ok := keyActionsMainGameLoop[key_pressed]; ok && b.game.state == 1 { // main game
			b.boardBeforeChange = b.board
			action(b)
			b.addNewRandomPieceIfBoardChanged()
		} else if b.game.state == 2 { // main menu
			action, ok = keyActionsMainMenu[key_pressed]
			if ok { // button is in map
				action(b)
			} else { // button is not, default behaviour
				b.game.state = 1
			}
		}

	} else if len(m.keys) == 0 {
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

func (b *Board) CloseGame() {
	b.game.shouldClose = true
}
