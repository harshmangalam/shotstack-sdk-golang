package edit

type TitleAssetStyle string
type TitleSize string

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

const (
	SizeXxSmall TitleSize = "xx-small"
	SizeXSmall  TitleSize = "x-small"
	SizeSmall   TitleSize = "small"
	SizeMedium  TitleSize = "medium"
	SizeXLarge  TitleSize = "x-large"
	SizeXxLarge TitleSize = "xx-large"
)

type TitleAsset struct {
	Type       string          `json:"type"`
	Text       string          `json:"text"`
	Style      TitleAssetStyle `json:"style,omitempty"`
	Color      string          `json:"color,omitempty"`
	Size       TitleSize       `json:"size,omitempty"`
	Background string          `json:"background,omitempty"`
	Position   string          `json:"position,omitempty"`
	Offset     *Offset         `json:"offset,omitempty"`
}

func NewTitleAsset() *TitleAsset {
	t := new(TitleAsset)
	t.Type = "title"
	return t
}

func (t *TitleAsset) SetText(text string) *TitleAsset {
	t.Text = text
	return t
}

func (t *TitleAsset) SetStyle(style TitleAssetStyle) *TitleAsset {
	t.Style = style
	return t
}

func (t *TitleAsset) SetColor(color string) *TitleAsset {
	t.Color = color
	return t
}

func (t *TitleAsset) SetSize(size TitleSize) *TitleAsset {
	t.Size = size
	return t
}

func (t *TitleAsset) SetBackground(background string) *TitleAsset {
	t.Background = background
	return t
}

func (t *TitleAsset) SetPosition(position string) *TitleAsset {
	t.Position = position
	return t
}

func (t *TitleAsset) SetOffset(offset *Offset) *TitleAsset {
	t.Offset = offset
	return t
}
