package twenty48

import (
	"github.com/andersjosef/2048/twenty48/animations"
	"github.com/hajimehoshi/ebiten/v2"
)

type Animation interface {
	Draw(screen *ebiten.Image)
	IsAnimating() bool
}

func NewAnimation(g *Router) Animation {
	deps := animations.Deps{
		Board:        g.Board,
		EventHandler: g.EventBus,
	}

	return animations.New(deps)
}
