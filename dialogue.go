package main

import "fmt"

const (
	dialogueBoxWidth   = screenWidth
	dialogueBoxHeight  = 200
	dialogueBoxOffsetY = -20
)

type ID string

type DialogueNode struct {
	Speaker          string
	Text             string
	Choice1, Choice2 *Choice
}

type Choice struct {
	Text   string
	NextID ID
}

type DialogueSystem struct {
	Content map[ID]*DialogueNode
	Current ID
	Box     *Rectangle
}

func NewDialogueSystem() *DialogueSystem {
	rx, ry := BottomCenter.Offset(0, -20).Position(float32(dialogueBoxWidth), float32(dialogueBoxHeight))
	return &DialogueSystem{
		Content: map[ID]*DialogueNode{"first": &DialogueNode{
			Speaker: "Vinicius",
			Text:    "Feliz aniversario!",
			Choice1: &Choice{
				Text:   "",
				NextID: "second",
			},
			Choice2: nil,
		}, "second": &DialogueNode{
			Speaker: "Clara",
			Text:    "Voce lembrou!",
			Choice1: nil,
			Choice2: nil,
		}},
		Current: "first",
		Box: &Rectangle{ // dialogue box
			w:   dialogueBoxWidth,
			h:   dialogueBoxHeight,
			x:   int(rx),
			y:   int(ry),
			col: BLACK,
		},
	}
}

func (ds *DialogueSystem) Choose(choice int) error {
	if choice == 1 {
		currentNode := ds.Content[ds.Current]
		if currentNode.Choice1 == nil {
			ds.Current = "first" // for now we loop around
			return nil
		}
		ds.Current = currentNode.Choice1.NextID
		return nil
	} else if choice == 2 {
		currentNode := ds.Content[ds.Current]
		ds.Current = currentNode.Choice2.NextID
		return nil
	}
	return fmt.Errorf("must be 1 or 2")
}
