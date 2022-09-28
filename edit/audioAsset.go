package edit

type AudioAssetType struct {
	Src    string      `json:"src"`
	Trim   int         `json:"trim"`
	Volume float32     `json:"volume"`
	Effect AudioEffect `json:"effect"`
}

func NewAudioAsset() *AudioAssetType {
	a := new(AudioAssetType)
	return a
}

func (a *AudioAssetType) SetAudioAssetSrc(src string) *AudioAssetType {
	a.Src = src
	return a
}
func (a *AudioAssetType) SetAudioAssetTrim(trim int) *AudioAssetType {
	a.Trim = trim
	return a
}

func (a *AudioAssetType) SetAudioAssetVolume(vol float32) *AudioAssetType {
	a.Volume = vol
	return a
}

func (a *AudioAssetType) SetAudioAssetEffect(effect AudioEffect) *AudioAssetType {
	a.Effect = effect
	return a
}
