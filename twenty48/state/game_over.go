package state

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type GameOver struct {
	D interface {
		DrawGameOver(*ebiten.Image)
	}
}

func (s *GameOver) Enter() {}

func (s *GameOver) Exit() {}

func (s *GameOver) Update() error { return nil }

func (s *GameOver) Draw(screen *ebiten.Image) {
	s.D.DrawGameOver(screen)
}
