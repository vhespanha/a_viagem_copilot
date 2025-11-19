// Package game implements the core game logic and state management.
package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/vhespanha/a_viagem/internal/dialogue"
	"github.com/vhespanha/a_viagem/internal/ui"
)

// Game represents the main game state and implements the ebiten.Game interface.
type Game struct {
	dialogue         *dialogue.Dialogue
	lastCheckPointID dialogue.ID
	ui               *ui.UI
}

// New creates and initializes a new game.
func New() *Game {
	g := &Game{
		ui:               ui.New(),
		dialogue:         dialogue.New(),
		lastCheckPointID: dialogue.NodeIDFirst,
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
		g.handleChoiceClick(click.ChoiceIndex)

	case ui.ClickDialogue:
		g.dialogue.Advance()
		g.showCurrentNode()
	}

	return nil
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

// Draw renders the game screen.
func (g *Game) Draw(screen *ebiten.Image) {
	ui.Draw(g.ui, screen)
}

// Layout defines the logical screen size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(ui.ScreenWidth), int(ui.ScreenHeight)
}
