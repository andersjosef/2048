package theme

import (
	"image/color"
)

type ThemePicker struct {
	themeSlice []Theme
	index      uint

	currentTheme Theme
}

func NewThemePicker() *ThemePicker {
	tp := &ThemePicker{}

	tp.themeSlice = []Theme{
		greenTheme,
		manMachineTheme,
		darkTheme,
		monoTheme,
		lightTheme,
	}

	tp.currentTheme = tp.themeSlice[0]

	return tp
}

func (tp *ThemePicker) GetCurrentTheme() Theme {
	return tp.currentTheme
}

func (tp *ThemePicker) IncrementCurrentTheme() Theme {
	tp.index++
	tp.currentTheme = tp.themeSlice[tp.index%uint(len(tp.themeSlice))]
	return tp.currentTheme
}

type Theme struct {
	Name                  string
	ColorText             color.RGBA
	ColorBorder           color.RGBA
	ColorBackgroundTile   color.RGBA
	ColorScreenBackground color.RGBA
	ColorMap              map[int][4]uint8
}

var lightTheme = Theme{
	Name:                  "Light",
	ColorText:             color.RGBA{110, 93, 71, 255},
	ColorBorder:           color.RGBA{194, 182, 169, 255},
	ColorBackgroundTile:   color.RGBA{204, 192, 179, 255},
	ColorScreenBackground: color.RGBA{212, 200, 182, 255},
	ColorMap: map[int][4]uint8{
		2:     {238, 228, 218, 255},
		4:     {237, 224, 200, 255},
		8:     {242, 177, 121, 255},
		16:    {245, 149, 99, 255},
		32:    {255, 104, 69, 255},
		64:    {246, 94, 59, 255},
		128:   {237, 207, 114, 255},
		256:   {237, 205, 100, 255},
		512:   {237, 204, 97, 255},
		1024:  {237, 200, 80, 255},
		2048:  {237, 197, 63, 255},
		4096:  {149, 189, 126, 255},
		8192:  {107, 127, 95, 255},
		16384: {247, 104, 104, 255},
		-1:    {255, 255, 255, 255},
	},
}

var greenTheme = Theme{
	Name:                  "Green",
	ColorText:             color.RGBA{110, 93, 71, 255},
	ColorBorder:           color.RGBA{154, 142, 129, 255},
	ColorBackgroundTile:   color.RGBA{164, 152, 139, 255},
	ColorScreenBackground: color.RGBA{0, 100, 102, 255},
	ColorMap: map[int][4]uint8{
		2:     {218, 208, 198, 255},
		4:     {217, 204, 180, 255},
		8:     {222, 157, 101, 255},
		16:    {225, 129, 79, 255},
		32:    {235, 84, 49, 255},
		64:    {226, 74, 39, 255},
		128:   {217, 187, 94, 255},
		256:   {217, 185, 80, 255},
		512:   {217, 184, 77, 255},
		1024:  {217, 180, 60, 255},
		2048:  {217, 177, 43, 255},
		4096:  {129, 169, 106, 255},
		8192:  {87, 107, 75, 255},
		16384: {227, 84, 84, 255},
		-1:    {255, 255, 255, 255},
	},
}
var manMachineTheme = Theme{
	Name:                  "Man-Machine",
	ColorText:             color.RGBA{79, 82, 45, 255},
	ColorBorder:           color.RGBA{255, 105, 180, 255},
	ColorBackgroundTile:   color.RGBA{40, 10, 50, 255},
	ColorScreenBackground: color.RGBA{15, 0, 35, 255},
	ColorMap: map[int][4]uint8{
		2:     {255, 235, 205, 255},
		4:     {255, 200, 150, 255},
		8:     {255, 150, 100, 255},
		16:    {255, 100, 50, 255},
		32:    {255, 160, 0, 255},
		64:    {255, 210, 0, 255},
		128:   {255, 240, 150, 255},
		256:   {200, 255, 150, 255},
		512:   {150, 255, 200, 255},
		1024:  {100, 200, 255, 255},
		2048:  {50, 150, 255, 255},
		4096:  {0, 100, 255, 255},
		8192:  {100, 0, 200, 255},
		16384: {200, 0, 150, 255},
		-1:    {255, 255, 255, 255},
	},
}

var monoTheme = Theme{
	Name:                  "Mono",
	ColorText:             color.RGBA{50, 50, 50, 255},
	ColorBorder:           color.RGBA{150, 150, 150, 255},
	ColorBackgroundTile:   color.RGBA{220, 220, 220, 255},
	ColorScreenBackground: color.RGBA{205, 205, 205, 255},
	ColorMap: map[int][4]uint8{
		2:     {238, 238, 238, 255},
		4:     {200, 200, 200, 255},
		8:     {170, 170, 170, 255},
		16:    {140, 140, 140, 255},
		32:    {110, 110, 110, 255},
		64:    {80, 80, 80, 255},
		128:   {60, 60, 60, 255},
		256:   {40, 40, 40, 255},
		512:   {20, 20, 20, 255},
		1024:  {0, 100, 200, 255},
		2048:  {0, 80, 160, 255},
		4096:  {0, 60, 120, 255},
		8192:  {0, 40, 80, 255},
		16384: {0, 20, 40, 255},
		-1:    {255, 255, 255, 255},
	},
}

var darkTheme = Theme{
	Name:                  "Dark",
	ColorText:             color.RGBA{220, 220, 230, 255},
	ColorBorder:           color.RGBA{90, 90, 100, 255},
	ColorBackgroundTile:   color.RGBA{30, 30, 40, 255},
	ColorScreenBackground: color.RGBA{15, 15, 20, 255},
	ColorMap: map[int][4]uint8{
		2:     {25, 59, 76, 255},
		4:     {76, 42, 25, 255},
		8:     {25, 76, 76, 255},
		16:    {76, 25, 76, 255},
		32:    {51, 76, 25, 255},
		64:    {76, 25, 25, 255},
		128:   {33, 25, 76, 255},
		256:   {76, 59, 25, 255},
		512:   {25, 76, 42, 255},
		1024:  {76, 34, 25, 255},
		2048:  {76, 25, 51, 255},
		4096:  {25, 50, 76, 255},
		8192:  {25, 76, 59, 255},
		16384: {76, 68, 25, 255},
		-1:    {255, 255, 255, 255},
	},
}

func GetColor(colorList [4]uint8) color.RGBA {
	c := color.RGBA{colorList[0], colorList[1], colorList[2], colorList[3]}
	return c
}
