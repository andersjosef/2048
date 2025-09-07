package board_view

import "github.com/hajimehoshi/ebiten/v2"

func (b *BoardView) DrawBackgoundBoard(screen *ebiten.Image) {
	screen.DrawImage(b.emptyBoard, b.opts)
}

func (b *BoardView) GetTile(v int) (img *ebiten.Image, ok bool) {
	img, ok = b.tiles[v]
	return img, ok
}
