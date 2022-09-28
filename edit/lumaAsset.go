package edit

type LumaAsset struct {
	Type string `json:"type"`
	Src  string `json:"src"`
	Trim int    `json:"trim,omitempty"`
}

func NewLumaAsset() *LumaAsset {
	l := new(LumaAsset)
	l.Type = "luma"
	return l
}

func (l *LumaAsset) SetSrc(src string) *LumaAsset {
	l.Src = src
	return l
}

func (l *LumaAsset) SetTrim(trim int) *LumaAsset {
	l.Trim = trim
	return l
}
