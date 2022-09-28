package edit

type AudioAssetType struct {
	Type   string      `json:"type"`
	Src    string      `json:"src"`
	Trim   int         `json:"trim,omitempty"`
	Volume float32     `json:"volume,omitempty"`
	Effect AudioEffect `json:"effect,omitempty"`
}

func NewAudioAsset() *AudioAssetType {
	a := new(AudioAssetType)
	a.Type = "audio"
	return a
}

func (a *AudioAssetType) SetSrc(src string) *AudioAssetType {
	a.Src = src
	return a
}
func (a *AudioAssetType) SetTrim(trim int) *AudioAssetType {
	a.Trim = trim
	return a
}

func (a *AudioAssetType) SetVolume(vol float32) *AudioAssetType {
	a.Volume = vol
	return a
}

func (a *AudioAssetType) SetEffect(effect AudioEffect) *AudioAssetType {
	a.Effect = effect
	return a
}
