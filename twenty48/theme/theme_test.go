package theme

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitFontsSuccess(t *testing.T) {
	scale := 1.0
	fontSet := InitFonts(scale)
	assert.NotNil(t, fontSet, "font set should not be nil")

	// Individual font faces
	assert.NotNil(t, fontSet.Normal, "Normal font should not be nil")
	assert.NotNil(t, fontSet.Smaller, "Smaller font should not be nil")
	assert.NotNil(t, fontSet.Mini, "Mini font should not be nil")
	assert.NotNil(t, fontSet.Big, "Big font should not be nil")

}

func TestInitFontsScaling(t *testing.T) {
	scale := 2.0 // Simulate high DPI

	fontSet := InitFonts(scale)
	assert.NotNil(t, fontSet)

}
