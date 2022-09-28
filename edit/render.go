package edit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	shotstack "github.com/harshmangalam/shotstack-sdk-golang"
)

// edit

func (e *Edit) PostRender(config *shotstack.Config) interface{} {
	Url := fmt.Sprintf("https://api.shotstack.io/%v/render", config.Env)
	jsonData, err := json.MarshalIndent(e, "", "   ")

	fmt.Println(string(jsonData))

	if err != nil {
		panic(err)
	}

	headers := map[string][]string{
		"Content-Type": {"application/json"},
		"Accept":       {"application/json"},
		"x-api-key":    {config.ApiKey},
	}

	data := bytes.NewBuffer(jsonData)
	req, err := http.NewRequest("POST", Url, data)

	if err != nil {
		panic(err)
	}
	req.Header = headers

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	stringResp := string(bodyBytes)
	return stringResp

}
