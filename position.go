package main

type Anchor struct {
	AnchorX, AnchorY float32
	OffsetX, OffsetY float32
}

func (a Anchor) Position(objW, objH float32) (float32, float32) {
	x := (ScreenWidth-objW)*a.AnchorX + a.OffsetX
	y := (ScreenHeight-objH)*a.AnchorY + a.OffsetY
	return x, y
}

func (a Anchor) Offset(x, y float32) Anchor {
	return Anchor{a.AnchorX, a.AnchorY, a.OffsetX + x, a.OffsetY + y}
}

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