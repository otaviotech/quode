package repository

import (
	"context"
	"time"
)

type CreateSubjectData struct {
	ID          string
	Name        string
	Description string
	CreatedAt   time.Time
}

type SubjectRepositoryInterface interface {
	Create(ctx context.Context, input CreateSubjectData) error
}
