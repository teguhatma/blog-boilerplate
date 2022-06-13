package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/teguhatma/blog-boilerplate/request"
	shttp "github.com/teguhatma/blog-boilerplate/server/http"
	"github.com/teguhatma/blog-boilerplate/service/tag"
)

type tagController struct {
	service tag.Service
}

func (c *tagController) RegisterRoutes(router *mux.Router) {
	router.Handle("/api/v1/tag", shttp.AppHandler(c.createTag)).Methods(http.MethodPost)
	router.Handle("/api/v1/tag/{id:[0-9]+}", shttp.AppHandler(c.getTag)).Methods(http.MethodGet)
	router.Handle("/api/v1/tag/{id:[0-9]+}", shttp.AppHandler(c.deleteTag)).Methods(http.MethodDelete)
	router.Handle("/api/v1/tags", shttp.AppHandler(c.listTag)).Methods(http.MethodGet)
	router.Handle("/api/v1/tag/{id:[0-9]+}", shttp.AppHandler(c.updateTag)).Methods(http.MethodPut)
}

func (c *tagController) createTag(r *http.Request) (*shttp.Response, error) {
	var req request.TagRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errResponse(err)
	}

	res, err := c.service.CreateTag(context.Background(), req.Name)
	if err != nil {
		return nil, errResponse(err)
	}

	return &shttp.Response{
		Data:       res,
		StatusCode: http.StatusCreated,
	}, nil
}

func (c *tagController) getTag(r *http.Request) (*shttp.Response, error) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return nil, errResponse(err)
	}

	res, err := c.service.GetTag(context.Background(), id)
	if err != nil {
		return nil, errResponse(err)
	}

	return &shttp.Response{
		Data:       res,
		StatusCode: http.StatusOK,
	}, nil
}

func (c *tagController) deleteTag(r *http.Request) (*shttp.Response, error) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return nil, errResponse(err)
	}

	if err := c.service.DeleteTag(context.Background(), id); err != nil {
		return nil, errResponse(err)
	}

	return &shttp.Response{
		Data:       "OK",
		StatusCode: http.StatusOK,
	}, nil
}

func (c *tagController) listTag(r *http.Request) (*shttp.Response, error) {
	tags, err := c.service.ListTag(context.Background())
	if err != nil {
		return nil, errResponse(err)
	}

	return &shttp.Response{
		Data:       tags,
		StatusCode: http.StatusOK,
	}, nil
}

func (c *tagController) updateTag(r *http.Request) (*shttp.Response, error) {
	var req request.TagRequest

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return nil, errResponse(err)
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errResponse(err)
	}

	res, err := c.service.UpdateTag(context.Background(), id, req.Name)
	if err != nil {
		return nil, errResponse(err)
	}

	return &shttp.Response{
		Data:       res,
		StatusCode: http.StatusCreated,
	}, nil
}

func NewTagController(router *mux.Router, service tag.Service) *tagController {
	Controller := &tagController{
		service: service,
	}
	Controller.RegisterRoutes(router)
	return Controller
}
