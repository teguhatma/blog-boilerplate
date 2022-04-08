package controller

import (
	"net/http"

	fe "github.com/teguhatma/blog-boilerplate/errors"
	shttp "github.com/teguhatma/blog-boilerplate/server/http"
)

func NewResponseErrorWithCause(code fe.Code, err error, message string, statusCode int) error {
	return shttp.ResponseError{
		FError:     fe.NewWithCause(code, err, message),
		StatusCode: statusCode,
	}
}

func errResponse(err error) error {
	if v, ok := err.(fe.FError); ok {
		switch v.Code() {
		case "NOT_FOUND":
			return NewResponseErrorWithCause(v.Code(), err, v.Message(), http.StatusNotFound)
		case "BAD_MESSAGE":
			return NewResponseErrorWithCause(v.Code(), err, v.Message(), http.StatusBadRequest)
		default:
			return NewResponseErrorWithCause(v.Code(), err, v.Message(), http.StatusInternalServerError)
		}
	}
	return NewResponseErrorWithCause(fe.INTERNAL_ERROR, err, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
}
