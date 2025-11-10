// Package dialogue provides the dialogue system for managing conversation
// nodes and choices.
package dialogue

import (
	"errors"
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
}

// New creates and initializes a new dialogue system.
func New() *Dialogue {
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
			NodeIDFirst: {
				Speaker: "character",
				Text:    "bla bla bla",
				Choices: nil,
				NextID:  NodeIDSecond,
			},
			NodeIDSecond: {
				Speaker:  "character",
				Text:     "bla bla bla??",
				Unlocked: true,
				Choices:  &[]Choice{firstChoice, secondChoice},
				NextID:   NodeIDFirst,
			},
		},
		CurrentID: NodeIDFirst,
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
