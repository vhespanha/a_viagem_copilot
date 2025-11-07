package main

// Vec2 represents a 2D vector with float32 components.
// This type reduces the need for casting between int, float32, and float64.
type Vec2 struct {
	X, Y float32
}

// ToInt converts the vector to integer coordinates for pixel-perfect positioning.
func (v Vec2) ToInt() (int, int) {
	return int(v.X), int(v.Y)
}

// Add returns the sum of two vectors.
func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{v.X + other.X, v.Y + other.Y}
}

// Scale returns the vector scaled by a factor.
func (v Vec2) Scale(factor float32) Vec2 {
	return Vec2{v.X * factor, v.Y * factor}
}

// Rect represents a rectangle with position and size using Vec2.
type Rect struct {
	Pos  Vec2
	Size Vec2
}

// NewRect creates a new rectangle from position and size.
func NewRect(x, y, width, height float32) Rect {
	return Rect{
		Pos:  Vec2{x, y},
		Size: Vec2{width, height},
	}
}

// Contains checks if the given integer coordinates are within the rectangle's bounds.
func (r Rect) Contains(x, y int) bool {
	fx, fy := float32(x), float32(y)
	return fx >= r.Pos.X && fx < r.Pos.X+r.Size.X &&
		fy >= r.Pos.Y && fy < r.Pos.Y+r.Size.Y
}

// ContainsFloat checks if the given float coordinates are within the rectangle's bounds.
func (r Rect) ContainsFloat(x, y float32) bool {
	return x >= r.Pos.X && x < r.Pos.X+r.Size.X &&
		y >= r.Pos.Y && y < r.Pos.Y+r.Size.Y
}