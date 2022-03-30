package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/teguhatma/blog-boilerplate/cmd"
	"github.com/teguhatma/blog-boilerplate/request"
	"github.com/teguhatma/blog-boilerplate/service/user"
)

type controller struct {
	service user.Service
}

func (c *controller) RegisterRoutes(router *mux.Router) {
	router.Handle("/api/v1/users/{username}", cmd.AppHandler(c.getUser)).Methods(http.MethodGet)
	router.Handle("/api/v1/users", cmd.AppHandler(c.createUser)).Methods(http.MethodPost)
}

func (c *controller) getUser(r *http.Request) (*cmd.Response, error) {
	username := mux.Vars(r)["username"]

	res, err := c.service.GetUser(context.Background(), username)
	if err != nil {
		return nil, err
	}

	return &cmd.Response{
		Data:       res,
		StatusCode: http.StatusCreated,
	}, nil
}

func (c *controller) createUser(r *http.Request) (*cmd.Response, error) {
	var req request.UserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	user, err := c.service.CreateUser(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &cmd.Response{
		Data:       user,
		StatusCode: http.StatusCreated,
	}, nil
}

func NewUserController(router *mux.Router, service user.Service) *controller {
	Controller := &controller{
		service: service,
	}
	Controller.RegisterRoutes(router)
	return Controller
}
