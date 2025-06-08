package twenty48

import (
	"fmt"
	"math/rand"

	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/shadertools"
	"github.com/andersjosef/2048/twenty48/theme"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// The sizes for the board that can be scaled up and down with window size changes
type Sizes struct {
	board      *Board
	tileSize   float32
	bordersize float32
	startPosX  float32
	startPosY  float32

	baseTileSize   float32
	baseBorderSize float32
}

func InitSizes(b *Board) *Sizes {
	const (
		BASE_TILESIZE   float32 = float32(co.LOGICAL_WIDTH) / 6.4
		BASE_BORDERSIZE float32 = BASE_TILESIZE / 25
		START_POS_X     float32 = float32((co.LOGICAL_WIDTH - (co.BOARDSIZE * int(BASE_TILESIZE))) / 2)
		START_POS_Y     float32 = float32((co.LOGICAL_HEIGHT - (co.BOARDSIZE * int(BASE_TILESIZE))) / 2)
	)

	dpiScale := ebiten.Monitor().DeviceScaleFactor()
	sfb := &Sizes{
		baseTileSize:   BASE_TILESIZE,
		baseBorderSize: BASE_BORDERSIZE,
		board:          b,
		tileSize:       BASE_TILESIZE * float32(dpiScale),
		bordersize:     BASE_BORDERSIZE * float32(dpiScale),
		startPosX:      START_POS_X * float32(dpiScale),
		startPosY:      START_POS_Y * float32(dpiScale),
	}
	return sfb
}

func (s *Sizes) scaleBoard() {
	scale := s.board.game.scale
	dpiScale := ebiten.Monitor().DeviceScaleFactor()

	s.tileSize = s.baseTileSize * float32(scale) * float32(dpiScale)
	s.bordersize = s.baseBorderSize * float32(scale) * float32(dpiScale)

	s.startPosX = float32((s.board.game.screenControl.actualWidth - (co.BOARDSIZE * int(s.tileSize))) / 2)
	s.startPosY = float32((s.board.game.screenControl.actualHeight - (co.BOARDSIZE * int(s.tileSize))) / 2)

	newOpt := &ebiten.DrawImageOptions{}
	newOpt.GeoM.Translate(float64(s.startPosX), float64(s.startPosY))
	s.board.boardImageOptions = newOpt

	s.board.createBoardImage()
}

type Board struct {
	board             [co.BOARDSIZE][co.BOARDSIZE]int // 2d array for the board :)
	game              *Game
	sizes             *Sizes
	boardBeforeChange [co.BOARDSIZE][co.BOARDSIZE]int
	boardImage        *ebiten.Image
	boardImageOptions *ebiten.DrawImageOptions

	boardForEndScreen *ebiten.Image
}

func NewBoard(g *Game) (*Board, error) {

	b := &Board{}
	b.sizes = InitSizes(b)

	b.game = g

	// add the two start pieces
	for i := 0; i < 2; i++ {
		b.randomNewPiece()
	}

	// create baordImage
	b.createBoardImage()

	return b, nil

}

func (b *Board) initBoardForEndScreen() {
	b.boardForEndScreen = ebiten.NewImage(b.game.screenControl.actualWidth, b.game.screenControl.actualHeight)
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
			b.DrawTile(b.boardForEndScreen, b.sizes.startPosX, b.sizes.startPosY, x, y, b.board[y][x], 0, 0)
		}
	}
	if !b.game.gameOver {
		screen.DrawImage(b.boardForEndScreen, &ebiten.DrawImageOptions{})

	} else {
		newImage, isDone := shadertools.GetImageFadeOut(b.boardForEndScreen)
		if isDone {
			// After animation go to game over state
			b.game.state = co.StateGameOver
		}
		screen.DrawImage(newImage, &ebiten.DrawImageOptions{})
	}
}

