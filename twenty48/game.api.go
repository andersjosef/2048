package twenty48

import (
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/theme"
)

func (g *Game) GetScore() int {
	return g.score
}

func (g *Game) GetState() co.GameState {
	return g.state
}

func (g *Game) SetState(gs co.GameState) {
	g.state = gs
}

func (g Game) IsGameOver() bool {
	return g.gameOver
}

func (g *Game) GetPreviousState() co.GameState {
	return g.previousState
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
