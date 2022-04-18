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
	router.Handle("/api/v1/tags", shttp.AppHandler(c.createTag)).Methods(http.MethodPost)
	router.Handle("/api/v1/tags/{id:[0-9]+}", shttp.AppHandler(c.getTag)).Methods(http.MethodGet)
	router.Handle("/api/v1/tags/{id:[0-9]+}", shttp.AppHandler(c.deleteTag)).Methods(http.MethodDelete)
	router.Handle("/api/v1/tags", shttp.AppHandler(c.listTag)).Queries("limit", "{limit}").Queries("offset", "{offset}").Methods(http.MethodGet)
	router.Handle("/api/v1/tags/{id:[0-9]+}", shttp.AppHandler(c.updateTag)).Methods(http.MethodPut)
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

	res, err := c.service.GetTag(context.Background(), int64(id))
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

	if err := c.service.DeleteTag(context.Background(), int64(id)); err != nil {
		return nil, errResponse(err)
	}

	return &shttp.Response{
		Data:       "OK",
		StatusCode: http.StatusOK,
	}, nil
}

func (c *tagController) listTag(r *http.Request) (*shttp.Response, error) {
	a := r.FormValue("limit")
	b := r.FormValue("offset")

	limit, offset, err := convertToInt32(a, b)
	if err != nil {
		return nil, errResponse(err)
	}

	tags, err := c.service.ListTag(context.Background(), limit, offset)
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

	res, err := c.service.UpdateTag(context.Background(), int64(id), req.Name)
	if err != nil {
		return nil, errResponse(err)
	}

	return &shttp.Response{
		Data:       res,
		StatusCode: http.StatusCreated,
	}, nil
}

func convertToInt32(l, o string) (int32, int32, error) {
	limit, err := strconv.Atoi(l)
	if err != nil {
		return -1, -1, err
	}
	offset, err := strconv.Atoi(o)
	if err != nil {
		return -1, -1, err
	}

	a := int32(limit)
	b := int32(offset)

	return a, b, nil
}

func NewTagController(router *mux.Router, service tag.Service) *tagController {
	Controller := &tagController{
		service: service,
	}
	Controller.RegisterRoutes(router)
	return Controller
}
