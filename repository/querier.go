// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package repository

import (
	"context"
)

type Querier interface {
	CreateContact(ctx context.Context, arg CreateContactParams) (Contact, error)
	CreateEntries(ctx context.Context, arg CreateEntriesParams) (Entry, error)
	CreateTag(ctx context.Context, name string) (Tag, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteContact(ctx context.Context, id int64) error
	DeleteEntries(ctx context.Context, id int64) error
	DeleteTag(ctx context.Context, id int64) error
	GetContact(ctx context.Context, id int64) (Contact, error)
	GetEntries(ctx context.Context, id int64) (Entry, error)
	GetTag(ctx context.Context, id int64) (Tag, error)
	GetUser(ctx context.Context, username string) (User, error)
	GetUsers(ctx context.Context, arg GetUsersParams) ([]User, error)
	ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error)
	ListTags(ctx context.Context, arg ListTagsParams) ([]Tag, error)
	UpdateContact(ctx context.Context, arg UpdateContactParams) (Contact, error)
	UpdateEntries(ctx context.Context, arg UpdateEntriesParams) (Entry, error)
	UpdateTag(ctx context.Context, arg UpdateTagParams) (Tag, error)
}

var _ Querier = (*Queries)(nil)
