package state

import (
	"github.com/andersjosef/2048/twenty48/theme"
	"github.com/hajimehoshi/ebiten/v2"
)

type DepsInstructions struct {
	G interface {
		DrawUI(*ebiten.Image)
		GetCurrentTheme() theme.Theme
	}
	Menu interface {
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
	screen.Fill(s.D.G.GetCurrentTheme().ColorScreenBackground)
	s.D.Menu.Draw(screen)
	s.D.G.DrawUI(screen)
}
