package twenty48

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type MyInput struct {
	keys []ebiten.Key
}

var m = &MyInput{}

func (m *MyInput) UpdateInput() error {
	m.keys = inpututil.AppendPressedKeys(m.keys[:0])
	return nil
}

func (m *MyInput) DrawInput(screen *ebiten.Image) {
	var keyStrs []string
	var keyNames []string
	for _, k := range m.keys {
		keyStrs = append(keyStrs, k.String())
		if name := ebiten.KeyName(k); name != "" {
			keyNames = append(keyNames, name)
		}
		fmt.Println(m.keys)
	}

	// Use bitmapfont.Face instead of ebitenutil.DebugPrint, since some key names might not be printed with DebugPrint.
	text.Draw(screen, strings.Join(keyStrs, ", ")+"\n"+strings.Join(keyNames, ", "), mplusBigFont, 4, 12, color.Black)
}
