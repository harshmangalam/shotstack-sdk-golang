package edit

type QueuedResponseData struct {
	Message string `json:"message"`
	Id      string `json:"id"`
}
type QueuedResponse struct {
	Success  bool   `json:"success"`
	Message  string `json:"message"`
	Response QueuedResponseData
}

type RenderResponseData struct {
	Id         string `json:"id"`
	Owner      string `json:"owner"`
	Plan       string `json:"plan,omitempty"`
	Status     string `json:"status"`
	Error      string `json:"error,omitempty"`
	Duration   int    `json:"duration,omitempty"`
	RenderTime int    `json:"renderTime,omitempty"`
	Url        string `json:"url,omitempty"`
	Poster     string `json:"poster,omitempty"`
	Thumbnail  string `json:"thumbnail,omitempty"`
	Data       Edit   `json:"data,omitempty"`
	Created    string `json:"created,omitempty"`
	Updated    string `json:"updated,omitempty"`
}
type RenderResponse struct {
	Success  bool               `json:"success"`
	Message  string             `json:"message"`
	Response RenderResponseData `json:"response"`
}
