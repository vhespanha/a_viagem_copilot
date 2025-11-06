package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	Blue  = color.RGBA{0, 0, 0xff, 0xff}
	Red   = color.RGBA{0xff, 0, 0, 0xff}
	Green = color.RGBA{0, 0xff, 0, 0xff}
)

type Rectangle struct {
	size, x, y int
	col        color.RGBA
}

func (r *Rectangle) IsClicked(cx, cy int) bool {
	return cx >= r.x && cx < r.x+r.size && cy >= r.y && cy < r.y+r.size
}

func (r *Rectangle) SwapColor() {
	if r.col == Green {
		r.col = Blue
	} else {
		r.col = Green
	}
}

type Game struct {
	rect *Rectangle
}

func newGame() *Game {
	rx, ry := Center.Position(100, 100)
	rectangle := &Rectangle{
		size: 100,
		x:    int(rx),
		y:    int(ry),
		col:  Green,
	}
	return &Game{rectangle}
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		cx, cy := ebiten.CursorPosition()
		if g.rect.IsClicked(cx, cy) {
			g.rect.SwapColor()
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	vector.FillRect(screen, float32(g.rect.x), float32(g.rect.y), float32(g.rect.size), float32(g.rect.size), g.rect.col, false)

	cx, cy := ebiten.CursorPosition()
	var r float32
	r = 20
	vector.FillCircle(screen, float32(cx), float32(cy), r, Red, false)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
