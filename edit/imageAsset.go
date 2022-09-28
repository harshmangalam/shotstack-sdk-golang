package edit

type ImageAssetType struct {
	Type string `json:"type"`
	Src  string `json:"src"`
	Crop *Crop  `json:"crop,omitempty"`
}

func NewImageAsset() *ImageAssetType {
	i := new(ImageAssetType)
	i.Type = "image"
	return i
}

func (i *ImageAssetType) SetImageAssetSrc(src string) *ImageAssetType {
	i.Src = src
	return i
}
func (i *ImageAssetType) SetImageAssetCrop(crop *Crop) *ImageAssetType {
	i.Crop = crop
	return i
}
