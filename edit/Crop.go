package edit

type Crop struct {
	Top    float32 `json:"top"`
	Bottom float32 `json:"bottom"`
	Left   float32 `json:"left"`
	Right  float32 `jight:"right"`
}

func NewCrop() *Crop {
	return new(Crop)
}

func (c *Crop) SetTop(top float32) *Crop {
	c.Top = top
	return c
}

func (c *Crop) SetBottom(bottom float32) *Crop {
	c.Bottom = bottom
	return c
}
func (c *Crop) SetLeft(left float32) *Crop {
	c.Left = left
	return c
}
func (c *Crop) SetRight(right float32) *Crop {
	c.Right = right
	return c
}
