package edit

type MediaFileType string
type MediaResolution string
type MediaAspectRatio string
type MediaQuality string

const (
	Mp4 MediaFileType = "mp4"
	Gif MediaFileType = "gif"
	Jpg MediaFileType = "jpg"
	Png MediaFileType = "png"
	Bmp MediaFileType = "bmp"
	Mp3 MediaFileType = "mp3"
)

const (
	ResolutionPreview MediaResolution = "preview"
	ResolutionMobile  MediaResolution = "mobile"
	ResolutionSd      MediaResolution = "sd"
	ResolutionHd      MediaResolution = "hd"
	Resolution1080    MediaResolution = "1080"
)

const (
	DefaultRatio    MediaAspectRatio = "16:9"
	RevDefaultRatio MediaAspectRatio = "9:16"
	SquareRatio     MediaAspectRatio = "1:1"
	ShortRatio      MediaAspectRatio = "4:5"
	LegacyRatio     MediaAspectRatio = "4:3"
)

const (
	Low    MediaQuality = "low"
	Medium MediaQuality = "medium"
	High   MediaQuality = "high"
)

type Size struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}
type Poster struct {
	Capture int `json:"capture"`
}
type Range struct {
	Start  float32 `json:"start"`
	Length float32 `json:"length"`
}
type Thumbnail struct {
	Capture int `json:"capture"`
	Scale   int `json:"scale"`
}

// default shotstack complete later
type Destination struct{}

type Output struct {
	Format       MediaFileType    `json:"format"`
	Resolution   MediaResolution  `json:"resolution"`
	AspectRatio  MediaAspectRatio `json:"aspectRatio"`
	Size         *Size            `json:"size"`
	Fps          float32          `json:"fps"`
	ScaleTo      string           `json:"scaleTo"`
	Quality      MediaQuality     `json:"quality"`
	Repeat       bool             `json:"repeat"`
	Range        *Range           `json:"range"`
	Poster       *Poster          `json:"poster"`
	Thumbnail    *Thumbnail       `json:"thumbnail"`
	Destinations *[]Destination   `json:"destinations"`
}

func NewOutput() *Output {
	return new(Output)
}

func (o *Output) SetFormat(format MediaFileType) *Output {
	o.Format = format
	return o
}

func (o *Output) SetResolution(resol MediaResolution) *Output {
	o.Resolution = resol
	return o
}

func (o *Output) SetAspectRatio(ratio MediaAspectRatio) *Output {
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

func (o *Output) SetQuality(quality MediaQuality) *Output {
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

// size

func NewSize() *Size {
	return new(Size)
}

func (s *Size) SetWidth(width int) *Size {
	s.Width = width
	return s
}

func (s *Size) SetHeight(height int) *Size {
	s.Height = height
	return s
}

// range

func NewRange() *Range {
	return new(Range)
}

func (r *Range) SetStart(start float32) *Range {
	r.Start = start
	return r
}

func (r *Range) SetLength(length float32) *Range {
	r.Length = length
	return r
}

func NewPoster() *Poster {
	return new(Poster)
}

func (p *Poster) SetCapture(capture int) *Poster {
	p.Capture = capture
	return p
}

// thumbnail

func NewThumbnail() *Thumbnail {
	return new(Thumbnail)
}
func (t *Thumbnail) SetCapture(capture int) *Thumbnail {
	t.Capture = capture
	return t
}

func (t *Thumbnail) SetScale(scale int) *Thumbnail {
	t.Scale = scale
	return t
}
