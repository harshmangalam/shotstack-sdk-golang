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

type OutputSize struct {
	Width  int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
}
type Poster struct {
	Capture float32 `json:"capture"`
}
type Range struct {
	Start  float32 `json:"start,omitempty"`
	Length float32 `json:"length,omitempty"`
}
type Thumbnail struct {
	Capture float32 `json:"capture"`
	Scale   float32 `json:"scale"`
}

type Destination struct {
	Provider string `json:"provider"`
	Exclude  bool   `json:"exclude,omitempty"`
}

type Output struct {
	Format       MediaFileType    `json:"format"`
	Resolution   MediaResolution  `json:"resolution,omitempty"`
	AspectRatio  MediaAspectRatio `json:"aspectRatio,omitempty"`
	Size         *OutputSize      `json:"size,omitempty"`
	Fps          float32          `json:"fps,omitempty"`
	ScaleTo      MediaResolution  `json:"scaleTo,omitempty"`
	Quality      MediaQuality     `json:"quality,omitempty"`
	Repeat       bool             `json:"repeat,omitempty"`
	Range        *Range           `json:"range,omitempty"`
	Poster       *Poster          `json:"poster,omitempty"`
	Thumbnail    *Thumbnail       `json:"thumbnail,omitempty"`
	Destinations *[]Destination   `json:"destinations,omitempty"`
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

func (o *Output) SetSize(size *OutputSize) *Output {
	o.Size = size
	return o
}

func (o *Output) SetFps(fps float32) *Output {
	o.Fps = fps
	return o
}

func (o *Output) SetScaleTo(scaleTo MediaResolution) *Output {
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

func NewSize() *OutputSize {
	return new(OutputSize)
}

func (s *OutputSize) SetWidth(width int) *OutputSize {
	s.Width = width
	return s
}

func (s *OutputSize) SetHeight(height int) *OutputSize {
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

func (p *Poster) SetCapture(capture float32) *Poster {
	p.Capture = capture
	return p
}

// thumbnail

func NewThumbnail() *Thumbnail {
	return new(Thumbnail)
}
func (t *Thumbnail) SetCapture(capture float32) *Thumbnail {
	t.Capture = capture
	return t
}

func (t *Thumbnail) SetScale(scale float32) *Thumbnail {
	t.Scale = scale
	return t
}

func NewDestination() *Destination {
	return new(Destination)
}

func (d *Destination) SetProvider(provider string) *Destination {
	d.Provider = provider
	return d
}

func (d *Destination) SetExclude(exclude bool) *Destination {
	d.Exclude = exclude
	return d
}
