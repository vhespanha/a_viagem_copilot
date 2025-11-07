package main

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 1920
	screenHeight = 1080
)

const (
	fontSizeNormal = 24
	fontSizeBig    = 32
	lineSpacing    = 1.5
)

type Rectangle struct {
	w, h, x, y int
	col        color.RGBA
}

func (r *Rectangle) IsOnTop(cx, cy int) bool {
	return cx >= r.x && cx < r.x+r.w && cy >= r.y && cy < r.y+r.h
}

type Game struct {
	fonts          *Fonts
	dialogueSystem *DialogueSystem
}

type Fonts struct {
	source *text.GoTextFaceSource
	normal *text.GoTextFace
	big    *text.GoTextFace
}

func newFonts() *Fonts {
	faceSource, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}

	faceNormal = &text.GoTextFace{
		Source: faceSource,
		Size:   fontSizeNormal,
	}
	faceBig = &text.GoTextFace{
		Source: faceSource,
		Size:   fontSizeBig,
	}

	return &Fonts{
		source: faceSource,
		normal: faceNormal,
		big:    faceBig,
	}
}

func newGame() *Game {
	return &Game{newFonts(), NewDialogueSystem()}
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		cx, cy := ebiten.CursorPosition()
		if g.dialogueSystem.Box.IsOnTop(cx, cy) {
			const firstChoiceOption = 1
			err := g.dialogueSystem.Choose(firstChoiceOption)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	vector.FillRect(screen, float32(g.dialogueSystem.Box.x), float32(g.dialogueSystem.Box.y),
		float32(g.dialogueSystem.Box.w), float32(g.dialogueSystem.Box.h), g.dialogueSystem.Box.col, false)

	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(g.dialogueSystem.Box.x), float64(g.dialogueSystem.Box.y))
	op.LineSpacing = g.fonts.normal.Size * lineSpacing

	currentNode := g.dialogueSystem.Content[g.dialogueSystem.Current]
	dialogue := currentNode.Text
	text.Draw(screen, dialogue, g.fonts.normal, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}