package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/vhespanha/a_viagem/internal/geometry"
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

func allocChoiceBounds(n int) []*geometry.Rect {
	available := float32(boxWidth - choicesPadding*2)
	x := float32(0 + choicesPadding)
	y := float32(ScreenHeight - choicesHeight - choicesPadding)
	a := make([]*geometry.Rect, 0, 2)
	height := float32(choicesHeight)

	switch n {
	case 1:
		a[0] = geometry.NewRect(
			x, y,
			available, height,
		)
	case 2:
		width := (available - choicesGap) / 2
		a[0] = geometry.NewRect(
			x, y,
			width, height,
		)
		a[1] = geometry.NewRect(
			(x + width + choicesGap), y,
			width, height,
		)
	default:
		return nil
	}
	return a
}
