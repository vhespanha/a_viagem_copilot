package main

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// Game represents the main game state and implements the ebiten.Game interface.
type Game struct {
	fonts          *Fonts
	dialogueSystem *DialogueSystem
}

// Fonts holds the font resources used for rendering text.
type Fonts struct {
	source *text.GoTextFaceSource
	normal *text.GoTextFace
	big    *text.GoTextFace
}

const (
	fontSizeNormal = 24
	fontSizeBig    = 32
	lineSpacing    = 1.5
)

const (
	screenWidth  = 1920
	screenHeight = 1080
)

func newFonts() *Fonts {
	faceSource, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}

	return &Fonts{
		source: faceSource,
		normal: createFace(faceSource, fontSizeNormal),
		big:    createFace(faceSource, fontSizeBig),
	}
}

func createFace(source *text.GoTextFaceSource, size float64) *text.GoTextFace {
	return &text.GoTextFace{
		Source: source,
		Size:   size,
	}
}

func newGame() *Game {
	return &Game{
		fonts:          newFonts(),
		dialogueSystem: NewDialogueSystem(),
	}
}

// Update handles game logic updates and processes input.
func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return g.handleMouseClick()
	}
	return nil
}

func (g *Game) handleMouseClick() error {
	cx, cy := ebiten.CursorPosition()
	if g.dialogueSystem.Box.Contains(cx, cy) {
		const firstChoiceOption = 1
		return g.dialogueSystem.Choose(firstChoiceOption)
	}
	return nil
}

// Draw renders the game screen.
func (g *Game) Draw(screen *ebiten.Image) {
	g.drawDialogueBox(screen)
	g.drawDialogueText(screen)
}

func (g *Game) drawDialogueBox(screen *ebiten.Image) {
	box := g.dialogueSystem.Box
	boxColor := g.dialogueSystem.BoxColor
	vector.FillRect(screen,
		box.Pos.X, box.Pos.Y,
		box.Size.X, box.Size.Y,
		boxColor, false)
}

func (g *Game) drawDialogueText(screen *ebiten.Image) {
	box := g.dialogueSystem.Box
	currentNode := g.dialogueSystem.Content[g.dialogueSystem.Current]

	op := &text.DrawOptions{}
	// it's ok to cast here, since it's a one off thing
	op.GeoM.Translate(float64(box.Pos.X), float64(box.Pos.Y))
	op.LineSpacing = g.fonts.normal.Size * lineSpacing

	text.Draw(screen, currentNode.Text, g.fonts.normal, op)
}

// Layout defines the logical screen size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
