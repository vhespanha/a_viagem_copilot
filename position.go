package main

// Anchor represents a position relative to an anchor point.
// AnchorX and AnchorY define the anchor point (0-1 range), while OffsetX and
// OffsetY provide additional pixel-based positioning.
type Anchor struct {
	AnchorX, AnchorY float32
	OffsetX, OffsetY float32
}

// Position calculates the absolute position for an object with given
// dimensions within a container rectangle.
func (a Anchor) Position(objWidth, objHeight, containerWidth, containerHeight float32) Vec2 {
	x := (containerWidth-objWidth)*a.AnchorX + a.OffsetX
	y := (containerHeight-objHeight)*a.AnchorY + a.OffsetY
	return Vec2{x, y}
}

// PositionInScreen calculates the absolute screen position for an object with
// given dimensions.
func (a Anchor) PositionInScreen(objWidth, objHeight float32) Vec2 {
	return a.Position(objWidth, objHeight, float32(screenWidth), float32(screenHeight))
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
	pos := anchor.PositionInScreen(width, height)
	return NewRect(pos.X, pos.Y, width, height)
}

// CenterTextInRect calculates the position to center text both horizontally
// and vertically within a rectangle.
func CenterTextInRect(textWidth, textHeight float32, rect *Rect) Vec2 {
	x := rect.Pos.X + (rect.Size.X-textWidth)/2
	y := rect.Pos.Y + (rect.Size.Y-textHeight)/2
	return Vec2{x, y}
}

// CenterTextHorizontally calculates the X position to center text horizontally
// in a rectangle, keeping the specified Y position.
func CenterTextHorizontally(textWidth float32, rect *Rect, y float32) Vec2 {
	x := rect.Pos.X + (rect.Size.X-textWidth)/2
	return Vec2{x, y}
}

// CenterTextVertically calculates the Y position to center text vertically in
// a rectangle, keeping the specified X position.
func CenterTextVertically(textHeight float32, rect *Rect, x float32) Vec2 {
	y := rect.Pos.Y + (rect.Size.Y-textHeight)/2
	return Vec2{x, y}
}
