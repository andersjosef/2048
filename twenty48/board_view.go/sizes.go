package board_view

import (
	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/eventhandler"
	"github.com/andersjosef/2048/twenty48/shadertools"
	"github.com/hajimehoshi/ebiten/v2"
)

type SizesDeps struct {
	EventHandler interface {
		Register(eventType eventhandler.EventType, handler func(eventhandler.Event))
		Dispatch()
		Emit(event eventhandler.Event)
	}
	ScreenControl
}

// The sizes for the board that can be scaled up and down with window size changes
type Sizes struct {
	d SizesDeps

	tileSize   float32
	bordersize float32
	startPosX  float32
	startPosY  float32

	baseTileSize   float32
	baseBorderSize float32
}

func InitSizes(b *BoardView, d SizesDeps) *Sizes {
	const (
		BASE_TILESIZE   float32 = float32(co.LOGICAL_WIDTH) / 6.4
		BASE_BORDERSIZE float32 = BASE_TILESIZE / 25
		START_POS_X     float32 = float32((co.LOGICAL_WIDTH - (co.BOARDSIZE * int(BASE_TILESIZE))) / 2)
		START_POS_Y     float32 = float32((co.LOGICAL_HEIGHT - (co.BOARDSIZE * int(BASE_TILESIZE))) / 2)
	)

	dpiScale := ebiten.Monitor().DeviceScaleFactor()
	sfb := &Sizes{
		d:              d,
		baseTileSize:   BASE_TILESIZE,
		baseBorderSize: BASE_BORDERSIZE,
		tileSize:       BASE_TILESIZE * float32(dpiScale),
		bordersize:     BASE_BORDERSIZE * float32(dpiScale),
		startPosX:      START_POS_X * float32(dpiScale),
		startPosY:      START_POS_Y * float32(dpiScale),
	}

	sfb.d.EventHandler.Register(
		eventhandler.EventScreenChanged,
		func(evt eventhandler.Event) {
			sfb.scaleBoard()
			// Scale boardview after scaling values
			sfb.d.EventHandler.Emit(eventhandler.Event{
				Type: eventhandler.EventScaleBoardView,
			})
			val := int(sfb.baseTileSize)
			shadertools.UpdateScaleNoiseImage(val, val)
		},
	)

	return sfb
}

func (s *Sizes) GetStartPos() (x, y float32) {
	return s.startPosX, s.startPosY
}

func (s *Sizes) scaleBoard() {
	scale := s.d.ScreenControl.GetScale()
	dpiScale := ebiten.Monitor().DeviceScaleFactor()

	s.tileSize = s.baseTileSize * float32(scale) * float32(dpiScale)
	s.bordersize = s.baseBorderSize * float32(scale) * float32(dpiScale)
	width, height := s.d.ScreenControl.GetActualSize()
	s.startPosX = float32((width - (co.BOARDSIZE * int(s.tileSize))) / 2)
	s.startPosY = float32((height - (co.BOARDSIZE * int(s.tileSize))) / 2)

}

func (b *BoardView) scaleBoard() {
	newOpt := &ebiten.DrawImageOptions{}
	x, y := b.sizes.GetStartPos()
	newOpt.GeoM.Translate(float64(x), float64(y))
	b.boardOpts = newOpt
	b.CreateBoardImage()
}

func (s *Sizes) TileSize() float32 {
	return s.tileSize
}

func (s *Sizes) StartPos() (x, y float32) {
	return s.startPosX, s.startPosY
}
