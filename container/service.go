package container

import (
	"github.com/teguhatma/blog-boilerplate/service/tag"
	"github.com/teguhatma/blog-boilerplate/service/user"
)

var userService user.Service
var tagService tag.Service

func getUserService() (user.Service, error) {
	if userService == nil {
		repo, err := getRepository()
		if err != nil {
			return nil, err
		}
		userService = user.NewService(repo)
	}
	return userService, nil
}

func getTagService() (tag.Service, error) {
	if tagService == nil {
		repo, err := getRepository()
		if err != nil {
			return nil, err
		}
		tagService = tag.NewService(repo)
	}
	return tagService, nil
}
