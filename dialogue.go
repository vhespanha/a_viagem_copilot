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
	Speaker          string
	Text             string
	Choice1, Choice2 *Choice
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
	return &Dialogue{
		Content: map[ID]*DialogueNode{nodeIDFirst: &DialogueNode{
			Speaker: "Vinicius",
			Text:    "Feliz aniversário!",
			Choice1: &Choice{
				Text:   "",
				NextID: nodeIDSecond,
			},
			Choice2: nil,
		}, nodeIDSecond: &DialogueNode{
			Speaker: "Clara",
			Text:    "Você lembrou!",
			Choice1: nil,
			Choice2: nil,
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
		if currentNode.Choice1 == nil {
			d.Current = nodeIDFirst // for now we loop around
			return nil
		}
		d.Current = currentNode.Choice1.NextID
		return nil
	case 2:
		if currentNode.Choice2 == nil {
			return fmt.Errorf("choice 2 is not available")
		}
		d.Current = currentNode.Choice2.NextID
		return nil
	default:
		return fmt.Errorf("must be 1 or 2")
	}
}
