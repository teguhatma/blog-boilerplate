package controller

import (
	"net/http"

	"github.com/lib/pq"
	"github.com/teguhatma/blog-boilerplate/errors"
	shttp "github.com/teguhatma/blog-boilerplate/server/http"
)

func errResponse(err error) *shttp.ErrorResponse {
	if val, ok := err.(errors.FError); ok {
		switch val.Code() {
		case http.StatusInternalServerError:
			return &shttp.ErrorResponse{
				Error:      val.Cause().Error(),
				StatusCode: int(val.Code()),
			}
		case http.StatusNotFound:
			return &shttp.ErrorResponse{
				Error:      val.Cause().Error(),
				StatusCode: int(val.Code()),
			}
		}
	}

	if val, ok := err.(*pq.Error); ok {
		switch val.Code.Name() {
		case "unique_violation":
		case "foreign_key_violation":
			return &shttp.ErrorResponse{
				Error:      err,
				StatusCode: http.StatusBadRequest,
			}
		}
	}
	return nil
}