// draws one tile of the game with everything background, number, color, etc.
func (b *Board) DrawTile(screen *ebiten.Image, startX, startY float32, x, y int, value int, movDistX, movDistY float32) {
	var (
		xpos float32 = (startX + float32(x)*b.sizes.tileSize + movDistX*b.sizes.tileSize)
		ypos float32 = (startY + float32(y)*b.sizes.tileSize + movDistY*b.sizes.tileSize)
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
	var sizeBorder float32 = (float32(b.sizes.tileSize) + b.sizes.bordersize)
	var sizeInside float32 = (b.sizes.tileSize - b.sizes.bordersize)

	vector.DrawFilledRect(screen, xpos, ypos,
		sizeBorder, sizeBorder, b.game.currentTheme.ColorBorder, false) //outer
	vector.DrawFilledRect(screen, xpos+b.sizes.bordersize, ypos+b.sizes.bordersize,
		sizeInside, sizeInside, b.game.currentTheme.ColorBackgroundTile, false) // inner
}

// background of a number, since they have colors
func (b *Board) DrawNumberBackground(screen *ebiten.Image, startX, startY float32, y, x int, val [4]uint8, movDistX, movDistY float32) {
	var (
		xpos      float32 = (startX + float32(x)*b.sizes.tileSize + b.sizes.bordersize + movDistX*b.sizes.tileSize)
		ypos      float32 = (startY + float32(y)*b.sizes.tileSize + b.sizes.bordersize + movDistY*b.sizes.tileSize)
		size_tile float32 = (float32(b.sizes.tileSize) - b.sizes.bordersize)
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
	if int(text.Advance(msg, fontSet.Big)) > int(b.sizes.tileSize) {
		// fontUsed = mplusNormalFontSmaller
		fontUsed = fontSet.Smaller
		textHeight = -(fontSet.Smaller.Metrics().VAscent + fontSet.Smaller.Metrics().VDescent)
		dx = (float32(int(text.Advance(msg, fontSet.Smaller)) + int(b.sizes.bordersize)))
		dy = float32(textHeight)
	}

	var (
		textPosX int = int(xpos + (b.sizes.bordersize/2 + b.sizes.tileSize/2) - dx/2)
		textPosY int = int(ypos + (b.sizes.bordersize/2 + b.sizes.tileSize/2) + dy/2)
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
		sizeX int = int(float64((co.BOARDSIZE * int(b.sizes.tileSize)) + (int(b.sizes.bordersize) * 2)))
		sizeY     = sizeX
	)
	b.boardImage = ebiten.NewImage(sizeX, sizeY)
	for y := 0; y < co.BOARDSIZE; y++ {
		for x := 0; x < co.BOARDSIZE; x++ {
			b.DrawBorderBackground(b.boardImage, float32(x)*b.sizes.tileSize, float32(y)*b.sizes.tileSize)
		}

	}
	b.boardImageOptions = &ebiten.DrawImageOptions{}
	b.boardImageOptions.GeoM.Translate(float64(b.sizes.startPosX), float64(b.sizes.startPosY))

	// Will update the size of it for screensize changes
	b.initBoardForEndScreen()
}

// Check if its gameOver
func (b *Board) isGameOver() bool {
	// Check if there are any empty spaces left, meaning its possible to play
	for i := 0; i < co.BOARDSIZE; i++ {
		for j := 0; j < co.BOARDSIZE; j++ {
			if b.board[i][j] == 0 {
				return false
			}
		}
	}

	// Check for vertical merges
	for i := 0; i < co.BOARDSIZE-1; i++ {
		for j := 0; j < co.BOARDSIZE; j++ {
			if b.board[i][j] == b.board[i+1][j] {
				return false
			}
		}
	}

	// Check for horisontal merges
	for i := 0; i < co.BOARDSIZE; i++ {
		for j := 0; j < co.BOARDSIZE-1; j++ {
			if b.board[i][j] == b.board[i][j+1] {
				return false
			}
		}
	}
	return true

}
