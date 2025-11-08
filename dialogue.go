package main

import (
	"errors"
	"image/color"
)

const (
	dialogueBoxWidth    = screenWidth
	dialogueBoxHeight   = 200
	dialogueBoxOffsetY  = -20
	dialogueBoxChoicesY = 40 // TODO: figure out if this is sane
)

var (
	ErrorDoesntKnow = errors.New("didn't know answer")
)

// Dialogue node IDs
const (
	nodeIDFirst  ID = "first"
	nodeIDSecond ID = "second"
)

// ID is a unique identifier for dialogue nodes.
type ID string

// DialogueNode represents a single node in the dialogue tree.
type DialogueNode struct {
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
	Nodes     map[ID]*DialogueNode
	CurrentID ID
	Box       *Rect
	BoxColor  color.RGBA
}

// NewDialogue creates and initializes a new dialogue system.
func NewDialogue() *Dialogue {
	firstChoice := Choice{
		Text:    "yes",
		Correct: true,
	}
	secondChoice := Choice{
		Text:    "no",
		Correct: false,
	}
	return &Dialogue{
		Nodes: map[ID]*DialogueNode{nodeIDFirst: &DialogueNode{
			Speaker: "character",
			Text:    "bla bla bla",
			Choices: nil,
			NextID:  nodeIDSecond,
		}, nodeIDSecond: &DialogueNode{
			Speaker:  "character",
			Text:     "bla bla bla??",
			Unlocked: true,
			Choices:  &[]Choice{firstChoice, secondChoice},
			NextID:   nodeIDFirst,
		}},
		CurrentID: nodeIDFirst,
		Box: PositionRect(BottomCenter.Offset(0, dialogueBoxOffsetY),
			dialogueBoxWidth, dialogueBoxHeight),
		BoxColor: Black,
	}
}

// Choose processes a dialogue choice and advances to the next node.
// Returns an error if the choice number is invalid.
func (n *DialogueNode) Choose(choice int) (bool, error) {
	if n.Unlocked {
		return (*n.Choices)[choice].Correct, nil
	}
	return false, ErrorDoesntKnow
}

func (d *Dialogue) AdvanceDialogueNode() {
	d.CurrentID = d.Nodes[d.CurrentID].NextID
}

func (d *Dialogue) GetCurrentNode() *DialogueNode {
	return d.Nodes[d.CurrentID]
}
