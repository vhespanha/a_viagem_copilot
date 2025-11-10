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
	fonts            *ui.Fonts
	dialogue         *dialogue.Dialogue
	seenEvents       *map[string]bool
	lastCheckPointID dialogue.ID
	view             *ui.View
}

// New creates and initializes a new game.
func New() *Game {
	return &Game{
		fonts:            ui.NewFonts(),
		view:             ui.NewView(),
		dialogue:         dialogue.New(),
		seenEvents:       new(map[string]bool),
		lastCheckPointID: dialogue.NodeIDFirst,
	}
}

// Update handles game logic updates and processes input.
func (g *Game) Update() error {
	if !inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return nil
	}

	cx, cy := ebiten.CursorPosition()
	if g.view.ContainsFullScreenWidget(cx, cy) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}

	if g.view.IsOnDeathScreen() {
		g.jumpToLastCheckPoint()
		g.view.ToggleDeathScreen()
		return nil
	}

	node := g.dialogue.GetCurrentNode()

	if node.Choices != nil {
		return g.handleChoiceClick(cx, cy)
	}

	if g.view.ContainsDialogueBox(cx, cy) {
		g.dialogue.Advance()
	}

	return nil
}

func (g *Game) handleChoiceClick(cx, cy int) error {
	node := g.dialogue.GetCurrentNode()
	contains, choice := g.view.ContainsDialogueChoice(cx, cy)
	if contains {
		correct, err := node.Choose(choice)
		if err != nil || !correct {
			g.view.ToggleDeathScreen()
			return nil
		}
		g.dialogue.Advance()

	}
	return nil
}

func (g *Game) jumpToLastCheckPoint() {
	g.dialogue.CurrentID = g.lastCheckPointID
}

// Draw renders the game screen.
func (g *Game) Draw(screen *ebiten.Image) {
	g.view.DrawFullScreenWidget(screen)
	if g.view.IsOnDeathScreen() {
		ui.DrawDeathScreen(screen, g.fonts.Big)
		return
	}

	node := g.dialogue.GetCurrentNode()

	g.view.DrawDialogueBox(screen)
	g.view.DrawDialogueText(screen, g.fonts, node.Text)

	choicesText := make([]string, 0)
	if node.Choices != nil {
		if node.Unlocked {
			for _, choice := range *node.Choices {
				choicesText = append(choicesText, choice.Text)
			}
		} else {
			choicesText[0] = ui.DontKnowString
		}
	}
	g.view.DrawDialogueChoices(screen, g.fonts.Normal, choicesText)
}

// Layout defines the logical screen size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(ui.ScreenWidth), int(ui.ScreenHeight)
}
