package edit

type TitleAssetType struct {
	Text  string          `json:"text"`
	Style TitleAssetStyle `json:"style"`
}

func NewTitleAsset() *TitleAssetType {
	return new(TitleAssetType)
}

func (t *TitleAssetType) SetText(text string) *TitleAssetType {
	t.Text = text
	return t
}

func (t *TitleAssetType) SetStyle(style TitleAssetStyle) *TitleAssetType {
	t.Style = style
	return t
}
