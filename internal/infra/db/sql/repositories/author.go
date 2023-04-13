package repositories

import (
	"context"
	"database/sql"

	"github.com/otaviotech/quode/internal/application/repository"
	"github.com/otaviotech/quode/internal/infra/db/sql/db_gen"
)

type AuthorRepository struct {
	q *db_gen.Queries
}

func NewAuthorRepository(db *sql.DB) *AuthorRepository {
	return &AuthorRepository{
		q: db_gen.New(db),
	}
}

func (r *AuthorRepository) Create(ctx context.Context, input repository.CreateAuthorData) error {
	return r.q.CreateAuthor(ctx, db_gen.CreateAuthorParams{
		ID:        input.ID,
		Name:      input.Name,
		Bio:       input.Bio,
		CreatedAt: input.CreatedAt,
	})
}
