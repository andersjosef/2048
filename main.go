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
	game, err := twenty48.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(twenty48.SCREENWIDTH, twenty48.SCREENHEIGHT)
	ebiten.SetWindowTitle("2048")
	err = ebiten.RunGame(game)
	if err != nil {
		log.Fatal(err)
	}
}
