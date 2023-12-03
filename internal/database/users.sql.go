// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: users.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, name, api_key)
VALUES (
    $1,
    $2,
    encode(sha256(random()::text::bytea), 'hex')
  )
RETURNING id, name, created_at, updated_at, api_key
`

type CreateUserParams struct {
	ID   uuid.UUID
	Name string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.ID, arg.Name)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ApiKey,
	)
	return i, err
}
