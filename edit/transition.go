package edit

type Transition struct {
	In  TransitionType `json:"in,omitempty"`
	Out TransitionType `json:"out,omitempty"`
}

func NewTransition() *Transition {
	return new(Transition)
}

func (t *Transition) SetIn(in TransitionType) *Transition {
	t.In = in
	return t
}

func (t *Transition) SetOut(out TransitionType) *Transition {
	t.Out = out
	return t
}
