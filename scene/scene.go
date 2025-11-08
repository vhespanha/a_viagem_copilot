// Package scene provides scene definitions and background management.
package scene

import "github.com/hajimehoshi/ebiten/v2"

type Scene struct {
	ID              string
	StartDialogueID string
	BackgroundImg   *ebiten.Image
}

// probably gonna hand-roll all scenes manually here
