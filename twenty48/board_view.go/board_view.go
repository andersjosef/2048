package board_view

import (
	"fmt"

	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/shadertools"
	"github.com/andersjosef/2048/twenty48/theme"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type BoardView struct {
	d BoardViewDeps

	sizes             *Sizes
	boardImage        *ebiten.Image
	boardImageOptions *ebiten.DrawImageOptions
	boardForEndScreen *ebiten.Image
}

func NewBoardView(d BoardViewDeps) *BoardView {
	bv := &BoardView{
		d: d,
	}
	// TODO: to be moved to its own place since not related to logic of board
	bv.sizes = InitSizes(bv, SizesDeps{
		EventHandler:  d.EventHandler,
		ScreenControl: d.ScreenControl,
	})

	// create boardImage
	bv.CreateBoardImage()

	return bv
}

func (b *BoardView) CreateBoardImage() {
	var (
		sizeX int = int(float64((co.BOARDSIZE * int(b.sizes.tileSize)) + (int(b.sizes.bordersize) * 2)))
		sizeY     = sizeX
	)
	b.boardImage = ebiten.NewImage(sizeX, sizeY)
	for y := range co.BOARDSIZE {
		for x := range co.BOARDSIZE {
			b.DrawBorderBackground(b.boardImage, float32(x)*b.sizes.tileSize, float32(y)*b.sizes.tileSize)
		}

	}
	b.boardImageOptions = &ebiten.DrawImageOptions{}
	b.boardImageOptions.GeoM.Translate(float64(b.sizes.startPosX), float64(b.sizes.startPosY))

	// Will update the size of it for screensize changes
	b.initBoardForEndScreen()
}

func (b *BoardView) Draw(screen *ebiten.Image) {
	// draw the backgroundimage of the game
	b.boardForEndScreen.DrawImage(b.boardImage, b.boardImageOptions)

	matrix := b.d.Board.CurMatrixSnapshot()
	// draw tiles
	for y := range len(matrix) {
		for x := range len(matrix[0]) {
			b.DrawTile(b.boardForEndScreen, b.sizes.startPosX, b.sizes.startPosY, x, y, matrix[y][x], 0, 0)
		}
	}
	if !b.d.IsGameOver() {
		screen.DrawImage(b.boardForEndScreen, &ebiten.DrawImageOptions{})

	} else {
		b.DrawBoardFadeOut(screen)
	}
}

func (b *BoardView) initBoardForEndScreen() {
	width, height := b.d.GetActualSize()
	b.boardForEndScreen = ebiten.NewImage(width, height)
}

func (b *BoardView) DrawBoardFadeOut(screen *ebiten.Image) bool {
	newImage, isDone := shadertools.GetImageFadeOut(b.boardForEndScreen)
	if isDone {
		return true
	}
	screen.DrawImage(newImage, &ebiten.DrawImageOptions{})
	return false
}

// draws one tile of the game with everything background, number, color, etc.
func (b *BoardView) DrawTile(screen *ebiten.Image, startX, startY float32, x, y int, value int, movDistX, movDistY float32) {
	var (
		xpos float32 = (startX + float32(x)*b.sizes.tileSize + movDistX*b.sizes.tileSize)
		ypos float32 = (startY + float32(y)*b.sizes.tileSize + movDistY*b.sizes.tileSize)
	)

	if value != 0 {
		// Set tile color to default color
		colorMap := b.d.Theme.Current().ColorMap

		val, ok := colorMap[value] // checks if num in map, if it is make the background else draw normal

		if ok { // If the key exists draw the coresponding color background
			b.DrawNumberBackground(screen, startX, startY, y, x, val, movDistX, movDistY)
		}
		b.DrawText(screen, xpos, ypos, x, y, value)
	}
}

func (b *BoardView) DrawBorderBackground(screen *ebiten.Image, xpos, ypos float32) {
	var sizeBorder float32 = (float32(b.sizes.tileSize) + b.sizes.bordersize)
	var sizeInside float32 = (b.sizes.tileSize - b.sizes.bordersize)

	vector.DrawFilledRect(screen, xpos, ypos,
		sizeBorder, sizeBorder, b.d.Theme.Current().ColorBorder, false) //outer
	vector.DrawFilledRect(screen, xpos+b.sizes.bordersize, ypos+b.sizes.bordersize,
		sizeInside, sizeInside, b.d.Theme.Current().ColorBackgroundTile, false) // inner
}

// background of a number, since they have colors
func (b *BoardView) DrawNumberBackground(screen *ebiten.Image, startX, startY float32, y, x int, val [4]uint8, movDistX, movDistY float32) {
	var (
		xpos      float32 = (startX + float32(x)*b.sizes.tileSize + b.sizes.bordersize + movDistX*b.sizes.tileSize)
		ypos      float32 = (startY + float32(y)*b.sizes.tileSize + b.sizes.bordersize + movDistY*b.sizes.tileSize)
		size_tile float32 = (float32(b.sizes.tileSize) - b.sizes.bordersize)
	)
	vector.DrawFilledRect(screen, xpos, ypos,
		size_tile, size_tile, theme.GetColor(val), false) // tiles
}

func (b *BoardView) DrawText(screen *ebiten.Image, xpos, ypos float32, x, y int, value int) {
	fontSet := b.d.Fonts()
	msg := fmt.Sprintf("%v", value)

	var fontUsed *text.GoTextFace
	if float32(text.Advance(msg, fontSet.Big)) > b.sizes.tileSize {
		fontUsed = fontSet.Smaller
	} else {
		fontUsed = fontSet.Normal
	}

	width, height := text.Measure(msg, fontUsed, 0)

	dx := float32(width)
	dy := float32(height)

	textPosX := int(xpos + (b.sizes.bordersize/2 + b.sizes.tileSize/2) - dx/2)
	textPosY := int(ypos + (b.sizes.bordersize/2 + b.sizes.tileSize/2) - dy/2)

	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(textPosX), float64(textPosY))
	op.ColorScale.ScaleWithColor(b.d.Theme.Current().ColorText)
	text.Draw(screen, msg, fontUsed, op)
}
