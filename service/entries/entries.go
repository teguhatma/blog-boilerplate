package entries

import (
	"context"
	"database/sql"

	fe "github.com/teguhatma/blog-boilerplate/errors"
	"github.com/teguhatma/blog-boilerplate/repository"
	"github.com/teguhatma/blog-boilerplate/request"
	"github.com/teguhatma/blog-boilerplate/response"
)

type Service interface {
	CreateEntry(context.Context, request.CreateEntryRequest) (*response.EntryResponse, error)
}

type service struct {
	repo repository.Querier
}

func NewService(repo repository.Querier) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateEntry(ctx context.Context, req request.CreateEntryRequest) (*response.EntryResponse, error) {
	user, err := s.repo.GetUser(ctx, req.Owner)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fe.NewWithCause(fe.NOT_FOUND, err, "User Not Found")
		}
		return nil, fe.NewWithCause(fe.INTERNAL_ERROR, err, "Get User")
	}

	tag, err := s.repo.GetTag(ctx, req.TagName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fe.NewWithCause(fe.NOT_FOUND, err, "Tag Not Found")
		}
		return nil, fe.NewWithCause(fe.INTERNAL_ERROR, err, "Get Tag")
	}

	arg := repository.CreateEntriesParams{
		Owner:    user.Username,
		TagName:  tag.Name,
		Blog:     req.Blog,
		Title:    req.Title,
		ReadTime: req.ReadTime,
	}

	entry, err := s.repo.CreateEntries(ctx, arg)
	if err != nil {
		return nil, fe.NewWithCause(fe.INTERNAL_ERROR, err, "Create Entry")
	}

	res := domainToResponse(entry)
	return res, nil
}

func domainToResponse(repo repository.Entry) *response.EntryResponse {
	return &response.EntryResponse{
		Owner:    repo.Owner,
		TagName:  repo.TagName,
		Blog:     repo.Blog,
		Title:    repo.Title,
		ReadTime: repo.ReadTime,
	}
}
