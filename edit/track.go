package edit

type Track struct {
	Clips *[]Clip `json:"clips"`
}

func NewTrack() *Track {
	t := new(Track)
	return t
}

func (t *Track) SetClips(clips *[]Clip) *Track {
	t.Clips = clips
	return t
}
