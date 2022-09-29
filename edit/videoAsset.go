package edit

type VideoAsset struct {
	Type   string  `json:"type"`
	Src    string  `json:"src"`
	Trim   int     `json:"trim,omitempty"`
	Volume float32 `json:"volume,omitempty"`
	Crop   *Crop   `json:"crop,omitempty"`
}

func NewVideoAsset() *VideoAsset {
	v := new(VideoAsset)
	v.Type = "video"
	return v
}
func (v *VideoAsset) SetSrc(src string) *VideoAsset {
	v.Src = src
	return v
}

func (v *VideoAsset) SetTrim(trim int) *VideoAsset {
	v.Trim = trim
	return v
}

func (v *VideoAsset) SetVolume(vol float32) *VideoAsset {
	v.Volume = vol
	return v
}

func (v *VideoAsset) SetCrop(crop *Crop) *VideoAsset {
	v.Crop = crop
	return v
}
