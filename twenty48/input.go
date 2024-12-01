package twenty48

import (
	"math"

	"github.com/andersjosef/2048/twenty48/shadertools"
	"github.com/andersjosef/2048/twenty48/theme"
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
		ebiten.KeyR:          ResetGame,
		ebiten.KeyF:          ToggleFullScreen,
		ebiten.KeyEscape:     CloseGame,
		ebiten.KeyQ:          SwitchDefaultDarkMode,
		ebiten.KeyI:          toggleInfo,
	},
	StateMainMenu: { // Menu
		ebiten.KeyEscape: CloseGame,
		ebiten.KeyF:      ToggleFullScreen,
		ebiten.KeyQ:      SwitchDefaultDarkMode,
		ebiten.KeyI:      toggleInfo,
	},
	StateInstructions: { // Instructions
		ebiten.KeyEscape: CloseGame,
		ebiten.KeyF:      ToggleFullScreen,
		ebiten.KeyQ:      SwitchDefaultDarkMode,
		ebiten.KeyI:      toggleInfo,
	},
	StateGameOver: { // Game Over
		ebiten.KeyEscape: CloseGame,
		ebiten.KeyF:      ToggleFullScreen,
		ebiten.KeyQ:      SwitchDefaultDarkMode,
		ebiten.KeyI:      toggleInfo,
		ebiten.KeyR:      ResetGame,
	},
}

func (m *Input) UpdateInput(b *Board) error {
	// Keyboard and Mouse input handling
	if m.game.buttonManager.checkButtons() {
		return nil
	}
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

func (i *Input) handleMouseInput(b *Board) {
	// Can left, right or wheel click
	var pressed bool = ebiten.IsMouseButtonPressed(ebiten.MouseButton0) ||
		ebiten.IsMouseButtonPressed(ebiten.MouseButton1) ||
		ebiten.IsMouseButtonPressed(ebiten.MouseButton2)

	// Cursor movement updates
	if pressed {
		if b.game.state == StateMainMenu { // If in main menu click will trigger game state
			b.game.state = StateRunning
		} else { // If not in menu update only end cursor coordinate
			i.endCursorPos[0], i.endCursorPos[1] = ebiten.CursorPosition()
		}
	} else { // If not clicking: update both values
		i.resetMouseState()
	}

	// Check if delta movements is large enough to trigger move
	if i.shoulTriggerMove() && !i.justMoved {
		i.performMove()
		i.justMoved = true
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

func (i *Input) performMove() {
	dx := i.endCursorPos[0] - i.startCursorPos[0]
	dy := i.endCursorPos[1] - i.startCursorPos[1]
	if i.game.gameOver {
		return
	}

	if math.Abs(float64(dx)) > math.Abs(float64(dy)) { // X-axis largest
		if dx > 0 {
			i.moveRight()
		} else {
			i.moveLeft()
		}
	} else { // Y-axis largest
		if dy > 0 {
			i.moveDown()
		} else {
			i.moveUp()
		}
	}

}

////////////////////////////////////////////////
//				Actions						  //
////////////////////////////////////////////////

///// Utilities //////

func ResetGame(i *Input) {
	i.game.board.board = [BOARDSIZE][BOARDSIZE]int{}
	i.game.board.game.score = 0
	i.game.board.randomNewPiece()
	i.game.board.randomNewPiece()
	i.game.board.game.state = StateMainMenu // Swap to main menu
	shadertools.ResetTimesMapsDissolve()
	i.game.menu.titleInFullView = false
	i.game.gameOver = false
}

func CloseGame(i *Input) {
	i.game.board.game.shouldClose = true
}

func ToggleFullScreen(i *Input) {
	if i.game.screenControl.fullscreen {
		ebiten.SetFullscreen(false)
		shadertools.UpdateNoiseImage(50, 50)
		i.game.buttonManager.buttonKeyMap["II"].UpdatePos(SCREENWIDTH-20, 20)
		i.game.screenControl.fullscreen = false
	} else {
		ebiten.SetFullscreen(true)
		shadertools.UpdateNoiseImage(100, 100)
		newScreenLength, _ := ebiten.Monitor().Size()
		i.game.buttonManager.buttonKeyMap["II"].UpdatePos(newScreenLength-20, 20)
		i.game.screenControl.fullscreen = true
	}
	i.game.menu.UpdateDynamicText()
	i.game.menu.titleImage = i.game.menu.initTitle()
	i.game.board.initBoardForEndScreen()
	i.game.screenSizeChanged = true
}

func SwitchDefaultDarkMode(i *Input) {
	i.game.darkMode = !i.game.darkMode

	if i.game.darkMode { // DARK MODE
		i.game.currentTheme = theme.DarkTheme
	} else { // DEFAULT MODE
		i.game.currentTheme = theme.DefaultTheme

	}
	i.game.board.createBoardImage()
	i.game.menu.UpdateDynamicText()
}

///// Main game logic action /////

func (i *Input) moveRight() {
	if i.game.gameOver {
		return
	}
	i.game.board.moveRight()
}
func (i *Input) moveLeft() {
	if i.game.gameOver {
		return
	}
	i.game.board.moveLeft()
}
func (i *Input) moveUp() {
	if i.game.gameOver {
		return
	}
	i.game.board.moveUp()
}
func (i *Input) moveDown() {
	if i.game.gameOver {
		return
	}
	i.game.board.moveDown()
}

///// Menu Logic /////

func toggleInfo(i *Input) {

	switch i.game.state {
	case StateMainMenu:
		i.game.state = StateInstructions
		i.game.previousState = StateMainMenu
	case StateRunning:
		i.game.state = StateInstructions
		i.game.previousState = StateRunning
	case StateInstructions:
		i.game.state = i.game.previousState
	}

}
