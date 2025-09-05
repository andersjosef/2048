package main

import (
	"log"

	"github.com/andersjosef/2048/app"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	a, err := app.NewApp()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowTitle("2048")
	err = ebiten.RunGame(a)
	if err != nil {
		log.Fatal(err)
	}
}
