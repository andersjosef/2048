package state

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type DepsInstructions struct {
	Menu interface {
		Draw(*ebiten.Image)
	}
	Buttons interface {
		Draw(*ebiten.Image)
	}
}

type Instructions struct {
	D DepsInstructions
}

func (s *Instructions) Enter() {}

func (s *Instructions) Exit() {}

func (s *Instructions) Update() error {
	return nil
}

func (s *Instructions) Draw(screen *ebiten.Image) {
	s.D.Menu.Draw(screen)
	s.D.Buttons.Draw(screen)
}
