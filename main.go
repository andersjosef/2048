package main

import (
	"log"

	"github.com/andersjosef/2048/twenty48"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game, err := twenty48.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowTitle("2048")
	err = ebiten.RunGame(game)
	if err != nil {
		log.Fatal(err)
	}
}
