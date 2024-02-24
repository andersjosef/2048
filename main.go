package main

import (
	"log"
	"math/rand"
	"mygame/twenty48"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	game, _ := twenty48.NewGame()
	ebiten.SetWindowSize(twenty48.SCREENWIDTH, twenty48.SCREENHEIGHT)
	ebiten.SetWindowTitle("2048")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
