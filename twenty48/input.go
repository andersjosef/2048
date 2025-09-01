package twenty48

import (
	"math"

	"github.com/andersjosef/2048/twenty48/board"
	"github.com/andersjosef/2048/twenty48/commands"
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/eventhandler"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Input struct {
	game              *Game
	Cmds              commands.Commands
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

	keyActions map[co.GameState]map[ebiten.Key]func()
}

func InitInput(g *Game, cmds commands.Commands) *Input {
	var i = &Input{game: g, Cmds: cmds}
	i.touchInput = newTouchInput(i)

	i.movementThreshold = 20 // Set how much the mouse has to move to reappear

	i.game.Register(
		eventhandler.EventScreenChanged,
		func(_ eventhandler.Event) {
			i.updatePauseButtonLocation()
		},
	)

	i.addKeyActions()

	return i
}

type ActionFunc func(*Input)

const MOVE_THRESHOLD = 100 // Delta distance needed to trigger a move

func (i *Input) addKeyActions() {

	// buttons
	i.keyActions = map[co.GameState]map[ebiten.Key]func(){
		co.StateRunning: { // Main loop
			ebiten.KeyArrowRight: i.Cmds.MoveRight,
			ebiten.KeyD:          i.Cmds.MoveRight,
			ebiten.KeyArrowLeft:  i.Cmds.MoveLeft,
			ebiten.KeyA:          i.Cmds.MoveLeft,
			ebiten.KeyArrowUp:    i.Cmds.MoveUp,
			ebiten.KeyW:          i.Cmds.MoveUp,
			ebiten.KeyArrowDown:  i.Cmds.MoveDown,
			ebiten.KeyS:          i.Cmds.MoveDown,
			ebiten.KeyR:          i.Cmds.ResetGame,
			ebiten.KeyF:          i.Cmds.ToggleFullscreen,
			ebiten.KeyEscape:     i.Cmds.CloseGame,
			ebiten.KeyQ:          i.Cmds.ToggleTheme,
			ebiten.KeyI:          i.Cmds.ToggleInfo,
			ebiten.KeyMinus:      i.Cmds.ScaleUp,
			ebiten.KeyPeriod:     i.Cmds.ScaleDown,
		},
		co.StateMainMenu: { // Menu
			ebiten.KeyEscape: i.Cmds.CloseGame,
			ebiten.KeyF:      i.Cmds.ToggleFullscreen,
			ebiten.KeyQ:      i.Cmds.ToggleTheme,
			ebiten.KeyI:      i.Cmds.ToggleInfo,
			ebiten.KeyMinus:  i.Cmds.ScaleUp,
			ebiten.KeyPeriod: i.Cmds.ScaleDown,
		},
		co.StateInstructions: { // Instructions
			ebiten.KeyEscape: i.Cmds.CloseGame,
			ebiten.KeyF:      i.Cmds.ToggleFullscreen,
			ebiten.KeyQ:      i.Cmds.ToggleTheme,
			ebiten.KeyI:      i.Cmds.ToggleInfo,
			ebiten.KeyMinus:  i.Cmds.ScaleUp,
			ebiten.KeyPeriod: i.Cmds.ScaleDown,
			ebiten.KeyR:      i.Cmds.ResetGame,
		},
		co.StateGameOver: { // Game Over
			ebiten.KeyEscape: i.Cmds.CloseGame,
			ebiten.KeyF:      i.Cmds.ToggleFullscreen,
			ebiten.KeyQ:      i.Cmds.ToggleTheme,
			ebiten.KeyI:      i.Cmds.ToggleInfo,
			ebiten.KeyMinus:  i.Cmds.ScaleUp,
			ebiten.KeyPeriod: i.Cmds.ScaleDown,
			ebiten.KeyR:      i.Cmds.ResetGame,
		},
	}
}

func (i *Input) UpdateInput() error {
	// Keyboard and Mouse input handling
	if i.game.buttonManager.checkButtons() {
		return nil
	}
	i.handleKeyboardInput()
	i.handleMouseInput()
	i.touchInput.TouchUpdate()
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
		if actionMap, ok := i.keyActions[i.game.state]; ok { // Check if actionmap exist for current game state
			if action, exists := actionMap[key_pressed]; exists { // Take snapshot of the board and do action
				action()
			} else if i.game.state == co.StateMainMenu { // If button is not in map and state is main menu
				i.game.state = co.StateRunning
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
		if i.game.state == co.StateMainMenu { // If in main menu click will trigger game state
			i.game.state = co.StateRunning
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
	i.game.Emit(eventhandler.Event{
		Type: eventhandler.EventResetGame,
	})
}

func CloseGame(i *Input) {
	i.game.shouldClose = true
}

func ToggleFullScreen(i *Input) {
	i.game.screenControl.ToggleFullScreen()
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
	i.game.board.CreateBoardImage()
	i.game.menu.UpdateDynamicText()
}

///// Main game logic action /////

func (i *Input) moveRight() {
	if i.game.gameOver {
		return
	}
	i.game.board.Move(board.Right)

}
func (i *Input) moveLeft() {
	if i.game.gameOver {
		return
	}
	i.game.board.Move(board.Left)
}
func (i *Input) moveUp() {
	if i.game.gameOver {
		return
	}
	i.game.board.Move(board.Up)
}
func (i *Input) moveDown() {
	if i.game.gameOver {
		return
	}
	i.game.board.Move(board.Down)
}

///// Menu Logic /////

func toggleInfo(i *Input) {
	switch i.game.state {
	case co.StateMainMenu:
		i.game.state = co.StateInstructions
		i.game.previousState = co.StateMainMenu
	case co.StateRunning:
		i.game.state = co.StateInstructions
		i.game.previousState = co.StateRunning
	case co.StateInstructions:
		i.game.state = i.game.previousState
	}

}

// Helper function for updating the pause button location
// When changing screen size
func (i *Input) updatePauseButtonLocation() {
	width, _ := i.game.GetActualSize()
	i.game.buttonManager.buttonKeyMap["II"].UpdatePos(width-20, 20)
}
