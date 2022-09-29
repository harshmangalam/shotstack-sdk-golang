package edit

type ClipFilter string
type ClipEffect string
type ClipFit string
type Position string

const (
	Top         Position = "top"
	TopRight    Position = "topRight"
	Right       Position = "right"
	BottomRight Position = "bottomRight"
	Bottom      Position = "bottom"
	BottomLeft  Position = "bottomLeft"
	Left        Position = "left"
	TopLeft     Position = "topLeft"
	Center      Position = "center"
)

const (
	Boost     ClipFilter = "boost"
	Contrast  ClipFilter = "contrast"
	Darken    ClipFilter = "darken"
	Greyscale ClipFilter = "greyscale"
	Lighten   ClipFilter = "lighten"
	Muted     ClipFilter = "muted"
	Negative  ClipFilter = "negative"
)

const (
	ZoomIn     ClipEffect = "zoomIn"
	ZoomOut    ClipEffect = "zoomOut"
	SlideLeft  ClipEffect = "slideLeft"
	SlideRight ClipEffect = "slideRight"
	SlideUp    ClipEffect = "slideUp"
	SlideDown  ClipEffect = "slideDown"
)

const (
	FitCover   ClipFit = "cover"
	FitContain ClipFit = "contain"
	FitCrop    ClipFit = "crop"
	FitNone    ClipFit = "none"
)

type Clip struct {
	Asset      any         `json:"asset"`
	Start      int         `json:"start"`
	Length     int         `json:"length"`
	Fit        ClipFit     `json:"fit,omitempty"`
	Scale      int         `json:"scale,omitempty"`
	Position   Position    `json:"position,omitempty"`
	Offset     *Offset     `json:"offset,omitempty"`
	Transition *Transition `json:"transition,omitempty"`
	Effect     ClipEffect  `json:"effect,omitempty"`
	Filter     ClipFilter  `json:"filter,omitempty"`
	Opacity    int8        `json:"opacity,omitempty"`
	Transform  *Transform  `json:"transform,omitempty"`
}

func NewClip() *Clip {
	c := new(Clip)
	return c
}

func (c *Clip) SetAsset(asset any) *Clip {
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
