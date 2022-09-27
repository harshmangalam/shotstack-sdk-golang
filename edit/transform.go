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

func NewTransform() *Transform {
	return new(Transform)
}

func (t *Transform) SetFlip(flip *FlipTransform) *Transform {
	t.Flip = flip
	return t
}

func (t *Transform) SetRotate(rotate *RotateTransform) *Transform {
	t.Rotate = rotate
	return t
}

func (t *Transform) SetSkew(skew *SkewTransform) *Transform {
	t.Skew = skew
	return t
}
