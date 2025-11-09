package geometry

// Rect represents a rectangle with position and size using Vec2.
type Rect struct {
	Pos  Vec2
	Size Vec2
}

// NewRect returns a pointer to a new rectangle from position and size.
func NewRect(x, y, width, height float32) *Rect {
	return &Rect{
		Pos:  Vec2{x, y},
		Size: Vec2{width, height},
	}
}

// Contains checks if the given integer coordinates are within the rectangle's
// bounds.
func (r *Rect) Contains(x, y int) bool {
	fx, fy := float32(x), float32(y)
	return fx >= r.Pos.X && fx < r.Pos.X+r.Size.X &&
		fy >= r.Pos.Y && fy < r.Pos.Y+r.Size.Y
}

// ContainsFloat checks if the given float coordinates are within the
// rectangle's bounds.
func (r *Rect) ContainsFloat(x, y float32) bool {
	return x >= r.Pos.X && x < r.Pos.X+r.Size.X &&
		y >= r.Pos.Y && y < r.Pos.Y+r.Size.Y
}
