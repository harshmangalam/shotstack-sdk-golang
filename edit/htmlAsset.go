package edit

type HtmlAssetType struct {
	Html       string
	Css        string
	Width      int
	Height     int
	Background string
	Position   Position
}

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
