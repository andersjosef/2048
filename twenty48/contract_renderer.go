package twenty48

import (
	"github.com/andersjosef/2048/twenty48/renderer"
)

func NewRenderer(g *Systems) *renderer.Renderer {
	deps := renderer.Deps{
		Animation: g.animation,
		BoardView: g.BoardView,
	}
	return renderer.New(deps)
}
