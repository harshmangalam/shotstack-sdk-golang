package shotstacksdkgolang

type QueuedResponseData struct {
	Message string `json:"message"`
	Id      string `json:"id"`
}
type QueuedResponse struct {
	Success  string `json:"success"`
	Message  string `json:"message"`
	Response QueuedResponseData
}
