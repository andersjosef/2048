package buttons

import (
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/input"
	"github.com/andersjosef/2048/twenty48/theme"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Deps struct {
	ScreenControl
	Input
	Utils

	GetFontSet func() theme.FontSet
	GetState   func() co.GameState
}

type ScreenControl interface {
	GetActualSize() (x, y int)
}

type Input interface {
	CheckTapped() bool
	GetTaps() []input.Tap
	ClearTaps()
}

type Utils interface {
	DrawDoubleText(screen *ebiten.Image, message string, xpos int, ypos int, offset float64, fontUsed *text.GoTextFace, isCentered bool)
}
