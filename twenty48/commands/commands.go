package commands

import (
	"github.com/andersjosef/2048/twenty48/board"
	"github.com/andersjosef/2048/twenty48/eventhandler"
	"github.com/hajimehoshi/ebiten/v2"
)

type Commands struct {
	MoveLeft, MoveRight, MoveUp, MoveDown  func()
	ResetGame, ToggleFullscreen, CloseGame func()
	ToggleTheme, ToggleInfo                func()
	ScaleUp, ScaleDown                     func()
}

func BuildCommands(d Deps) Commands {
	return Commands{
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
		CloseGame:        func() { d.SetCloseGame(true) },

		ToggleTheme: func() {
			d.IncrementCurrentTheme()
			d.Board.CreateBoardImage()
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
	}
}
