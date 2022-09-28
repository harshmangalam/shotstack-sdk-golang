package edit

type ImageAssetType struct {
	Src  string
	Crop *Crop
}

func NewImageAsset() *ImageAssetType {
	i := new(ImageAssetType)
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
