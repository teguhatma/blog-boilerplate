package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/teguhatma/blog-boilerplate/repository"
	"github.com/teguhatma/blog-boilerplate/request"
	"github.com/teguhatma/blog-boilerplate/response"
	"github.com/teguhatma/blog-boilerplate/utils"
)

type Service interface {
	GetUser(context.Context, string) (*response.UserResponse, error)
	CreateUser(context.Context, request.UserRequest) (*response.UserResponse, error)
}

type service struct {
	repo repository.Querier
}

func NewService(repo repository.Querier) Service {
	return &service{
		repo: repo,
	}
}

func (service *service) GetUser(ctx context.Context, username string) (*response.UserResponse, error) {
	user, err := service.repo.GetUser(ctx, username)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	res := mapToResponse(user)
	return res, nil
}

func (service *service) CreateUser(ctx context.Context, request request.UserRequest) (*response.UserResponse, error) {
	req, err := mapToRepository(request)
	if err != nil {
		return nil, err
	}

	user, err := service.repo.CreateUser(ctx, *req)
	if err != nil {
		return nil, err
	}

	res := mapToResponse(user)

	return res, nil
}

func mapToResponse(res repository.User) *response.UserResponse {
	return &response.UserResponse{
		ID:        res.ID.Int64,
		Username:  res.Username,
		FullName:  res.FullName,
		Email:     res.Email,
		UpdatedAt: res.UpdatedAt,
		CreatedAt: res.CreatedAt,
	}
}

func mapToRepository(req request.UserRequest) (*repository.CreateUserParams, error) {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	return &repository.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
	}, nil
}
