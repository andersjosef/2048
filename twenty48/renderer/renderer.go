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

func (r *Renderer) DrawDoubleText(screen *ebiten.Image, message string, xpos int, ypos int, offset int, fontUsed *text.GoTextFace, isCentered bool) {

	scale := r.scale

	// Calculate text dimensions
	textWidth := int(text.Advance(message, fontUsed))
	textHeight := -int(fontUsed.Metrics().VAscent + fontUsed.Metrics().VDescent)

	// Scale the position
	textPosX := int(scale) * xpos
	textPosY := int(scale) * ypos

	// Adjust for centering
	if isCentered {
		textPosX -= textWidth / 2  // Center horizontally
		textPosY += textHeight / 2 // Center vertically
	} else {
		textPosY -= textHeight / 4 // Center vertically
	}

	// Draw shadow (black text)
	shadowOpt := &text.DrawOptions{}
	shadowOpt.GeoM.Translate(float64(textPosX), float64(textPosY))
	shadowOpt.ColorScale.ScaleWithColor(color.Black)

	text.Draw(screen, message, fontUsed, shadowOpt)

	// Draw main text (white text) with offset
	mainOpt := &text.DrawOptions{}
	mainOpt.GeoM.Translate(
		float64(textPosX-int(float64(offset)*scale)),
		float64(textPosY-int(float64(offset)*scale)))
	mainOpt.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen, message, fontUsed,
		mainOpt)
}
