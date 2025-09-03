package twenty48

import (
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/theme"
)

func (g *Game) GetScore() int {
	return g.score
}

func (g *Game) GetState() co.GameState {
	return g.d.FSM.Current()
}

func (g *Game) SetState(gs co.GameState) {
	g.d.FSM.Switch(gs)
}

func (g Game) IsGameOver() bool {
	return g.d.IsGameOver()
}

func (g *Game) GetPreviousState() co.GameState {
	return g.d.FSM.Previous()
}

func (g *Game) GetCurrentTheme() theme.Theme {
	return g.currentTheme
}

func (g *Game) GetFontSet() theme.FontSet {
	return *g.fontSet
}

func (g *Game) ScreenControl() ScreenControl {
	return g.screenControl
}
