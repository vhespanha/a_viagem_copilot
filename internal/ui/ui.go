package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/vhespanha/a_viagem/internal/geometry"
)

// ElementID identifies different UI elements
type ElementID string

// Element represents a clickable UI element
type Element struct {
	ID      ElementID
	Bounds  *geometry.Rect
	Visible bool
	Data    any // element-specific data
}

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

	// point-and-click elements
	clickableElements []Element
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
		clickableElements:  make([]Element, 0),
	}
}

// Draw renders the UI
func Draw(ui *UI, screen *ebiten.Image) {
	if ui.deathScreenVisible {
		drawDeathScreen(screen, ui.deathScreenBounds, ui.bigFont)
		return
	}

	// draw point-and-click elements
	for _, elem := range ui.clickableElements {
		if elem.Visible {
			drawClickableElement(screen, &elem, ui.normalFont)
		}
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

	// check dialogue choices first (higher priority than dialogue box)
	for i, bounds := range ui.choiceBounds {
		if bounds != nil && bounds.Contains(cx, cy) {
			return ClickResult{Type: ClickChoice, ChoiceIndex: i}
		}
	}

	// check point-and-click elements
	for i, elem := range ui.clickableElements {
		if elem.Visible && elem.Bounds != nil && elem.Bounds.Contains(cx, cy) {
			return ClickResult{Type: ClickElement, ElementIndex: i, ElementID: elem.ID}
		}
	}

	// dialogue box - only clickable when no choices are shown
	if ui.dialogueVisible && len(ui.choices) == 0 && ui.dialogueBox.Contains(cx, cy) {
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
	ClickElement // for point-and-click elements
)

// ClickResult contains information about what was clicked
type ClickResult struct {
	Type         ClickType
	ChoiceIndex  int       // only relevant for ClickChoice
	ElementIndex int       // only relevant for ClickElement
	ElementID    ElementID // only relevant for ClickElement
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

// AddElement adds a clickable element to the UI
func AddElement(ui *UI, id ElementID, bounds *geometry.Rect, data any) {
	ui.clickableElements = append(ui.clickableElements, Element{
		ID:      id,
		Bounds:  bounds,
		Visible: true,
		Data:    data,
	})
}

// RemoveElement removes a clickable element by ID
func RemoveElement(ui *UI, id ElementID) {
	for i, elem := range ui.clickableElements {
		if elem.ID == id {
			ui.clickableElements = append(
				ui.clickableElements[:i],
				ui.clickableElements[i+1:]...)
			return
		}
	}
}

// ClearElements removes all clickable elements
func ClearElements(ui *UI) {
	ui.clickableElements = make([]Element, 0)
}

// SetElementVisible sets the visibility of an element
func SetElementVisible(ui *UI, id ElementID, visible bool) {
	for i := range ui.clickableElements {
		if ui.clickableElements[i].ID == id {
			ui.clickableElements[i].Visible = visible
			return
		}
	}
}

// GetElement returns an element by ID, or nil if not found
func GetElement(ui *UI, id ElementID) *Element {
	for i := range ui.clickableElements {
		if ui.clickableElements[i].ID == id {
			return &ui.clickableElements[i]
		}
	}
	return nil
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
