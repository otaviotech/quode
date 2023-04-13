package repository

import (
	"context"
	"time"
)

type CreateAuthorData struct {
	ID        string
	Name      string
	Bio       string
	CreatedAt time.Time
}

type AuthorRepositoryInterface interface {
	Create(ctx context.Context, input CreateAuthorData) error
}
