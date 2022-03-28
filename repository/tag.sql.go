// Code generated by sqlc. DO NOT EDIT.
// source: tag.sql

package db

import (
	"context"
)

const createTag = `-- name: CreateTag :one
INSERT INTO tag (
   name 
) VALUES (
    $1
) RETURNING id, name, updated_at, created_at
`

func (q *Queries) CreateTag(ctx context.Context, name string) (Tag, error) {
	row := q.queryRow(ctx, q.createTagStmt, createTag, name)
	var i Tag
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const deleteTag = `-- name: DeleteTag :exec
DELETE FROM tag
WHERE id = $1
`

func (q *Queries) DeleteTag(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteTagStmt, deleteTag, id)
	return err
}

const listTag = `-- name: ListTag :many
SELECT id, name, updated_at, created_at FROM tag
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListTagParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListTag(ctx context.Context, arg ListTagParams) ([]Tag, error) {
	rows, err := q.query(ctx, q.listTagStmt, listTag, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Tag{}
	for rows.Next() {
		var i Tag
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.UpdatedAt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTag = `-- name: UpdateTag :one
UPDATE tag
SET name = $2
WHERE id = $1
RETURNING id, name, updated_at, created_at
`

type UpdateTagParams struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdateTag(ctx context.Context, arg UpdateTagParams) (Tag, error) {
	row := q.queryRow(ctx, q.updateTagStmt, updateTag, arg.ID, arg.Name)
	var i Tag
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}
