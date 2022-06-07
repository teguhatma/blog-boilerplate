// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: users.sql

package repository

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    username,
    hashed_password,
    full_name,
    email
) VALUES (
    $1, $2, $3, $4
) RETURNING id, username, hashed_password, full_name, email, updated_at, created_at
`

type CreateUserParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.queryRow(ctx, q.createUserStmt, createUser,
		arg.Username,
		arg.HashedPassword,
		arg.FullName,
		arg.Email,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, username, hashed_password, full_name, email, updated_at, created_at FROM users 
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.queryRow(ctx, q.getUserStmt, getUser, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, username, hashed_password, full_name, email, updated_at, created_at FROM users
ORDER BY id
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.query(ctx, q.listUsersStmt, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.HashedPassword,
			&i.FullName,
			&i.Email,
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
