package container

import (
	"github.com/gorilla/mux"
	c "github.com/teguhatma/blog-boilerplate/controller"
)

func getUserController(router *mux.Router) error {
	service, err := getUserService()
	if err != nil {
		return err
	}

	c.NewUserController(router, service)
	return nil
}

func getTagController(router *mux.Router) error {
	service, err := getTagService()
	if err != nil {
		return err
	}

	c.NewTagController(router, service)
	return nil
}
