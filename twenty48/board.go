package twenty48

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const TILESIZE float32 = 50
const BORDERSIZE float32 = 2

var color_map = map[int][4]uint8{
	2:  {238, 228, 218, 255},
	4:  {237, 224, 200, 255},
	8:  {242, 177, 121, 255},
	16: {245, 149, 99, 255},
	32: {255, 104, 69, 255},
	64: {246, 94, 59, 255},
	-1: {255, 255, 255, 255},
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
	fmt.Println(b.board)

	return b, nil
}

func (b *Board) randomNewPiece() {
	var x, y int = len(b.board), len(b.board[0])

	var posFound bool = false
	var count int

	for !posFound && (count < x*y) {
		var pos_x, pos_y int = rand.Intn(x), rand.Intn(y)
		if b.board[pos_x][pos_y] == 0 {
			b.board[pos_x][pos_y] = 2 // atm hardcoded to always give 2
			break
		}
		count++
	}
}

func (b *Board) drawBoard(screen *ebiten.Image) {
	var start_pos_x, start_pos_y float32 = 0, 0
	var ()

	for y := 0; y < len(b.board); y++ {
		for x := 0; x < len(b.board[0]); x++ {
			// border
			vector.DrawFilledRect(screen, start_pos_x+float32(x)*TILESIZE, start_pos_y+float32(y)*TILESIZE,
				float32(TILESIZE)+BORDERSIZE*2, float32(TILESIZE)+BORDERSIZE*2, b.color_border, false) //border
			// inner
			vector.DrawFilledRect(screen, start_pos_x+float32(x)*TILESIZE+BORDERSIZE, start_pos_y+float32(y)*TILESIZE+BORDERSIZE,
				float32(TILESIZE), float32(TILESIZE), b.color_backgorund_tile, false) // tiles
			if b.board[y][x] != 0 {
				val, ok := color_map[b.board[y][x]] // checks if num in map, if it is make the background else draw normal
				// If the key exists
				if ok {
					vector.DrawFilledRect(screen, start_pos_x+float32(x)*TILESIZE+BORDERSIZE, start_pos_y+float32(y)*TILESIZE+BORDERSIZE,
						float32(TILESIZE), float32(TILESIZE), getColor(val), false) // tiles
				} else {

				}
				text.Draw(screen, fmt.Sprintf("%v", b.board[y][x]), mplusNormalFont, int(start_pos_x+float32(x)*TILESIZE+BORDERSIZE+10), int(start_pos_y+float32(y)*TILESIZE+BORDERSIZE)+int(TILESIZE-10),
					color.Black) // letters
			}
		}
	}

}
