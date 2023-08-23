package edit

import (
	"encoding/json"

	shotstack "github.com/dougbarrett/shotstack-sdk-golang"
)

type Template struct {
	Merge *[]MergeField `json:"merge,omitempty"`
	ID    string        `json:"id,omitempty"`
}

func NewTemplate() *Template {
	t := new(Template)
	return t

}

func (e *Template) SetMerge(merge *[]MergeField) *Template {
	e.Merge = merge
	return e
}

func (e *Template) SetID(id string) *Template {
	e.ID = id
	return e
}

func (e *Template) PostRender(config *shotstack.Config) (*QueuedResponse, error) {

	res, err := shotstack.NewRequest().SetMethod(shotstack.POST).SetPath("templates/render").SetConfig(config).SetData(e).Send()

	if err != nil {
		return nil, err
	}

	respData := new(QueuedResponse)
	json.Unmarshal(res, respData)
	return respData, nil

}
