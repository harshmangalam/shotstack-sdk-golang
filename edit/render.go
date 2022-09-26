package edit

type SoundTrackEffect string
type Disk string

const (
	FadeIn        SoundTrackEffect = "fadeIn"
	FadeOut       SoundTrackEffect = "fadeOut"
	FadeInFadeOut SoundTrackEffect = "fadeInFadeOut"
)

const (
	Local Disk = "local"
	Mount Disk = "mount"
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
	Disk     Disk // "local" | "mount"
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

// create new edit instance
func NewEdit() *Edit {
	return &Edit{}
}

// set edit timeline
func (e *Edit) SetTimeline(timeline *Timeline) *Edit {
	return &Edit{Timeline: timeline}
}

// set edit output
func (e *Edit) SetOutput(output *Output) *Edit {
	return &Edit{Output: output}
}

// set edit merges
func (e *Edit) SetMerges(merges *[]Merge) *Edit {
	return &Edit{Merges: merges}
}

// set edit callback
func (e *Edit) SetCallback(callback string) *Edit {
	return &Edit{Callback: callback}
}

// set edit disk
func (e *Edit) SetDisk(disk Disk) *Edit {
	return &Edit{Disk: disk}
}
