package main

import (
	"log"
	"mygame/twenty48"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game, _ := twenty48.NewGame()
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
