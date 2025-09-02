package twenty48

import (
	"github.com/andersjosef/2048/twenty48/commands"
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/input"
)

func NewInput(g *Game, cmds commands.Commands) *input.Input {
	deps := input.Deps{
		EventHandler:  g.eventBus,
		Buttons:       g.buttonManager,
		ScreenControl: g.screenControl,

		Cmds:       cmds,
		GetState:   func() co.GameState { return g.GetState() },
		SetState:   func(gs co.GameState) { g.state = gs },
		IsGameOver: func() bool { return g.gameOver },
	}
	return input.New(deps)
}
