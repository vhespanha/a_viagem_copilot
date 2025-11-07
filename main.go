package main

import (
	"errors"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := newGame()
	if err := ebiten.RunGame(game); err != nil {
		if !errors.Is(err, ebiten.Termination) {
			log.Fatal(err)
		}
	}
}
