package twenty48

import (
	"github.com/andersjosef/2048/twenty48/animations"
)

func NewAnimation(g *Systems) *animations.Animation {
	deps := animations.Deps{
		Board:        g.Board,
		BoardView:    g.BoardView,
		EventHandler: g.EventBus,
	}

	return animations.New(deps)
}
