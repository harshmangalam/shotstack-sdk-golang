package shotstacksdkgolang

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ReqMethod string

const (
	POST ReqMethod = "POST"
	GET  ReqMethod = "GET"
)

type Request struct {
	Method ReqMethod `json:"method"`
	Data   any       `json:"data"`
	Path   string    `json:"params"`
	Config *Config   `json:"config"`
}

func PostRequest(req *Request) ([]byte, error) {
	BaseUrl := fmt.Sprintf("https://api.shotstack.io/%v", req.Config.ApiKey)

	switch req.Method {
	case POST:
		Url := fmt.Sprintf(BaseUrl + req.Path)
		jsonData, err := json.MarshalIndent(req.Data, "", "   ")
		if err != nil {
			return nil, err
		}
		fmt.Println(string(jsonData))

		headers := map[string][]string{
			"Content-Type": {"application/json"},
			"Accept":       {"application/json"},
			"x-api-key":    {req.Config.ApiKey},
		}

		buff := bytes.NewBuffer(jsonData)
		httpReq, err := http.NewRequest(string(POST), Url, buff)

		if err != nil {
			return nil, err
		}
		httpReq.Header = headers

		client := &http.Client{}
		resp, err := client.Do(httpReq)

		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		bodyBytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return nil, err
		}

		return bodyBytes, nil

	default:
		err := errors.New("request method is invalid")
		return nil, err
	}

}
