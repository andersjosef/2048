package theme

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type FontSet struct {
	Normal  text.GoTextFace
	Smaller text.GoTextFace
	Mini    text.GoTextFace
	Big     text.GoTextFace
}

const (
	dpi           float64 = 72
	fontSizeBig   int     = 54
	fontSize      int     = 50
	fontSizeSmall int     = 35
	fontSizeMini  int     = 25
)

func InitFonts(scale float64) (*FontSet, error) {
	mplusFaceSource, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}

	initializeFont := func(size int) text.GoTextFace {
		face := text.GoTextFace{
			Source: mplusFaceSource,
			Size:   float64(size) * scale,
		}
		if err != nil {
			log.Fatal(err)
		}
		return face
	}
	return &FontSet{
		Normal:  initializeFont(fontSize),
		Smaller: initializeFont(fontSizeSmall),
		Mini:    initializeFont(fontSizeMini),
		Big:     initializeFont(fontSizeBig),
	}, nil
}
