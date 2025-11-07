package main

import (
	"fmt"
	"image/color"
)

const (
	dialogueBoxWidth   = screenWidth
	dialogueBoxHeight  = 200
	dialogueBoxOffsetY = -20
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

// DialogueSystem manages the dialogue state and progression.
type DialogueSystem struct {
	Content  map[ID]*DialogueNode
	Current  ID
	Box      *Rect
	BoxColor color.RGBA
}

// NewDialogueSystem creates and initializes a new dialogue system.
func NewDialogueSystem() *DialogueSystem {
	rect := PositionRect(BottomCenter.Offset(0, dialogueBoxOffsetY),
		dialogueBoxWidth, dialogueBoxHeight)
	return &DialogueSystem{
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
		Current:  nodeIDFirst,
		Box:      &rect,
		BoxColor: Black,
	}
}

// Choose processes a dialogue choice and advances to the next node.
// Returns an error if the choice number is invalid.
func (ds *DialogueSystem) Choose(choice int) error {
	currentNode := ds.Content[ds.Current]

	switch choice {
	case 1:
		if currentNode.Choice1 == nil {
			ds.Current = nodeIDFirst // for now we loop around
			return nil
		}
		ds.Current = currentNode.Choice1.NextID
		return nil
	case 2:
		if currentNode.Choice2 == nil {
			return fmt.Errorf("choice 2 is not available")
		}
		ds.Current = currentNode.Choice2.NextID
		return nil
	default:
		return fmt.Errorf("must be 1 or 2")
	}
}
