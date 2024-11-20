package twenty48

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type MyInput struct {
	keys              []ebiten.Key
	keyIsBeingPressed bool

	// Cursor positions
	startCursorPos [2]int
	endCursorPos   [2]int
	justMoved      bool // To make sure only one move is done
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

	// Mouse
	m.MouseInput(b)

	// Keyboard
	if len(m.keys) > 0 && !m.keyIsBeingPressed {
		m.keyIsBeingPressed = true
		key_pressed := m.keys[len(m.keys)-1]

		// fmt.Println(key_pressed)
		if action, ok := keyActionsMainGameLoop[key_pressed]; ok && b.game.state == StateRunning { // main game
			b.boardBeforeChange = b.board
			action(b)
			b.addNewRandomPieceIfBoardChanged()
		} else if b.game.state == StateMainMenu { // main menu
			action, ok = keyActionsMainMenu[key_pressed]
			if ok { // button is in map
				action(b)
			} else { // button is not, default behaviour
				b.game.state = StateRunning
			}
		}

	} else if len(m.keys) == 0 {
		m.keyIsBeingPressed = false
	}
	return nil
}

// Moving pieces by moving the mouse
func (m *MyInput) MouseInput(b *Board) {

	// Can left, right or wheel click
	var pressed bool = ebiten.IsMouseButtonPressed(ebiten.MouseButton0) || ebiten.IsMouseButtonPressed(ebiten.MouseButton1) || ebiten.IsMouseButtonPressed(ebiten.MouseButton2)

	// Cursor movement updates
	if pressed {
		if b.game.state == StateMainMenu { // If in main menu click will trigger game state
			b.game.state = StateRunning
		} else { // If not in menu update only end cursor coordinate
			m.endCursorPos[0], m.endCursorPos[1] = ebiten.CursorPosition()
		}
	} else { // If not clicking: update both values
		m.justMoved = false
		m.startCursorPos[0], m.startCursorPos[1] = ebiten.CursorPosition()
		m.endCursorPos[0], m.endCursorPos[1] = ebiten.CursorPosition()
	}

	threshold := 100 // Delta distance needed to trigger a move
	dx := m.endCursorPos[0] - m.startCursorPos[0]
	dy := m.endCursorPos[1] - m.startCursorPos[1]

	// Check if delta movements is large enough to trigger move
	if (int(math.Abs(float64(dx))) > threshold || int(math.Abs(float64(dy))) > threshold) && !m.justMoved {
		b.boardBeforeChange = b.board
		if math.Abs(float64(dx)) > math.Abs(float64(dy)) { // X-axis largest
			if dx > 0 {
				b.moveRight()
			} else {
				b.moveLeft()
			}
		} else { // Y-axis largest
			if dy > 0 {
				b.moveDown()
			} else {
				b.moveUp()
			}
		}
		b.addNewRandomPieceIfBoardChanged()
		m.justMoved = true
	}
}

func (b *Board) ResetGame() {
	b.board = [BOARDSIZE][BOARDSIZE]int{}
	b.game.score = 0
	b.randomNewPiece()
	b.game.state = StateMainMenu // swap to main menu
}

func (b *Board) CloseGame() {
	b.game.shouldClose = true
}
