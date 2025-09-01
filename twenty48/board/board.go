package board

import (
	"fmt"
	"math/rand"

	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/eventhandler"
	"github.com/andersjosef/2048/twenty48/shadertools"
	"github.com/andersjosef/2048/twenty48/shared"
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

	sfb.board.d.Register(
		eventhandler.EventScreenChanged,
		func(evt eventhandler.Event) {
			sfb.scaleBoard()
			val := int(sfb.baseTileSize)
			shadertools.UpdateScaleNoiseImage(val, val)
		},
	)

	return sfb
}

func (s *Sizes) scaleBoard() {
	scale := s.board.d.GetScale()
	dpiScale := ebiten.Monitor().DeviceScaleFactor()

	s.tileSize = s.baseTileSize * float32(scale) * float32(dpiScale)
	s.bordersize = s.baseBorderSize * float32(scale) * float32(dpiScale)

	width, height := s.board.d.GetActualSize()

	s.startPosX = float32((width - (co.BOARDSIZE * int(s.tileSize))) / 2)
	s.startPosY = float32((height - (co.BOARDSIZE * int(s.tileSize))) / 2)

	newOpt := &ebiten.DrawImageOptions{}
	newOpt.GeoM.Translate(float64(s.startPosX), float64(s.startPosY))
	s.board.boardImageOptions = newOpt

	s.board.CreateBoardImage()
}

type Board struct {
	matrix             [co.BOARDSIZE][co.BOARDSIZE]int // 2d array for the board :)
	matrixBeforeChange [co.BOARDSIZE][co.BOARDSIZE]int
	d                  Deps
	sizes              *Sizes
	boardImage         *ebiten.Image
	boardImageOptions  *ebiten.DrawImageOptions

	boardForEndScreen *ebiten.Image
}

func New(d Deps) (*Board, error) {

	b := &Board{}
	b.d = d
	b.sizes = InitSizes(b)

	// add the two start pieces
	for range 2 {
		b.randomNewPiece()
	}

	// create boardImage
	b.CreateBoardImage()

	b.registerEvents()

	return b, nil

}

func (b *Board) registerEvents() {
	b.d.Register(
		eventhandler.EventResetGame,
		func(_ eventhandler.Event) {
			b.matrix = [co.BOARDSIZE][co.BOARDSIZE]int{}
			b.randomNewPiece()
			b.randomNewPiece()

		},
	)
	b.d.Register(
		eventhandler.EventMoveMade,
		func(e eventhandler.Event) {
			moveData, ok := e.Data.(shared.MoveData)
			if !ok {
				return
			}
			b.matrix = moveData.NewBoard
			b.addNewRandomPieceIfBoardChanged()
			b.d.SetGameOver(b.isGameOver())
		},
	)
}

func (b *Board) initBoardForEndScreen() {
	width, height := b.d.GetActualSize()
	b.boardForEndScreen = ebiten.NewImage(width, height)
}

func (b *Board) randomNewPiece() {

	var x, y int = len(b.matrix), len(b.matrix[0])

	// Will start at a random position, then check every available spot after
	// until all tiles are checked
	var count int = rand.Intn(x * y)
	for ; count < count+x*y-1; count++ {
		var posX int = count % x
		var posY int = (count / y) % y
		if b.matrix[posX][posY] == 0 {
			if rand.Float32() > 0.16 {
				b.matrix[posX][posY] = 2 // 84%
			} else {
				b.matrix[posX][posY] = 4 // 16% chance of 4 spawning
			}
			break
		}
	}
}

func (b *Board) Draw(screen *ebiten.Image) {
	// draw the backgroundimage of the game
	b.boardForEndScreen.DrawImage(b.boardImage, b.boardImageOptions)

	// draw tiles
	for y := range len(b.matrix) {
		for x := range len(b.matrix[0]) {
			b.DrawTile(b.boardForEndScreen, b.sizes.startPosX, b.sizes.startPosY, x, y, b.matrix[y][x], 0, 0)
		}
	}
	if !b.d.IsGameOver() {
		screen.DrawImage(b.boardForEndScreen, &ebiten.DrawImageOptions{})

	} else {
		newImage, isDone := shadertools.GetImageFadeOut(b.boardForEndScreen)
		if isDone {
			// After animation go to game over state
			b.d.SetGameState(co.StateGameOver)
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
		colorMap := b.d.GetCurrentTheme().ColorMap

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
		sizeBorder, sizeBorder, b.d.GetCurrentTheme().ColorBorder, false) //outer
	vector.DrawFilledRect(screen, xpos+b.sizes.bordersize, ypos+b.sizes.bordersize,
		sizeInside, sizeInside, b.d.GetCurrentTheme().ColorBackgroundTile, false) // inner
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
	fontSet := b.d.GetCurrentFontSet()
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
	op.ColorScale.ScaleWithColor(b.d.GetCurrentTheme().ColorText)
	text.Draw(screen, msg, fontUsed, op)
}

// the functions for adding a random piece if the board is
func (b *Board) addNewRandomPieceIfBoardChanged() {
	if b.matrixBeforeChange != b.matrix { // there will only be a new piece if it is a change
		b.randomNewPiece()
	}
}

func (b *Board) CreateBoardImage() {
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

// Check if its gameOver
func (b *Board) isGameOver() bool {
	// Check if there are any empty spaces left, meaning its possible to play
	for i := range co.BOARDSIZE {
		for j := range co.BOARDSIZE {
			if b.matrix[i][j] == 0 {
				return false
			}
		}
	}

	// Check for vertical merges
	for i := range co.BOARDSIZE - 1 {
		for j := range co.BOARDSIZE {
			if b.matrix[i][j] == b.matrix[i+1][j] {
				return false
			}
		}
	}

	// Check for horisontal merges
	for i := range co.BOARDSIZE {
		for j := range co.BOARDSIZE - 1 {
			if b.matrix[i][j] == b.matrix[i][j+1] {
				return false
			}
		}
	}
	return true

}
