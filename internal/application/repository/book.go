package repository

import (
	"context"
	"time"
)

type CreateBookData struct {
	ID        string
	ISBN      string
	Title     string
	Authors   []string
	Subjects  []string
	Year      int
	Pages     int
	CreatedAt time.Time
}

type BookRepositoryInterface interface {
	Create(ctx context.Context, data CreateBookData) error
}
