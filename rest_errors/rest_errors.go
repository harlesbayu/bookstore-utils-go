package rest_errors

import (
	"errors"
	"net/http"
)

func NewError(msg string) error {
	return errors.New(msg)
}

type RestErr struct {
	Message string `JSON:"message"`
	Status  int    `JSON:"status"`
	Error   string `JSON:"error"`
	Cause []interface{} `JSON:"cause"`
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewInternalServerError(message string, err error) *RestErr {
	result := &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
	if err != nil {
		result.Cause = append(result.Cause, err.Error())
	}
	return  result
}
