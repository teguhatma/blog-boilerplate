// Code generated by sqlc. DO NOT EDIT.
// source: entries.sql

package repository

import (
	"context"
)

const createEntries = `-- name: CreateEntries :one
INSERT INTO entries (
    owner,
    tag_name,
    blog,
    title,
    read_time
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING id, owner, tag_name, blog, title, read_time, updated_at, created_at
`

type CreateEntriesParams struct {
	Owner    string `json:"owner"`
	TagName  string `json:"tag_name"`
	Blog     string `json:"blog"`
	Title    string `json:"title"`
	ReadTime string `json:"read_time"`
}

func (q *Queries) CreateEntries(ctx context.Context, arg CreateEntriesParams) (Entry, error) {
	row := q.queryRow(ctx, q.createEntriesStmt, createEntries,
		arg.Owner,
		arg.TagName,
		arg.Blog,
		arg.Title,
		arg.ReadTime,
	)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.TagName,
		&i.Blog,
		&i.Title,
		&i.ReadTime,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const deleteEntries = `-- name: DeleteEntries :exec
DELETE FROM entries
WHERE id = $1
`

func (q *Queries) DeleteEntries(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteEntriesStmt, deleteEntries, id)
	return err
}

const getEntries = `-- name: GetEntries :one
SELECT id, owner, tag_name, blog, title, read_time, updated_at, created_at FROM entries
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetEntries(ctx context.Context, id int64) (Entry, error) {
	row := q.queryRow(ctx, q.getEntriesStmt, getEntries, id)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.TagName,
		&i.Blog,
		&i.Title,
		&i.ReadTime,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const listEntries = `-- name: ListEntries :many
SELECT id, owner, tag_name, blog, title, read_time, updated_at, created_at FROM entries
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListEntriesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error) {
	rows, err := q.query(ctx, q.listEntriesStmt, listEntries, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Entry{}
	for rows.Next() {
		var i Entry
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.TagName,
			&i.Blog,
			&i.Title,
			&i.ReadTime,
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

const updateEntries = `-- name: UpdateEntries :one
UPDATE entries
SET owner = $2, blog = $3, title = $4, read_time = $5
WHERE id = $1
RETURNING id, owner, tag_name, blog, title, read_time, updated_at, created_at
`

type UpdateEntriesParams struct {
	ID       int64  `json:"id"`
	Owner    string `json:"owner"`
	Blog     string `json:"blog"`
	Title    string `json:"title"`
	ReadTime string `json:"read_time"`
}

func (q *Queries) UpdateEntries(ctx context.Context, arg UpdateEntriesParams) (Entry, error) {
	row := q.queryRow(ctx, q.updateEntriesStmt, updateEntries,
		arg.ID,
		arg.Owner,
		arg.Blog,
		arg.Title,
		arg.ReadTime,
	)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.TagName,
		&i.Blog,
		&i.Title,
		&i.ReadTime,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}
