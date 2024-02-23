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

type Board struct {
	board                 [4][4]int // 2d array for the board :)
	color_border          color.RGBA
	color_backgorund_tile color.RGBA
}

func NewBoard() (*Board, error) {

	b := &Board{}

	b.color_border = color.RGBA{200, 200, 200, 255}
	b.color_backgorund_tile = color.RGBA{255, 255, 255, 255}

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
				text.Draw(screen, fmt.Sprintf("%v", b.board[y][x]), mplusNormalFont, int(start_pos_x+float32(x)*TILESIZE+BORDERSIZE+10), int(start_pos_y+float32(y)*TILESIZE+BORDERSIZE)+int(TILESIZE-10),
					color.Black) // letters
			}
		}
	}

}
