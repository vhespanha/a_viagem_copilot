// Package game implements the core game logic and state management.
package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/vhespanha/a_viagem/internal/dialogue"
	"github.com/vhespanha/a_viagem/internal/ui"
)

// GameState represents the current mode of the game
type GameState int

const (
	StateDialogue GameState = iota
	StatePointAndClick
	StatePaused
)

// Game represents the main game state and implements the ebiten.Game interface.
type Game struct {
	dialogue         *dialogue.Dialogue
	lastCheckPointID dialogue.ID
	ui               *ui.UI
	state            GameState
}

// New creates and initializes a new game.
func New() *Game {
	g := &Game{
		ui:               ui.New(),
		dialogue:         dialogue.New(),
		lastCheckPointID: dialogue.NodeIDFirst,
		state:            StateDialogue, // start in dialogue mode
	}
	g.showCurrentNode()
	return g
}

// Update handles game logic updates and processes input.
func (g *Game) Update() error {
	if !inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return nil
	}

	cx, cy := ebiten.CursorPosition()
	click := ui.HandleClick(g.ui, cx, cy)

	switch click.Type {
	case ui.ClickFullscreen:
		ui.ToggleFullscreen()

	case ui.ClickDeathScreen:
		g.jumpToLastCheckPoint()
		ui.HideDeathScreen(g.ui)
		g.showCurrentNode()

	case ui.ClickChoice:
		if g.state == StateDialogue {
			g.handleChoiceClick(click.ChoiceIndex)
		}

	case ui.ClickDialogue:
		if g.state == StateDialogue {
			g.dialogue.Advance()
			g.showCurrentNode()
		}

	case ui.ClickElement:
		if g.state == StatePointAndClick {
			g.handleElementClick(click.ElementID)
		}
	}

	return nil
}

func (g *Game) handleElementClick(elementID ui.ElementID) {
	// Handle point-and-click element clicks
	// Game logic can be implemented here
	// For now, just a placeholder
}

func (g *Game) handleChoiceClick(choiceIndex int) {
	node := g.dialogue.GetCurrentNode()
	if node.Choices == nil {
		return
	}

	correct, err := node.Choose(choiceIndex)
	if err != nil || !correct {
		ui.ShowDeathScreen(g.ui)
		return
	}

	g.dialogue.Advance()
	g.showCurrentNode()
}

func (g *Game) jumpToLastCheckPoint() {
	g.dialogue.CurrentID = g.lastCheckPointID
}

func (g *Game) showCurrentNode() {
	node := g.dialogue.GetCurrentNode()

	// build choices text
	var choicesText []string
	if node.Choices != nil {
		if node.Unlocked {
			for _, choice := range *node.Choices {
				choicesText = append(choicesText, choice.Text)
			}
		} else {
			choicesText = []string{"Não sei..."}
		}
	}

	ui.ShowDialogue(g.ui, node.Text, choicesText)
}

// SetState changes the game state and manages UI transitions
func (g *Game) SetState(newState GameState) {
	g.state = newState

	switch newState {
	case StateDialogue:
		// Show dialogue UI, hide point-and-click elements
		ui.ClearElements(g.ui)
		g.showCurrentNode()

	case StatePointAndClick:
		// Hide dialogue, show point-and-click elements
		ui.HideDialogue(g.ui)
		// Game can add clickable elements here using ui.AddElement

	case StatePaused:
		// Pause state management
	}
}

// GetState returns the current game state
func (g *Game) GetState() GameState {
	return g.state
}

// Draw renders the game screen.
func (g *Game) Draw(screen *ebiten.Image) {
	ui.Draw(g.ui, screen)
}

// Layout defines the logical screen size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(ui.ScreenWidth), int(ui.ScreenHeight)
}
