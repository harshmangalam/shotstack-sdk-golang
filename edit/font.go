package edit

type Font struct {
	Src string
}

func NewFont() *Font {
	f := new(Font)
	return f
}

func (f *Font) SetSrc(src string) *Font {
	f.Src = src
	return f
}
