package twenty48

import (
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/theme"
)

func (g *Systems) GetScore() int {
	return g.Core.Score()
}

func (g *Systems) GetState() co.GameState {
	return g.d.FSM.Current()
}

func (g *Systems) SetState(gs co.GameState) {
	g.d.FSM.Switch(gs)
}

func (g Systems) IsGameOver() bool {
	return g.d.IsGameOver()
}

func (g *Systems) GetPreviousState() co.GameState {
	return g.d.FSM.Previous()
}

func (g *Systems) GetCurrentTheme() theme.Theme {
	return g.Theme.Current()
}

func (g *Systems) GetFontSet() theme.FontSet {
	return *g.Theme.Fonts()
}

func (g *Systems) ScreenControl() ScreenControl {
	return g.screenControl
}
