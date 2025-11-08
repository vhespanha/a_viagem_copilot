package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/vhespanha/tour_clara/dialogue"
	"github.com/vhespanha/tour_clara/geometry"
)

const (
	DontKnowChoiceString = "Não sei."
	WrongChoiceString    = "Você escolheu errado.."
)

// DrawFilledRect draws a filled rectangle on the screen with the specified color.
func DrawFilledRect(screen *ebiten.Image, rect *geometry.Rect, color color.RGBA) {
	vector.FillRect(screen, rect.Pos.X, rect.Pos.Y, rect.Size.X, rect.Size.Y, color, false)
}

// CreateTextDrawOptions creates text drawing options with the specified position and line spacing.
func CreateTextDrawOptions(x, y, lineSpacing float32) *text.DrawOptions {
	options := &text.DrawOptions{}
	options.GeoM.Translate(float64(x), float64(y))
	options.LineSpacing = float64(lineSpacing)
	return options
}

// DrawDialogueBox draws the dialogue box on the screen.
func DrawDialogueBox(screen *ebiten.Image, d *dialogue.Dialogue) {
	DrawFilledRect(screen, d.Box, d.BoxColor)
}

// DrawDialogueText draws the dialogue text on the screen.
func DrawDialogueText(screen *ebiten.Image, d *dialogue.Dialogue, fonts *Fonts) {
	box := d.Box
	currentNode := d.GetCurrentNode()

	textOptions := CreateTextDrawOptions(box.Pos.X, box.Pos.Y-dialogue.ChoicesY, float32(fonts.Normal.Size*LineSpacing))
	text.Draw(screen, currentNode.Text, fonts.Normal, textOptions)
}

// DrawDialogueChoices draws the dialogue choices on the screen.
func DrawDialogueChoices(screen *ebiten.Image, d *dialogue.Dialogue, fonts *Fonts, choiceRects []*geometry.Rect) {
	currentNode := d.GetCurrentNode()
	if currentNode.Choices == nil {
		return
	}

	currentChoices := *currentNode.Choices
	if !currentNode.Unlocked {
		// draw "don't know" option
		drawSingleChoice(screen, fonts, choiceRects[0], DontKnowChoiceString)
		return
	}

	// draw all unlocked choices
	for i, choice := range currentChoices {
		drawSingleChoice(screen, fonts, choiceRects[i], choice.Text)
	}
}

func drawSingleChoice(screen *ebiten.Image, fonts *Fonts, rect *geometry.Rect, choiceText string) {
	DrawFilledRect(screen, rect, Gray)
	textWeight, textHeight := text.Measure(choiceText, fonts.Normal, fonts.Normal.Size*LineSpacing)
	centeredTextPos := geometry.CenterTextInRect(float32(textWeight), float32(textHeight), rect)
	choiceTextOptions := CreateTextDrawOptions(centeredTextPos.X, centeredTextPos.Y, float32(fonts.Normal.Size*LineSpacing))
	text.Draw(screen, choiceText, fonts.Normal, choiceTextOptions)
}

// DrawDeathScreen draws the death/wrong choice screen.
func DrawDeathScreen(screen *ebiten.Image, fonts *Fonts, screenWidth, screenHeight int) {
	deathScreenRect := geometry.NewRect(0, 0, float32(screenWidth), float32(screenHeight))
	DrawFilledRect(screen, deathScreenRect, Gray)
	textWeight, textHeight := text.Measure(WrongChoiceString, fonts.Big, fonts.Big.Size*LineSpacing)
	centeredTextPos := geometry.CenterTextInRect(float32(textWeight), float32(textHeight), deathScreenRect)
	deathScreenTextOption := CreateTextDrawOptions(centeredTextPos.X, centeredTextPos.Y, float32(fonts.Big.Size*LineSpacing))
	text.Draw(screen, WrongChoiceString, fonts.Big, deathScreenTextOption)
}
