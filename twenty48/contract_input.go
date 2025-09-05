package twenty48

import (
	"github.com/andersjosef/2048/twenty48/commands"
	"github.com/andersjosef/2048/twenty48/input"
)

func NewInput(g *Systems, cmds *commands.Commands) *input.Input {
	deps := input.Deps{
		EventHandler:  g.EventBus,
		Buttons:       g.Buttons,
		ScreenControl: g.screenControl,

		Cmds:  cmds,
		State: g,
	}
	return input.New(deps)
}
