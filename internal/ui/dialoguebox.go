package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/vhespanha/a_viagem/internal/geometry"
)

const (
	boxWidth       = ScreenWidth
	boxHeight      = ScreenHeight / 4
	textBoxScale   = 0.95
	choicesHeight  = boxHeight / 6
	choicesGap     = 8
	choicesPadding = 12
	dontKnowString = "Não sei..."
)

type DialogueBox struct {
	baseElement
	textBounds          *geometry.Rect
	normalFont, bigFont *text.GoTextFace
	dialogue            *string
	choices             []ChoiceButton
}

func (d *DialogueBox) Draw(screen *ebiten.Image) {
	drawFilledRect(screen, d.bounds, Gray)
	if d.dialogue != nil {
		drawDialogueText(screen, d.normalFont, d.dialogue, d.bounds)
	}
	if len(d.choices) != 0 {
		for _, choice := range d.choices {
			choice.Draw(screen)
		}
	}
}

func drawDialogueText(screen *ebiten.Image, font *text.GoTextFace, dialogue *string, bounds *geometry.Rect) {
	to := createTextDrawOptions(
		bounds.Pos.X, bounds.Pos.Y,
		font.Size*LineSpacing,
	)
	text.Draw(screen, *dialogue, font, to)
}

func NewDialogueBox(normalFont, bigFont *text.GoTextFace) *DialogueBox {
	ob := geometry.PositionRect(
		geometry.BottomCenter,
		boxWidth,
		boxWidth,
		ScreenWidth,
		ScreenHeight,
	)
	ib := ob.Scale(textBoxScale)

	return &DialogueBox{
		baseElement: baseElement{
			id:       DialogueBoxID,
			bounds:   ob,
			isActive: true,
		},
		textBounds: ib,
		choices:    make([]ChoiceButton, 0, 2),
		normalFont: normalFont,
		bigFont:    bigFont,
	}
}

type ChoiceButton struct {
	baseElement
	label string
	font  *text.GoTextFace
}

func (b *ChoiceButton) Draw(screen *ebiten.Image) {
	drawFilledRect(screen, b.bounds, Gray)
	drawCenteredText(screen, b.label, b.bounds, b.font)
}

func newChoiceButton(
	id ElementID,
	label string,
	bounds *geometry.Rect,
	font *text.GoTextFace,
) *ChoiceButton {
	return &ChoiceButton{
		baseElement: baseElement{
			id:       id,
			bounds:   bounds,
			isActive: true,
		},
		label: label,
		font:  font,
	}
}

func (d *DialogueBox) newChoices(choices []string, font *text.GoTextFace) []*ChoiceButton {
	n := len(choices)
	rects := allocChoiceBounds(n)
	a := make([]*ChoiceButton, n)
	switch n {
	case 1:
		newChoiceButton(ChoiceDontKnowID, dontKnowString, rects[0], font)
		return a
	case 2:
		newChoiceButton(ChoiceOneID, choices[0], rects[0], font)
		newChoiceButton(ChoiceTwoID, choices[1], rects[1], font)
		return a
	}
	return nil
}
