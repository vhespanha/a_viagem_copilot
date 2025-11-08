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

// Game represents the main game state and implements the ebiten.Game interface.
type Game struct {
	fonts    *Fonts
	dialogue *Dialogue
	// clickables []*Clickable
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
		fonts:    newFonts(),
		dialogue: NewDialogue(),
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
	cursorX, cursorY := ebiten.CursorPosition()
	if g.dialogue.Box.Contains(cursorX, cursorY) {
		const firstChoiceOption = 1
		return g.dialogue.Choose(firstChoiceOption)
	}
	return nil
}

// Draw renders the game screen.
func (g *Game) Draw(screen *ebiten.Image) {
	g.drawDialogueBox(screen)
	g.drawDialogueText(screen)
	g.drawDialogueChoices(screen)
}

func (g *Game) drawDialogueBox(screen *ebiten.Image) {
	box := g.dialogue.Box
	boxColor := g.dialogue.BoxColor
	drawFilledRect(screen, box, boxColor)
}

func (g *Game) drawDialogueText(screen *ebiten.Image) {
	box := g.dialogue.Box
	currentNode := g.getCurrentDialogueNode()

	textOptions := g.createTextDrawOptions(box.Pos.X, box.Pos.Y-dialogueBoxChoicesY)
	text.Draw(screen, currentNode.Text, g.fonts.normal, textOptions)
}

func (g *Game) drawDialogueChoices(screen *ebiten.Image) {
	currentNode := g.getCurrentDialogueNode()
	if currentNode.Choice1 == nil {
		return
	}

	choicesRect := g.createChoicesRect()
	drawFilledRect(screen, choicesRect, Gray)

	textWeight, textHeight := text.Measure(currentNode.Choice1.Text, g.fonts.normal, g.fonts.normal.Size*lineSpacing)

	centeredTextPos := CenterTextInRect(float32(textWeight), float32(textHeight), choicesRect)

	choiceTextOptions := g.createTextDrawOptions(centeredTextPos.X, centeredTextPos.Y)
	text.Draw(screen, currentNode.Choice1.Text, g.fonts.normal, choiceTextOptions)
}

// getCurrentDialogueNode returns the currently active dialogue node.
func (g *Game) getCurrentDialogueNode() *DialogueNode {
	return g.dialogue.Content[g.dialogue.Current]
}

// drawFilledRect draws a filled rectangle on the screen with the specified color.
func drawFilledRect(screen *ebiten.Image, rect *Rect, color color.RGBA) {
	vector.FillRect(screen, rect.Pos.X, rect.Pos.Y, rect.Size.X, rect.Size.Y, color, false)
}

// createTextDrawOptions creates text drawing options with the specified position and line spacing.
func (g *Game) createTextDrawOptions(x, y float32) *text.DrawOptions {
	options := &text.DrawOptions{}
	options.GeoM.Translate(float64(x), float64(y))
	options.LineSpacing = g.fonts.normal.Size * lineSpacing
	return options
}

// createChoicesRect creates the rectangle for dialogue choices based on the dialogue box.
func (g *Game) createChoicesRect() *Rect {
	box := g.dialogue.Box
	return NewRect(0, box.Pos.Y+box.Size.Y-dialogueBoxChoicesY, box.Size.X, dialogueBoxChoicesY)
}

// Layout defines the logical screen size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
