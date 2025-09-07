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
	b.initBoardForEndScreen()

	// Empty tiles
	for y := range h {
		for x := range l {
			b.drawBorderBackground(
				b.emptyBoard,
				float32(x)*tileSize,
				float32(y)*tileSize,
			)
		}
	}

	// The color tiles
	b.tiles = make(map[int]*ebiten.Image, 15)
	themeSnap := b.d.Theme.Current()
	var textOps text.DrawOptions
	for v, rgba := range b.d.Theme.Current().ColorMap {
		innerSize := int(tileSize - borderSize)
		img := ebiten.NewImage(innerSize, innerSize)

		// Background
		img.Fill(theme.GetColor(rgba))

		// Text
		msg := strconv.Itoa(v)
		font := b.pickFont(msg, tileSize)
		width, height := text.Measure(msg, font, 0)
		tx := float64((tileSize-float32(width))/2 - borderSize/2)
		ty := float64((tileSize-float32(height))/2 - borderSize/2)

		textOps.GeoM.Reset()
		textOps.GeoM.Translate(tx, ty)
		textOps.ColorScale.Reset()
		textOps.ColorScale.ScaleWithColor(themeSnap.ColorText)
		text.Draw(img, msg, font, &textOps)

		// Store
		b.tiles[v] = img
	}

	b.opts.GeoM.Reset() // Set the new opts
	x, y := b.d.Layout.StartPos()
	b.opts.GeoM.Translate(float64(x), float64(y))

}

func (b *BoardView) Draw(screen *ebiten.Image) {
	tileSize, borderSize := b.d.Layout.TileSize(), b.d.Layout.BorderSize()
	startX, startY := b.d.Layout.GetStartPos()

	// Empty board
	b.BoardSnapshot.DrawImage(b.emptyBoard, b.opts)

	// Tiles and numbers
	mat := b.d.Board.CurMatrixSnapshot()
	length, height := b.d.Board.GetBoardDimentions()
	for y := range height {
		for x := range length {
			val := mat[y][x]
			if val == 0 {
				continue
			}

			// Draw tile from map
			if img, ok := b.tiles[val]; ok {
				var o ebiten.DrawImageOptions
				o.GeoM.Translate(float64(startX+float32(x)*tileSize+borderSize), float64(startY+float32(y)*tileSize+borderSize))
				b.BoardSnapshot.DrawImage(img, &o)
			}
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

func (b *BoardView) drawBorderBackground(img *ebiten.Image, xpos, ypos float32) {
	tileSize := b.d.Layout.TileSize()
	borderSize := b.d.Layout.BorderSize()

	sizeBorder := tileSize + borderSize
	sizeInside := tileSize - borderSize

	vector.DrawFilledRect(img, xpos, ypos,
		sizeBorder, sizeBorder, b.d.Theme.Current().ColorBorder, false) //outer
	vector.DrawFilledRect(img, xpos+borderSize, ypos+borderSize,
		sizeInside, sizeInside, b.d.Theme.Current().ColorBackgroundTile, false) // inner
}
