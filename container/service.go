package container

import (
	"github.com/teguhatma/blog-boilerplate/service/user"
)

var userService user.Service

func getUserService() (user.Service, error) {
	if userService == nil {
		repo, err := getUserRepository()
		if err != nil {
			return nil, err
		}
		userService = user.NewService(repo)
	}
	return userService, nil
}
