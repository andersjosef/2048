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
	fontSize               int = 50
	fontSizeSmall          int = 35
)

func initText() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72 // Try adjusting this value for high-res displays
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(fontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	mplusNormalFontSmaller, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(fontSizeSmall),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	mplusBigFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    50,
		DPI:     dpi,
		Hinting: font.HintingFull, // Use quantization to save glyph cache images.
	})
	if err != nil {
		log.Fatal(err)
	}

	// Adjust the line height.
	mplusBigFont = text.FaceWithLineHeight(mplusBigFont, 54)
}
