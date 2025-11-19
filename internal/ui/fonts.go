package ui

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	fontSizeNormal = 24
	fontSizeBig    = 32
	lineSpacing    = 1.5
)

// faces holds the font resources used for rendering text.
type faces struct {
	source *text.GoTextFaceSource
	normal *text.GoTextFace
	big    *text.GoTextFace
}

// newFaces creates and initializes font resources.
func newFaces() *faces {
	faceSource, err := text.NewGoTextFaceSource(
		bytes.NewReader(fonts.MPlus1pRegular_ttf),
	)
	if err != nil {
		log.Fatal(err)
	}

	return &faces{
		source: faceSource,
		normal: createFace(faceSource, fontSizeNormal),
		big:    createFace(faceSource, fontSizeBig),
	}
}

func createFace(source *text.GoTextFaceSource, size float64) *text.GoTextFace {
	return &text.GoTextFace{
		Source: source,
		Size:   size,
	}
}
