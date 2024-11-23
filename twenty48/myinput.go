package twenty48

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Input struct {
	game              *Game
	keys              []ebiten.Key
	keyIsBeingPressed bool

	// Cursor positions
	startCursorPos [2]int
	endCursorPos   [2]int
	justMoved      bool // To make sure only one move is done
}

func InitInput(g *Game) *Input {
	var m = &Input{game: g}

	return m
}

type ActionFunc func(*Input)

const MOVE_THRESHOLD = 100 // Delta distance needed to trigger a move

// buttons
var keyActions = map[GameState]map[ebiten.Key]ActionFunc{
	StateRunning: { // Main loop
		ebiten.KeyArrowRight: (*Input).moveRight,
		ebiten.KeyD:          (*Input).moveRight,
		ebiten.KeyArrowLeft:  (*Input).moveLeft,
		ebiten.KeyA:          (*Input).moveLeft,
		ebiten.KeyArrowUp:    (*Input).moveUp,
		ebiten.KeyW:          (*Input).moveUp,
		ebiten.KeyArrowDown:  (*Input).moveDown,
		ebiten.KeyS:          (*Input).moveDown,
		ebiten.KeyR:          (*Input).ResetGame,
		ebiten.KeyF:          (*Input).ToggleFullScreen,
		ebiten.KeyEscape:     (*Input).CloseGame,
		ebiten.KeyQ:          (*Input).SwitchDefaultDarkMode,
	},
	StateMainMenu: { // Menu
		ebiten.KeyEscape: (*Input).CloseGame,
		ebiten.KeyF:      (*Input).ToggleFullScreen,
		ebiten.KeyQ:      (*Input).SwitchDefaultDarkMode,
		ebiten.KeyI:      (*Input).toggleInfo,
	},
}

func (m *Input) UpdateInput(b *Board) error {
	// Keyboard and Mouse input handling
	m.game.buttonManager.checkButtons()
	m.handleKeyboardInput(b)
	m.handleMouseInput(b)
	return nil
}

func (i *Input) handleKeyboardInput(b *Board) error {
	i.keys = inpututil.AppendPressedKeys(i.keys[:0])

	// Take key and prevent retriggering
	if len(i.keys) > 0 && !i.keyIsBeingPressed {
		i.keyIsBeingPressed = true
		key_pressed := i.keys[len(i.keys)-1]

		// Get the appropriate action map based on the current game state
		if actionMap, ok := keyActions[b.game.state]; ok { // Check if actionmap exist for current game state
			if action, exists := actionMap[key_pressed]; exists { // Take snapshot of the board and do action
				action(i)
			} else if b.game.state == StateMainMenu { // If button is not in map and state is main menu
				b.game.state = StateRunning
			}
		}

	} else if len(i.keys) == 0 {
		i.keyIsBeingPressed = false
	}
	return nil
}

func (m *Input) handleMouseInput(b *Board) {
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

func (m *Input) shoulTriggerMove() bool {
	dx := m.endCursorPos[0] - m.startCursorPos[0]
	dy := m.endCursorPos[1] - m.startCursorPos[1]

	return int(math.Abs(float64(dx))) > MOVE_THRESHOLD || int(math.Abs(float64(dy))) > MOVE_THRESHOLD
}

func (m *Input) resetMouseState() {
	m.justMoved = false
	m.startCursorPos[0], m.startCursorPos[1] = ebiten.CursorPosition()
	m.endCursorPos[0], m.endCursorPos[1] = ebiten.CursorPosition()
}

func (m *Input) performMove(b *Board) {
	dx := m.endCursorPos[0] - m.startCursorPos[0]
	dy := m.endCursorPos[1] - m.startCursorPos[1]

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

}

////////////////////////////////////////////////
//				Actions						  //
////////////////////////////////////////////////

func (i *Input) ResetGame() {
	i.game.board.board = [BOARDSIZE][BOARDSIZE]int{}
	i.game.board.game.score = 0
	i.game.board.randomNewPiece()
	i.game.board.randomNewPiece()
	i.game.menu.state = MenuStateMain       // Main menu screen in menu
	i.game.board.game.state = StateMainMenu // Swap to main menu
}

func (i *Input) CloseGame() {
	i.game.board.game.shouldClose = true
}

// Main game logic action

func (i *Input) moveRight() {
	i.game.board.moveRight()
}
func (i *Input) moveLeft() {
	i.game.board.moveLeft()
}
func (i *Input) moveUp() {
	i.game.board.moveUp()
}
func (i *Input) moveDown() {
	i.game.board.moveDown()
}

// Menu Logic

func (i *Input) toggleInfo() {
	switch i.game.menu.state {
	case MenuStateMain:
		i.game.menu.state = MenuStateInstructions
	case MenuStateInstructions:
		i.game.menu.state = MenuStateMain

	}

}

func testForButtonAction(i *Input) {
	fmt.Println("Button pressed!!!! and action triggered")
}
