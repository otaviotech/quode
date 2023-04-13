package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/otaviotech/quode/internal/application/repository"
	"github.com/otaviotech/quode/internal/infra/db/sql/db_gen"
)

type BookRepository struct {
	db *sql.DB
	q  *db_gen.Queries
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{
		db: db,
		q:  db_gen.New(db),
	}
}

func (r *BookRepository) Create(ctx context.Context, data repository.CreateBookData) error {
	tx, err := r.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	qtx := r.q.WithTx(tx)

	err = qtx.CreateBook(ctx, db_gen.CreateBookParams{
		ID:        data.ID,
		Title:     data.Title,
		Isbn:      data.ISBN,
		Year:      int32(data.Year),
		Pages:     int32(data.Pages),
		CreatedAt: data.CreatedAt,
	})

	if err != nil {
		rbErr := tx.Rollback()

		if rbErr != nil {
			return fmt.Errorf("error creating book: %v, error rolling back: %v", err, rbErr)
		}

		return err
	}

	err = r.connectAuthors(ctx, tx, qtx, data.ID, data.Authors)
	if err != nil {
		return err
	}

	err = r.connectSubjects(ctx, tx, qtx, data.ID, data.Subjects)
	if err != nil {
		return err
	}

	err = tx.Commit()

	if err != nil {
		return err
	}

	return nil
}

func (r *BookRepository) connectAuthors(ctx context.Context, tx *sql.Tx, qtx *db_gen.Queries, bookID string, authors []string) error {
	for _, aid := range authors {
		err := qtx.ConnectBookAuthor(ctx, db_gen.ConnectBookAuthorParams{
			BookID:   bookID,
			AuthorID: aid,
		})

		if err != nil {
			rbErr := tx.Rollback()

			if rbErr != nil {
				return fmt.Errorf("error connecting book to author: %v, error rolling back: %v", err, rbErr)
			}

			return err
		}
	}

	return nil
}

func (r *BookRepository) connectSubjects(ctx context.Context, tx *sql.Tx, qtx *db_gen.Queries, bookID string, subjects []string) error {
	for _, sid := range subjects {
		err := qtx.ConnectBookSubject(ctx, db_gen.ConnectBookSubjectParams{
			BookID:    bookID,
			SubjectID: sid,
		})

		if err != nil {
			rbErr := tx.Rollback()

			if rbErr != nil {
				return fmt.Errorf("error connecting book to subject: %v, error rolling back: %v", err, rbErr)
			}

			return err
		}
	}

	return nil
}
