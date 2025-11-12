package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/vhespanha/a_viagem/internal/geometry"
)

const deathScreenText = "Você escolheu errado..."

type DeathScreen struct {
	baseElement
	font *text.GoTextFace
}

func newDeathScreen(font *text.GoTextFace) *DeathScreen {
	return &DeathScreen{
		baseElement: baseElement{
			id:       DeathScreenID,
			bounds:   geometry.NewRect(0, 0, 1920, 1080),
			isActive: false,
		},
		font: font,
	}
}

func (d *DeathScreen) Draw(screen *ebiten.Image) {
	drawFilledRect(screen, d.bounds, Gray)
	drawCenteredText(screen, deathScreenText, d.bounds, d.font)
}

const fullScreenButtonSize float32 = 40

type FullScreenButton struct {
	baseElement
}

func newFullScreenButton() *FullScreenButton {
	return &FullScreenButton{
		baseElement: baseElement{
			id: FullScreenButtonID,
			bounds: geometry.PositionRect(
				geometry.TopRight.Offset(
					fullScreenButtonSize,
					fullScreenButtonSize,
				),
				fullScreenButtonSize, fullScreenButtonSize,
				ScreenWidth, ScreenHeight,
			),
			isActive: false,
		},
	}
}

func (b *FullScreenButton) Draw(screen *ebiten.Image) {
	drawFilledRect(screen, b.bounds, Gray)
}
