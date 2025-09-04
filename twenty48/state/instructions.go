package state

import (
	"github.com/andersjosef/2048/twenty48/theme"
	"github.com/hajimehoshi/ebiten/v2"
)

type DepsInstructions struct {
	G interface {
		DrawMenu(*ebiten.Image)
		DrawUI(*ebiten.Image)
		GetCurrentTheme() theme.Theme
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
	s.D.G.DrawMenu(screen)
	s.D.G.DrawUI(screen)
}
