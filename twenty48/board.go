package twenty48

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const TILESIZE float32 = float32(SCREENWIDTH) / 6.4
const BORDERSIZE float32 = TILESIZE / 25

var (
	startPosX float32 = float32((SCREENWIDTH - (BOARDSIZE * int(TILESIZE))) / 2)
	startPosY float32 = float32((SCREENHEIGHT - (BOARDSIZE * int(TILESIZE))) / 2)
)

type Board struct {
	board               [BOARDSIZE][BOARDSIZE]int // 2d array for the board :)
	colorBorder         color.RGBA
	colorBackgroundTile color.RGBA
	game                *Game
	boardBeforeChange   [BOARDSIZE][BOARDSIZE]int
	boardImage          *ebiten.Image
	boardImageOptions   *ebiten.DrawImageOptions
}

func NewBoard(g *Game) (*Board, error) {

	b := &Board{}

	// border and background colors
	if g.darkMode { // INIT in DARK MODE
		b.colorBorder = colorBorderDarkMode
		b.colorBackgroundTile = colorBackgroundTileDarkMode

	} else { // INIT in DEFAULT MODE
		b.colorBorder = colorBorderDefault
		b.colorBackgroundTile = colorBackgroundTileDefault
	}
	b.game = g
	// add the two start pieces
	for i := 0; i < 2; i++ {
		b.randomNewPiece()
	}

	// create baordImage
	b.createBoardImage()

	return b, nil

}

func (b *Board) randomNewPiece() {

	var x, y int = len(b.board), len(b.board[0])

	for count := 0; count < x*y; count++ {
		var posX, posY int = rand.Intn(x), rand.Intn(y)
		if b.board[posX][posY] == 0 {
			if rand.Float32() > 0.16 {
				b.board[posX][posY] = 2 // 84%
			} else {
				b.board[posX][posY] = 4 // 16% chance of 4 spawning
			}
			break
		}
	}
}

func (b *Board) drawBoard(screen *ebiten.Image) {
	// draw the backgroundimage of the game
	screen.DrawImage(b.boardImage, b.boardImageOptions)

	// draw tiles
	for y := 0; y < len(b.board); y++ {
		for x := 0; x < len(b.board[0]); x++ {
			b.DrawTile(screen, startPosX, startPosY, x, y, b.board[y][x], 0, 0)
		}
	}

}

// draws one tile of the game with everything background, number, color, etc.
func (b *Board) DrawTile(screen *ebiten.Image, startX, startY float32, x, y int, value int, movDistX, movDistY float32) {
	var (
		xpos float32 = (startX + float32(x)*TILESIZE + movDistX*TILESIZE) * float32(b.game.scale)
		ypos float32 = (startY + float32(y)*TILESIZE + movDistY*TILESIZE) * float32(b.game.scale)
	)

	if value != 0 {
		// Set tile color to default color
		colorMap := colorMapDefault

		// change tile colors to dark if darkmode is activated
		if b.game.darkMode {
			colorMap = colorMapDarkMode
		}
		val, ok := colorMap[value] // checks if num in map, if it is make the background else draw normal

		if ok { // If the key exists draw the coresponding color background
			b.DrawNumberBackground(screen, startX, startY, y, x, val, movDistX, movDistY)
		}
		b.DrawText(screen, xpos, ypos, x, y, value)
	}
}

func (b *Board) DrawBorderBackground(screen *ebiten.Image, xpos, ypos float32) {
	xpos *= float32(b.game.scale)
	ypos *= float32(b.game.scale)
	var sizeBorder float32 = (float32(TILESIZE) + BORDERSIZE) * float32(b.game.scale)
	var sizeInside float32 = (TILESIZE - BORDERSIZE) * float32(b.game.scale)

	vector.DrawFilledRect(screen, xpos, ypos,
		sizeBorder, sizeBorder, b.colorBorder, false) //outer
	vector.DrawFilledRect(screen, xpos+BORDERSIZE*float32(b.game.scale), ypos+BORDERSIZE*float32(b.game.scale),
		sizeInside, sizeInside, b.colorBackgroundTile, false) // inner
}

// background of a number, since they have colors
func (b *Board) DrawNumberBackground(screen *ebiten.Image, startX, startY float32, y, x int, val [4]uint8, movDistX, movDistY float32) {
	var (
		xpos      float32 = (startX + float32(x)*TILESIZE + BORDERSIZE + movDistX*TILESIZE) * float32(b.game.scale)
		ypos      float32 = (startY + float32(y)*TILESIZE + BORDERSIZE + movDistY*TILESIZE) * float32(b.game.scale)
		size_tile float32 = (float32(TILESIZE) - BORDERSIZE) * float32(b.game.scale)
	)
	vector.DrawFilledRect(screen, xpos, ypos,
		size_tile, size_tile, getColor(val), false) // tiles
}

func (b *Board) DrawText(screen *ebiten.Image, xpos, ypos float32, x, y int, value int) {
	// draw the number to the screen
	msg := fmt.Sprintf("%v", value)
	fontUsed := mplusNormalFont

	var (
		dx float32 = float32(text.BoundString(mplusBigFont, msg).Dx())
		dy float32 = float32(text.BoundString(mplusBigFont, msg).Dy())
	)

	// check for text with first font is too large for it and swap
	if text.BoundString(mplusBigFont, msg).Dx() > int(TILESIZE*float32(b.game.scale)) {
		fontUsed = mplusNormalFontSmaller
		dx = (float32(text.BoundString(mplusNormalFontSmaller, msg).Dx() + int(BORDERSIZE)))
		dy = float32(text.BoundString(mplusNormalFontSmaller, msg).Dy())
	}

	var (
		textPosX int = int(xpos + (BORDERSIZE/2+TILESIZE/2)*float32(b.game.scale) - dx/2)
		textPosY int = int(ypos + (BORDERSIZE/2+TILESIZE/2)*float32(b.game.scale) + dy/2)
	)

	text.Draw(screen, msg, fontUsed,
		textPosX,
		textPosY,
		colorText)
}

// the functions for adding a random piece if the board is
func (b *Board) addNewRandomPieceIfBoardChanged() {
	if b.boardBeforeChange != b.board { // there will only be a new piece if it is a change
		b.randomNewPiece()
	}
}

func (b *Board) createBoardImage() {
	var (
		scale float64 = b.game.scale
		sizeX int     = int(float64((BOARDSIZE*int(TILESIZE))+(int(BORDERSIZE)*2)) * scale)
		sizeY         = sizeX
	)
	b.boardImage = ebiten.NewImage(sizeX, sizeY)
	for y := 0; y < BOARDSIZE; y++ {
		for x := 0; x < BOARDSIZE; x++ {
			b.DrawBorderBackground(b.boardImage, float32(x)*TILESIZE, float32(y)*TILESIZE)
		}

	}
	b.boardImageOptions = &ebiten.DrawImageOptions{}
	b.boardImageOptions.GeoM.Translate(float64(startPosX)*scale, float64(startPosY)*scale)
}

func (i *Input) SwitchDefaultDarkMode() {
	i.game.darkMode = !i.game.darkMode

	if i.game.darkMode { // DARK MODE
		i.game.board.colorBorder = colorBorderDarkMode
		i.game.board.colorBackgroundTile = colorBackgroundTileDarkMode
		i.game.board.createBoardImage()
	} else { // DEFAULT MODE
		i.game.board.colorBorder = colorBorderDefault
		i.game.board.colorBackgroundTile = colorBackgroundTileDefault
		i.game.board.createBoardImage()

	}
}
