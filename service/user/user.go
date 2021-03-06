package user

import (
	"context"
	"database/sql"
	"fmt"

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
	UpdateUser(context.Context, string, request.UpdateUserRequest) (*response.UserResponse, error)
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
	// var usr repository.User 
	req, err := mapToRepository(request)
	if err != nil {
		return nil, fe.NewWithCause(fe.BAD_MESSAGE, err, "Map Request to Domain")
	}

	users, _ := service.repo.ListUsers(ctx)
	if err := checkDuplicateUser(users, request); err != nil {
		return nil, fe.NewWithCause(fe.DUPLICATE, err, fmt.Sprintf("cannot duplicate %s", err.Error()))
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
	if err != nil {
		return nil, fe.NewWithCause(fe.INTERNAL_ERROR, err, "Create Token") 
	}

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

func (service *service) UpdateUser(ctx context.Context, username string, request request.UpdateUserRequest) (*response.UserResponse, error) {
	user, err := service.repo.GetUser(ctx, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fe.NewWithCause(fe.NOT_FOUND, err, "User Not Found")
		}
		return nil, fe.NewWithCause(fe.INTERNAL_ERROR, err, "Get User")
	}

	args := repository.UpdateUserParams{
		ID:       user.ID,
		Username: request.Username,
		FullName: request.FullName,
		Email:    request.Email,
	}

	user, err = service.repo.UpdateUser(ctx, args)
	if err != nil {
		return nil, fe.NewWithCause(fe.INTERNAL_ERROR, err, "Update User")
	}

	res := domainToResponse(user)
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

func checkDuplicateUser(users []repository.User, req request.UserRequest) error {
	for _, user := range users {
		if user.Username == req.Username {
			return fmt.Errorf("username")
		}

		if user.Email == req.Email {
			return fmt.Errorf("email")	
		}
	}
	return nil
}