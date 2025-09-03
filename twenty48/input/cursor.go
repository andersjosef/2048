package input

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type CursorVisibility struct {
	isHidden  bool
	last      [2]int
	threshold float64 // If mouse is moved beyond this show again
}

func NewCursorVisibility(th float64) *CursorVisibility {
	return &CursorVisibility{threshold: th}
}

func (c *CursorVisibility) MaybeShow() {
	if !c.isHidden {
		return
	}

	x, y := ebiten.CursorPosition()
	if math.Abs(float64(c.last[0]-x)) > c.threshold ||
		math.Abs(float64(c.last[1]-y)) > c.threshold {
		ebiten.SetCursorMode(ebiten.CursorModeVisible)
		c.isHidden = false
	}
}

func (c *CursorVisibility) Hide() {
	if c.isHidden {
		return
	}

	c.last[0], c.last[1] = ebiten.CursorPosition()
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	c.isHidden = true
}
