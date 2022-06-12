package response

import "net/http"

type Response struct {
	Status  bool        `json:"status"`
	Errors  interface{} `json:"errors,omitempty"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (response *Response) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func Build(data interface{}, message string, errors interface{}, status bool) *Response {
	return &Response{
		Status:  status,
		Errors:  errors,
		Message: message,
		Data:    data,
	}
}
