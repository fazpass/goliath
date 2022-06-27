package response

import "net/http"

type Response struct {
	Status  bool        `json:"status"`
	Errors  interface{} `json:"errors,omitempty"`
	Message string      `json:"message"`
	Code    string      `json:"code"`
	Data    interface{} `json:"data,omitempty"`
}

func (response *Response) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func Build(response *Response) *Response {
	return response
}
