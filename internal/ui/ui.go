package ui

import "github.com/hajimehoshi/ebiten/v2"

type UI struct {
	OnClick func(id ElementID) // defined by game logic?

	elements []element
	fonts    *Fonts
	screen   *ebiten.Image
}

func New(screen *ebiten.Image) *UI {
	return &UI{
		OnClick: nil, // maybe implemented on logic layer

		elements: nil, // empty for now
		fonts:    NewFonts(),
		screen:   screen,
	}
}

func (ui *UI) HandleClick(cx, cy int) {
	for _, el := range ui.elements {
		if el.Contains(cx, cy) && ui.OnClick != nil {
			ui.OnClick(el.ID())
		}
		return
	}
}

func (ui *UI) Draw() {
	for _, el := range ui.elements {
		if el.IsActive() && el.Draw != nil {
			el.Draw(ui.screen)
		}
	}
}
