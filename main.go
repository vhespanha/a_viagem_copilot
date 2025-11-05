package main

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const sampleText = "The quick brown fox jumps over the lazy dog."

var (
	mplusFaceSource *text.GoTextFaceSource
	mplusNormalFace *text.GoTextFace
)

const (
	screenWidth  = 1920
	screenHeight = 1080
)

type Game struct{}

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s

	mplusNormalFace = &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   24,
	}
}

func (g *Game) Update() error {
	return nil
}

func drawDialogueBox(msg string, screen *ebiten.Image) {
	gray := color.RGBA{0x80, 0x80, 0x80, 0xff}
	w, h := text.Measure(msg, mplusNormalFace, mplusNormalFace.Size*1.5)
	var x, y = (screenWidth / 2) - (w / 2), screenHeight - h
	vector.FillRect(screen, float32(x), float32(y), float32(w), float32(h), gray, false)
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	op.LineSpacing = mplusNormalFace.Size * 1.5
	text.Draw(screen, msg, mplusNormalFace, op)
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawDialogueBox(sampleText, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Text (Ebitengine Demo)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
