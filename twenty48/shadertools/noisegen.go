package shadertools

import (
	"image/color"

	"github.com/aquilax/go-perlin"
	"github.com/hajimehoshi/ebiten/v2"
)

func generateNoiseImage(width, height int) *ebiten.Image {
	img := ebiten.NewImage(width, height)
	var (
		alpha       = 2.0 // Frequency
		beta        = 2.0 // Persistence
		n     int32 = 3   // Octaves
		scale       = 0.1 // Adjust this for zoom
	)

	perlin := perlin.NewPerlin(alpha, beta, n, 42)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Generate Perlin noise value
			noise := perlin.Noise2D(float64(x)*scale, float64(y)*scale)

			// Normalize noise to a 0-255 range
			normalized := uint8((noise + 1) * 127.5)

			// Set pixel color (grayscale)
			col := color.RGBA{normalized, normalized, normalized, 255}
			img.Set(x, y, col)
		}
	}

	return img
}

func UpdateNewNoiseImage(newWidth, newHeight int) {
	noiseImage = generateNoiseImage(newWidth, newHeight)
}

func UpdateScaleNoiseImage(newWidth, newHeight int) {
	img := ebiten.NewImage(newWidth, newHeight)
	oldX, oldY := noiseImage.Bounds().Dx(), noiseImage.Bounds().Dy()
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(float64(newWidth)/float64(oldX), float64(newHeight)/float64(oldY))
	img.DrawImage(noiseImage, opts)
	noiseImage = img
}
