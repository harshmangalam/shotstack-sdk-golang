package edit

type HtmlAssetType struct {
	Type       string   `json:"type"`
	Html       string   `json:"html"`
	Css        string   `json:"css,omitempty"`
	Width      int      `json:"width,omitempty"`
	Height     int      `json:"height,omitempty"`
	Background string   `json:"background,omitempty"`
	Position   Position `json:"position,omitempty"`
}

func NewHtmlAsset() *HtmlAssetType {
	h := new(HtmlAssetType)
	h.Type = "html"
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
