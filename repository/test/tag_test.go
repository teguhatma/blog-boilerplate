package test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/teguhatma/blog-boilerplate/repository"
	"github.com/teguhatma/blog-boilerplate/utils"
)

func createRandomTag(t *testing.T) repository.Tag {
	name := utils.RandomString(5)
	tag, err := testQueries.CreateTag(context.Background(), name)
	require.NoError(t, err)
	require.NotEmpty(t, tag)

	require.Equal(t, name, tag.Name)

	require.True(t, tag.UpdatedAt.IsZero())
	require.NotZero(t, tag.CreatedAt)

	return tag
}

func TestCreateTag(t *testing.T) {
	createRandomTag(t)
}

func TestListTag(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomTag(t)
	}

	arg := repository.ListTagParams{
		Limit:  5,
		Offset: 0,
	}

	tags, err := testQueries.ListTag(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, tags, 5)

	for _, tag := range tags {
		require.NotEmpty(t, tag)
	}
}

func TestGetTag(t *testing.T) {
	tag1 := createRandomTag(t)

	tag2, err := testQueries.GetTag(context.Background(), tag1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, tag2)

	require.Equal(t, tag1.Name, tag2.Name)

	require.WithinDuration(t, tag1.UpdatedAt, tag2.UpdatedAt, time.Second)
	require.WithinDuration(t, tag1.CreatedAt, tag2.CreatedAt, time.Second)
}

func TestUpdateTag(t *testing.T) {
	tag1 := createRandomTag(t)

	arg := repository.UpdateTagParams{
		ID:   tag1.ID,
		Name: utils.RandomUsername(),
	}

	tag2, err := testQueries.UpdateTag(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, tag2)

	require.Equal(t, tag1.ID, tag2.ID)
	require.NotEqual(t, tag1.Name, tag2.Name)

	require.WithinDuration(t, tag1.UpdatedAt, tag2.UpdatedAt, time.Second)
	require.WithinDuration(t, tag1.CreatedAt, tag2.CreatedAt, time.Second)
}

func TestDeleteTag(t *testing.T) {
	tag := createRandomTag(t)

	err := testQueries.DeleteTag(context.Background(), tag.ID)
	require.NoError(t, err)
}
