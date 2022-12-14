package edit

type Offset struct {
	X float32 `json:"x,omitempty"`
	Y float32 `json:"y,omitempty"`
}

func NewOffset() *Offset {
	return new(Offset)
}

func (o *Offset) SetX(x float32) *Offset {
	o.X = x
	return o
}

func (o *Offset) SetY(y float32) *Offset {
	o.Y = y
	return o
}
