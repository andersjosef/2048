package twenty48

import (
	"github.com/andersjosef/2048/twenty48/buttons"
	"github.com/andersjosef/2048/twenty48/commands"
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/theme"
)

func NewButtonManager(g *Router, cmds *commands.Commands) *buttons.ButtonManager {
	deps := buttons.Deps{
		ScreenControl: g.screenControl,
		Input:         g.Input,
		Utils:         g.utils,
		EventHandler:  g.EventBus,

		GetFontSet: func() theme.FontSet { return g.GetFontSet() },
		GetState:   func() co.GameState { return g.GetState() },
	}

	return buttons.NewButtonManager(deps, cmds)
}
