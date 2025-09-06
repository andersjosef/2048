package board_view

import (
	"strconv"

	"github.com/andersjosef/2048/twenty48/eventhandler"
	"github.com/andersjosef/2048/twenty48/shadertools"
	"github.com/andersjosef/2048/twenty48/theme"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type BoardView struct {
	d BoardViewDeps

	emptyBoard *ebiten.Image
	opts       *ebiten.DrawImageOptions
	tiles      map[int]*ebiten.Image

	BoardSnapshot *ebiten.Image // For making it dissapear in the game over
	endOpts       *ebiten.DrawImageOptions
}

func NewBoardView(d BoardViewDeps) *BoardView {
	bv := &BoardView{
		d:       d,
		opts:    &ebiten.DrawImageOptions{},
		endOpts: &ebiten.DrawImageOptions{},
	}

	// create boardImage
	// bv.CreateBoardImage()
	bv.RebuildBoard()

	bv.d.EventHandler.Register(
		eventhandler.EventScreenChanged,
		func(eventhandler.Event) {
			bv.d.Layout.Recalculate()
			bv.RebuildBoard()
		},
	)

	bv.d.EventHandler.Register(
		eventhandler.EventThemeChanged,
		func(eventhandler.Event) {
			bv.d.Layout.Recalculate()
			bv.RebuildBoard()
		},
	)
	return bv
}

func (b *BoardView) RebuildBoard() {
	tileSize, borderSize := b.d.Layout.TileSize(), b.d.Layout.BorderSize()

	// Background image
	l, h := b.d.Board.GetBoardDimentions()
	sizeX := int(float32(l)*tileSize + 2*borderSize)
	sizeY := int(float32(h)*tileSize + 2*borderSize)
	b.emptyBoard = ebiten.NewImage(sizeX, sizeY)

	// Empty tiles
	for y := range h {
		for x := range l {
			b.DrawBorderBackground(
				b.emptyBoard,
				float32(x)*tileSize,
				float32(y)*tileSize,
			)
		}
	}

	// The color tiles
	b.tiles = make(map[int]*ebiten.Image, 15)
	for v, rgba := range b.d.Theme.Current().ColorMap {
		innerSize := int(tileSize - borderSize)
		img := ebiten.NewImage(innerSize, innerSize)
		img.Fill(theme.GetColor(rgba))
		b.tiles[v] = img
	}

	b.opts.GeoM.Reset() // Set the new opts
	x, y := b.d.Layout.StartPos()
	b.opts.GeoM.Translate(float64(x), float64(y))

	b.initBoardForEndScreen()
}

// func (b *BoardView) scaleBoard() {
// 	newOpt := &ebiten.DrawImageOptions{}
// 	x, y := b.d.Layout.GetStartPos()
// 	newOpt.GeoM.Translate(float64(x), float64(y))
// 	b.opts = newOpt
// 	b.rebuildBoard()
// }

func (b *BoardView) Draw(screen *ebiten.Image) {
	// // Draw onto the snapshot so it contains both the empty board and tiles
	// b.BoardSnapshot.DrawImage(b.emptyBoard, b.opts)
	// b.drawTiles(b.BoardSnapshot)
	// screen.DrawImage(b.BoardSnapshot, &ebiten.DrawImageOptions{})

	themeSnap := b.d.Theme.Current()
	tileSize, borderSize := b.d.Layout.TileSize(), b.d.Layout.BorderSize()
	startX, startY := b.d.Layout.GetStartPos()

	// Empty board
	b.BoardSnapshot.DrawImage(b.emptyBoard, b.opts)

	// Tiles and numbers
	mat := b.d.Board.CurMatrixSnapshot()
	length, height := b.d.Board.GetBoardDimentions()
	var textOps text.DrawOptions
	for y := range height {
		for x := range length {
			val := mat[y][x]
			if val == 0 {
				continue
			}

			// Colored tile
			if img, ok := b.tiles[val]; ok {
				var o ebiten.DrawImageOptions
				o.GeoM.Translate(float64(startX+float32(x)*tileSize+borderSize), float64(startY+float32(y)*tileSize+borderSize))
				b.BoardSnapshot.DrawImage(img, &o)
			}

			// Text
			msg := strconv.Itoa(val)
			font := b.pickFont(msg, tileSize)
			width, height := text.Measure(msg, font, 0)
			tx := float64(startX + float32(x)*tileSize + borderSize/2 + (tileSize-float32(width))/2)
			ty := float64(startY + float32(y)*tileSize + borderSize/2 + (tileSize-float32(height))/2)

			textOps.GeoM.Reset()
			textOps.GeoM.Translate(tx, ty)
			textOps.ColorScale.Reset()
			textOps.ColorScale.ScaleWithColor(themeSnap.ColorText)
			text.Draw(b.BoardSnapshot, msg, font, &textOps)

		}
	}
	screen.DrawImage(b.BoardSnapshot, b.endOpts)
}

func (b *BoardView) pickFont(s string, size float32) *text.GoTextFace {
	fontSet := b.d.Fonts()
	var fontUsed *text.GoTextFace
	if float32(text.Advance(s, fontSet.Big)) > size {
		fontUsed = fontSet.Smaller
	} else {
		fontUsed = fontSet.Normal
	}

	return fontUsed
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
	startX, startY := b.d.Layout.StartPos()
	tileSize := b.d.Layout.TileSize()
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

func (b *BoardView) DrawBorderBackground(img *ebiten.Image, xpos, ypos float32) {
	tileSize := b.d.Layout.TileSize()
	borderSize := b.d.Layout.BorderSize()

	sizeBorder := tileSize + borderSize
	sizeInside := tileSize - borderSize

	vector.DrawFilledRect(img, xpos, ypos,
		sizeBorder, sizeBorder, b.d.Theme.Current().ColorBorder, false) //outer
	vector.DrawFilledRect(img, xpos+borderSize, ypos+borderSize,
		sizeInside, sizeInside, b.d.Theme.Current().ColorBackgroundTile, false) // inner
}

// background of a number, since they have colors
func (b *BoardView) DrawNumberBackground(screen *ebiten.Image, startX, startY float32, y, x int, val [4]uint8, movDistX, movDistY float32) {
	tileSize := b.d.Layout.TileSize()
	borderSize := b.d.Layout.BorderSize()

	xpos := startX + float32(x)*tileSize + borderSize + movDistX*tileSize
	ypos := startY + float32(y)*tileSize + borderSize + movDistY*tileSize
	size_tile := tileSize - borderSize

	vector.DrawFilledRect(screen, xpos, ypos,
		size_tile, size_tile, theme.GetColor(val), false) // tiles
}

func (b *BoardView) DrawText(screen *ebiten.Image, xpos, ypos float32, x, y int, value int) {
	msg := strconv.Itoa(value)

	tileSize := b.d.Layout.TileSize()
	borderSize := b.d.Layout.BorderSize()

	// var fontUsed *text.GoTextFace
	// if float32(text.Advance(msg, fontSet.Big)) > tileSize {
	// 	fontUsed = fontSet.Smaller
	// } else {
	// 	fontUsed = fontSet.Normal
	// }
	fontUsed := b.pickFont(msg, tileSize)

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
