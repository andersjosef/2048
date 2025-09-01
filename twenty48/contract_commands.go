package twenty48

import (
	"github.com/andersjosef/2048/twenty48/commands"
	co "github.com/andersjosef/2048/twenty48/constants"
)

func NewCommands(g *Game) commands.Commands {
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
	}
	return commands.BuildCommands(deps)
}
