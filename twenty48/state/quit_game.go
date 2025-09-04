package state

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// ATM only quits game, but can be expanded later
type QuitGame struct {
}

func (s *QuitGame) Enter() {}

func (s *QuitGame) Exit() {}

func (s *QuitGame) Update() error {
	return ebiten.Termination
}

func (s *QuitGame) Draw(screen *ebiten.Image) {
}
