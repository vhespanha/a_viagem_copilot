package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct{}

func newGame() *Game {
	return &Game{}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	blue := color.RGBA{0, 0, 0xff, 0xff}
	var w, h float32
	w, h = 100.0, 100.0
	x, y := Center.Position(w, h)
	vector.FillRect(screen, x, y, w, h, blue, false)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
