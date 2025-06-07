package twenty48

import (
	"math"

	"github.com/andersjosef/2048/twenty48/shadertools"
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

	touchInput *TouchInput

	isHidingMouse     bool
	lastPosMouse      [2]int
	movementThreshold float64 // If mouse is moved beyond this show again
}

func InitInput(g *Game) *Input {
	var i = &Input{game: g}
	i.touchInput = newTouchInput(i)

	i.movementThreshold = 20 // Set how much the mouse has to move to reappear

	return i
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
		ebiten.KeyQ:          toggleTheme,
		ebiten.KeyI:          toggleInfo,
		ebiten.KeyMinus:      ScaleWindowUp,
		ebiten.KeyPeriod:     ScaleWindowDown,
	},
	StateMainMenu: { // Menu
		ebiten.KeyEscape: CloseGame,
		ebiten.KeyF:      ToggleFullScreen,
		ebiten.KeyQ:      toggleTheme,
		ebiten.KeyI:      toggleInfo,
		ebiten.KeyMinus:  ScaleWindowUp,
		ebiten.KeyPeriod: ScaleWindowDown,
	},
	StateInstructions: { // Instructions
		ebiten.KeyEscape: CloseGame,
		ebiten.KeyF:      ToggleFullScreen,
		ebiten.KeyQ:      toggleTheme,
		ebiten.KeyI:      toggleInfo,
		ebiten.KeyMinus:  ScaleWindowUp,
		ebiten.KeyPeriod: ScaleWindowDown,
	},
	StateGameOver: { // Game Over
		ebiten.KeyEscape: CloseGame,
		ebiten.KeyF:      ToggleFullScreen,
		ebiten.KeyQ:      toggleTheme,
		ebiten.KeyI:      toggleInfo,
		ebiten.KeyR:      ResetGame,
		ebiten.KeyMinus:  ScaleWindowUp,
		ebiten.KeyPeriod: ScaleWindowDown,
	},
}

func (m *Input) UpdateInput(b *Board) error {
	// Keyboard and Mouse input handling
	if m.game.buttonManager.checkButtons() {
		return nil
	}
	m.handleKeyboardInput()
	m.handleMouseInput()
	m.touchInput.TouchUpdate()
	return nil
}

func (i *Input) handleKeyboardInput() error {
	i.keys = inpututil.AppendPressedKeys(i.keys[:0])

	// Take key and prevent retriggering
	if len(i.keys) > 0 && !i.keyIsBeingPressed {
		i.checkForMakingCursorHidden()
		i.keyIsBeingPressed = true
		key_pressed := i.keys[len(i.keys)-1]

		// Get the appropriate action map based on the current game state
		if actionMap, ok := keyActions[i.game.state]; ok { // Check if actionmap exist for current game state
			if action, exists := actionMap[key_pressed]; exists { // Take snapshot of the board and do action
				action(i)
			} else if i.game.state == StateMainMenu { // If button is not in map and state is main menu
				i.game.state = StateRunning
			}
		}

	} else if len(i.keys) == 0 {
		i.keyIsBeingPressed = false
	}
	return nil
}

func (i *Input) handleMouseInput() {
	i.checkForMakingCursorVisible()

	// Can left, right or wheel click
	var pressed bool = ebiten.IsMouseButtonPressed(ebiten.MouseButton0) ||
		ebiten.IsMouseButtonPressed(ebiten.MouseButton1) ||
		ebiten.IsMouseButtonPressed(ebiten.MouseButton2)

	// Cursor movement updates
	if pressed {
		if i.game.state == StateMainMenu { // If in main menu click will trigger game state
			i.game.state = StateRunning
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

	i.SelectMoveDelta(dx, dy)
}

func (i *Input) SelectMoveDelta(dx, dy int) {
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
		i.game.screenControl.fullscreen = false
		i.screenChanging()
	} else {
		ebiten.SetFullscreen(true)
		i.game.screenControl.fullscreen = true
		i.screenChanging()
	}
	i.game.screenSizeChanged = true
}

// Helper function for toggle screen
// Contains everything that is the same for full screen and windowed
func (i *Input) screenChanging() {
	i.game.screenControl.UpdateActualDimentions()
	i.game.board.sizes.scaleBoard()
	i.game.menu.initTitle()
	i.updatePauseButtonLocation()
	shadertools.UpdateScaleNoiseImage()
}

// Helper functions for toggeling mouse being displayed or not
func (i *Input) checkForMakingCursorVisible() {
	if i.isHidingMouse {
		lastX := float64(i.lastPosMouse[0])
		lastY := float64(i.lastPosMouse[1])

		x, y := ebiten.CursorPosition()

		if math.Abs(lastX-float64(x)) > i.movementThreshold ||
			math.Abs(lastY-float64(y)) > i.movementThreshold {
			ebiten.SetCursorMode(ebiten.CursorModeVisible)
			i.isHidingMouse = false
		}
	}
}

func (i *Input) checkForMakingCursorHidden() {
	if !i.isHidingMouse {
		i.lastPosMouse[0], i.lastPosMouse[1] = ebiten.CursorPosition()
		ebiten.SetCursorMode(ebiten.CursorModeHidden)
		i.isHidingMouse = true
	}
}

func toggleTheme(i *Input) {
	i.game.currentTheme = i.game.themePicker.IncrementCurrentTheme()
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

func ScaleWindowUp(i *Input) {
	// Oob Check to stop the growth somewhere
	mw, mh := ebiten.Monitor().Size()
	ww, wh := ebiten.WindowSize()
	if ww >= mw || wh >= mh {
		return
	}
	i.game.scale++
	ScaleWindow(i)
}

func ScaleWindowDown(i *Input) {
	if i.game.scale > 1 {
		i.game.scale--
		ScaleWindow(i)
	}
}

// Helper function for scaling image, contains what is equal for up and down
func ScaleWindow(i *Input) {
	i.game.screenControl.UpdateActualDimentions()
	i.game.updateFonts()
	i.game.board.sizes.scaleBoard()
	i.game.menu.initTitle()
	i.updatePauseButtonLocation()
	i.game.buttonManager.UpdateFontsForButtons()
	ebiten.SetWindowSize(LOGICAL_WIDTH*int(i.game.scale), LOGICAL_HEIGHT*int(i.game.scale))
	i.centerWindow()

}

// Will center the image to the new size
// when scaling the screen up and down
func (i *Input) centerWindow() {
	mw, mh := ebiten.Monitor().Size()
	ww, wh := ebiten.WindowSize()
	ebiten.SetWindowPosition(mw/2-ww/2, mh/2-wh/2)
}

// Helper function for updating the pause button location
// When changing screen size
func (i *Input) updatePauseButtonLocation() {
	i.game.buttonManager.buttonKeyMap["II"].UpdatePos(i.game.screenControl.actualWidth-20, 20)
}
