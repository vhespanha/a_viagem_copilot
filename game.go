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
	fonts            *Fonts
	dialogue         *Dialogue
	seenEvents       *map[string]bool
	lastCheckPointID ID
	isOnDeathScreen  bool

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

const (
	DontKnowChoiceString = "Não sei."
	WrongChoiceString    = "Você escolheu errado.."
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
		fonts:            newFonts(),
		dialogue:         NewDialogue(),
		seenEvents:       new(map[string]bool),
		lastCheckPointID: nodeIDFirst,
		isOnDeathScreen:  false,
	}
}

// Update handles game logic updates and processes input.
func (g *Game) Update() error {
	if !inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return nil
	}
	if g.isOnDeathScreen {
		g.jumpToLastCheckPoint()
		g.isOnDeathScreen = false
	}

	cx, cy := ebiten.CursorPosition()
	node := g.getCurrentDialogueNode()

	if node.Choices != nil {
		return g.handleChoiceClick(cx, cy)
	}

	if g.dialogue.Box.Contains(cx, cy) {
		g.dialogue.AdvanceDialogueNode()
	}

	return nil
}

func (g *Game) handleChoiceClick(cx, cy int) error {
	choiceRects := g.createChoiceRects(len(*g.dialogue.GetCurrentNode().Choices), 8, 12)
	for i, rect := range choiceRects {
		if rect.Contains(cx, cy) {
			correct, err := g.dialogue.GetCurrentNode().Choose(i)
			if err != nil || !correct {
				g.isOnDeathScreen = true
				return nil
			}
			g.dialogue.AdvanceDialogueNode()
		}
	}
	return nil
}

func (g *Game) jumpToLastCheckPoint() {
	g.dialogue.CurrentID = g.lastCheckPointID
}

// Draw renders the game screen.
func (g *Game) Draw(screen *ebiten.Image) {
	if g.isOnDeathScreen {
		g.drawDeathScreen(screen)
	} else {
		g.drawDialogueBox(screen)
		g.drawDialogueText(screen)
		g.drawDialogueChoices(screen)
	}

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

func (g *Game) drawDeathScreen(screen *ebiten.Image) {
	deathScreenRect := NewRect(0, 0, screenWidth, screenHeight)
	drawFilledRect(screen, deathScreenRect, Gray)
	textWeight, textHeight := text.Measure(WrongChoiceString, g.fonts.big, g.fonts.big.Size*lineSpacing)
	centeredTextPos := CenterTextInRect(float32(textWeight), float32(textHeight), deathScreenRect)
	deathScreenTextOption := g.createTextDrawOptions(centeredTextPos.X, centeredTextPos.Y)
	text.Draw(screen, WrongChoiceString, g.fonts.big, deathScreenTextOption)
}

func (g *Game) drawDialogueChoices(screen *ebiten.Image) {
	if g.getCurrentDialogueNode().Choices != nil {
		currentChoices := (*g.getCurrentDialogueNode().Choices)
		unlocked := g.getCurrentDialogueNode().Unlocked
		if unlocked {
			rects := g.createChoiceRects(len(currentChoices), 8, 12)
			for i := 0; i < len(currentChoices); i++ {
				drawFilledRect(screen, rects[i], Gray)
				textWeight, textHeight := text.Measure(currentChoices[i].Text, g.fonts.normal, g.fonts.normal.Size*lineSpacing)
				centeredTextPos := CenterTextInRect(float32(textWeight), float32(textHeight), rects[i])
				choiceTextOptions := g.createTextDrawOptions(centeredTextPos.X, centeredTextPos.Y)
				text.Draw(screen, currentChoices[i].Text, g.fonts.normal, choiceTextOptions)
			}
			return
		}
		// draw idk option
		DontKnowRect := g.createChoiceRects(1, 8, 12)
		drawFilledRect(screen, DontKnowRect[0], Gray)
		textWeight, textHeight := text.Measure(DontKnowChoiceString, g.fonts.normal, g.fonts.normal.Size*lineSpacing)
		centeredTextPos := CenterTextInRect(float32(textWeight), float32(textHeight), DontKnowRect[0])
		choiceTextOptions := g.createTextDrawOptions(centeredTextPos.X, centeredTextPos.Y)
		text.Draw(screen, DontKnowChoiceString, g.fonts.normal, choiceTextOptions)
		return
	}
}

// getCurrentDialogueNode returns the currently active dialogue node.
func (g *Game) getCurrentDialogueNode() *DialogueNode {
	return g.dialogue.Nodes[g.dialogue.CurrentID]
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

// createChoiceRects returns a slice of Rects that evenly split the horizontal
// space of the dialogue box for N choices. If n == 1 the single rect will
// take the full available width (minus optional padding). You can pass a gap
// (space between choices) and padding (space left and right inside the box).
func (g *Game) createChoiceRects(n int, gap, padding float32) []*Rect {
	box := g.dialogue.Box
	if n <= 0 {
		return nil
	}

	// compute available horizontal width inside the box after padding
	available := box.Size.X - padding*2
	if available < 0 {
		available = 0
	}

	// y and height of choice rects (keeps the same vertical location as before)
	y := box.Pos.Y + box.Size.Y - dialogueBoxChoicesY
	h := dialogueBoxChoicesY

	// if there is only one choice, give it the full available width
	if n == 1 {
		x := box.Pos.X + padding
		return []*Rect{NewRect(x, y, available, float32(h))}
	}

	// subtract total gaps between N items: (n-1) * gap
	totalGaps := gap * float32(n-1)
	usable := available - totalGaps
	if usable < 0 {
		// not enough space for the requested gap; clamp to zero width items
		usable = 0
	}

	width := usable / float32(n)

	// create rects positioned left-to-right
	rects := make([]*Rect, n)
	for i := 0; i < n; i++ {
		x := box.Pos.X + padding + float32(i)*(width+gap)
		rects[i] = NewRect(x, y, width, float32(h))
	}
	return rects
}

// Layout defines the logical screen size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
