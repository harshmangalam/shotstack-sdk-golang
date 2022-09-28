package edit

type Merge struct {
	Find    string `json:"find"`
	Replace any    `json:"replace"`
}

func NewMerge() *Merge {
	return new(Merge)
}

func (m *Merge) SetFind(find string) *Merge {
	m.Find = find
	return m
}

func (m *Merge) SetReplace(replace any) *Merge {
	m.Replace = replace
	return m
}
