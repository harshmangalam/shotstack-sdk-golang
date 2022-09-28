package edit

type TitleAssetStyle string

const (
	Minimal     TitleAssetStyle = "minimal"
	Blockbuster TitleAssetStyle = "blockbuster"
	Vogue       TitleAssetStyle = "vogue"
	Sketchy     TitleAssetStyle = "sketchy"
	Skinny      TitleAssetStyle = "skinny"
	Chunk       TitleAssetStyle = "chunk"
	ChunkLight  TitleAssetStyle = "chunkLight"
	Marker      TitleAssetStyle = "marker"
	Future      TitleAssetStyle = "future"
	Subtitle    TitleAssetStyle = "subtitle"
)

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
