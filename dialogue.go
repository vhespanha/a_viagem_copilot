package main

import (
	"fmt"
	"image/color"
)

const (
	dialogueBoxWidth    = screenWidth
	dialogueBoxHeight   = 200
	dialogueBoxOffsetY  = -20
	dialogueBoxChoicesY = 40 // TODO: figure out if this is sane
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
	Speaker string
	Text    string
	Choices []*Choice
}

// Choice represents a player's dialogue choice leading to another node.
type Choice struct {
	Text   string
	NextID ID
}

// Dialogue manages the dialogue state and progression.
type Dialogue struct {
	Content  map[ID]*DialogueNode
	Current  ID
	Box      *Rect
	BoxColor color.RGBA
}

// NewDialogue creates and initializes a new dialogue system.
func NewDialogue() *Dialogue {
	firstChoice := &Choice{
		Text:   "Claro que sim!",
		NextID: nodeIDSecond,
	}
	secondChoice := &Choice{
		Text:   "Claro que nao!",
		NextID: nodeIDSecond,
	}
	return &Dialogue{
		Content: map[ID]*DialogueNode{nodeIDFirst: &DialogueNode{
			Speaker: "Vinicius",
			Text:    "Feliz aniversário!",
			Choices: []*Choice{firstChoice, secondChoice},
		}, nodeIDSecond: &DialogueNode{
			Speaker: "Clara",
			Text:    "Você lembrou!",
			Choices: []*Choice{},
		}},
		Current: nodeIDFirst,
		Box: PositionRect(BottomCenter.Offset(0, dialogueBoxOffsetY),
			dialogueBoxWidth, dialogueBoxHeight),
		BoxColor: Black,
	}
}

// Choose processes a dialogue choice and advances to the next node.
// Returns an error if the choice number is invalid.
func (d *Dialogue) Choose(choice int) error {
	currentNode := d.Content[d.Current]

	switch choice {
	case 1:
		if len(currentNode.Choices) == 0 {
			d.Current = nodeIDFirst // for now we loop around
			return nil
		}
		d.Current = currentNode.Choices[0].NextID
		return nil
	case 2:
		if len(currentNode.Choices) == 1 {
			return fmt.Errorf("choice 2 is not available")
		}
		d.Current = currentNode.Choices[1].NextID
		return nil
	default:
		return fmt.Errorf("must be 1 or 2")
	}
}
