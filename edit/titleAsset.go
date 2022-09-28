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

type TitleAssetType struct {
	Type       string          `json:"type"`
	Text       string          `json:"text"`
	Style      TitleAssetStyle `json:"style,omitempty"`
	Color      string          `json:"color,omitempty"`
	Size       TitleSize       `json:"size,omitempty"`
	Background string          `json:"background,omitempty"`
	Position   string          `json:"position,omitempty"`
	Offset     *Offset         `json:"offset,omitempty"`
}

func NewTitleAsset() *TitleAssetType {
	t := new(TitleAssetType)
	t.Type = "title"
	return t
}

func (t *TitleAssetType) SetText(text string) *TitleAssetType {
	t.Text = text
	return t
}

func (t *TitleAssetType) SetStyle(style TitleAssetStyle) *TitleAssetType {
	t.Style = style
	return t
}

func (t *TitleAssetType) SetColor(color string) *TitleAssetType {
	t.Color = color
	return t
}

func (t *TitleAssetType) SetSize(size TitleSize) *TitleAssetType {
	t.Size = size
	return t
}

func (t *TitleAssetType) SetBackground(background string) *TitleAssetType {
	t.Background = background
	return t
}

func (t *TitleAssetType) SetPosition(position string) *TitleAssetType {
	t.Position = position
	return t
}

func (t *TitleAssetType) SetOffset(offset *Offset) *TitleAssetType {
	t.Offset = offset
	return t
}
