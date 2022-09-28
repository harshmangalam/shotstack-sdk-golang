package edit

type HtmlAssetType struct {
	Html       string   `json:"html"`
	Css        string   `json:"css"`
	Width      int      `json:"width"`
	Height     int      `json:"height"`
	Background string   `json:"background"`
	Position   Position `json:"position"`
}

func NewHtmlAsset() *HtmlAssetType {
	h := new(HtmlAssetType)
	return h
}

func (h *HtmlAssetType) SetHtml(html string) *HtmlAssetType {
	h.Html = html
	return h
}

func (h *HtmlAssetType) SetCss(css string) *HtmlAssetType {
	h.Css = css
	return h
}

func (h *HtmlAssetType) SetHeight(height int) *HtmlAssetType {
	h.Height = height
	return h
}

func (h *HtmlAssetType) SetWidth(width int) *HtmlAssetType {
	h.Width = width
	return h
}

func (h *HtmlAssetType) SetPosition(pos Position) *HtmlAssetType {
	h.Position = pos
	return h
}
