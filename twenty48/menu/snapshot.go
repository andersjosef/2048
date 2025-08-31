package menu

import (
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/theme"
)

type Snapshot struct {
	State         co.GameState
	PreviousState co.GameState
	CurrentTheme  theme.Theme
	Fonts         theme.FontSet
	Score         int
	Widht, Height int
	IsFullScreen  bool
}
