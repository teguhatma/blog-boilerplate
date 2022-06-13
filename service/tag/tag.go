package tag

import (
	"context"
	"database/sql"

	fe "github.com/teguhatma/blog-boilerplate/errors"
	"github.com/teguhatma/blog-boilerplate/repository"
	"github.com/teguhatma/blog-boilerplate/response"
)

type Service interface {
	CreateTag(ctx context.Context, name string) (*response.TagResponse, error)
	DeleteTag(ctx context.Context, id int) error
	GetTag(ctx context.Context, id int) (*response.TagResponse, error)
	ListTag(ctx context.Context) ([]*response.TagResponse, error)
	UpdateTag(ctx context.Context, id int, name string) (*response.TagResponse, error)
}

type service struct {
	repo repository.Querier
}

func NewService(repo repository.Querier) Service {
	return &service{
		repo: repo,
	}
}

func (service *service) CreateTag(ctx context.Context, name string) (*response.TagResponse, error) {
	tag, err := service.repo.CreateTag(ctx, name)
	if err != nil {
		return nil, fe.NewWithCause(fe.INTERNAL_ERROR, err, "Create Tag")
	}
	res := domainToResponse(tag)
	return res, nil
}

func (service *service) DeleteTag(ctx context.Context, id int) error {
	err := service.repo.DeleteTag(ctx, int64(id))
	if err != nil {
		return fe.NewWithCause(fe.INTERNAL_ERROR, err, "Delete Tag")
	}

	return nil
}

func (service *service) GetTag(ctx context.Context, id int) (*response.TagResponse, error) {
	tag, err := service.repo.GetTag(ctx, int64(id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fe.NewWithCause(fe.NOT_FOUND, err, "Tag Not Found")
		}
		return nil, fe.NewWithCause(fe.INTERNAL_ERROR, err, "Get Tag")
	}

	res := domainToResponse(tag)
	return res, nil
}

func (service *service) ListTag(ctx context.Context) ([]*response.TagResponse, error) {
	tags, err := service.repo.ListTags(ctx)
	if err != nil {
		return nil, fe.NewWithCause(fe.INTERNAL_ERROR, err, "List Tag")
	}

	res := domainToResponses(tags)

	return res, nil
}

func (service *service) UpdateTag(ctx context.Context, id int, name string) (*response.TagResponse, error) {
	arg := repository.UpdateTagParams{
		ID:   int64(id),
		Name: name,
	}

	tag, err := service.repo.UpdateTag(ctx, arg)
	if err != nil {
		return nil, fe.NewWithCause(fe.INTERNAL_ERROR, err, "Update Tag")
	}

	res := domainToResponse(tag)
	return res, nil
}

func domainToResponse(res repository.Tag) *response.TagResponse {
	return &response.TagResponse{
		ID:        res.ID,
		Name:      res.Name,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}
}

func domainToResponses(tags []repository.Tag) []*response.TagResponse {
	var response []*response.TagResponse

	for _, tag := range tags {
		res := domainToResponse(tag)
		response = append(response, res)
	}

	return response
}
