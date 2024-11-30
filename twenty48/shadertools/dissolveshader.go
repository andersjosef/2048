// Copyright 2020 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Modifications:
// - Transformed the shader into a dissolving effect module.
// - Added support for fade-in and fade-out animations.
// - Introduced noise caching and resize optimization.

package shadertools

import (
	_ "embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed dissolve_shader_logic.go
	dissolve_shader_logic []byte
)

var (
	dissolve     *ebiten.Shader
	timesFadeIn  map[int]float32
	timesFadeOut map[int]float32
	noiseCache   map[[2]int]*ebiten.Image
	noiseImage   *ebiten.Image
	idCounter    int
	imageToId    map[*ebiten.Image]int
	emptyImage   = ebiten.NewImage(1, 1)
)

const FadeDuration = 60

func init() {
	var err error
	dissolve, err = ebiten.NewShader(dissolve_shader_logic)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize times maps
	ResetTimesMapsDissolve()

	noiseImage = generateNoiseImage()

	// Counter ID
	imageToId = make(map[*ebiten.Image]int)
}

func getImageId(image *ebiten.Image) int {
	if id, exists := imageToId[image]; exists {
		return id
	}

	id := idCounter
	imageToId[image] = id
	idCounter++
	return id

}

// Resets the times maps
func ResetTimesMapsDissolve() {
	timesFadeIn = make(map[int]float32)
	timesFadeOut = make(map[int]float32)
	noiseCache = make(map[[2]int]*ebiten.Image)
	idCounter = 0
	imageToId = make(map[*ebiten.Image]int)
}

func getResizedNoiseImage(w, h int) *ebiten.Image {
	key := [2]int{w, h}
	if noise, exists := noiseCache[key]; exists {
		return noise
	}
	resized := resizeImage(noiseImage, w, h)
	noiseCache[key] = resized
	return resized

}

// Function to fade out (dissolve) the image
func GetImageFadeOut(image *ebiten.Image) (newImage *ebiten.Image, isDone bool) {
	id := getImageId(image)

	// Retrieve the time value for this image, or initialize it
	if _, exists := timesFadeOut[id]; !exists {
		timesFadeOut[id] = 0 // Initialize time for this image
	}

	time := timesFadeOut[id]

	// If its already faded out dont bother with rendering
	if time >= FadeDuration {
		return emptyImage, true
	}

	return applyDissolveShader(image, time), false
}

// Function to fade in (reverse dissolve) the image
func GetImageFadeIn(image *ebiten.Image) (newImage *ebiten.Image, isDone bool) {
	id := getImageId(image)

	// Retrieve the time value for this image, or initialize it
	if _, exists := timesFadeIn[id]; !exists {
		timesFadeIn[id] = FadeDuration // Initialize time for this image to fully faded out state
	}

	// If image is fully shown, dont bother rendering
	if timesFadeIn[id] <= 0 {
		return image, true
	}
	time := timesFadeIn[id]

	return applyDissolveShader(image, time), false
}

// Send work to GPU
func applyDissolveShader(image *ebiten.Image, time float32) *ebiten.Image {
	w, h := image.Bounds().Dx(), image.Bounds().Dy()
	// Check cache for noise image in right size, if not create new
	noise := getResizedNoiseImage(w, h)
	newImage := ebiten.NewImage(w, h)
	op := &ebiten.DrawRectShaderOptions{
		Uniforms: map[string]any{
			"Time": time / FadeDuration,
		},
	}
	op.Images[0] = image
	op.Images[1] = noise
	newImage.DrawRectShader(w, h, dissolve, op)
	return newImage
}

//// UPDATES ////

// Fade in update
func updateFadeIn() {
	for key := range timesFadeIn {
		if timesFadeIn[key] > 0 {
			timesFadeIn[key]--
		} else {
			timesFadeIn[key] = 0
		}
	}

}

// Update fade out times map
func updateFadeOut() {
	// Fade out update
	for key := range timesFadeOut {
		if timesFadeOut[key] < FadeDuration {
			timesFadeOut[key]++
		} else {
			timesFadeOut[key] = FadeDuration
		}
	}
}
