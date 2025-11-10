package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/vhespanha/a_viagem/internal/geometry"
)

const (
	boxWidth       = ScreenWidth
	boxHeight      = ScreenHeight / 4
	textBoxWidth   = boxWidth * 0.95
	textBoxHeight  = boxHeight * 0.95
	choicesHeight  = boxHeight / 6
	choicesGap     = 8
	choicesPadding = 12
)

type box struct {
	container     *geometry.Rect
	textContainer *geometry.Rect
	choices       []*geometry.Rect
}

func newDialogueBox() *box {
	outer := geometry.PositionRect(
		geometry.BottomCenter,
		boxWidth,
		boxHeight,
		ScreenWidth,
		ScreenHeight,
	)
	innerPos := geometry.CenterInRect(textBoxWidth, textBoxHeight, outer)
	inner := geometry.NewRect(
		innerPos.X, innerPos.Y,
		textBoxWidth, textBoxHeight-choicesHeight-choicesPadding,
	)
	return &box{
		container:     outer,
		textContainer: inner,
	}
}

// DrawDialogueBox draws the dialogue box on the screen.
func (v *View) DrawDialogueBox(s *ebiten.Image) {
	DrawFilledRect(s, v.dialogueBox.container, Gray)
}

// DrawDialogueText draws the dialogue text on the screen.
func (v *View) DrawDialogueText(s *ebiten.Image, fonts *Fonts, dialogueText string) {
	box := v.dialogueBox

	textOptions := CreateTextDrawOptions(
		box.textContainer.Pos.X,
		box.textContainer.Pos.Y,
		fonts.Normal.Size*LineSpacing,
	)
	text.Draw(s, dialogueText, fonts.Normal, textOptions)
}

func (v *View) DrawDialogueChoices(s *ebiten.Image, font *text.GoTextFace, choicesText []string) {
	v.dialogueBox.updateChoiceRects(len(choicesText))

	for i, choice := range choicesText {
		drawSingleChoice(s, font, v.dialogueBox.choices[i], choice)
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

	centeredTextPos := geometry.CenterInRect(
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

// UpdateChoiceRects returns a slice of Rects that evenly split the horizontal
// space of the dialogue box for N choices. If n == 1 the single rect will
// take the full available width (minus optional padding). You can pass a gap
// (space between choices) and padding (space left and right inside the box).
func (b *box) updateChoiceRects(n int) {
	if n <= 0 {
		return
	}

	// compute available horizontal width inside the box after padding
	available := b.container.Size.X - choicesPadding*2
	if available < 0 {
		available = 0
	}

	// y and height of choice rects (keeps the same vertical location as
	// before)
	y := b.container.Pos.Y + b.container.Size.Y - choicesHeight
	h := choicesHeight

	// if there is only one choice, give it the full available width
	if n == 1 {
		x := b.container.Pos.X + choicesPadding
		b.choices = []*geometry.Rect{
			geometry.NewRect(x, y, available, float32(h)),
		}
	}

	// subtract total gaps between N items: (n-1) * gap
	totalGaps := choicesGap * float32(n-1)
	usable := available - totalGaps
	if usable < 0 {
		// not enough space for the requested gap; clamp to zero width
		// items
		usable = 0
	}

	width := usable / float32(n)

	// create rects positioned left-to-right
	rects := make([]*geometry.Rect, n)
	for i := 0; i < n; i++ {
		x := b.container.Pos.X + choicesPadding + float32(i)*(width+choicesGap)
		rects[i] = geometry.NewRect(x, y, width, float32(h))
	}
	b.choices = rects
}

func (v *View) ContainsDialogueBox(cx, cy int) bool {
	return v.dialogueBox.container.Contains(cx, cy)
}

func (v *View) ContainsDialogueChoice(cx, cy int) (bool, int) {
	for i, rect := range v.dialogueBox.choices {
		if rect.Contains(cx, cy) {
			return true, i
		}
	}
	return false, 0
}
