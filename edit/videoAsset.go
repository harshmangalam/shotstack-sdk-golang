package edit

type VideoAsset struct {
	Type   string
	Src    string  `json:"src"`
	Trim   int     `json:"trim"`
	Volume float32 `json:"volume"`
	Crop   *Crop   `json:"crop"`
}

func NewVideoAsset() *VideoAsset {
	v := new(VideoAsset)
	v.Type = "video"
	return v
}
func (v *VideoAsset) SetVideoAssetSrc(src string) *VideoAsset {
	v.Src = src
	return v
}

func (v *VideoAsset) SetVideoAssetTrim(trim int) *VideoAsset {
	v.Trim = trim
	return v
}

func (v *VideoAsset) SetVideoAssetVolume(vol float32) *VideoAsset {
	v.Volume = vol
	return v
}

func (v *VideoAsset) SetVideoAssetCrop(crop *Crop) *VideoAsset {
	v.Crop = crop
	return v
}
