package twenty48

import (
	"log"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	mplusNormalFont        font.Face
	mplusNormalFontSmaller font.Face
	mplusBigFont           font.Face
)

const (
	dpi           float64 = 72 // Try adjusting this value for high-res displays
	fontSize      int     = 50
	fontSizeSmall int     = 35
)

func initText(g *Game) {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	initializeFont := func(size int, g *Game) font.Face {
		face, err := opentype.NewFace(tt, &opentype.FaceOptions{
			Size:    float64(size) * g.scale,
			DPI:     dpi,
			Hinting: font.HintingFull,
		})
		if err != nil {
			log.Fatal(err)
		}
		return face
	}

	mplusNormalFont = initializeFont(fontSize, g)
	mplusNormalFontSmaller = initializeFont(fontSizeSmall, g)
	mplusBigFont = text.FaceWithLineHeight(initializeFont(fontSize, g), 1.08)
}
