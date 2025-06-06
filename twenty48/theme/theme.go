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
		DarkTheme,
		ManMachineTheme,
		LightTheme,
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

var LightTheme = Theme{
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

var DarkTheme = Theme{
	Name:                  "Dark",
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
var ManMachineTheme = Theme{
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

func GetColor(colorList [4]uint8) color.RGBA {
	c := color.RGBA{colorList[0], colorList[1], colorList[2], colorList[3]}
	return c
}
