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
		BoardView:     g.BoardView,
		EventHandler:  g.EventBus,
		ScreenControl: g.screenControl,
		FSM:           g.d.FSM,

		IncrementCurrentTheme: func() { // Change this
			g.Theme.NextFont()

			// g.Board.CreateBoardImage()
			g.BoardView.CreateBoardImage()
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
			width, _ := g.screenControl.GetActualSize()
			g.Buttons.UpdatePosForButton("II", width-20, 20)

			g.Buttons.UpdateFontsForButtons()
			ebiten.SetWindowSize(co.LOGICAL_WIDTH*int(g.screenControl.GetScale()), co.LOGICAL_HEIGHT*int(g.screenControl.GetScale()))
			centerWindow()
		},
	}
	return commands.BuildCommands(deps)
}
