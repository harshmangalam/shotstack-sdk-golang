package edit

type VideoAssetType struct {
	Type   string
	Src    string  `json:"src"`
	Trim   int     `json:"trim"`
	Volume float32 `json:"volume"`
	Crop   *Crop   `json:"crop"`
}

func NewVideoAsset() *VideoAssetType {
	v := new(VideoAssetType)
	v.Type = "video"
	return v
}
func (v *VideoAssetType) SetVideoAssetSrc(src string) *VideoAssetType {
	v.Src = src
	return v
}

func (v *VideoAssetType) SetVideoAssetTrim(trim int) *VideoAssetType {
	v.Trim = trim
	return v
}

func (v *VideoAssetType) SetVideoAssetVolume(vol float32) *VideoAssetType {
	v.Volume = vol
	return v
}

func (v *VideoAssetType) SetVideoAssetCrop(crop *Crop) *VideoAssetType {
	v.Crop = crop
	return v
}
