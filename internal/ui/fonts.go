package ui

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	FontSizeNormal = 24
	FontSizeBig    = 32
	LineSpacing    = 1.5
)

// Fonts holds the font resources used for rendering text.
type Fonts struct {
	Source *text.GoTextFaceSource
	Normal *text.GoTextFace
	Big    *text.GoTextFace
}

// NewFonts creates and initializes font resources.
func NewFonts() *Fonts {
	faceSource, err := text.NewGoTextFaceSource(
		bytes.NewReader(fonts.MPlus1pRegular_ttf),
	)
	if err != nil {
		log.Fatal(err)
	}

	return &Fonts{
		Source: faceSource,
		Normal: createFace(faceSource, FontSizeNormal),
		Big:    createFace(faceSource, FontSizeBig),
	}
}

func createFace(source *text.GoTextFaceSource, size float64) *text.GoTextFace {
	return &text.GoTextFace{
		Source: source,
		Size:   size,
	}
}
