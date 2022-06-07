package user

import (
	"context"
	"database/sql"

	fe "github.com/teguhatma/blog-boilerplate/errors"
	"github.com/teguhatma/blog-boilerplate/repository"
	"github.com/teguhatma/blog-boilerplate/request"
	"github.com/teguhatma/blog-boilerplate/response"
	"github.com/teguhatma/blog-boilerplate/token"
	"github.com/teguhatma/blog-boilerplate/utils"
)

type Service interface {
	GetUser(context.Context, string) (*response.UserResponse, error)
	CreateUser(context.Context, request.UserRequest) (*response.UserResponse, error)
	LoginUser(context.Context, request.LoginUserRequest) (*response.LoginUserResponse, error)
	GetUsers(context.Context) ([]*response.UserResponse, error)
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
		if err == sql.ErrNoRows {
			return nil, fe.NewWithCause(fe.NOT_FOUND, err, "User Not Found")
		}
		return nil, fe.NewWithCause(fe.INTERNAL_ERROR, err, "Get User")
	}

	res := domainToResponse(user)
	return res, nil
}

func (service *service) CreateUser(ctx context.Context, request request.UserRequest) (*response.UserResponse, error) {
	req, err := mapToRepository(request)
	if err != nil {
		return nil, fe.NewWithCause(fe.BAD_MESSAGE, err, "Map Request to Domain")
	}

	user, err := service.repo.CreateUser(ctx, *req)
	if err != nil {
		return nil, fe.NewWithCause(fe.INTERNAL_ERROR, err, "Create User")
	}

	res := domainToResponse(user)

	return res, nil
}

func (service *service) LoginUser(ctx context.Context, req request.LoginUserRequest) (*response.LoginUserResponse, error) {
	user, err := service.repo.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fe.NewWithCause(fe.NOT_FOUND, err, "User Not Found")
		}
		return nil, fe.NewWithCause(fe.INTERNAL_ERROR, err, "Get User")
	}

	err = utils.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		return nil, fe.NewWithCause(fe.BAD_MESSAGE, err, "Check Password")
	}

	symmectricKey := utils.GenerateNewID()
	tokenMaker, err := token.NewPasetoMaker(symmectricKey)

	// TODO 1 => use environment variable
	accessToken, err := tokenMaker.CreateToken(user.Username, 15)
	if err != nil {
		return nil, fe.NewWithCause(fe.INTERNAL_ERROR, err, "Create Token")
	}

	res := mapToLoginResponse(accessToken, user)
	return res, nil
}

func (service *service) GetUsers(ctx context.Context) ([]*response.UserResponse, error) {
	users, err := service.repo.ListUsers(ctx)
	if err != nil {
		return nil, fe.NewWithCause(fe.INTERNAL_ERROR, err, "Get Users")
	}

	res := domainToResponses(users)
	return res, nil
}

func mapToLoginResponse(accessToken string, user repository.User) *response.LoginUserResponse {
	userRes := domainToResponse(user)

	return &response.LoginUserResponse{
		AccessToken: accessToken,
		User:        *userRes,
	}
}

func domainToResponse(res repository.User) *response.UserResponse {
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

func domainToResponses(users []repository.User) []*response.UserResponse {
	var response []*response.UserResponse

	for _, user := range users {
		res := domainToResponse(user)
		response = append(response, res)
	}

	return response
}
