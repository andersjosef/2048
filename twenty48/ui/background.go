package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Background struct{ Color func() color.RGBA }

func (b Background) Draw(screen *ebiten.Image) {
	screen.Fill(b.Color())
}
