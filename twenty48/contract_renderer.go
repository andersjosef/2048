package twenty48

import (
	"github.com/andersjosef/2048/twenty48/renderer"
	"github.com/hajimehoshi/ebiten/v2"
)

type Renderer interface {
	Draw(*ebiten.Image)
}

func NewRenderer(g *Game) Renderer {
	deps := renderer.Deps{
		Animation: g.animation,
		Base:      g.Board,
	}
	return renderer.New(deps)
}
