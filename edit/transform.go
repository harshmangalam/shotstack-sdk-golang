package edit

type RotateTransform struct {
	angle int
}

type SkewTransform struct {
	X float32
	Y float32
}

type FlipTransform struct {
	Horizontal bool
	Vertical   bool
}
type Transform struct {
	Rotate *RotateTransform
	Skew   *SkewTransform
	Flip   *FlipTransform
}
