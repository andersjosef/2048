package shadertools

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func Update() {
	// Update the times maps
	updateFadeOut()
	updateFadeIn()
}

// Returns a new Image with new size
func resizeImage(source *ebiten.Image, newWidth, newHeight int) *ebiten.Image {
	// Create a new empty image with the target size
	newImage := ebiten.NewImage(newWidth, newHeight)

	// Calculate scaling factors
	scaleX := float64(newWidth) / float64(source.Bounds().Dx())
	scaleY := float64(newHeight) / float64(source.Bounds().Dy())

	// Create options to draw the scaled image
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scaleX, scaleY)

	// Draw the original image onto the new image using the scaling options
	newImage.DrawImage(source, op)

	return newImage
}
