package repositories

import (
	"context"
	"database/sql"

	"github.com/otaviotech/quode/internal/application/repository"
	"github.com/otaviotech/quode/internal/infra/db/sql/db_gen"
)

type SubjectRepository struct {
	q *db_gen.Queries
}

func NewSubjectRepository(db *sql.DB) *SubjectRepository {
	return &SubjectRepository{
		q: db_gen.New(db),
	}
}

func (r *SubjectRepository) Create(ctx context.Context, input repository.CreateSubjectData) error {
	return r.q.CreateSubject(ctx, db_gen.CreateSubjectParams{
		ID:          input.ID,
		Name:        input.Name,
		Description: input.Description,
		CreatedAt:   input.CreatedAt,
	})
}
