package renderer

import "github.com/hajimehoshi/ebiten/v2"

type Deps struct {
	Animation
	BoardView
}

type Animation interface {
	Draw(*ebiten.Image)
	IsAnimating() bool
}

type BoardView interface {
	Draw(*ebiten.Image)
}
