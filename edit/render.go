package edit

type SoundTrackEffect string

const (
	FadeIn        SoundTrackEffect = "fadeIn"
	FadeOut       SoundTrackEffect = "fadeOut"
	FadeInFadeOut SoundTrackEffect = "fadeInFadeOut"
)

type SoundTrack struct {
	Src    string
	Effect SoundTrackEffect
	Volume string
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
	Asset    Asset
	Start    int
	Length   int
	Fit      string
	Scale    int
	Position string
	Offset   Offset
	Transition
	Effect  string
	Filter  string
	Opacity int
	Transform
}
type Track struct {
	Clips []Clip
}

type Timeline struct {
	SoundTrack
	Background string
	fonts      []Font
	tracks     []Track
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
	Disk     string // "local" | "mount"
}

func NewEdit() *Edit {
	return &Edit{}
}

func NewTimeline() *Timeline {
	return &Timeline{}
}

func NewOutput() *Output {
	return &Output{}
}

func NewMerges() *[]Merge {
	return &[]Merge{}
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

func (e *Edit) SetDisk(disk string) *Edit {
	return &Edit{Disk: disk}
}
