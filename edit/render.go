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

type Clip struct {
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
