package edit

type HtmlAsset struct {
	Type       string   `json:"type"`
	Html       string   `json:"html"`
	Css        string   `json:"css,omitempty"`
	Width      int      `json:"width,omitempty"`
	Height     int      `json:"height,omitempty"`
	Background string   `json:"background,omitempty"`
	Position   Position `json:"position,omitempty"`
}

func NewHtmlAsset() *HtmlAsset {
	h := new(HtmlAsset)
	h.Type = "html"
	return h
}

func (h *HtmlAsset) SetHtml(html string) *HtmlAsset {
	h.Html = html
	return h
}

func (h *HtmlAsset) SetCss(css string) *HtmlAsset {
	h.Css = css
	return h
}

func (h *HtmlAsset) SetHeight(height int) *HtmlAsset {
	h.Height = height
	return h
}

func (h *HtmlAsset) SetWidth(width int) *HtmlAsset {
	h.Width = width
	return h
}

func (h *HtmlAsset) SetBackground(background string) *HtmlAsset {
	h.Background = background
	return h
}

func (h *HtmlAsset) SetPosition(pos Position) *HtmlAsset {
	h.Position = pos
	return h
}
