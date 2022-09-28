package edit

type SoundTrack struct {
	Src    string      `json:"src"`
	Effect AudioEffect `json:"effect"`
	Volume float32     `json:"volume"`
}

func NewSoundTrack() *SoundTrack {
	s := new(SoundTrack)
	return s
}

func (s *SoundTrack) SetSrc(src string) *SoundTrack {
	s.Src = src
	return s
}

func (s *SoundTrack) SetEffect(effect AudioEffect) *SoundTrack {
	s.Effect = effect
	return s
}

func (s *SoundTrack) SetVolume(volume float32) *SoundTrack {
	s.Volume = volume
	return s
}
