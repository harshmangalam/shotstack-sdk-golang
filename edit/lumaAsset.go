package edit

type LumaAssetType struct {
	Src  string
	Trim int
}

func NewLumaAsset() *LumaAssetType {
	return new(LumaAssetType)
}

func (l *LumaAssetType) SetSrc(src string) *LumaAssetType {
	l.Src = src
	return l
}

func (l *LumaAssetType) SetTrim(trim int) *LumaAssetType {
	l.Trim = trim
	return l
}
