package twenty48

import (
	"fmt"
	"math/rand"

	"github.com/andersjosef/2048/twenty48/shadertools"
	"github.com/andersjosef/2048/twenty48/theme"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const TILESIZE float32 = float32(SCREENWIDTH) / 6.4
const BORDERSIZE float32 = TILESIZE / 25

var (
	startPosX float32 = float32((SCREENWIDTH - (BOARDSIZE * int(TILESIZE))) / 2)
	startPosY float32 = float32((SCREENHEIGHT - (BOARDSIZE * int(TILESIZE))) / 2)
)

type Board struct {
	board             [BOARDSIZE][BOARDSIZE]int // 2d array for the board :)
	game              *Game
	boardBeforeChange [BOARDSIZE][BOARDSIZE]int
	boardImage        *ebiten.Image
	boardImageOptions *ebiten.DrawImageOptions

	boardForEndScreen *ebiten.Image
}

func NewBoard(g *Game) (*Board, error) {

	b := &Board{}

	b.game = g

	b.initBoardForEndScreen()

	// add the two start pieces
	for i := 0; i < 2; i++ {
		b.randomNewPiece()
	}

	// create baordImage
	b.createBoardImage()

	return b, nil

}

func (b *Board) initBoardForEndScreen() {
	realW, realH := b.game.screenControl.GetRealWidthHeight()
	scale := int(b.game.scale)
	b.boardForEndScreen = ebiten.NewImage(realW*scale, realH*scale)
}

func (b *Board) randomNewPiece() {

	var x, y int = len(b.board), len(b.board[0])

	// Will start at a random position, then check every available spot after
	// until all tiles are checked
	var count int = rand.Intn(x * y)
	for ; count < count+x*y-1; count++ {
		var posX int = count % x
		var posY int = (count / y) % y
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
	b.boardForEndScreen.DrawImage(b.boardImage, b.boardImageOptions)

	// draw tiles
	for y := 0; y < len(b.board); y++ {
		for x := 0; x < len(b.board[0]); x++ {
			b.DrawTile(b.boardForEndScreen, startPosX, startPosY, x, y, b.board[y][x], 0, 0)
		}
	}
	if !b.game.gameOver {
		screen.DrawImage(b.boardForEndScreen, &ebiten.DrawImageOptions{})

	} else {
		newImage, isDone := shadertools.GetImageFadeOut(b.boardForEndScreen)
		if isDone {
			// After animation go to game over state
			b.game.state = StateGameOver
		}
		screen.DrawImage(newImage, &ebiten.DrawImageOptions{})
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
		colorMap := b.game.currentTheme.ColorMap

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
		sizeBorder, sizeBorder, b.game.currentTheme.ColorBorder, false) //outer
	vector.DrawFilledRect(screen, xpos+BORDERSIZE*float32(b.game.scale), ypos+BORDERSIZE*float32(b.game.scale),
		sizeInside, sizeInside, b.game.currentTheme.ColorBackgroundTile, false) // inner
}

// background of a number, since they have colors
func (b *Board) DrawNumberBackground(screen *ebiten.Image, startX, startY float32, y, x int, val [4]uint8, movDistX, movDistY float32) {
	var (
		xpos      float32 = (startX + float32(x)*TILESIZE + BORDERSIZE + movDistX*TILESIZE) * float32(b.game.scale)
		ypos      float32 = (startY + float32(y)*TILESIZE + BORDERSIZE + movDistY*TILESIZE) * float32(b.game.scale)
		size_tile float32 = (float32(TILESIZE) - BORDERSIZE) * float32(b.game.scale)
	)
	vector.DrawFilledRect(screen, xpos, ypos,
		size_tile, size_tile, theme.GetColor(val), false) // tiles
}

func (b *Board) DrawText(screen *ebiten.Image, xpos, ypos float32, x, y int, value int) {
	// draw the number to the screen
	fontSet := b.game.fontSet
	msg := fmt.Sprintf("%v", value)
	// fontUsed := mplusNormalFont
	fontUsed := fontSet.Normal
	textHeight := -(fontSet.Normal.Metrics().VAscent + fontSet.Normal.Metrics().VDescent)

	var (
		dx float32 = float32(text.Advance(msg, fontSet.Big))
		dy float32 = float32(textHeight)
	)

	// check for text with first font is too large for it and swap
	if int(text.Advance(msg, fontSet.Big)) > int(TILESIZE*float32(b.game.scale)) {
		// fontUsed = mplusNormalFontSmaller
		fontUsed = fontSet.Smaller
		textHeight = -(fontSet.Smaller.Metrics().VAscent + fontSet.Smaller.Metrics().VDescent)
		dx = (float32(int(text.Advance(msg, fontSet.Smaller)) + int(BORDERSIZE)))
		dy = float32(textHeight)
	}

	var (
		textPosX int = int(xpos + (BORDERSIZE/2+TILESIZE/2)*float32(b.game.scale) - dx/2)
		textPosY int = int(ypos + (BORDERSIZE/2+TILESIZE/2)*float32(b.game.scale) + dy/2)
	)

	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(textPosX), float64(textPosY))
	op.ColorScale.ScaleWithColor(b.game.currentTheme.ColorText)
	text.Draw(screen, msg, fontUsed, op)
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

// Check if its gameOver
func (b *Board) isGameOver() bool {
	// Check if there are any empty spaces left, meaning its possible to play
	for i := 0; i < BOARDSIZE; i++ {
		for j := 0; j < BOARDSIZE; j++ {
			if b.board[i][j] == 0 {
				return false
			}
		}
	}

	// Check for vertical merges
	for i := 0; i < BOARDSIZE-1; i++ {
		for j := 0; j < BOARDSIZE; j++ {
			if b.board[i][j] == b.board[i+1][j] {
				return false
			}
		}
	}

	// Check for horisontal merges
	for i := 0; i < BOARDSIZE; i++ {
		for j := 0; j < BOARDSIZE-1; j++ {
			if b.board[i][j] == b.board[i][j+1] {
				return false
			}
		}
	}
	return true

}
