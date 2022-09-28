package edit

import (
	"encoding/json"
	"fmt"

	shotstack "github.com/harshmangalam/shotstack-sdk-golang"
)

// edit

func (e *Edit) PostRender(config *shotstack.Config) (*shotstack.QueuedResponse, error) {

	res, err := shotstack.NewRequest().SetMethod(shotstack.POST).SetPath("/render").SetConfig(config).SetData(e).Send()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(res))
	respData := new(shotstack.QueuedResponse)
	json.Unmarshal(res, respData)
	return respData, nil

}
