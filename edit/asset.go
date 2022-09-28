package edit

type TitleAssetType struct {
	Text  string
	Style TitleAssetStyle
}

type HtmlAssetType struct {
	Html       string
	Css        string
	Width      int
	Height     int
	Background string
	Position   Position
}

type AudioAssetType struct {
	Src    string
	Trim   int
	Volume float32
	Effect AudioEffect
}

type LumaAssetType struct {
	Src  string
	Trim int
}
type Asset struct {
	Type       AssetType
	VideoAsset *VideoAssetType
	ImageAsset *ImageAssetType
	TitleAsset *TitleAssetType
	HtmlAsset  *HtmlAssetType
	AudioAsset *AudioAssetType
	LumaAsset  *LumaAssetType
}

func NewAsset() *Asset {
	a := new(Asset)
	return a
}

func (a *Asset) SetAssetType(assetType AssetType) *Asset {
	a.Type = assetType
	return a
}

func (a *Asset) SetVideoAsset(videoAsset *VideoAssetType) *Asset {
	a.VideoAsset = videoAsset
	return a
}

func (a *Asset) SetAudioAsset(audioAsset *AudioAssetType) *Asset {
	a.AudioAsset = audioAsset
	return a
}

func (a *Asset) SetImageAsset(imageAsset *ImageAssetType) *Asset {
	a.ImageAsset = imageAsset
	return a
}

func (a *Asset) SetHtmlAsset(htmlAsset *HtmlAssetType) *Asset {
	a.HtmlAsset = htmlAsset
	return a
}

func (a *Asset) SetTitleAsset(titleAsset *TitleAssetType) *Asset {
	a.TitleAsset = titleAsset
	return a
}

func (a *Asset) SetLumaAsset(lumaAsset *LumaAssetType) *Asset {
	a.LumaAsset = lumaAsset
	return a
}

// audio asset

func NewAudioAsset() *AudioAssetType {
	a := new(AudioAssetType)
	return a
}

func (a *AudioAssetType) SetAudioAssetSrc(src string) *AudioAssetType {
	a.Src = src
	return a
}
func (a *AudioAssetType) SetAudioAssetTrim(trim int) *AudioAssetType {
	a.Trim = trim
	return a
}

func (a *AudioAssetType) SetAudioAssetVolume(vol float32) *AudioAssetType {
	a.Volume = vol
	return a
}

func (a *AudioAssetType) SetAudioAssetEffect(effect AudioEffect) *AudioAssetType {
	a.Effect = effect
	return a
}

// html asset

func NewHtmlAsset() *HtmlAssetType {
	h := new(HtmlAssetType)
	return h
}

func (h *HtmlAssetType) SetHtmlAssetHtml(html string) *HtmlAssetType {
	h.Html = html
	return h
}

func (h *HtmlAssetType) SetHtmlAssetCss(css string) *HtmlAssetType {
	h.Css = css
	return h
}

func (h *HtmlAssetType) SetHtmlAssetHeight(height int) *HtmlAssetType {
	h.Height = height
	return h
}

func (h *HtmlAssetType) SetHtmlAssetWidth(width int) *HtmlAssetType {
	h.Width = width
	return h
}

func (h *HtmlAssetType) SetHtmlAssetPosition(pos Position) *HtmlAssetType {
	h.Position = pos
	return h
}
