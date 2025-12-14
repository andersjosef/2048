package animations

import (
	"fmt"
	"math"
	"time"

	co "github.com/andersjosef/2048/twenty48/constants"
	"github.com/andersjosef/2048/twenty48/eventhandler"
	"github.com/andersjosef/2048/twenty48/shared"
	"github.com/hajimehoshi/ebiten/v2"
)

type Animation struct {
	d               Deps
	isAnimating     bool
	deltas          []shared.MoveDelta
	currentDir      string
	animationLength float32 // Seconds
	startTime       time.Time
}

func New(d Deps) *Animation {
	a := &Animation{
		isAnimating:     false,
		d:               d,
		animationLength: 0.20, // Animation duration in seconds
	}

	a.d.Register(
		eventhandler.EventMoveMade,
		func(e eventhandler.Event) {
			moveData, ok := e.Data.(shared.MoveData)
			if !ok {
				return
			}

			a.play(moveData.MoveDeltas, moveData.Dir)
		},
	)
	return a
}

func (a *Animation) play(deltas []shared.MoveDelta, dir string) {
	a.deltas = deltas
	a.currentDir = dir
	a.isAnimating = true
	a.startTime = time.Now()
}

func (a *Animation) Draw(screen *ebiten.Image) {
	a.d.DrawBackgoundBoard(screen)

	elapsed := float32(time.Since(a.startTime).Seconds())
	progress := float32(min(float64(elapsed/a.animationLength), 1))

	// How far any tile would go if it had to move the full width/height
	fullDist := float32(co.BOARDSIZE-1) * progress

	var ImgOpts ebiten.DrawImageOptions
	startX, startY := a.d.Layout.GetStartPos()
	tileSize := a.d.Layout.TileSize()
	for _, d := range a.deltas {
		// Get tile image form boardview
		img, ok := a.d.BoardView.GetTile(d.ValueMoved)
		if !ok {
			continue
		}

		// Signed cell‐deltas
		dxCells := d.ToCol - d.FromCol
		dyCells := d.ToRow - d.FromRow

		dirX := float32(sign(dxCells))
		dirY := float32(sign(dyCells))

		// How many cells this tile needs in each axis
		needX := float32(math.Abs(float64(dxCells)))
		needY := float32(math.Abs(float64(dyCells)))

		// Cap fullDist at each axis need
		moveX := dirX * min(fullDist, needX)
		moveY := dirY * min(fullDist, needY)

		// Draw tiles
		x := float64(startX + float32(d.FromCol)*tileSize + a.d.BorderSize() + moveX*tileSize)
		y := float64(startY + float32(d.FromRow)*tileSize + a.d.BorderSize() + moveY*tileSize)
		ImgOpts.GeoM.Reset()
		ImgOpts.GeoM.Translate(x, y)
		screen.DrawImage(img, &ImgOpts)
	}

	if progress >= 1 {
		a.isAnimating = false
		fmt.Println("Animation done")
		if a.d.IsGameOver() {
			a.d.EventHandler.Emit(eventhandler.Event{
				Type: eventhandler.EventAnimationDoneGameOver,
			})
		}
	}
}

// Helper to get sign of an int
func sign(n int) int {
	switch {
	case n < 0:
		return -1
	case n > 0:
		return +1
	default:
		return 0
	}
}

func (a *Animation) IsAnimating() bool {
	return a.isAnimating
}
