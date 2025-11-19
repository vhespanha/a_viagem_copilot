package ui

import "github.com/hajimehoshi/ebiten/v2"

type UI struct {
	elements []element
	faces    *faces
}

func New() *UI {
	return &UI{
		elements: nil, // empty for now
		faces:    newFaces(),
	}
}

func (ui *UI) Draw(screen *ebiten.Image) {
	for _, el := range ui.elements {
		if el.Active() && el.Draw != nil {
			el.Draw(screen)
		}
	}
}

func (ui *UI) HandleClick(cx, cy int) {
	for _, el := range ui.elements {
		if el.Contains(cx, cy) && el.Active() {

		}
		return
	}
}

func (ui *UI) Command(cmd CommandID, data any) {
	switch cmd {
	case UpdateDialogueBox:
		ui.getElement(DialogueBoxID).Update(data)
	case HideDialogue:
		ui.getElement(DialogueBoxID).CustomClear()
	case ShowDeathScreen:
		ui.getElement(DeathScreenID).Activate()
	case HideDeathScreen:
		ui.getElement(DeathScreenID).Clear()
	case ToggleFullScreen:
		ui.getElement(FullScreenButtonID).Update(nil)
	}
}

func (ui *UI) getElement(id ElementID) element {
	for _, el := range ui.elements {
		if el.ID() == id {
			return el
		}
	}
	return nil
}
