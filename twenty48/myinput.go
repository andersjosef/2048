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

var m = &MyInput{}

type ActionFunc func(*Board)

const MOVE_THRESHOLD = 100 // Delta distance needed to trigger a move

// buttons
var keyActions = map[GameState]map[ebiten.Key]ActionFunc{
	StateRunning: { // Main loop
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
	},
	StateMainMenu: { // Menu
		ebiten.KeyEscape: (*Board).CloseGame,
		ebiten.KeyF:      (*Board).ToggleFullScreen,
		ebiten.KeyQ:      (*Board).SwitchDefaultDarkMode,
	},
}

func (m *MyInput) UpdateInput(b *Board) error {
	// Keyboard and Mouse input handling
	m.handleKeyboardInput(b)
	m.handleMouseInput(b)
	return nil
}

func (m *MyInput) handleKeyboardInput(b *Board) error {
	m.keys = inpututil.AppendPressedKeys(m.keys[:0])

	// Take key and prevent retriggering
	if len(m.keys) > 0 && !m.keyIsBeingPressed {
		m.keyIsBeingPressed = true
		key_pressed := m.keys[len(m.keys)-1]

		// Get the appropriate action map based on the current game state
		if actionMap, ok := keyActions[b.game.state]; ok { // Check if actionmap exist for current game state
			if action, exists := actionMap[key_pressed]; exists { // Take snapshot of the board and do action
				b.boardBeforeChange = b.board
				action(b)
				if b.game.state == StateRunning { // If its the main loop add a piece
					b.addNewRandomPieceIfBoardChanged()
				}
			} else if b.game.state == StateMainMenu { // If button is not in map and state is main menu
				b.game.state = StateRunning
			}
		}

	} else if len(m.keys) == 0 {
		m.keyIsBeingPressed = false
	}
	return nil
}

func (m *MyInput) handleMouseInput(b *Board) {
	// Can left, right or wheel click
	var pressed bool = ebiten.IsMouseButtonPressed(ebiten.MouseButton0) ||
		ebiten.IsMouseButtonPressed(ebiten.MouseButton1) ||
		ebiten.IsMouseButtonPressed(ebiten.MouseButton2)

	// Cursor movement updates
	if pressed {
		if b.game.state == StateMainMenu { // If in main menu click will trigger game state
			b.game.state = StateRunning
		} else { // If not in menu update only end cursor coordinate
			m.endCursorPos[0], m.endCursorPos[1] = ebiten.CursorPosition()
		}
	} else { // If not clicking: update both values
		m.resetMouseState()
	}

	// Check if delta movements is large enough to trigger move
	if m.shoulTriggerMove() && !m.justMoved {
		m.performMove(b)
		m.justMoved = true
	}
}

func (m *MyInput) shoulTriggerMove() bool {
	dx := m.endCursorPos[0] - m.startCursorPos[0]
	dy := m.endCursorPos[1] - m.startCursorPos[1]

	return int(math.Abs(float64(dx))) > MOVE_THRESHOLD || int(math.Abs(float64(dy))) > MOVE_THRESHOLD
}

func (m *MyInput) resetMouseState() {
	m.justMoved = false
	m.startCursorPos[0], m.startCursorPos[1] = ebiten.CursorPosition()
	m.endCursorPos[0], m.endCursorPos[1] = ebiten.CursorPosition()
}

func (m *MyInput) performMove(b *Board) {
	dx := m.endCursorPos[0] - m.startCursorPos[0]
	dy := m.endCursorPos[1] - m.startCursorPos[1]

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
