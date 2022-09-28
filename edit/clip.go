package edit

type Clip struct {
	Asset      *Asset      `json:"asset"`
	Start      int         `json:"start"`
	Length     int         `json:"length"`
	Fit        ClipFit     `json:"fit"`
	Scale      int         `json:"scale"`
	Position   Position    `json:"position"`
	Offset     *Offset     `json:"offset"`
	Transition *Transition `json:"transition"`
	Effect     ClipEffect  `json:"effect"`
	Filter     ClipFilter  `json:"filter"`
	Opacity    int8        `json:"opacity"`
	Transform  *Transform  `json:"transform"`
}

func NewClip() *Clip {
	c := new(Clip)
	return c
}

func (c *Clip) SetAsset(asset *Asset) *Clip {
	c.Asset = asset
	return c
}

func (c *Clip) SetStart(start int) *Clip {
	c.Start = start
	return c
}

func (c *Clip) SetLength(length int) *Clip {
	c.Length = length
	return c
}

func (c *Clip) SetFit(fit ClipFit) *Clip {
	c.Fit = fit
	return c
}

func (c *Clip) SetScale(scale int) *Clip {
	c.Scale = scale
	return c
}

func (c *Clip) SetPosition(pos Position) *Clip {
	c.Position = pos
	return c
}

func (c *Clip) SetOffset(offset *Offset) *Clip {
	c.Offset = offset
	return c
}

func (c *Clip) SetTransition(transition *Transition) *Clip {
	c.Transition = transition
	return c
}

func (c *Clip) SetEffect(effect ClipEffect) *Clip {
	c.Effect = effect
	return c
}

func (c *Clip) SetFilter(filter ClipFilter) *Clip {
	c.Filter = filter
	return c
}

func (c *Clip) SetOpacity(opacity int8) *Clip {
	c.Opacity = opacity
	return c
}

func (c *Clip) SetTransform(transform *Transform) *Clip {
	c.Transform = transform
	return c
}
