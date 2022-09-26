package edit

type EffectType string

const (
	FadeIn        EffectType = "fadeIn"
	FadeOut       EffectType = "fadeOut"
	FadeInFadeOut EffectType = "fadeInFadeOut"
)

type SoundTrack struct {
	Src    string
	Effect EffectType
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
	Timeline Timeline
	Output   Output
	Merge    []Merge
	Callback string
	Disk     string // "local" | "mount"
}
