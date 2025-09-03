package input

import (
	"github.com/andersjosef/2048/twenty48/commands"
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type KeyboardDeps struct {
	State interface {
		SetState(co.GameState)
		GetState() co.GameState
	}
	nav    interface{ Switch(co.GameState) }
	cmds   commands.Commands
	Cursor interface{ Hide() }
}

type KeyboardInput struct {
	d                 KeyboardDeps
	keys              []ebiten.Key
	keyIsBeingPressed bool

	keyActions   map[co.GameState]map[ebiten.Key]func()
	onUnhandeled map[co.GameState]func() // What happens an unregistered key is pressed in state
}

func NewKeyboardInput(d KeyboardDeps, nav *Navigator) *KeyboardInput {
	ki := &KeyboardInput{
		d: d,
	}
	ki.addKeyBindings()
	return ki
}

func (i *KeyboardInput) Update() {
	i.keys = inpututil.AppendPressedKeys(i.keys[:0])

	// Take key and prevent retriggering
	if len(i.keys) > 0 && !i.keyIsBeingPressed {
		i.d.Cursor.Hide()
		i.keyIsBeingPressed = true
		key_pressed := i.keys[len(i.keys)-1]

		// Get the appropriate action map based on the current game state
		currentState := i.d.State.GetState()
		if actionMap, ok := i.keyActions[currentState]; ok { // Check if actionmap exist for current game state
			if action, exists := actionMap[key_pressed]; exists { // Take snapshot of the board and do action
				action()
			} else if onUnhandledAction, exists := i.onUnhandeled[currentState]; exists {
				// What happens if key not in list is pressed for state
				onUnhandledAction()
			}
		}

	} else if len(i.keys) == 0 {
		i.keyIsBeingPressed = false
	}
}
