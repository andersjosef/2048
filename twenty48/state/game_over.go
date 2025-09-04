package state

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type GameOverDeps struct {
	Menu interface {
		DrawGameOver(*ebiten.Image)
	}
	Board interface{ DrawBoardFadeOut(*ebiten.Image) bool }
}

type GameOver struct {
	D        GameOverDeps
	animDone bool
}

func (s *GameOver) Enter() {
	s.animDone = false
}

func (s *GameOver) Exit() {}

func (s *GameOver) Update() error { return nil }

func (s *GameOver) Draw(screen *ebiten.Image) {
	if !s.animDone {
		s.animDone = s.D.Board.DrawBoardFadeOut(screen)
	} else {
		s.D.Menu.DrawGameOver(screen)
	}
}
