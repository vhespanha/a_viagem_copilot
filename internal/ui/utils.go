package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/vhespanha/a_viagem/internal/geometry"
)

const (
	choicesHeight  = ScreenHeight / 24
	choicesGap     = 8
	choicesPadding = 12
	dontKnowString = "Não sei..."
)

// drawFilledRect draws a filled rectangle on the screen with the specified color.
func drawFilledRect(s *ebiten.Image, rect *geometry.Rect, color color.RGBA) {
	vector.FillRect(
		s,
		rect.Pos.X, rect.Pos.Y,
		rect.Size.X, rect.Size.Y,
		color,
		false,
	)
}

// createTextDrawOptions creates text drawing options with the specified
// position and line spacing.
func createTextDrawOptions(x, y float32, lineSpacing float64) *text.DrawOptions {
	options := &text.DrawOptions{}
	options.GeoM.Translate(float64(x), float64(y))
	options.LineSpacing = lineSpacing
	return options
}

func drawCenteredText(
	screen *ebiten.Image,
	s string,
	bounds *geometry.Rect,
	font *text.GoTextFace,
) {
	tw, th := text.Measure(
		s,
		font,
		font.Size*lineSpacing,
	)
	centeredTextPos := geometry.CenterInRect(
		float32(tw), float32(th), bounds,
	)
	textOption := createTextDrawOptions(
		centeredTextPos.X, centeredTextPos.Y, font.Size*lineSpacing,
	)
	text.Draw(screen, s, font, textOption)
}

func drawDialogueBox(
	screen *ebiten.Image,
	bounds *geometry.Rect,
	font *text.GoTextFace,
	dialogue string,
) {
	drawFilledRect(screen, bounds, Gray)
	scaled := bounds.Scale(0.95)
	to := createTextDrawOptions(scaled.Pos.X, scaled.Pos.Y, font.Size*lineSpacing)
	text.Draw(screen, dialogue, font, to)
}

func drawChoices(
	screen *ebiten.Image,
	bounds []*geometry.Rect,
	choices []string,
	font *text.GoTextFace,
) {
	for i, choice := range choices {
		if i < len(bounds) && bounds[i] != nil {
			drawFilledRect(screen, bounds[i], Gray)
			drawCenteredText(screen, choice, bounds[i], font)
		}
	}
}

func drawDeathScreen(screen *ebiten.Image, bounds *geometry.Rect, font *text.GoTextFace) {
	drawFilledRect(screen, bounds, Gray)
	drawCenteredText(screen, "Você escolheu errado...", bounds, font)
}

func drawClickableElement(screen *ebiten.Image, elem *Element, font *text.GoTextFace) {
	// Default rendering: semi-transparent rect with ID as text
	// Games can customize by checking elem.Data
	drawFilledRect(screen, elem.Bounds, color.RGBA{0x3f, 0x3f, 0x3f, 0x7f})
	if elem.ID != "" {
		drawCenteredText(screen, string(elem.ID), elem.Bounds, font)
	}
}

func allocChoiceBounds(n int) []*geometry.Rect {
	if n == 0 {
		return nil
	}

	available := float32(ScreenWidth - choicesPadding*2)
	x := float32(choicesPadding)
	y := float32(ScreenHeight - choicesHeight - choicesPadding)
	height := float32(choicesHeight)

	a := make([]*geometry.Rect, n)
	switch n {
	case 1:
		a[0] = geometry.NewRect(x, y, available, height)
	case 2:
		width := (available - choicesGap) / 2
		a[0] = geometry.NewRect(x, y, width, height)
		a[1] = geometry.NewRect(x+width+choicesGap, y, width, height)
	}
	return a
}
