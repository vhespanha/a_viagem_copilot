package main

import (
	"errors"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	SCREEN_WIDTH  = 1920
	SCREEN_HEIGHT = 1080
)

var (
	faceSource *text.GoTextFaceSource
	faceNormal *text.GoTextFace
	faceBig    *text.GoTextFace
)

func main() {
	game := newGame()
	if err := ebiten.RunGame(game); err != nil {
		if !errors.Is(err, ebiten.Termination) {
			log.Fatal(err)
		}
	}
}