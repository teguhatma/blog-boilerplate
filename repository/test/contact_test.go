package test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/teguhatma/blog-boilerplate/repository"
	"github.com/teguhatma/blog-boilerplate/utils"
)

func createRandomContact(t *testing.T) repository.Contact {
	user := createRandomUser(t)
	arg := repository.CreateContactParams{
		Owner:   user.Username,
		Github:  sql.NullString{String: utils.RandomEmail(), Valid: true},
		Twitter: sql.NullString{String: utils.RandomEmail(), Valid: true},
	}
	contact, err := testQueries.CreateContact(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, contact)

	require.Equal(t, arg.Owner, contact.Owner)
	require.Equal(t, arg.Github, contact.Github)
	require.Equal(t, arg.Twitter, contact.Twitter)

	require.True(t, contact.UpdatedAt.IsZero())
	require.NotZero(t, contact.CreatedAt)

	return contact
}

func TestCreateContact(t *testing.T) {
	createRandomContact(t)
}

func TestDeleteContact(t *testing.T) {
	contact := createRandomContact(t)

	err := testQueries.DeleteContact(context.Background(), contact.ID)
	require.NoError(t, err)
}

func TestGetContact(t *testing.T) {
	contact1 := createRandomContact(t)

	contact2, err := testQueries.GetContact(context.Background(), contact1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, contact2)

	require.Equal(t, contact1.ID, contact2.ID)
	require.Equal(t, contact1.Owner, contact2.Owner)
	require.Equal(t, contact1.Github, contact2.Github)
	require.Equal(t, contact1.Twitter, contact2.Twitter)

	require.WithinDuration(t, contact1.UpdatedAt, contact2.UpdatedAt, time.Second)
	require.WithinDuration(t, contact1.CreatedAt, contact2.CreatedAt, time.Second)
}

func TestUpdateContact(t *testing.T) {
	contact1 := createRandomContact(t)

	arg := repository.UpdateContactParams{
		ID:      contact1.ID,
		Owner:   contact1.Owner,
		Github:  sql.NullString{String: utils.RandomEmail(), Valid: true},
		Twitter: sql.NullString{String: utils.RandomEmail(), Valid: true},
	}

	contact2, err := testQueries.UpdateContact(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, contact2)

	require.Equal(t, contact1.ID, contact2.ID)
	require.Equal(t, contact1.Owner, contact2.Owner)

	require.NotEqual(t, contact1.Github, contact2.Github)
	require.NotEqual(t, contact1.Twitter, contact2.Twitter)

	require.WithinDuration(t, contact1.UpdatedAt, contact2.UpdatedAt, time.Second)
	require.WithinDuration(t, contact1.CreatedAt, contact2.CreatedAt, time.Second)
}
