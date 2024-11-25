package renderer

import (
	"image/color"

	"github.com/andersjosef/2048/twenty48/theme"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Renderer struct {
	scale float64
}

func InitRenderer(fontSet *theme.FontSet, scale float64) *Renderer {
	r := &Renderer{
		scale: scale,
	}

	return r
}

func (r *Renderer) DrawDoubleText(screen *ebiten.Image, message string, xpos int, ypos int, offset float64, fontUsed *text.GoTextFace, isCentered bool) {

	scale := r.scale

	// Calculate text dimensions
	textWidth := text.Advance(message, fontUsed)
	textHeight := -(fontUsed.Metrics().VAscent + fontUsed.Metrics().VDescent)

	// Scale the position
	baseX := scale * float64(xpos)
	baseY := scale * float64(ypos)

	// Adjust for centering
	if isCentered {
		baseX -= textWidth / 2  // Center horizontally
		baseY += textHeight / 2 // Center vertically
	} else {
		baseY -= textHeight / 4 // Center vertically
	}

	// Set options for shadow text
	shadowOpt := &text.DrawOptions{}
	shadowOpt.GeoM.Translate(baseX, baseY)
	shadowOpt.ColorScale.ScaleWithColor(color.Black)

	// Set options for main text
	mainOpt := &text.DrawOptions{}
	mainOpt.GeoM.Translate(
		baseX-offset*scale,
		baseY-offset*scale)
	mainOpt.ColorScale.ScaleWithColor(color.White)

	// Draw shadow and main text
	text.Draw(screen, message, fontUsed, shadowOpt)
	text.Draw(screen, message, fontUsed, mainOpt)
}
