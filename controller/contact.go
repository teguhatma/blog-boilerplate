package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/teguhatma/blog-boilerplate/request"
	shttp "github.com/teguhatma/blog-boilerplate/server/http"
	"github.com/teguhatma/blog-boilerplate/service/contact"
)

type contactController struct {
	service contact.Service
}

func (c *contactController) RegisterRoutes(router *mux.Router) {
	router.Handle("/api/v1/contacts", shttp.AppHandler(c.createContact)).Methods(http.MethodPost)
}

func (c *contactController) createContact(r *http.Request) (*shttp.Response, error) {
	var req request.ContactRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errResponse(err)
	}

	res, err := c.service.CreateContact(context.Background(), &req)
	if err != nil {
		return nil, errResponse(err)
	}

	return &shttp.Response{
		Data:       res,
		StatusCode: http.StatusCreated,
	}, nil
}

func NewContactController(router *mux.Router, service contact.Service) *contactController {
	Controller := &contactController{
		service: service,
	}
	Controller.RegisterRoutes(router)
	return Controller
}