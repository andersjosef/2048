package commands

import (
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/core/board"
	"github.com/andersjosef/2048/twenty48/eventhandler"
	"github.com/hajimehoshi/ebiten/v2"
)

type Commands struct {
	MoveLeft, MoveRight, MoveUp, MoveDown  func()
	ResetGame, ToggleFullscreen, CloseGame func()
	ToggleTheme, ToggleInfo                func()
	ScaleUp, ScaleDown                     func()
	GoToRunning                            func()
	TurboMove, OneBeforeGameOver           func()
}

func BuildCommands(d Deps) *Commands {
	c := &Commands{
		MoveLeft:  func() { d.Board.Move(board.Left) },
		MoveRight: func() { d.Board.Move(board.Right) },
		MoveUp:    func() { d.Board.Move(board.Up) },
		MoveDown:  func() { d.Board.Move(board.Down) },

		ResetGame: func() {
			d.EventHandler.Emit(eventhandler.Event{
				Type: eventhandler.EventResetGame,
			})
		},
		ToggleFullscreen: d.ScreenControl.ToggleFullScreen,
		CloseGame:        func() { d.Switch(co.StateQuitGame) },

		ToggleTheme: func() {
			d.IncrementCurrentTheme()
			d.EventHandler.Emit(eventhandler.Event{
				Type: eventhandler.EventThemeChanged,
			})
		},
		ToggleInfo: func() { d.ToggleInfo() },

		ScaleUp: func() {
			// Oob Check to stop the growth somewhere
			mw, mh := ebiten.Monitor().Size()
			ww, wh := ebiten.WindowSize()
			if ww >= mw || wh >= mh {
				return
			}
			d.ScreenControl.IncrementScale()
			d.ScaleWindow()
		},
		ScaleDown: func() {
			if d.ScreenControl.DecrementScale() {
				d.ScaleWindow()
			}
		},
		GoToRunning: func() {
			d.FSM.Switch(co.StateRunning)
		},
		OneBeforeGameOver: func() {
			d.Board.SetBoard([4][4]int{
				{8, 2, 4, 128},
				{2048, 128, 2, 4},
				{8, 2, 4, 128},
				{4096, 128, 2, 2},
			})
		},
	}
	c.TurboMove = func() {
		for range 10 {
			c.MoveUp()
			c.MoveRight()
			c.MoveLeft()
			c.MoveDown()
		}
	}
	return c
}
