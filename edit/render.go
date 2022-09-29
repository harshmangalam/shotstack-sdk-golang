package edit

import (
	"encoding/json"

	shotstack "github.com/harshmangalam/shotstack-sdk-golang"
)

// edit

func (e *Edit) PostRender(config *shotstack.Config) (*QueuedResponse, error) {

	res, err := shotstack.NewRequest().SetMethod(shotstack.POST).SetPath("render").SetConfig(config).SetData(e).Send()

	if err != nil {
		return nil, err
	}

	respData := new(QueuedResponse)
	json.Unmarshal(res, respData)
	return respData, nil

}

func GetRender(id string, config *shotstack.Config) (*RenderResponse, error) {
	res, err := shotstack.NewRequest().SetMethod(shotstack.GET).SetPath("render" + "/" + id).SetConfig(config).Send()

	if err != nil {
		return nil, err
	}

	respData := new(RenderResponse)
	json.Unmarshal(res, respData)
	return respData, nil
}
