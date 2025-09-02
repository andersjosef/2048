package input

import (
	"math"

	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/eventhandler"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const MOVE_THRESHOLD = 100 // Delta distance needed to trigger a move

type Input struct {
	d                 Deps
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

func New(d Deps) *Input {
	var i = &Input{d: d}

	i.touchInput = newTouchInput(i)
	i.movementThreshold = 20 // Set how much the mouse has to move to reappear

	i.d.Register(
		eventhandler.EventScreenChanged,
		func(_ eventhandler.Event) {
			i.updatePauseButtonLocation()
		},
	)

	i.addKeyBindings()

	return i
}

func (i *Input) GiveButtons(b Buttons) {
	i.d.Buttons = b
}

// Keybindings
func (i *Input) addKeyBindings() {
	i.keyActions = map[co.GameState]map[ebiten.Key]func(){
		co.StateRunning: { // Main loop
			ebiten.KeyArrowRight: i.d.Cmds.MoveRight,
			ebiten.KeyD:          i.d.Cmds.MoveRight,
			ebiten.KeyArrowLeft:  i.d.Cmds.MoveLeft,
			ebiten.KeyA:          i.d.Cmds.MoveLeft,
			ebiten.KeyArrowUp:    i.d.Cmds.MoveUp,
			ebiten.KeyW:          i.d.Cmds.MoveUp,
			ebiten.KeyArrowDown:  i.d.Cmds.MoveDown,
			ebiten.KeyS:          i.d.Cmds.MoveDown,
			ebiten.KeyR:          i.d.Cmds.ResetGame,
			ebiten.KeyF:          i.d.Cmds.ToggleFullscreen,
			ebiten.KeyEscape:     i.d.Cmds.CloseGame,
			ebiten.KeyQ:          i.d.Cmds.ToggleTheme,
			ebiten.KeyI:          i.d.Cmds.ToggleInfo,
			ebiten.KeyMinus:      i.d.Cmds.ScaleUp,
			ebiten.KeyPeriod:     i.d.Cmds.ScaleDown,
		},
		co.StateMainMenu: { // Menu
			ebiten.KeyEscape: i.d.Cmds.CloseGame,
			ebiten.KeyF:      i.d.Cmds.ToggleFullscreen,
			ebiten.KeyQ:      i.d.Cmds.ToggleTheme,
			ebiten.KeyI:      i.d.Cmds.ToggleInfo,
			ebiten.KeyMinus:  i.d.Cmds.ScaleUp,
			ebiten.KeyPeriod: i.d.Cmds.ScaleDown,
		},
		co.StateInstructions: { // Instructions
			ebiten.KeyEscape: i.d.Cmds.CloseGame,
			ebiten.KeyF:      i.d.Cmds.ToggleFullscreen,
			ebiten.KeyQ:      i.d.Cmds.ToggleTheme,
			ebiten.KeyI:      i.d.Cmds.ToggleInfo,
			ebiten.KeyMinus:  i.d.Cmds.ScaleUp,
			ebiten.KeyPeriod: i.d.Cmds.ScaleDown,
			ebiten.KeyR:      i.d.Cmds.ResetGame,
		},
		co.StateGameOver: { // Game Over
			ebiten.KeyEscape: i.d.Cmds.CloseGame,
			ebiten.KeyF:      i.d.Cmds.ToggleFullscreen,
			ebiten.KeyQ:      i.d.Cmds.ToggleTheme,
			ebiten.KeyI:      i.d.Cmds.ToggleInfo,
			ebiten.KeyMinus:  i.d.Cmds.ScaleUp,
			ebiten.KeyPeriod: i.d.Cmds.ScaleDown,
			ebiten.KeyR:      i.d.Cmds.ResetGame,
		},
	}
}

func (i *Input) UpdateInput() error {
	// Keyboard and Mouse input handling
	if i.d.Buttons.CheckButtons() {
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
		if actionMap, ok := i.keyActions[i.d.GetState()]; ok { // Check if actionmap exist for current game state
			if action, exists := actionMap[key_pressed]; exists { // Take snapshot of the board and do action
				action()
			} else if i.d.GetState() == co.StateMainMenu { // If button is not in map and state is main menu
				i.d.SetState(co.StateRunning)
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
		if i.d.GetState() == co.StateMainMenu { // If in main menu click will trigger game state
			i.d.SetState(co.StateRunning)
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
	if i.d.IsGameOver() {
		return
	}

	i.SelectMoveDelta(dx, dy)
}

func (i *Input) SelectMoveDelta(dx, dy int) {
	if i.d.IsGameOver() {
		return
	}
	if math.Abs(float64(dx)) > math.Abs(float64(dy)) { // X-axis largest
		if dx > 0 {
			i.d.Cmds.MoveRight()
		} else {
			i.d.Cmds.MoveLeft()
		}
	} else { // Y-axis largest
		if dy > 0 {
			i.d.Cmds.MoveDown()
		} else {
			i.d.Cmds.MoveUp()
		}
	}

}

///// Utilities //////

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

// Helper function for updating the pause button location
// When changing screen size
// TODO: Move this
func (i *Input) updatePauseButtonLocation() {
	width, _ := i.d.ScreenControl.GetActualSize()
	i.d.Buttons.UpdatePosForButton("II", width-20, 20)
}
