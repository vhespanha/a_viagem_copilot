package main

import (
	"errors"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 1920
	ScreenHeight = 1080
)

func main() {
	game := newGame()
	if err := ebiten.RunGame(game); err != nil {
		if !errors.Is(err, ebiten.Termination) {
			log.Fatal(err)
		}
	}
}
