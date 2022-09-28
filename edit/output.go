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
