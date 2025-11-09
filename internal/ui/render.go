// Package ui provides user interface rendering functions and resources.
package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/vhespanha/a_viagem/internal/dialogue"
	"github.com/vhespanha/a_viagem/internal/geometry"
)

const (
	ScreenWidth  = 1920
	ScreenHeight = 1080
)

const (
	dontKnowString    = "Não sei."
	wrongChoiceString = "Você escolheu errado.."
)

const (
	dialogueBoxWidth   = 1920 // screen width
	dialogueBoxHeight  = 200
	dialogueBoxOffsetY = -20
	choicesY           = 40
)

// DrawFilledRect draws a filled rectangle on the screen with the specified color.
func DrawFilledRect(s *ebiten.Image, rect *geometry.Rect, color color.RGBA) {
	vector.FillRect(
		s,
		rect.Pos.X, rect.Pos.Y,
		rect.Size.X, rect.Size.Y,
		color,
		false,
	)
}

// CreateTextDrawOptions creates text drawing options with the specified
// position and line spacing.
func CreateTextDrawOptions(x, y float32, lineSpacing float64) *text.DrawOptions {
	options := &text.DrawOptions{}
	options.GeoM.Translate(float64(x), float64(y))
	options.LineSpacing = lineSpacing
	return options
}

// DrawDialogueBox draws the dialogue box on the screen.
func DrawDialogueBox(s *ebiten.Image, d *dialogue.Dialogue) {
	box := geometry.PositionRect(
		geometry.BottomCenter.Offset(0, dialogueBoxOffsetY),
		dialogueBoxWidth, dialogueBoxHeight,
		ScreenWidth, ScreenHeight)
	DrawFilledRect(s, box, Gray)
}

// DrawDialogueText draws the dialogue text on the screen.
func DrawDialogueText(s *ebiten.Image, d *dialogue.Dialogue, fonts *Fonts) {
	box := d.Box
	currentNode := d.GetCurrentNode()

	textOptions := CreateTextDrawOptions(
		box.Pos.X,
		box.Pos.Y-dialogue.ChoicesY,
		fonts.Normal.Size*LineSpacing,
	)
	text.Draw(s, currentNode.Text, fonts.Normal, textOptions)
}

// DrawDialogueChoices draws the dialogue choices on the screen.
func DrawDialogueChoices(
	s *ebiten.Image,
	d *dialogue.Dialogue,
	font *text.GoTextFace,
	rects []*geometry.Rect,
) {
	currentNode := d.GetCurrentNode()
	if currentNode.Choices == nil {
		return
	}

	currentChoices := *currentNode.Choices
	if !currentNode.Unlocked {
		// draw "don't know" option
		drawSingleChoice(s, font, rects[0], dontKnowString)
		return
	}

	// draw all unlocked choices
	for i, choice := range currentChoices {
		drawSingleChoice(s, font, rects[i], choice.Text)
	}
}

func drawSingleChoice(
	s *ebiten.Image,
	font *text.GoTextFace,
	rect *geometry.Rect,
	choiceText string,
) {
	DrawFilledRect(s, rect, Gray)

	textWeight, textHeight := text.Measure(
		choiceText,
		font,
		font.Size*LineSpacing,
	)

	centeredTextPos := geometry.CenterTextInRect(
		float32(textWeight),
		float32(textHeight),
		rect,
	)

	choiceTextOptions := CreateTextDrawOptions(
		centeredTextPos.X,
		centeredTextPos.Y,
		font.Size*LineSpacing,
	)

	text.Draw(s, choiceText, font, choiceTextOptions)
}

// DrawDeathScreen draws the death/wrong choice screen.
func DrawDeathScreen(s *ebiten.Image, font *text.GoTextFace, screenWidth, screenHeight int) {
	rect := geometry.NewRect(0, 0, float32(screenWidth), float32(screenHeight))
	DrawFilledRect(s, rect, Gray)

	textWeight, textHeight := text.Measure(
		wrongChoiceString,
		font,
		font.Size*LineSpacing,
	)

	centeredTextPos := geometry.CenterTextInRect(
		float32(textWeight),
		float32(textHeight),
		rect,
	)

	textOption := CreateTextDrawOptions(
		centeredTextPos.X,
		centeredTextPos.Y,
		font.Size*LineSpacing,
	)

	text.Draw(s, wrongChoiceString, font, textOption)
}
