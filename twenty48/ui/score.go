package ui

import (
	"fmt"
	"image/color"

	"github.com/andersjosef/2048/twenty48/theme"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type ScoreOverlayDeps struct {
	Fonts interface{ Fonts() *theme.FontSet }
	Score interface{ Score() int }
}

type ScoreOverlay struct {
	d ScoreOverlayDeps
}

func NewScoreOverlay(d ScoreOverlayDeps) *ScoreOverlay {
	return &ScoreOverlay{
		d: d,
	}
}

func (so *ScoreOverlay) DrawScore(screen *ebiten.Image) {
	myFont := so.d.Fonts.Fonts().Smaller

	margin := 10
	shadowOffsett := 2
	score_text := fmt.Sprintf("%v", so.d.Score.Score())

	getOpt := func(x, y float64, col color.Color) *text.DrawOptions {
		opt := &text.DrawOptions{}
		opt.GeoM.Translate(x, y)
		opt.ColorScale.ScaleWithColor(col)
		return opt
	}

	// TODO: Make colors part of the themes and modular, taken from themeManager
	shadowOpt := getOpt(float64((shadowOffsett + margin)), 0, color.Black)
	text.Draw(screen, score_text, myFont, shadowOpt)

	mainOpt := getOpt(float64(margin), 0, color.White)
	text.Draw(screen, score_text, myFont, mainOpt)
}
