// Package game implements the core game logic and state management.
package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/vhespanha/a_viagem/internal/dialogue"
	"github.com/vhespanha/a_viagem/internal/geometry"
	"github.com/vhespanha/a_viagem/internal/ui"
)

// Game represents the main game state and implements the ebiten.Game interface.
type Game struct {
	fonts            *ui.Fonts
	dialogue         *dialogue.Dialogue
	seenEvents       *map[string]bool
	lastCheckPointID dialogue.ID
	isOnDeathScreen  bool
}

// New creates and initializes a new game.
func New() *Game {
	return &Game{
		fonts:            ui.NewFonts(),
		dialogue:         dialogue.New(ui.ScreenWidth, ui.ScreenHeight),
		seenEvents:       new(map[string]bool),
		lastCheckPointID: dialogue.NodeIDFirst,
		isOnDeathScreen:  false,
	}
}

// Update handles game logic updates and processes input.
func (g *Game) Update() error {
	if !inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return nil
	}

	if g.isOnDeathScreen {
		g.jumpToLastCheckPoint()
		g.isOnDeathScreen = false
		return nil
	}

	cx, cy := ebiten.CursorPosition()
	node := g.dialogue.GetCurrentNode()

	if node.Choices != nil {
		return g.handleChoiceClick(cx, cy)
	}

	if g.dialogue.Box.Contains(cx, cy) {
		g.dialogue.Advance()
	}

	return nil
}

func (g *Game) handleChoiceClick(cx, cy int) error {
	node := g.dialogue.GetCurrentNode()
	numChoices := getNumChoices(node)

	choiceRects := createChoiceRects(g.dialogue.Box, numChoices, 8, 12)
	for i, rect := range choiceRects {
		if rect.Contains(cx, cy) {
			correct, err := node.Choose(i)
			if err != nil || !correct {
				g.isOnDeathScreen = true
				return nil
			}
			g.dialogue.Advance()
		}
	}
	return nil
}

func (g *Game) jumpToLastCheckPoint() {
	g.dialogue.CurrentID = g.lastCheckPointID
}

// Draw renders the game screen.
func (g *Game) Draw(screen *ebiten.Image) {
	if g.isOnDeathScreen {
		ui.DrawDeathScreen(screen, g.fonts.Big, ui.ScreenWidth, ui.ScreenHeight)
		return
	}

	ui.DrawDialogueBox(screen, g.dialogue)
	ui.DrawDialogueText(screen, g.dialogue, g.fonts)

	node := g.dialogue.GetCurrentNode()
	if node.Choices != nil {
		numChoices := getNumChoices(node)
		choiceRects := createChoiceRects(
			g.dialogue.Box,
			numChoices,
			8,
			12,
		)
		ui.DrawDialogueChoices(screen, g.dialogue, g.fonts.Normal, choiceRects)
	}
}

// Layout defines the logical screen size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ui.ScreenWidth, ui.ScreenHeight
}

// getNumChoices returns the number of choices to display for a dialogue node.
// Returns 1 for locked nodes (showing "don't know" option) or the actual
// number of choices for unlocked nodes.
func getNumChoices(node *dialogue.Node) int {
	if node.Unlocked && node.Choices != nil {
		return len(*node.Choices)
	}
	return 1 // default for "don't know" option
}

// createChoiceRects returns a slice of Rects that evenly split the horizontal
// space of the dialogue box for N choices. If n == 1 the single rect will
// take the full available width (minus optional padding). You can pass a gap
// (space between choices) and padding (space left and right inside the box).
func createChoiceRects(box *geometry.Rect, n int, gap, padding float32) []*geometry.Rect {
	if n <= 0 {
		return nil
	}

	// compute available horizontal width inside the box after padding
	available := box.Size.X - padding*2
	if available < 0 {
		available = 0
	}

	// y and height of choice rects (keeps the same vertical location as
	// before)
	y := box.Pos.Y + box.Size.Y - dialogue.ChoicesY
	h := dialogue.ChoicesY

	// if there is only one choice, give it the full available width
	if n == 1 {
		x := box.Pos.X + padding
		return []*geometry.Rect{
			geometry.NewRect(x, y, available, float32(h)),
		}
	}

	// subtract total gaps between N items: (n-1) * gap
	totalGaps := gap * float32(n-1)
	usable := available - totalGaps
	if usable < 0 {
		// not enough space for the requested gap; clamp to zero width
		// items
		usable = 0
	}

	width := usable / float32(n)

	// create rects positioned left-to-right
	rects := make([]*geometry.Rect, n)
	for i := 0; i < n; i++ {
		x := box.Pos.X + padding + float32(i)*(width+gap)
		rects[i] = geometry.NewRect(x, y, width, float32(h))
	}
	return rects
}
