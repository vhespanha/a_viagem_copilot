package main

type Anchor struct {
	Ax, Ay float32
}

func (a Anchor) Position(objW, objH float32) (float32, float32) {
	x := (ScreenWidth - objW) * a.Ax
	y := (ScreenHeight - objH) * a.Ay
	return x, y
}

var (
	TopLeft      = Anchor{0, 0}
	TopCenter    = Anchor{0.5, 0}
	TopRight     = Anchor{1, 0}
	CenterLeft   = Anchor{0, 0.5}
	Center       = Anchor{0.5, 0.5}
	CenterRight  = Anchor{1, 0.5}
	BottomLeft   = Anchor{0, 1}
	BottomCenter = Anchor{0.5, 1}
	BottomRight  = Anchor{1, 1}
)
