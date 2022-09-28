package edit

type Timeline struct {
	SoundTrack *SoundTrack `json:"soundtrack,omitempty"`
	Background string      `json:"background,omitempty"`
	Fonts      *[]Font     `json:"fonts,omitempty"`
	Tracks     *[]Track    `json:"tracks"`
	Cache      bool        `json:"cache,omitempty"`
}

func NewTimeline() *Timeline {
	return new(Timeline)

}

func (t *Timeline) SetSoundTrack(soundTrack *SoundTrack) *Timeline {
	t.SoundTrack = soundTrack
	return t
}

func (t *Timeline) SetBackground(background string) *Timeline {
	t.Background = background
	return t
}

func (t *Timeline) SetFonts(fonts *[]Font) *Timeline {
	t.Fonts = fonts
	return t
}

func (t *Timeline) SetTracks(tracks *[]Track) *Timeline {
	t.Tracks = tracks
	return t
}

func (t *Timeline) SetCache(cache bool) *Timeline {
	t.Cache = cache
	return t
}
