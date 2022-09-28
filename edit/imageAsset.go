package edit

type ImageAsset struct {
	Type string `json:"type"`
	Src  string `json:"src"`
	Crop *Crop  `json:"crop,omitempty"`
}

func NewImageAsset() *ImageAsset {
	i := new(ImageAsset)
	i.Type = "image"
	return i
}

func (i *ImageAsset) SetImageAssetSrc(src string) *ImageAsset {
	i.Src = src
	return i
}
func (i *ImageAsset) SetImageAssetCrop(crop *Crop) *ImageAsset {
	i.Crop = crop
	return i
}
