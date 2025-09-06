package layout

import (
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/hajimehoshi/ebiten/v2"
)

type SizesDeps struct {
	ScreenControl interface {
		GetScale() float64
		GetActualSize() (x, y int)
	}
}

// The sizes for the board that can be scaled up and down with window size changes
type Layout struct {
	d SizesDeps

	tileSize   float32
	bordersize float32
	startPosX  float32
	startPosY  float32

	baseTileSize   float32
	baseBorderSize float32
}

func New(d SizesDeps) *Layout {
	const (
		BASE_TILESIZE   float32 = float32(co.LOGICAL_WIDTH) / 6.4
		BASE_BORDERSIZE float32 = BASE_TILESIZE / 25
		START_POS_X     float32 = float32((co.LOGICAL_WIDTH - (co.BOARDSIZE * int(BASE_TILESIZE))) / 2)
		START_POS_Y     float32 = float32((co.LOGICAL_HEIGHT - (co.BOARDSIZE * int(BASE_TILESIZE))) / 2)
	)

	dpiScale := ebiten.Monitor().DeviceScaleFactor()
	sfb := &Layout{
		d:              d,
		baseTileSize:   BASE_TILESIZE,
		baseBorderSize: BASE_BORDERSIZE,
		tileSize:       BASE_TILESIZE * float32(dpiScale),
		bordersize:     BASE_BORDERSIZE * float32(dpiScale),
		startPosX:      START_POS_X * float32(dpiScale),
		startPosY:      START_POS_Y * float32(dpiScale),
	}

	return sfb
}

func (s *Layout) Recalculate() {
	scale := s.d.ScreenControl.GetScale()
	dpiScale := ebiten.Monitor().DeviceScaleFactor()

	s.tileSize = s.baseTileSize * float32(scale) * float32(dpiScale)
	s.bordersize = s.baseBorderSize * float32(scale) * float32(dpiScale)
	width, height := s.d.ScreenControl.GetActualSize()
	s.startPosX = float32((width - (co.BOARDSIZE * int(s.tileSize))) / 2)
	s.startPosY = float32((height - (co.BOARDSIZE * int(s.tileSize))) / 2)

}

func (s *Layout) BorderSize() float32 {
	return s.bordersize
}

func (s *Layout) GetStartPos() (x, y float32) {
	return s.startPosX, s.startPosY
}

func (s *Layout) TileSize() float32 {
	return s.tileSize
}

func (s *Layout) StartPos() (x, y float32) {
	return s.startPosX, s.startPosY
}
