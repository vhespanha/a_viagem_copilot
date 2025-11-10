// Package geometry provides 2D geometric primitives and positioning utilities.
package geometry

// Anchor represents a position relative to an anchor point.
// AnchorX and AnchorY define the anchor point (0-1 range), while OffsetX and
// OffsetY provide additional pixel-based positioning.
type Anchor struct {
	AnchorX, AnchorY float32
	OffsetX, OffsetY float32
}

// Position calculates the absolute position for an object with given
// dimensions within a container rectangle.
func (a Anchor) Position(
	objWidth, objHeight, containerWidth, containerHeight float32,
) Vec2 {
	x := (containerWidth-objWidth)*a.AnchorX + a.OffsetX
	y := (containerHeight-objHeight)*a.AnchorY + a.OffsetY
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
func PositionRect(anchor Anchor, width, height float32, screenWidth, screenHeight int) *Rect {
	pos := anchor.Position(width, height, float32(screenWidth), float32(screenHeight))
	return NewRect(pos.X, pos.Y, width, height)
}

// CenterInRect calculates the position to center a vector both horizontally
// and vertically within a rectangle.
func CenterInRect(width, height float32, rect *Rect) Vec2 {
	x := rect.Pos.X + (rect.Size.X-width)/2
	y := rect.Pos.Y + (rect.Size.Y-height)/2
	return Vec2{x, y}
}

// CenterHorizontally calculates the X position to center a vector
// horizontally in a rectangle, keeping the specified Y position.
func CenterHorizontally(width float32, rect *Rect, y float32) Vec2 {
	x := rect.Pos.X + (rect.Size.X-width)/2
	return Vec2{x, y}
}

// CenterVertically calculates the Y position to center a vector
// vertically in a rectangle, keeping the specified X position.
func CenterVertically(height float32, rect *Rect, x float32) Vec2 {
	y := rect.Pos.Y + (rect.Size.Y-height)/2
	return Vec2{x, y}
}
