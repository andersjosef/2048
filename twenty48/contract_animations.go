package twenty48

import (
	"github.com/andersjosef/2048/twenty48/animations"
	"github.com/hajimehoshi/ebiten/v2"
)

type Animation interface {
	Draw(screen *ebiten.Image)
	IsAnimating() bool
}

func NewAnimation(g *Game) Animation {
	deps := animations.Deps{
		Board:        g.board,
		EventHandler: g.eventBus,
	}

	return animations.New(deps)
}
