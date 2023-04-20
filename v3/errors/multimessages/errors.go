package multimessages

import (
	goliatherrors "github.com/fazpass/goliath/v3/errors"
)

const (
	DefaultMessage = "Error"
)

type Error struct {
	Errors     map[string]error
	HttpStatus int
	CaseCode   string
}

func New(httpStatus int, caseCode string, messages map[string]string) *Error {

	var errorsMap = make(map[string]error)

	for key, message := range messages {
		errorsMap[key] = goliatherrors.New(httpStatus, caseCode, message)
	}

	return &Error{
		Errors:     errorsMap,
		HttpStatus: httpStatus,
		CaseCode:   caseCode,
	}
}

func (e *Error) Get(key string) error {

	message, ok := e.Errors[key]
	if ok {
		return message
	}

	if key == "" {
		for _, err := range e.Errors {
			return err
		}
	}

	return goliatherrors.New(e.HttpStatus, e.CaseCode, DefaultMessage)
}
