package theme

import (
	"bytes"
	"log"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type FontSet struct {
	Normal  *text.GoTextFace
	Smaller *text.GoTextFace
	Mini    *text.GoTextFace
	Big     *text.GoTextFace
}

const (
	dpi           float64 = 72
	fontSizeBig   int     = 54
	fontSize      int     = 50
	fontSizeSmall int     = 35
	fontSizeMini  int     = 25
)

//go:embed fonts/Rubik-Regular.ttf
var rubikTTF []byte

func InitFonts(scale float64) *FontSet {
	faceSource, err := text.NewGoTextFaceSource(bytes.NewReader(rubikTTF))
	if err != nil {
		log.Fatal(err)
	}

	dpiScale := ebiten.Monitor().DeviceScaleFactor()

	initializeFont := func(size int) *text.GoTextFace {
		face := &text.GoTextFace{
			Source: faceSource,
			Size:   float64(size) * scale * dpiScale,
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
	}
}
