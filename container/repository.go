package container

import (
	"github.com/teguhatma/blog-boilerplate/repository"
)

var userRepository repository.Querier

func getUserRepository() (repository.Querier, error) {
	if userRepository == nil {
		getDB := getDatabase()
		userRepository = repository.New(getDB)
	}
	return userRepository, nil
}
