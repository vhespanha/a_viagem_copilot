// Package dialogue provides the dialogue system for managing conversation nodes and choices.
package dialogue

import (
	"errors"
	"image/color"

	"github.com/vhespanha/tour_clara/geometry"
)

const (
	BoxWidth   = 1920 // screen width
	BoxHeight  = 200
	BoxOffsetY = -20
	ChoicesY   = 40 // TODO: figure out if this is sane
)

var (
	ErrorDoesntKnow = errors.New("didn't know answer")
)

// ID is a unique identifier for dialogue nodes.
type ID string

// Predefined dialogue node IDs
const (
	NodeIDFirst  ID = "first"
	NodeIDSecond ID = "second"
)

// Node represents a single node in the dialogue tree.
type Node struct {
	Speaker  string
	Text     string
	Unlocked bool
	NextID   ID
	Choices  *[]Choice
}

// Choice represents a player's dialogue choice leading to another node.
type Choice struct {
	Text    string
	Correct bool
}

// Dialogue manages the dialogue state and progression.
type Dialogue struct {
	Nodes     map[ID]*Node
	CurrentID ID
	Box       *geometry.Rect
	BoxColor  color.RGBA
}

// New creates and initializes a new dialogue system.
func New(screenWidth, screenHeight int) *Dialogue {
	firstChoice := Choice{
		Text:    "yes",
		Correct: true,
	}
	secondChoice := Choice{
		Text:    "no",
		Correct: false,
	}
	return &Dialogue{
		Nodes: map[ID]*Node{
			NodeIDFirst: &Node{
				Speaker: "character",
				Text:    "bla bla bla",
				Choices: nil,
				NextID:  NodeIDSecond,
			},
			NodeIDSecond: &Node{
				Speaker:  "character",
				Text:     "bla bla bla??",
				Unlocked: true,
				Choices:  &[]Choice{firstChoice, secondChoice},
				NextID:   NodeIDFirst,
			},
		},
		CurrentID: NodeIDFirst,
		Box: geometry.PositionRect(
			geometry.BottomCenter.Offset(0, BoxOffsetY),
			BoxWidth, BoxHeight,
			screenWidth, screenHeight,
		),
		BoxColor: color.RGBA{0, 0, 0, 0xff}, // Black
	}
}

// Choose processes a dialogue choice and advances to the next node.
// Returns an error if the choice number is invalid.
func (n *Node) Choose(choice int) (bool, error) {
	if n.Unlocked {
		return (*n.Choices)[choice].Correct, nil
	}
	return false, ErrorDoesntKnow
}

// Advance moves to the next dialogue node.
func (d *Dialogue) Advance() {
	d.CurrentID = d.Nodes[d.CurrentID].NextID
}

// GetCurrentNode returns the currently active dialogue node.
func (d *Dialogue) GetCurrentNode() *Node {
	return d.Nodes[d.CurrentID]
}
