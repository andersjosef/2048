package twenty48

import (
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/theme"
)

func (g *Router) GetScore() int {
	return g.Core.Score()
}

func (g *Router) GetState() co.GameState {
	return g.d.FSM.Current()
}

func (g *Router) SetState(gs co.GameState) {
	g.d.FSM.Switch(gs)
}

func (g Router) IsGameOver() bool {
	return g.d.IsGameOver()
}

func (g *Router) GetPreviousState() co.GameState {
	return g.d.FSM.Previous()
}

func (g *Router) GetCurrentTheme() theme.Theme {
	return g.Theme.Current()
}

func (g *Router) GetFontSet() theme.FontSet {
	return *g.Theme.Fonts()
}

func (g *Router) ScreenControl() ScreenControl {
	return g.screenControl
}
