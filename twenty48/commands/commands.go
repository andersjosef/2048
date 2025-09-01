package commands

import (
	"github.com/andersjosef/2048/twenty48/board"
	"github.com/andersjosef/2048/twenty48/eventhandler"
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
			// d.Menu.UpdateDynamicText()
		},
		ToggleInfo: func() { d.ToggleInfo() },

		ScaleUp: func() {
			// // Oob Check to stop the growth somewhere
			// mw, mh := ebiten.Monitor().Size()
			// ww, wh := ebiten.WindowSize()
			// if ww >= mw || wh >= mh {
			// 	return
			// }
			// i.game.screenControl.IncrementScale()
			// ScaleWindow(i)
		},
		ScaleDown: func() {
			// if i.game.screenControl.DecrementScale() {
			// 	ScaleWindow(i)
			// }
		},
	}
}

// // Helper function for scaling image, contains what is equal for up and down
// func ScaleWindow() {
// 	i.game.updateFonts()
// 	i.game.board.ScaleBoard()
// 	i.game.menu.UpdateCenteredTitle()
// 	i.updatePauseButtonLocation()
// 	i.game.buttonManager.UpdateFontsForButtons()
// 	ebiten.SetWindowSize(co.LOGICAL_WIDTH*int(i.game.screenControl.GetScale()), co.LOGICAL_HEIGHT*int(i.game.screenControl.GetScale()))
// 	i.centerWindow()

// }

// // Will center the image to the new size
// // when scaling the screen up and down
// func (i *Input) centerWindow() {
// 	mw, mh := ebiten.Monitor().Size()
// 	ww, wh := ebiten.WindowSize()
// 	ebiten.SetWindowPosition(mw/2-ww/2, mh/2-wh/2)
// }
