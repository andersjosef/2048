package renderer

import "github.com/hajimehoshi/ebiten/v2"

type Renderer struct {
	d Deps
}

func New(d Deps) *Renderer {
	r := &Renderer{
		d: d,
	}

	return r
}

func (r *Renderer) Draw(screen *ebiten.Image) {
	if r.d.Animation.IsAnimating() {
		r.d.Animation.Draw(screen)
		return
	}

	r.d.BoardView.Draw(screen)
}
