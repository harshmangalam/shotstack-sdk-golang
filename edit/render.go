package edit

type AudioEffect string
type Disk string
type ClipFilter string
type ClipEffect string
type Position string
type ClipFit string
type TitleAssetStyle string

type AssetType string

const (
	FadeIn        AudioEffect = "fadeIn"
	FadeOut       AudioEffect = "fadeOut"
	FadeInFadeOut AudioEffect = "fadeInFadeOut"
)

const (
	Local Disk = "local"
	Mount Disk = "mount"
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
	FitCover   ClipFit = "cover"
	FitContain ClipFit = "contain"
	FitCrop    ClipFit = "crop"
	FitNone    ClipFit = "none"
)

const (
	VideoAsset AssetType = "video"
	ImageAsset AssetType = "image"
	TitleAsset AssetType = "title"
	HtmlAsset  AssetType = "html"
	AudioAsset AssetType = "audio"
	LumaAsset  AssetType = "luma"
)

const (
	Minimal     TitleAssetStyle = "minimal"
	Blockbuster TitleAssetStyle = "blockbuster"
	Vogue       TitleAssetStyle = "vogue"
	Sketchy     TitleAssetStyle = "sketchy"
	Skinny      TitleAssetStyle = "skinny"
	Chunk       TitleAssetStyle = "chunk"
	ChunkLight  TitleAssetStyle = "chunkLight"
	Marker      TitleAssetStyle = "marker"
	Future      TitleAssetStyle = "future"
	Subtitle    TitleAssetStyle = "subtitle"
)

type SoundTrack struct {
	Src    string
	Effect AudioEffect
	Volume int8
}

type Font struct {
	Src string
}

type Crop struct {
	Top    float32
	Bottom float32
	Left   float32
	Right  float32
}
type VideoAssetType struct {
	Src    string
	Trim   int
	Volume float32
	Crop   *Crop
}

type ImageAssetType struct {
	Src  string
	Crop *Crop
}

type TitleAssetType struct {
	Text  string
	Style TitleAssetStyle
}

type HtmlAssetType struct {
	Html       string
	Css        string
	Width      int
	Height     int
	Background string
	Position   Position
}

type AudioAssetType struct {
	Src    string
	Trim   int
	Volume float32
	Effect AudioEffect
}

type LumaAssetType struct {
	Src  string
	Trim int
}
type Asset struct {
	Type       AssetType
	VideoAsset *VideoAssetType
	ImageAsset *ImageAssetType
	TitleAsset *TitleAssetType
	HtmlAsset  *HtmlAssetType
	AudioAsset *AudioAssetType
	LumaAsset  *LumaAssetType
}

type Offset struct {
}

type Transform struct {
}

type Transition struct {
}

type Clip struct {
	Asset      *Asset
	Start      int
	Length     int
	Fit        ClipFit
	Scale      int
	Position   Position
	Offset     *Offset
	Transition *Transition
	Effect     ClipEffect
	Filter     ClipFilter
	Opacity    int8
	Transform  *Transform
}
type Track struct {
	Clips *[]Clip
}

type Timeline struct {
	SoundTrack *SoundTrack
	Background string
	Fonts      *[]Font
	Tracks     *[]Track
	Cache      bool
}

type Output struct {
}

type Merge struct {
}

type Edit struct {
	Timeline *Timeline
	Output   *Output
	Merges   *[]Merge
	Callback string
	Disk     Disk // "local" | "mount"
}

func NewOutput() *Output {
	return &Output{}
}

func NewMerges() *[]Merge {
	return &[]Merge{}
}

// edit
func NewEdit() *Edit {
	return &Edit{}
}

func (e *Edit) SetTimeline(timeline *Timeline) *Edit {
	return &Edit{Timeline: timeline}
}

func (e *Edit) SetOutput(output *Output) *Edit {
	return &Edit{Output: output}
}

func (e *Edit) SetMerges(merges *[]Merge) *Edit {
	return &Edit{Merges: merges}
}

func (e *Edit) SetCallback(callback string) *Edit {
	return &Edit{Callback: callback}
}

func (e *Edit) SetDisk(disk Disk) *Edit {
	return &Edit{Disk: disk}
}

// timeline

func NewTimeline() *Timeline {
	return &Timeline{}
}

func (t *Timeline) SetSoundTrack(soundTrack *SoundTrack) *Timeline {
	return &Timeline{SoundTrack: soundTrack}
}

func (t *Timeline) SetBackground(background string) *Timeline {
	return &Timeline{Background: background}
}

func (t *Timeline) SetFonts(fonts *[]Font) *Timeline {
	return &Timeline{Fonts: fonts}
}

