package edit

type MergeField struct {
	Find    string `json:"find"`
	Replace any    `json:"replace"`
}

func NewMergeField() *MergeField {
	return new(MergeField)
}

func (m *MergeField) SetFind(find string) *MergeField {
	m.Find = find
	return m
}

func (m *MergeField) SetReplace(replace any) *MergeField {
	m.Replace = replace
	return m
}
