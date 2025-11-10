package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/vhespanha/a_viagem/internal/geometry"
)

const (
	ScreenWidth  = 1920
	ScreenHeight = 1080
)

const (
	DontKnowString    = "Não sei."
	WrongChoiceString = "Você escolheu errado.."
)

const (
	FullScreenWidgetSize = 40
)

type View struct {
	dialogueBox      *box
	fullScreenWidget *geometry.Rect
	isOnDeathScreen  bool
}

func NewView() *View {
	return &View{
		dialogueBox: newDialogueBox(),
		fullScreenWidget: geometry.PositionRect(
			geometry.TopRight.Offset(-12, 12),
			FullScreenWidgetSize, FullScreenWidgetSize,
			ScreenWidth, ScreenHeight,
		),
	}
}

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

// DrawDeathScreen draws the death/wrong choice screen.
func DrawDeathScreen(s *ebiten.Image, font *text.GoTextFace) {
	rect := geometry.NewRect(0, 0, float32(ScreenWidth), float32(ScreenHeight))
	DrawFilledRect(s, rect, Gray)

	textWeight, textHeight := text.Measure(
		WrongChoiceString,
		font,
		font.Size*LineSpacing,
	)

	centeredTextPos := geometry.CenterInRect(
		float32(textWeight),
		float32(textHeight),
		rect,
	)

	textOption := CreateTextDrawOptions(
		centeredTextPos.X,
		centeredTextPos.Y,
		font.Size*LineSpacing,
	)

	text.Draw(s, WrongChoiceString, font, textOption)
}

func (v *View) IsOnDeathScreen() bool {
	return v.isOnDeathScreen
}

func (v *View) ToggleDeathScreen() {
	v.isOnDeathScreen = !v.isOnDeathScreen
}

func (v *View) DrawFullScreenWidget(s *ebiten.Image) {
	DrawFilledRect(s, v.fullScreenWidget, Gray)
}

func (v *View) ContainsFullScreenWidget(cx, cy int) bool {
	return v.fullScreenWidget.Contains(cx, cy)
}
