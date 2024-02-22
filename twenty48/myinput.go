package twenty48

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type MyInput struct {
	keys              []ebiten.Key
	keyIsBeingPressed bool
}

var m = &MyInput{
	keyIsBeingPressed: false,
}

func (m *MyInput) UpdateInput() error {
	m.keys = inpututil.AppendPressedKeys(m.keys[:0])
	return nil
}

// func (m *MyInput) DrawInput(screen *ebiten.Image) {
// 	var keyStrs []string
// 	var keyNames []string
// 	for _, k := range m.keys {
// 		keyStrs = append(keyStrs, k.String())
// 		if name := ebiten.KeyName(k); name != "" {
// 			keyNames = append(keyNames, name)
// 		}
// 		fmt.Println(m.keys)
// 	}

// 	// Use bitmapfont.Face instead of ebitenutil.DebugPrint, since some key names might not be printed with DebugPrint.
// 	text.Draw(screen, strings.Join(keyStrs, ", ")+"\n"+strings.Join(keyNames, ", "), mplusBigFont, 4, 12, color.Black)
// }

func (m *MyInput) DrawInput() {
	if len(m.keys) > 0 {
		if !m.keyIsBeingPressed {
			m.keyIsBeingPressed = true
			key_pressed := m.keys[len(m.keys)-1]
			// fmt.Println(key_pressed)
			switch fmt.Sprintf("%v", key_pressed) {
			case "D", "ArrowRight": // right
				fmt.Println("right")
			case "A", "ArrowLeft": // left
				fmt.Println("left")
			case "W", "ArrowUp":
				fmt.Println("up")
			case "S", "ArrowDown":
				fmt.Println("down")
			}
		}
	} else {
		m.keyIsBeingPressed = false
	}
}
