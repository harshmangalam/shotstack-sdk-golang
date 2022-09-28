package edit

type Size struct{}
type Poster struct{}
type Range struct{}
type Thumbnail struct{}
type Destination struct{}

type Output struct {
	Format       string
	Resolution   string
	AspectRatio  string
	Size         *Size
	Fps          float32
	ScaleTo      string
	Quality      string
	Repeat       bool
	Range        *Range
	Poster       *Poster
	Thumbnail    *Thumbnail
	Destinations *[]Destination
}

func NewOutput() *Output {
	return new(Output)
}

func (o *Output) SetFormat(format string) *Output {
	o.Format = format
	return o
}

func (o *Output) SetResolution(resol string) *Output {
	o.Resolution = resol
	return o
}

func (o *Output) SetAspectRatio(ratio string) *Output {
	o.AspectRatio = ratio
	return o
}

func (o *Output) SetSize(size *Size) *Output {
	o.Size = size
	return o
}

func (o *Output) SetFps(fps float32) *Output {
	o.Fps = fps
	return o
}

func (o *Output) SetScaleTo(scaleTo string) *Output {
	o.ScaleTo = scaleTo
	return o
}

func (o *Output) SetQuality(quality string) *Output {
	o.Quality = quality
	return o
}

func (o *Output) SetRepeat(repeat bool) *Output {
	o.Repeat = repeat
	return o
}

func (o *Output) SetRange(rang *Range) *Output {
	o.Range = rang
	return o
}

func (o *Output) SetPoster(poster *Poster) *Output {
	o.Poster = poster
	return o
}

func (o *Output) SetThumbnail(thumb *Thumbnail) *Output {
	o.Thumbnail = thumb
	return o
}

func (o *Output) SetDestinations(dest *[]Destination) *Output {
	o.Destinations = dest
	return o
}