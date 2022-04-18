package container

import (
	"github.com/gorilla/mux"
)

func CreateHTTPContainer(router *mux.Router) error {
	if err := getUserController(router); err != nil {
		return err
	}

	if err := getTagController(router); err != nil {
		return err
	}

	return nil
}
