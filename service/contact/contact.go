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
	GetContact(ctx context.Context, id int) (*response.ContactResponse, error)
	UpdateContact(ctx context.Context, id int, req *request.ContactRequest) (*response.ContactResponse, error)
	DeleteContact(ctx context.Context, id int) error
	GetAllContact(ctx context.Context) ([]*response.ContactResponse, error)
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
	user, err := service.repo.GetUser(ctx, req.Owner)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fe.NewWithCause(fe.NOT_FOUND, err, "User Not Found")
		}
		return nil, fe.NewWithCause(fe.INTERNAL_ERROR, err, "Get User")
	}

	arg := repository.CreateContactParams{
		Owner:   user.Username,
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

func (service *service) GetContact(ctx context.Context, id int) (*response.ContactResponse, error) {
	contact, err := service.repo.GetContact(ctx, int64(id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fe.NewWithCause(fe.NOT_FOUND, err, "Contact Not Found")
		}
		return nil, fe.NewWithCause(fe.INTERNAL_ERROR, err, "Get Contact")
	}

	res := domainToResponse(contact)
	return res, nil
}

func (service *service) UpdateContact(ctx context.Context, id int, req *request.ContactRequest) (*response.ContactResponse, error) {
	user, err := service.repo.GetUser(ctx, req.Owner)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fe.NewWithCause(fe.NOT_FOUND, err, "User Not Found")
		}
		return nil, fe.NewWithCause(fe.INTERNAL_ERROR, err, "Get User")
	}

	contact, err := service.repo.GetContact(ctx, int64(id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fe.NewWithCause(fe.NOT_FOUND, err, "Contact Not Found")
		}
		return nil, fe.NewWithCause(fe.INTERNAL_ERROR, err, "Get User")
	}

	arg := repository.UpdateContactParams{
		ID:      int64(id),
		Owner:   user.Username,
		Github:  sql.NullString{Valid: true, String: req.Github},
		Twitter: sql.NullString{Valid: true, String: req.Twitter},
	}

	contact, err = service.repo.UpdateContact(ctx, arg)
	if err != nil {
		return nil, fe.NewWithCause(fe.INTERNAL_ERROR, err, "Create Contact")
	}

	res := domainToResponse(contact)
	return res, nil
}

func (service *service) DeleteContact(ctx context.Context, id int) error {
	_, err := service.repo.GetContact(ctx, int64(id))
	if err != nil {
		if err == sql.ErrNoRows {
			return fe.NewWithCause(fe.NOT_FOUND, err, "Contact Not Found")
		}
		return fe.NewWithCause(fe.INTERNAL_ERROR, err, "Get User")
	}

	err = service.repo.DeleteContact(ctx, int64(id))
	if err != nil {
		return fe.NewWithCause(fe.INTERNAL_ERROR, err, "Create Contact")
	}

	return nil
}

func (service *service) GetAllContact(ctx context.Context) ([]*response.ContactResponse, error) {
	contacts, err := service.repo.GetAllContact(ctx)
	if err != nil {
		return nil, fe.NewWithCause(fe.INTERNAL_ERROR, err, "Get All Contact")
	}

	res := domainToResponses(contacts)
	return res, nil
}

func domainToResponse(res repository.Contact) *response.ContactResponse {
	return &response.ContactResponse{
		ID:      int(res.ID),
		Owner:   res.Owner,
		Github:  res.Github.String,
		Twitter: res.Twitter.String,
	}
}

func domainToResponses(contacts []repository.Contact) []*response.ContactResponse {
	var response []*response.ContactResponse

	for _, contact := range contacts {
		res := domainToResponse(contact)
		response = append(response, res)
	}

	return response
}
