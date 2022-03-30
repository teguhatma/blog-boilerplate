package controller

import (
	"net/http"

	shttp "github.com/teguhatma/blog-boilerplate/server/http"
	"github.com/teguhatma/blog-boilerplate/service"
)

func errResponse(err error) *shttp.Response {
	if val, ok := err.(service.FError); ok {
		switch val.Code() {
		case http.StatusInternalServerError:
			return &shttp.Response{
				Data:       err.Error(),
				StatusCode: int(val.Code()),
			}
		case http.StatusNotFound:
			return &shttp.Response{
				Data:       err.Error(),
				StatusCode: int(val.Code()),
			}
		}
	}
	return nil
}
