package utils

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Utils struct {
}

func New() *Utils {
	r := &Utils{}

	return r
}

func (r *Utils) DrawDoubleText(screen *ebiten.Image, message string, xpos int, ypos int, offset float64, fontUsed *text.GoTextFace, isCentered bool) {

	// Calculate text dimensions
	textWidth, textHeight := text.Measure(message, fontUsed, 0)

	baseX := float64(xpos)
	baseY := float64(ypos)

	// Adjust for centering
	if isCentered {
		baseX -= textWidth / 2  // Center horizontally
		baseY -= textHeight / 2 // Center vertically
	}

	// Set options for shadow text
	shadowOpt := &text.DrawOptions{}
	shadowOpt.GeoM.Translate(baseX, baseY)
	shadowOpt.ColorScale.ScaleWithColor(color.Black)

	// Set options for main text
	mainOpt := &text.DrawOptions{}
	mainOpt.GeoM.Translate(
		baseX-offset,
		baseY-offset)
	mainOpt.ColorScale.ScaleWithColor(color.White)

	// Draw shadow and main text
	text.Draw(screen, message, fontUsed, shadowOpt)
	text.Draw(screen, message, fontUsed, mainOpt)
}
