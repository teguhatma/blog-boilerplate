package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/teguhatma/blog-boilerplate/request"
	shttp "github.com/teguhatma/blog-boilerplate/server/http"
	"github.com/teguhatma/blog-boilerplate/service/user"
)

type controller struct {
	service user.Service
}

func (c *controller) RegisterRoutes(router *mux.Router) {
	router.Handle("/api/v1/users/{username}", shttp.AppHandler(c.getUser)).Methods(http.MethodGet)
	router.Handle("/api/v1/users", shttp.AppHandler(c.createUser)).Methods(http.MethodPost)
}

func (c *controller) getUser(r *http.Request) *shttp.Response {
	username := mux.Vars(r)["username"]

	res, err := c.service.GetUser(context.Background(), username)
	if err != nil {
		return errResponse(err)
	}

	return &shttp.Response{
		Data:       res,
		StatusCode: http.StatusOK,
	}
}

func (c *controller) createUser(r *http.Request) *shttp.Response {
	var req request.UserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil
	}

	user, err := c.service.CreateUser(context.Background(), req)
	if err != nil {
		return nil
	}

	return &shttp.Response{
		Data:       user,
		StatusCode: http.StatusCreated,
	}
}

func NewUserController(router *mux.Router, service user.Service) *controller {
	Controller := &controller{
		service: service,
	}
	Controller.RegisterRoutes(router)
	return Controller
}
