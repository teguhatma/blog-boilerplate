package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/teguhatma/blog-boilerplate/request"
	shttp "github.com/teguhatma/blog-boilerplate/server/http"
	"github.com/teguhatma/blog-boilerplate/service/entries"
)

type entryController struct {
	service entries.Service
}

func NewEntriesController(router *mux.Router, service entries.Service) *entryController {
	Controller := &entryController{
		service: service,
	}
	Controller.RegisterRoutes(router)
	return Controller
}

func (c *entryController) RegisterRoutes(router *mux.Router) {
	router.Handle("/api/v1/entry", shttp.AppHandler(c.createEntry)).Methods(http.MethodPost)

}

func (c *entryController) createEntry(r *http.Request) (*shttp.Response, error) {
	var req request.CreateEntryRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errResponse(err)
	}

	entry, err := c.service.CreateEntry(context.Background(), req)
	if err != nil {
		return nil, errResponse(err)
	}

	return &shttp.Response{
		Data:       entry,
		StatusCode: http.StatusCreated,
	}, nil
}
