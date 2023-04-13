// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: author.sql

package db_gen

import (
	"context"
	"time"
)

const createAuthor = `-- name: CreateAuthor :exec
INSERT INTO authors (id, name, bio, created_at) VALUES ($1, $2, $3, $4)
`

type CreateAuthorParams struct {
	ID        string
	Name      string
	Bio       string
	CreatedAt time.Time
}

func (q *Queries) CreateAuthor(ctx context.Context, arg CreateAuthorParams) error {
	_, err := q.db.ExecContext(ctx, createAuthor,
		arg.ID,
		arg.Name,
		arg.Bio,
		arg.CreatedAt,
	)
	return err
}