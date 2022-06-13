package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/teguhatma/blog-boilerplate/request"
	shttp "github.com/teguhatma/blog-boilerplate/server/http"
	"github.com/teguhatma/blog-boilerplate/service/contact"
)

type contactController struct {
	service contact.Service
}

func (c *contactController) RegisterRoutes(router *mux.Router) {
	router.Handle("/api/v1/contact", shttp.AppHandler(c.createContact)).Methods(http.MethodPost)
	router.Handle("/api/v1/contact/{id:[0-9]+}", shttp.AppHandler(c.getContact)).Methods(http.MethodGet)
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

func (c *contactController) getContact(r *http.Request) (*shttp.Response, error) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return nil, errResponse(err)
	}

	res, err := c.service.GetContact(context.Background(), id)
	if err != nil {
		return nil, errResponse(err)
	}

	return &shttp.Response{
		Data:       res,
		StatusCode: http.StatusOK,
	}, nil
}

func NewContactController(router *mux.Router, service contact.Service) *contactController {
	Controller := &contactController{
		service: service,
	}
	Controller.RegisterRoutes(router)
	return Controller
}
