package twenty48

import (
	"image/color"
)

var BEIGE = [4]uint8{232, 220, 202, 255}
var DARKMODE_BEIGE = [4]uint8{0, 100, 102, 255}

func getColor(colorList [4]uint8) color.RGBA {
	c := color.RGBA{colorList[0], colorList[1], colorList[2], colorList[3]}
	return c
}
