package twenty48

import (
	"github.com/andersjosef/2048/twenty48/commands"
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/eventhandler"
	"github.com/hajimehoshi/ebiten/v2"
)

// TODO: Change centerWindwo and Scale Window
func NewCommands(g *Systems) *commands.Commands {
	centerWindow := func() {
		mw, mh := ebiten.Monitor().Size()
		ww, wh := ebiten.WindowSize()
		ebiten.SetWindowPosition(mw/2-ww/2, mh/2-wh/2)
	}
	deps := commands.Deps{
		Board:         g.Board,
		EventHandler:  g.EventBus,
		ScreenControl: g.ScreenControl,
		FSM:           g.d.FSM,

		IncrementCurrentTheme: func() { // Change this
			g.Theme.NextFont()
			g.BoardView.RebuildBoard()
			g.Menu.UpdateDynamicText()
		},
		ToggleInfo: func() {
			switch g.GetState() {
			case co.StateMainMenu:
				g.d.FSM.Switch(co.StateInstructions)
			case co.StateRunning:
				g.d.FSM.Switch(co.StateInstructions)
			case co.StateInstructions:
				g.d.FSM.Switch(g.d.FSM.Previous())
			}
		},
		ScaleWindow: func() {
			g.EventBus.Emit(eventhandler.Event{
				Type: eventhandler.EventScreenChanged,
			})
			g.Theme.UpdateFonts()
			g.Menu.UpdateCenteredTitle()
			g.Buttons.UpdatePauseButtonLocation()

			g.Buttons.UpdateFontsForButtons()
			scale := g.ScreenControl.GetScale()
			newWindowWidth := int(float64(co.LOGICAL_WIDTH) * scale)
			newWindowHeight := int(float64(co.LOGICAL_HEIGHT) * scale)
			ebiten.SetWindowSize(newWindowWidth, newWindowHeight)
			centerWindow()
		},
	}
	return commands.BuildCommands(deps)
}
