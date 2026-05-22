package http_server

import (
	"errors"
)

var CannotDecodeRequest = errors.New("Cannot decode your request")

type APIError struct {
	Status  int
	Message string
}

func (err *APIError) Error() string {
	return err.Message
}
func MappingErrorsToStatus(err error) APIError {
	msg := err.Error()
	switch err {
	case CannotDecodeRequest:
		return APIError{Status: 400, Message: msg}
	default:
		return APIError{Status: 500, Message: msg}
	}
}
