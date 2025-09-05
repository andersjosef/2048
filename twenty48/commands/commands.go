package commands

import (
	"github.com/andersjosef/2048/twenty48/board"
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/eventhandler"
	"github.com/hajimehoshi/ebiten/v2"
)

type Commands struct {
	MoveLeft, MoveRight, MoveUp, MoveDown, TurboMove func()
	ResetGame, ToggleFullscreen, CloseGame           func()
	ToggleTheme, ToggleInfo                          func()
	ScaleUp, ScaleDown                               func()
	GoToRunning                                      func()
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
			d.BoardView.CreateBoardImage()
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
