package input

import (
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/hajimehoshi/ebiten/v2"
)

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

	ki.onUnhandeled = map[co.GameState]func(){
		co.StateMainMenu: func() { ki.d.State.SetState(co.StateRunning) },
	}
}
