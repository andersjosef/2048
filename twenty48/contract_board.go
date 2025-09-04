package twenty48

import (
	"github.com/andersjosef/2048/twenty48/board"
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/theme"
)

func NewBoard(g *Router) *board.Board {
	d := board.Deps{
		EventHandler:  g.EventBus,
		ScreenControl: g.screenControl,
		Core:          g.Core,
		SetGameOver: func(isGameOver bool) {
			if isGameOver {
				g.d.FSM.Switch(co.StateGameOver)
			}
		},
		SetGameState:      func(gs co.GameState) { g.SetState(gs) },
		IsGameOver:        func() bool { return g.IsGameOver() },
		GetCurrentTheme:   func() theme.Theme { return g.Theme.Current() },
		GetCurrentFontSet: func() theme.FontSet { return *g.Theme.Fonts() },
	}

	b, err := board.New(d)
	if err != nil {
		panic(err)
	}

	return b
}
