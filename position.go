package main

// Anchor represents a position on the screen relative to an anchor point.
// AnchorX and AnchorY define the anchor point (0-1 range), while OffsetX and OffsetY
// provide additional pixel-based positioning.
type Anchor struct {
	AnchorX, AnchorY float32
	OffsetX, OffsetY float32
}

// Position calculates the absolute screen position for an object with given dimensions.
func (a Anchor) Position(objW, objH float32) Vec2 {
	x := (float32(screenWidth)-objW)*a.AnchorX + a.OffsetX
	y := (float32(screenHeight)-objH)*a.AnchorY + a.OffsetY
	return Vec2{x, y}
}

// Offset returns a new Anchor with added offset values.
func (a Anchor) Offset(x, y float32) Anchor {
	return Anchor{a.AnchorX, a.AnchorY, a.OffsetX + x, a.OffsetY + y}
}

// Predefined anchor positions for common screen locations.
var (
	TopLeft      = Anchor{0, 0, 0, 0}
	TopCenter    = Anchor{0.5, 0, 0, 0}
	TopRight     = Anchor{1, 0, 0, 0}
	CenterLeft   = Anchor{0, 0.5, 0, 0}
	Center       = Anchor{0.5, 0.5, 0, 0}
	CenterRight  = Anchor{1, 0.5, 0, 0}
	BottomLeft   = Anchor{0, 1, 0, 0}
	BottomCenter = Anchor{0.5, 1, 0, 0}
	BottomRight  = Anchor{1, 1, 0, 0}
)

// PositionRect positions a rectangle using an anchor point.
func PositionRect(anchor Anchor, width, height float32) *Rect {
	pos := anchor.Position(width, height)
	return NewRect(pos.X, pos.Y, width, height)
}
