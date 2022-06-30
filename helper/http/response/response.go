package response

import "net/http"

type Response struct {
	Status  bool        `json:"status"`
	Errors  interface{} `json:"errors,omitempty"`
	Message string      `json:"message,omitempty"`
	Code    string      `json:"code,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

func (response *Response) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func Build(response *Response) *Response {
	return response
}
