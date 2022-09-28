package edit

type Transition struct {
	In  TransitionType `json:"in"`
	Out TransitionType `json:"out"`
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
