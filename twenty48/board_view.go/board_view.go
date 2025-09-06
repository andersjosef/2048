package board_view

import (
	"fmt"

	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/eventhandler"
	"github.com/andersjosef/2048/twenty48/shadertools"
	"github.com/andersjosef/2048/twenty48/theme"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type BoardView struct {
	d BoardViewDeps

	// sizes      *Sizes
	EmptyBoard *ebiten.Image
	boardOpts  *ebiten.DrawImageOptions

	BoardSnapshot *ebiten.Image // For making it dissapear in the game over
}

func NewBoardView(d BoardViewDeps) *BoardView {
	bv := &BoardView{
		d: d,
	}

	// create boardImage
	bv.CreateBoardImage()

	bv.d.EventHandler.Register(
		eventhandler.EventScaleBoardView,
		func(eventhandler.Event) {
			bv.scaleBoard()
		},
	)

	return bv
}

func (b *BoardView) CreateBoardImage() {
	sizeX := int(float64((co.BOARDSIZE * int(b.d.Sizes.TileSize())) + (int(b.d.Sizes.BorderSize()) * 2)))
	sizeY := sizeX

	b.EmptyBoard = ebiten.NewImage(sizeX, sizeY)
	length, height := b.d.Board.GetBoardDimentions()
	for y := range height {
		for x := range length {
			b.DrawBorderBackground(
				float32(x)*b.d.Sizes.TileSize(),
				float32(y)*b.d.Sizes.TileSize(),
			)
		}

	}
	x, y := b.d.Sizes.GetStartPos()
	b.boardOpts = &ebiten.DrawImageOptions{}
	b.boardOpts.GeoM.Translate(float64(x), float64(y))

	// Will update the size of it for screensize changes
	b.initBoardForEndScreen()
}

func (b *BoardView) scaleBoard() {
	newOpt := &ebiten.DrawImageOptions{}
	x, y := b.d.Sizes.GetStartPos()
	newOpt.GeoM.Translate(float64(x), float64(y))
	b.boardOpts = newOpt
	b.CreateBoardImage()
}

func (b *BoardView) Draw(screen *ebiten.Image) {
	// Draw onto the snapshot so it contains both the empty board and tiles
	b.BoardSnapshot.DrawImage(b.EmptyBoard, b.boardOpts)
	b.drawTiles(b.BoardSnapshot)
	screen.DrawImage(b.BoardSnapshot, &ebiten.DrawImageOptions{})
}

// Draw tiles
func (b *BoardView) drawTiles(img *ebiten.Image) {
	matrix := b.d.Board.CurMatrixSnapshot()
	length, height := b.d.Board.GetBoardDimentions()
	for y := range height {
		for x := range length {
			b.drawTile(
				img,
				x, y, matrix[y][x], 0, 0)
		}
	}

}

func (b *BoardView) initBoardForEndScreen() {
	width, height := b.d.GetActualSize()
	b.BoardSnapshot = ebiten.NewImage(width, height)
}

func (b *BoardView) DrawBoardFadeOut(screen *ebiten.Image) bool {
	newImage, isDone := shadertools.GetImageFadeOut(b.BoardSnapshot)
	if isDone {
		return true
	}
	screen.DrawImage(newImage, &ebiten.DrawImageOptions{})
	return false
}

// draws one tile of the game with everything background, number, color, etc.
func (b *BoardView) drawTile(screen *ebiten.Image, x, y int, value int, movDistX, movDistY float32) {
	startX, startY := b.d.Sizes.StartPos()
	tileSize := b.d.Sizes.TileSize()
	xpos := startX + (float32(x)+movDistX)*tileSize
	ypos := startY + (float32(y)+movDistY)*tileSize

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

func (b *BoardView) DrawBorderBackground(xpos, ypos float32) {
	tileSize := b.d.Sizes.TileSize()
	borderSize := b.d.Sizes.BorderSize()

	sizeBorder := tileSize + borderSize
	sizeInside := tileSize - borderSize

	screen := b.EmptyBoard

	vector.DrawFilledRect(screen, xpos, ypos,
		sizeBorder, sizeBorder, b.d.Theme.Current().ColorBorder, false) //outer
	vector.DrawFilledRect(screen, xpos+borderSize, ypos+borderSize,
		sizeInside, sizeInside, b.d.Theme.Current().ColorBackgroundTile, false) // inner
}

// background of a number, since they have colors
func (b *BoardView) DrawNumberBackground(screen *ebiten.Image, startX, startY float32, y, x int, val [4]uint8, movDistX, movDistY float32) {
	tileSize := b.d.Sizes.TileSize()
	borderSize := b.d.Sizes.BorderSize()

	xpos := startX + float32(x)*tileSize + borderSize + movDistX*tileSize
	ypos := startY + float32(y)*tileSize + borderSize + movDistY*tileSize
	size_tile := tileSize - borderSize

	vector.DrawFilledRect(screen, xpos, ypos,
		size_tile, size_tile, theme.GetColor(val), false) // tiles
}

func (b *BoardView) DrawText(screen *ebiten.Image, xpos, ypos float32, x, y int, value int) {
	fontSet := b.d.Fonts()
	msg := fmt.Sprintf("%v", value)

	tileSize := b.d.Sizes.TileSize()
	borderSize := b.d.Sizes.BorderSize()

	var fontUsed *text.GoTextFace
	if float32(text.Advance(msg, fontSet.Big)) > tileSize {
		fontUsed = fontSet.Smaller
	} else {
		fontUsed = fontSet.Normal
	}

	width, height := text.Measure(msg, fontUsed, 0)

	dx := float32(width)
	dy := float32(height)

	textPosX := int(xpos + (borderSize/2 + tileSize/2) - dx/2)
	textPosY := int(ypos + (borderSize/2 + tileSize/2) - dy/2)

	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(textPosX), float64(textPosY))
	op.ColorScale.ScaleWithColor(b.d.Theme.Current().ColorText)
	text.Draw(screen, msg, fontUsed, op)
}
