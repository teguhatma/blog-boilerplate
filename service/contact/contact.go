package contact

import (
	"context"
	"database/sql"

	fe "github.com/teguhatma/blog-boilerplate/errors"
	"github.com/teguhatma/blog-boilerplate/repository"
	"github.com/teguhatma/blog-boilerplate/request"
	"github.com/teguhatma/blog-boilerplate/response"
)

type Service interface {
	CreateContact(ctx context.Context, req *request.ContactRequest) (*response.ContactResponse, error)
}

type service struct {
	repo repository.Querier
}

func NewService(repo repository.Querier) Service {
	return &service{
		repo: repo,
	}
}

func (service *service) CreateContact(ctx context.Context, req *request.ContactRequest) (*response.ContactResponse, error) {
	arg := repository.CreateContactParams{
		Owner:   "ha",
		Github:  sql.NullString{Valid: true, String: req.Github},
		Twitter: sql.NullString{Valid: true, String: req.Twitter},
	}

	contact, err := service.repo.CreateContact(ctx, arg)
	if err != nil {
		return nil, fe.NewWithCause(fe.INTERNAL_ERROR, err, "Create Contact")
	}

	res := domainToResponse(contact)
	return res, nil
}

func domainToResponse(res repository.Contact) *response.ContactResponse {
	return &response.ContactResponse{
		ID:      int(res.ID),
		Github:  res.Github.String,
		Twitter: res.Twitter.String,
	}
}
