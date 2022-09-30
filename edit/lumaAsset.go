package edit

type LumaAsset struct {
	Type string  `json:"type"`
	Src  string  `json:"src"`
	Trim float32 `json:"trim,omitempty"`
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

func (l *LumaAsset) SetTrim(trim float32) *LumaAsset {
	l.Trim = trim
	return l
}
