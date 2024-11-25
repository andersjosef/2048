package theme

import (
	"log"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type FontSet struct {
	Normal  font.Face
	Smaller font.Face
	Mini    font.Face
	Big     font.Face
}

const (
	dpi           float64 = 72
	fontSize      int     = 50
	fontSizeSmall int     = 35
	fontSizeMini  int     = 25
)

func InitFonts(scale float64) (*FontSet, error) {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		return nil, err
	}

	initializeFont := func(size int) font.Face {
		face, err := opentype.NewFace(tt, &opentype.FaceOptions{
			Size:    float64(size) * scale,
			DPI:     dpi,
			Hinting: font.HintingFull,
		})
		if err != nil {
			log.Fatal(err)
		}
		return face
	}
	return &FontSet{
		Normal:  initializeFont(fontSize),
		Smaller: initializeFont(fontSizeSmall),
		Mini:    initializeFont(fontSizeMini),
		Big:     text.FaceWithLineHeight(initializeFont(fontSize), 1.08),
	}, nil
}
