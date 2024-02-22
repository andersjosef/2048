package twenty48

import (
	"image/color"
)

var BEIGE = [...]uint8{232, 220, 202, 255}

func getColor(colorList [4]uint8) color.RGBA {
	c := color.RGBA{colorList[0], colorList[1], colorList[2], colorList[3]}
	return c
}
