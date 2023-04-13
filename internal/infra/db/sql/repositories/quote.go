package repositories

import (
	"context"
	"database/sql"

	"github.com/otaviotech/quode/internal/application/repository"
	"github.com/otaviotech/quode/internal/infra/db/sql/db_gen"
)

type QuoteRepository struct {
	q *db_gen.Queries
}

func NewQuoteRepository(db *sql.DB) *QuoteRepository {
	return &QuoteRepository{
		q: db_gen.New(db),
	}
}

func (r *QuoteRepository) Create(ctx context.Context, data repository.CreateQuoteData) error {
	return r.q.CreateQuote(ctx, db_gen.CreateQuoteParams{
		ID:        data.ID,
		BookID:    data.BookID,
		Page:      int32(data.Page),
		Content:   data.Content,
		CreatedAt: data.CreatedAt,
	})
}

func (r *QuoteRepository) List(ctx context.Context, data repository.ListQuotesData) (*repository.ListQuotesResult, error) {
	quotes, err := r.q.ListQuotes(ctx, db_gen.ListQuotesParams{
		Limit:  int32(data.Pagination.Limit),
		Offset: int32(data.Pagination.Offset),
	})

	if err != nil {
		return nil, err
	}

	var result repository.ListQuotesResult

	count := len(quotes)

	if count > 0 {
		count = int(quotes[0].FullCount)
	}

	result.Total = count

	for _, q := range quotes {
		result.Data = append(result.Data, repository.ListQuotesResultItem{
			ID:       q.ID,
			BookID:   q.BookID,
			BookName: q.BookTitle,
			Page:     int(q.Page),
			Content:  q.Content,
		})
	}

	return &result, nil
}
