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
	return &Track{}
}

func (t *Track) SetClips(clips *[]Clip) *Track {
	return &Track{Clips: clips}
}

// clip
func NewClip() *Clip {
	return &Clip{}
}

func (c *Clip) SetAsset(asset *Asset) *Clip {
	return &Clip{Asset: asset}
}

func (c *Clip) SetStart(start int) *Clip {
	return &Clip{Start: start}
}

func (c *Clip) SetLength(length int) *Clip {
	return &Clip{Length: length}
}

func (c *Clip) SetFit(fit ClipFit) *Clip {
	return &Clip{Fit: fit}
}

func (c *Clip) SetScale(scale int) *Clip {
	return &Clip{Scale: scale}
}

func (c *Clip) SetPosition(pos Position) *Clip {
	return &Clip{Position: pos}
}

func (c *Clip) SetOffset(offset *Offset) *Clip {
	return &Clip{Offset: offset}
}

func (c *Clip) SetTransition(transition *Transition) *Clip {
	return &Clip{Transition: transition}
}

func (c *Clip) SetEffect(effect ClipEffect) *Clip {
	return &Clip{Effect: effect}
}

func (c *Clip) SetFilter(filter ClipFilter) *Clip {
	return &Clip{Filter: filter}
}

func (c *Clip) SetOpacity(opacity int8) *Clip {
	return &Clip{Opacity: opacity}
}

func (c *Clip) SetTransform(transform *Transform) *Clip {
	return &Clip{Transform: transform}
}

//  asset

func NewAsset() *Asset {
	return &Asset{}
}

func (a *Asset) SetAssetType(assetType AssetType) *Asset {
	return &Asset{Type: assetType}
}

func (a *Asset) SetVideoAsset(videoAssetType *VideoAssetType) *Asset {
	return &Asset{VideoAsset: videoAssetType}
}

func (a *Asset) SetAudioAsset(audioAsset *AudioAssetType) *Asset {
	return &Asset{AudioAsset: audioAsset}
}

func (a *Asset) SetImageAsset(imageAsser *ImageAssetType) *Asset {
	return &Asset{ImageAsset: imageAsser}
}

func (a *Asset) SetHtmlAsset(htmlAsset *HtmlAssetType) *Asset {
	return &Asset{HtmlAsset: htmlAsset}
}

func (a *Asset) SetTitleAsset(titleAsset *TitleAssetType) *Asset {
	return &Asset{TitleAsset: titleAsset}
}

func (a *Asset) SetLumaAsset(LumaAsset *LumaAssetType) *Asset {
	return &Asset{LumaAsset: LumaAsset}
}

// video asset

func (v *VideoAssetType) SetVideoAssetSrc(src string) *VideoAssetType {
	return &VideoAssetType{Src: src}
}

func (v *VideoAssetType) SetVideoAssetTrim(trim int) *VideoAssetType {
	return &VideoAssetType{Trim: trim}
}

func (v *VideoAssetType) SetVideoAssetVolume(vol float32) *VideoAssetType {
	return &VideoAssetType{Volume: vol}
}

func (v *VideoAssetType) SetVideoAssetCrop(crop *Crop) *VideoAssetType {
	return &VideoAssetType{Crop: crop}
}

// audio asset

func NewAudio() *AudioAssetType {
	return &AudioAssetType{}
}

func (*AudioAssetType) SetAudioSrc(src string) *AudioAssetType {
	return &AudioAssetType{Src: src}
}
func (*AudioAssetType) SetAudioTrim(trim int) *AudioAssetType {
	return &AudioAssetType{Trim: trim}
}

func (*AudioAssetType) SetAudioVolume(vol float32) *AudioAssetType {
	return &AudioAssetType{Volume: vol}
}

func (*AudioAssetType) SetAudioEffect(effect AudioEffect) *AudioAssetType {
	return &AudioAssetType{Effect: effect}
}

// offset

// transition

// transform
