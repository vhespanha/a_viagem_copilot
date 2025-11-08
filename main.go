package main

import (
	"errors"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/vhespanha/tour_clara/game"
)

func main() {
	g := game.New()
	if err := ebiten.RunGame(g); err != nil {
		if !errors.Is(err, ebiten.Termination) {
			log.Fatal(err)
		}
	}
}
