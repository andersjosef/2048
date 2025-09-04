package twenty48

import (
	"github.com/andersjosef/2048/twenty48/commands"
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/hajimehoshi/ebiten/v2"
)

// TODO: Change centerWindwo and Scale Window
func NewCommands(g *Game) *commands.Commands {
	centerWindow := func() {
		mw, mh := ebiten.Monitor().Size()
		ww, wh := ebiten.WindowSize()
		ebiten.SetWindowPosition(mw/2-ww/2, mh/2-wh/2)
	}
	deps := commands.Deps{
		Board:         g.Board,
		EventHandler:  g.EventBus,
		ScreenControl: g.screenControl,
		FSM:           g.d.FSM,

		IncrementCurrentTheme: func() { // Change this
			g.Theme.NextFont()

			g.Board.CreateBoardImage()
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
			g.updateFonts()
			g.Board.ScaleBoard()
			g.Menu.UpdateCenteredTitle()

			width, _ := g.screenControl.GetActualSize()
			g.buttonManager.UpdatePosForButton("II", width-20, 20)

			g.buttonManager.UpdateFontsForButtons()
			ebiten.SetWindowSize(co.LOGICAL_WIDTH*int(g.screenControl.GetScale()), co.LOGICAL_HEIGHT*int(g.screenControl.GetScale()))
			centerWindow()
		},
	}
	return commands.BuildCommands(deps)
}
