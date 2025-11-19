package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/vhespanha/a_viagem/internal/geometry"
)

// UI holds all UI state and resources
type UI struct {
	// fonts
	normalFont *text.GoTextFace
	bigFont    *text.GoTextFace

	// dialogue box state
	dialogueText    string
	dialogueVisible bool
	dialogueBox     *geometry.Rect
	choices         []string
	choiceBounds    []*geometry.Rect

	// death screen state
	deathScreenVisible bool
	deathScreenBounds  *geometry.Rect

	// fullscreen button
	fullscreenBounds *geometry.Rect
}

// New creates a new UI
func New() *UI {
	faces := newFaces()
	return &UI{
		normalFont:         faces.normal,
		bigFont:            faces.big,
		dialogueBox:        makeDialogueBox(),
		deathScreenBounds:  geometry.NewRect(0, 0, ScreenWidth, ScreenHeight),
		fullscreenBounds:   makeFullscreenButton(),
		dialogueVisible:    false,
		deathScreenVisible: false,
	}
}

// Draw renders the UI
func Draw(ui *UI, screen *ebiten.Image) {
	if ui.deathScreenVisible {
		drawDeathScreen(screen, ui.deathScreenBounds, ui.bigFont)
		return
	}

	if ui.dialogueVisible {
		drawDialogueBox(screen, ui.dialogueBox, ui.normalFont, ui.dialogueText)
		drawChoices(screen, ui.choiceBounds, ui.choices, ui.normalFont)
	}

	// fullscreen button always visible
	drawFilledRect(screen, ui.fullscreenBounds, Gray)
}

// HandleClick processes mouse clicks and returns what was clicked
func HandleClick(ui *UI, cx, cy int) ClickResult {
	// death screen takes priority
	if ui.deathScreenVisible && ui.deathScreenBounds.Contains(cx, cy) {
		return ClickResult{Type: ClickDeathScreen}
	}

	// fullscreen button
	if ui.fullscreenBounds.Contains(cx, cy) {
		return ClickResult{Type: ClickFullscreen}
	}

	// check dialogue choices
	for i, bounds := range ui.choiceBounds {
		if bounds != nil && bounds.Contains(cx, cy) {
			return ClickResult{Type: ClickChoice, ChoiceIndex: i}
		}
	}

	// dialogue box
	if ui.dialogueVisible && ui.dialogueBox.Contains(cx, cy) {
		return ClickResult{Type: ClickDialogue}
	}

	return ClickResult{Type: ClickNone}
}

// ClickType represents what was clicked
type ClickType int

const (
	ClickNone ClickType = iota
	ClickDialogue
	ClickChoice
	ClickDeathScreen
	ClickFullscreen
)

// ClickResult contains information about what was clicked
type ClickResult struct {
	Type        ClickType
	ChoiceIndex int // only relevant for ClickChoice
}

// ShowDialogue displays dialogue text with optional choices
func ShowDialogue(ui *UI, text string, choices []string) {
	ui.dialogueVisible = true
	ui.dialogueText = text
	ui.choices = choices
	ui.choiceBounds = allocChoiceBounds(len(choices))
}

// HideDialogue hides the dialogue box
func HideDialogue(ui *UI) {
	ui.dialogueVisible = false
	ui.dialogueText = ""
	ui.choices = nil
	ui.choiceBounds = nil
}

// ShowDeathScreen shows the death screen
func ShowDeathScreen(ui *UI) {
	ui.deathScreenVisible = true
}

// HideDeathScreen hides the death screen
func HideDeathScreen(ui *UI) {
	ui.deathScreenVisible = false
}

// ToggleFullscreen toggles fullscreen mode
func ToggleFullscreen() {
	ebiten.SetFullscreen(!ebiten.IsFullscreen())
}

func makeDialogueBox() *geometry.Rect {
	return geometry.PositionRect(
		geometry.BottomCenter,
		ScreenWidth,
		ScreenHeight/4,
		ScreenWidth,
		ScreenHeight,
	)
}

func makeFullscreenButton() *geometry.Rect {
	size := float32(40)
	return geometry.PositionRect(
		geometry.TopRight.Offset(-size, size),
		size, size,
		ScreenWidth, ScreenHeight,
	)
}
