package geometry

// Vec2 represents a 2D vector with float32 components.
// This type reduces the need for casting between int, float32, and float64.
type Vec2 struct {
	X, Y float32
}

// ToInt converts the vector to integer coordinates for pixel-perfect
// positioning.
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
