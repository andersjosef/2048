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

func initText() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	initializeFont := func(size int) font.Face {
		face, err := opentype.NewFace(tt, &opentype.FaceOptions{
			Size:    float64(size),
			DPI:     dpi,
			Hinting: font.HintingFull,
		})
		if err != nil {
			log.Fatal(err)
		}
		return face
	}

	mplusNormalFont = initializeFont(fontSize)
	mplusNormalFontSmaller = initializeFont(fontSizeSmall)
	mplusBigFont = text.FaceWithLineHeight(initializeFont(fontSize), 1.08)
}
