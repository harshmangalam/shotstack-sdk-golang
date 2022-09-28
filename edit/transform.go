package edit

type RotateTransform struct {
	Angle int
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

// rotate

func NewRotateTransform() *RotateTransform {
	return new(RotateTransform)
}

func (r *RotateTransform) SetAngle(angle int) *RotateTransform {
	r.Angle = angle
	return r
}

// skew

func NewSkewTransform() *SkewTransform {
	return new(SkewTransform)
}

func (s *SkewTransform) SetX(x float32) *SkewTransform {
	s.X = x
	return s
}

func (s *SkewTransform) SetY(y float32) *SkewTransform {
	s.Y = y
	return s
}

// flip

func NewFlipTransform() *FlipTransform {
	return new(FlipTransform)
}

func (f *FlipTransform) SetHorizontal(h bool) *FlipTransform {
	f.Horizontal = h
	return f
}

func (f *FlipTransform) SetVertical(v bool) *FlipTransform {
	f.Vertical = v
	return f
}
