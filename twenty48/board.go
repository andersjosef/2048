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

var color_text = color.RGBA{110, 93, 71, 255}

// colors for different numbers
var color_map = map[int][4]uint8{
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
	8192:  {87, 107, 75, 255},
	16384: {247, 104, 104, 255},
	-1:    {255, 255, 255, 255},
}

type Board struct {
	board                 [BOARDSIZE][BOARDSIZE]int // 2d array for the board :)
	color_border          color.RGBA
	color_backgorund_tile color.RGBA
}

func NewBoard() (*Board, error) {

	b := &Board{}

	// border and background colors
	b.color_border = color.RGBA{194, 182, 169, 255}
	b.color_backgorund_tile = color.RGBA{204, 192, 179, 255}

	for i := 0; i < 2; i++ {
		b.randomNewPiece()
	}

	return b, nil
}

func (b *Board) randomNewPiece() {
	var x, y int = len(b.board), len(b.board[0])

	var posFound bool = false
	var count int

	for !posFound && (count < x*y) {
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
	var start_pos_x, start_pos_y float32 = float32((SCREENWIDTH / 2) - (BOARDSIZE*int(TILESIZE))/2), float32((SCREENHEIGHT / 2) - (BOARDSIZE*int(TILESIZE))/2)

	for y := 0; y < len(b.board); y++ {
		for x := 0; x < len(b.board[0]); x++ {
			var xpos, ypos float32 = start_pos_x + float32(x)*TILESIZE, start_pos_y + float32(y)*TILESIZE
			// border
			vector.DrawFilledRect(screen, xpos, ypos,
				float32(TILESIZE)+BORDERSIZE*2, float32(TILESIZE)+BORDERSIZE*2, b.color_border, false) //border
			// inner
			vector.DrawFilledRect(screen, xpos+BORDERSIZE, ypos+BORDERSIZE, // + bordersize so you can see the border size on the left
				float32(TILESIZE), float32(TILESIZE), b.color_backgorund_tile, false) // tiles
			if b.board[y][x] != 0 {
				val, ok := color_map[b.board[y][x]] // checks if num in map, if it is make the background else draw normal
				// If the key exists
				if ok {
					vector.DrawFilledRect(screen, start_pos_x+float32(x)*TILESIZE+BORDERSIZE, start_pos_y+float32(y)*TILESIZE+BORDERSIZE,
						float32(TILESIZE), float32(TILESIZE), getColor(val), false) // tiles
				}
				// draw the number to the screen
				msg := fmt.Sprintf("%v", b.board[y][x])

				fontUsed := mplusNormalFont
				var (
					dx float32 = float32(text.BoundString(mplusBigFont, msg).Dx())
					dy float32 = float32(text.BoundString(mplusBigFont, msg).Dy())
				)
				if text.BoundString(mplusBigFont, msg).Dx() > int(TILESIZE)+int(BORDERSIZE) {
					fontUsed = mplusNormalFontSmaller
					dx = float32(text.BoundString(mplusNormalFontSmaller, msg).Dx())
					dy = float32(text.BoundString(mplusNormalFontSmaller, msg).Dy())
				}

				var (
					xpos int = int(xpos + BORDERSIZE/2 + TILESIZE/2 - dx/2)
					ypos int = int(ypos + BORDERSIZE/2 + TILESIZE/2 + dy/2)
				)
				text.Draw(screen, msg, fontUsed,
					xpos,
					ypos,
					color_text)
			}
		}
	}

}

func (b *Board) addNewRandomPieceIfBoardChanged(board_before_change [BOARDSIZE][BOARDSIZE]int) {
	if board_before_change != b.board { // there will only be a new piece if it is a change
		b.randomNewPiece()
	}
}
