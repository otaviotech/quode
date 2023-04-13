package repository

import (
	"context"
	"time"

	"github.com/otaviotech/quode/pkg/dbutil"
)

type QuoteRepositoryInterface interface {
	Create(ctx context.Context, data CreateQuoteData) error
	List(ctx context.Context, data ListQuotesData) (*ListQuotesResult, error)
}

type CreateQuoteData struct {
	ID        string
	BookID    string
	Page      int
	Content   string
	CreatedAt time.Time
}

type ListQuotesData struct {
	Pagination dbutil.Pagination
}

type ListQuotesResultItem struct {
	ID       string
	BookID   string
	BookName string
	Page     int
	Content  string
}

type ListQuotesResult = dbutil.PaginatedResult[ListQuotesResultItem]
