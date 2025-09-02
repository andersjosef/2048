package input

import (
	"github.com/andersjosef/2048/twenty48/commands"
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type KeyboardDeps struct {
	GetState                   func() co.GameState
	SetState                   func(co.GameState)
	CheckForMakingCursorHidden func()
	cmds                       commands.Commands
}

type KeyboardInput struct {
	d                 KeyboardDeps
	keys              []ebiten.Key
	keyIsBeingPressed bool

	keyActions map[co.GameState]map[ebiten.Key]func()
}

func NewKeyboardInput(d KeyboardDeps) *KeyboardInput {
	ki := &KeyboardInput{d: d}
	ki.addKeyBindings()
	return ki
}

// Keybindings
func (ki *KeyboardInput) addKeyBindings() {
	ki.keyActions = map[co.GameState]map[ebiten.Key]func(){
		co.StateRunning: { // Main loop
			ebiten.KeyArrowRight: ki.d.cmds.MoveRight,
			ebiten.KeyD:          ki.d.cmds.MoveRight,
			ebiten.KeyArrowLeft:  ki.d.cmds.MoveLeft,
			ebiten.KeyA:          ki.d.cmds.MoveLeft,
			ebiten.KeyArrowUp:    ki.d.cmds.MoveUp,
			ebiten.KeyW:          ki.d.cmds.MoveUp,
			ebiten.KeyArrowDown:  ki.d.cmds.MoveDown,
			ebiten.KeyS:          ki.d.cmds.MoveDown,
			ebiten.KeyR:          ki.d.cmds.ResetGame,
			ebiten.KeyF:          ki.d.cmds.ToggleFullscreen,
			ebiten.KeyEscape:     ki.d.cmds.CloseGame,
			ebiten.KeyQ:          ki.d.cmds.ToggleTheme,
			ebiten.KeyI:          ki.d.cmds.ToggleInfo,
			ebiten.KeyMinus:      ki.d.cmds.ScaleUp,
			ebiten.KeyPeriod:     ki.d.cmds.ScaleDown,
		},
		co.StateMainMenu: { // Menu
			ebiten.KeyEscape: ki.d.cmds.CloseGame,
			ebiten.KeyF:      ki.d.cmds.ToggleFullscreen,
			ebiten.KeyQ:      ki.d.cmds.ToggleTheme,
			ebiten.KeyI:      ki.d.cmds.ToggleInfo,
			ebiten.KeyMinus:  ki.d.cmds.ScaleUp,
			ebiten.KeyPeriod: ki.d.cmds.ScaleDown,
		},
		co.StateInstructions: { // Instructions
			ebiten.KeyEscape: ki.d.cmds.CloseGame,
			ebiten.KeyF:      ki.d.cmds.ToggleFullscreen,
			ebiten.KeyQ:      ki.d.cmds.ToggleTheme,
			ebiten.KeyI:      ki.d.cmds.ToggleInfo,
			ebiten.KeyMinus:  ki.d.cmds.ScaleUp,
			ebiten.KeyPeriod: ki.d.cmds.ScaleDown,
			ebiten.KeyR:      ki.d.cmds.ResetGame,
		},
		co.StateGameOver: { // Game Over
			ebiten.KeyEscape: ki.d.cmds.CloseGame,
			ebiten.KeyF:      ki.d.cmds.ToggleFullscreen,
			ebiten.KeyQ:      ki.d.cmds.ToggleTheme,
			ebiten.KeyI:      ki.d.cmds.ToggleInfo,
			ebiten.KeyMinus:  ki.d.cmds.ScaleUp,
			ebiten.KeyPeriod: ki.d.cmds.ScaleDown,
			ebiten.KeyR:      ki.d.cmds.ResetGame,
		},
	}
}

func (i *KeyboardInput) Update() {
	i.keys = inpututil.AppendPressedKeys(i.keys[:0])

	// Take key and prevent retriggering
	if len(i.keys) > 0 && !i.keyIsBeingPressed {
		i.d.CheckForMakingCursorHidden()
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
}
