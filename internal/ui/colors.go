// Package ui provides user interface rendering functions and resources.
package ui

import "image/color"

// Common color constants for use throughout the application.
var (
	Red   = color.RGBA{0xff, 0, 0, 0xff}
	Green = color.RGBA{0, 0xff, 0, 0xff}
	Blue  = color.RGBA{0, 0, 0xff, 0xff}
	White = color.RGBA{0xff, 0xff, 0xff, 0xff}
	Black = color.RGBA{0, 0, 0, 0xff}
	Gray  = color.RGBA{0x7f, 0x7f, 0x7f, 0x7f}
)
