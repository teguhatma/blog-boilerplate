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
	router.Handle("/api/v1/user/{username}", shttp.AppHandler(c.getUser)).Methods(http.MethodGet)
	router.Handle("/api/v1/user", shttp.AppHandler(c.createUser)).Methods(http.MethodPost)
	router.Handle("/api/v1/user/login", shttp.AppHandler(c.loginUser)).Methods(http.MethodPost)
	router.Handle("/api/v1/users", shttp.AppHandler(c.getUsers)).Methods(http.MethodGet)
	router.Handle("/api/v1/user/{username}", shttp.AppHandler(c.updateUser)).Methods(http.MethodPut)
}

func (c *controller) getUser(r *http.Request) (*shttp.Response, error) {
	username := mux.Vars(r)["username"]

	res, err := c.service.GetUser(context.Background(), username)
	if err != nil {
		return nil, errResponse(err)
	}

	return &shttp.Response{
		Data:       res,
		StatusCode: http.StatusOK,
	}, nil
}

func (c *controller) createUser(r *http.Request) (*shttp.Response, error) {
	var req request.UserRequest
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errResponse(err)
	}
	
	if err := structValidator(&req); err != nil {
		return nil, errResponse(err)
	}
	
	user, err := c.service.CreateUser(context.Background(), req)
	if err != nil {
		return nil, errResponse(err)
	}

	return &shttp.Response{
		Data:       user,
		StatusCode: http.StatusCreated,
	}, nil
}

func (c *controller) loginUser(r *http.Request) (*shttp.Response, error) {
	var req request.LoginUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errResponse(err)
	}

	user, err := c.service.LoginUser(context.Background(), req)
	if err != nil {
		return nil, errResponse(err)
	}

	return &shttp.Response{
		Data:       user,
		StatusCode: http.StatusOK,
	}, nil
}

func (c *controller) getUsers(r *http.Request) (*shttp.Response, error) {
	users, err := c.service.GetUsers(context.Background())
	if err != nil {
		return nil, errResponse(err)
	}

	return &shttp.Response{
		Data:       users,
		StatusCode: http.StatusOK,
	}, nil
}

func (c *controller) updateUser(r *http.Request) (*shttp.Response, error) {
	var req request.UpdateUserRequest
	username := mux.Vars(r)["username"]

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errResponse(err)
	}

	id, err := c.service.UpdateUser(context.Background(), username, req)
	if err != nil {
		return nil, errResponse(err)
	}

	return &shttp.Response{
		Data:       id,
		StatusCode: http.StatusOK,
	}, nil
}

func NewUserController(router *mux.Router, service user.Service) *controller {
	Controller := &controller{
		service: service,
	}
	Controller.RegisterRoutes(router)
	return Controller
}
