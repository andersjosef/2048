package twenty48

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const TILESIZE float32 = 100
const BORDERSIZE float32 = TILESIZE / 25

var (
	color_text          = color.RGBA{110, 93, 71, 255}
	start_pos_x float32 = float32((SCREENWIDTH - (BOARDSIZE * int(TILESIZE))) / 2)
	start_pos_y float32 = float32((SCREENHEIGHT - (BOARDSIZE * int(TILESIZE))) / 2)

	colorBorderDefault         = color.RGBA{194, 182, 169, 255}
	colorBackgroundTileDefault = color.RGBA{204, 192, 179, 255}
)

// colors for different numbers DEFAULT/LIGHT MODE
var colorMapDefault = map[int][4]uint8{
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
}

// colors for different numbers DARK MODE
var colorMapDarkMode = map[int][4]uint8{
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
}

type Board struct {
	board                 [BOARDSIZE][BOARDSIZE]int // 2d array for the board :)
	color_border          color.RGBA
	color_background_tile color.RGBA
	game                  *Game
	board_before_change   [BOARDSIZE][BOARDSIZE]int
	board_image           *ebiten.Image
	board_image_options   *ebiten.DrawImageOptions
}

func NewBoard(g *Game) (*Board, error) {

	b := &Board{}

	// border and background colors
	b.color_border = colorBorderDefault
	b.color_background_tile = colorBackgroundTileDefault
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
	var count int

	for count < x*y {
		var pos_x, pos_y int = rand.Intn(x), rand.Intn(y)
		if b.board[pos_x][pos_y] == 0 {
			if rand.Float32() > 0.16 {
				b.board[pos_x][pos_y] = 2 // 84%
			} else {
				b.board[pos_x][pos_y] = 4 // 16% chance of 4 spawning
			}
			break
		}
		count++
	}
}

func (b *Board) drawBoard(screen *ebiten.Image) {
	// draw the backgroundimage of the game
	screen.DrawImage(b.board_image, b.board_image_options)

	// draw tiles
	for y := 0; y < len(b.board); y++ {
		for x := 0; x < len(b.board[0]); x++ {
			b.DrawTile(screen, start_pos_x, start_pos_y, x, y, b.board[y][x], 0, 0)
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
		sizeBorder, sizeBorder, b.color_border, false) //outer
	vector.DrawFilledRect(screen, xpos+BORDERSIZE*float32(b.game.scale), ypos+BORDERSIZE*float32(b.game.scale),
		sizeInside, sizeInside, b.color_background_tile, false) // inner
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
		color_text)
}

// the functions for adding a random piece if the board is
func (b *Board) addNewRandomPieceIfBoardChanged() {
	if b.board_before_change != b.board { // there will only be a new piece if it is a change
		b.randomNewPiece()
	}
}

func (b *Board) createBoardImage() {
	var (
		scale  float64 = b.game.scale
		size_x int     = int(float64((BOARDSIZE*int(TILESIZE))+(int(BORDERSIZE)*2)) * scale)
		size_y         = size_x
	)
	b.board_image = ebiten.NewImage(size_x, size_y)
	for y := 0; y < BOARDSIZE; y++ {
		for x := 0; x < BOARDSIZE; x++ {
			b.DrawBorderBackground(b.board_image, float32(x)*TILESIZE, float32(y)*TILESIZE)
		}

	}
	b.board_image_options = &ebiten.DrawImageOptions{}
	b.board_image_options.GeoM.Translate(float64(start_pos_x)*scale, float64(start_pos_y)*scale)
}

func (b *Board) SwitchDefaultDarkMode() {
	b.game.darkMode = !b.game.darkMode

	if b.game.darkMode {
		fmt.Println("In darkmode!")
		fmt.Println(b.game.darkMode)
	} else {
		fmt.Println("In defautlmode!")
		fmt.Println(b.game.darkMode)

	}
}
