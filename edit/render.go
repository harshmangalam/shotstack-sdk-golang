package edit

type SoundTrackEffect string
type Disk string
type ClipFilter string

const (
	FadeIn        SoundTrackEffect = "fadeIn"
	FadeOut       SoundTrackEffect = "fadeOut"
	FadeInFadeOut SoundTrackEffect = "fadeInFadeOut"
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

type SoundTrack struct {
	Src    string
	Effect SoundTrackEffect
	Volume int8
}

type Font struct {
	Src string
}

type Asset struct {
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
	Fit        string
	Scale      int
	Position   string
	Offset     *Offset
	Transition *Transition
	Effect     string
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

func (s *SoundTrack) SetEffect(effect SoundTrackEffect) *SoundTrack {
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

func (c *Clip) SetFit(fit string) *Clip {
	return &Clip{Fit: fit}
}

func (c *Clip) SetScale(scale int) *Clip {
	return &Clip{Scale: scale}
}

func (c *Clip) SetPosition(pos string) *Clip {
	return &Clip{Position: pos}
}

func (c *Clip) SetOffset(offset *Offset) *Clip {
	return &Clip{Offset: offset}
}

func (c *Clip) SetTransition(transition *Transition) *Clip {
	return &Clip{Transition: transition}
}

func (c *Clip) SetEffect(effect string) *Clip {
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
