package edit

type Effect string

const (
	FadeIn        Effect = "fadeIn"
	FadeOut       Effect = "fadeOut"
	FadeInFadeOut Effect = "fadeInFadeOut"
)

type SoundTrack struct {
	src    string
	effect Effect
	volume string
}

type Font struct {
}

type Track struct {
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
