package errors

type Error struct {
	CaseCode       string
	HttpStatus     int
	Message        string
	DefaultMessage string
	Data           interface{}
}

func New(httpStatus int, caseCode string, message string) error {
	return &Error{
		CaseCode:   caseCode,
		HttpStatus: httpStatus,
		Message:    message,
		Data:       nil,
	}
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) GetHttpStatus() int {
	return e.HttpStatus
}
