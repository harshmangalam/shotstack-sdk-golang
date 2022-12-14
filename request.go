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

func NewRequest() *Request {
	return new(Request)
}

func (req *Request) SetMethod(method ReqMethod) *Request {
	req.Method = method
	return req
}
func (req *Request) SetData(data any) *Request {
	req.Data = data
	return req
}

func (req *Request) SetPath(path string) *Request {
	req.Path = path
	return req
}

func (req *Request) SetConfig(config *Config) *Request {
	req.Config = config
	return req
}

func (req *Request) Send() ([]byte, error) {
	Url := fmt.Sprintf("https://api.shotstack.io/%v/%v", req.Config.Env, req.Path)
	headers := map[string][]string{
		"Content-Type": {"application/json"},
		"Accept":       {"application/json"},
		"x-api-key":    {req.Config.ApiKey},
	}
	client := &http.Client{}

	switch req.Method {
	case GET:
		httpReq, err := http.NewRequest(http.MethodGet, Url, nil)

		if err != nil {
			return nil, err
		}
		httpReq.Header = headers
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

	case POST:

		jsonData, err := json.MarshalIndent(req.Data, "", "   ")
		if err != nil {
			return nil, err
		}
		fmt.Println(string(jsonData))

		buff := bytes.NewBuffer(jsonData)
		httpReq, err := http.NewRequest(http.MethodPost, Url, buff)

		if err != nil {
			return nil, err
		}
		httpReq.Header = headers

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
