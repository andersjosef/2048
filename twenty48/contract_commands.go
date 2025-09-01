package twenty48

import (
	"github.com/andersjosef/2048/twenty48/commands"
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/hajimehoshi/ebiten/v2"
)

// Change centerWindwo and Scale Window
func NewCommands(g *Game) commands.Commands {
	centerWindow := func() {
		mw, mh := ebiten.Monitor().Size()
		ww, wh := ebiten.WindowSize()
		ebiten.SetWindowPosition(mw/2-ww/2, mh/2-wh/2)
	}
	deps := commands.Deps{
		Board:         g.board,
		EventHandler:  g.eventBus,
		ScreenControl: g.screenControl,

		SetCloseGame: func(b bool) { g.shouldClose = b },
		IncrementCurrentTheme: func() { // Change this
			g.currentTheme = g.themePicker.IncrementCurrentTheme()
			g.board.CreateBoardImage()
			g.menu.UpdateDynamicText()
		},
		ToggleInfo: func() {
			switch g.state {
			case co.StateMainMenu:
				g.state = co.StateInstructions
				g.previousState = co.StateMainMenu
			case co.StateRunning:
				g.state = co.StateInstructions
				g.previousState = co.StateRunning
			case co.StateInstructions:
				g.state = g.previousState
			}
		},
		ScaleWindow: func() {
			g.updateFonts()
			g.board.ScaleBoard()
			g.menu.UpdateCenteredTitle()
			g.input.updatePauseButtonLocation()
			g.buttonManager.UpdateFontsForButtons()
			ebiten.SetWindowSize(co.LOGICAL_WIDTH*int(g.screenControl.GetScale()), co.LOGICAL_HEIGHT*int(g.screenControl.GetScale()))
			centerWindow()
		},
	}
	return commands.BuildCommands(deps)
}