func (t *Timeline) SetTracks(tracks *[]Track) *Timeline {
	return &Timeline{Tracks: tracks}
}

func (t *Timeline) SetCache(cache bool) *Timeline {
	return &Timeline{Cache: cache}
}

//  sound track

func NewSoundTrack() *SoundTrack {
	return &SoundTrack{}
}

func (s *SoundTrack) SetSource(src string) *SoundTrack {
	return &SoundTrack{Src: src}
}

func (s *SoundTrack) SetEffect(effect AudioEffect) *SoundTrack {
	return &SoundTrack{Effect: effect}
}

func (s *SoundTrack) SetVolume(volume int8) *SoundTrack {
	return &SoundTrack{Volume: volume}
}

// font

func NewFont() *Font {
	return &Font{}
}

func (f *Font) SetSrc(src string) *Font {
	return &Font{Src: src}
}

// track

func NewTrack() *Track {
	t := new(Track)
	return t
}

func (t *Track) SetClips(clips *[]Clip) *Track {
	t.Clips = clips
	return t
}

// clip
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

//  asset

func NewAsset() *Asset {
	a := new(Asset)
	return a
}

func (a *Asset) SetAssetType(assetType AssetType) *Asset {
	a.Type = assetType
	return a
}

func (a *Asset) SetVideoAsset(videoAsset *VideoAssetType) *Asset {
	a.VideoAsset = videoAsset
	return a
}

func (a *Asset) SetAudioAsset(audioAsset *AudioAssetType) *Asset {
	a.AudioAsset = audioAsset
	return a
}

func (a *Asset) SetImageAsset(imageAsset *ImageAssetType) *Asset {
	a.ImageAsset = imageAsset
	return a
}

func (a *Asset) SetHtmlAsset(htmlAsset *HtmlAssetType) *Asset {
	a.HtmlAsset = htmlAsset
	return a
}

func (a *Asset) SetTitleAsset(titleAsset *TitleAssetType) *Asset {
	a.TitleAsset = titleAsset
	return a
}

func (a *Asset) SetLumaAsset(lumaAsset *LumaAssetType) *Asset {
	a.LumaAsset = lumaAsset
	return a
}

// video asset

func NewVideoAsset() *VideoAssetType {
	v := new(VideoAssetType)
	return v
}
func (v *VideoAssetType) SetVideoAssetSrc(src string) *VideoAssetType {
	v.Src = src
	return v
}

func (v *VideoAssetType) SetVideoAssetTrim(trim int) *VideoAssetType {
	v.Trim = trim
	return v
}

func (v *VideoAssetType) SetVideoAssetVolume(vol float32) *VideoAssetType {
	v.Volume = vol
	return v
}

func (v *VideoAssetType) SetVideoAssetCrop(crop *Crop) *VideoAssetType {
	v.Crop = crop
	return v
}

// audio asset

func NewAudioAsset() *AudioAssetType {
	a := new(AudioAssetType)
	return a
}

func (a *AudioAssetType) SetAudioAssetSrc(src string) *AudioAssetType {
	a.Src = src
	return a
}
func (a *AudioAssetType) SetAudioAssetTrim(trim int) *AudioAssetType {
	a.Trim = trim
	return a
}

func (a *AudioAssetType) SetAudioAssetVolume(vol float32) *AudioAssetType {
	a.Volume = vol
	return a
}

func (a *AudioAssetType) SetAudioAssetEffect(effect AudioEffect) *AudioAssetType {
	a.Effect = effect
	return a
}

// image asset

func NewImageAsset() *ImageAssetType {
	i := new(ImageAssetType)
	return i
}

func (i *ImageAssetType) SetImageAssetSrc(src string) *ImageAssetType {
	i.Src = src
	return i
}
func (i *ImageAssetType) SetImageAssetCrop(crop *Crop) *ImageAssetType {
	i.Crop = crop
	return i
}

// html asset

func NewHtmlAsset() *HtmlAssetType {
	h := new(HtmlAssetType)
	return h
}

func (h *HtmlAssetType) SetHtmlAssetHtml(html string) *HtmlAssetType {
	h.Html = html
	return h
}

func (h *HtmlAssetType) SetHtmlAssetCss(css string) *HtmlAssetType {
	h.Css = css
	return h
}

func (h *HtmlAssetType) SetHtmlAssetHeight(height int) *HtmlAssetType {
	h.Height = height
	return h
}

func (h *HtmlAssetType) SetHtmlAssetWidth(width int) *HtmlAssetType {
	h.Width = width
	return h
}

func (h *HtmlAssetType) SetHtmlAssetPosition(pos Position) *HtmlAssetType {
	h.Position = pos
	return h
}

// offset

// transition

// transform
