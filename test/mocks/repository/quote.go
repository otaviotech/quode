package repository_test

import (
	"context"

	"github.com/otaviotech/quode/internal/application/repository"
	"github.com/stretchr/testify/mock"
)

type QuoteRepositoryMock struct {
	mock.Mock
}

func (q *QuoteRepositoryMock) Create(ctx context.Context, input repository.CreateQuoteData) error {
	args := q.Called(ctx, input)
	return args.Error(0)
}

func (q *QuoteRepositoryMock) List(ctx context.Context, input repository.ListQuotesData) (*repository.ListQuotesResult, error) {
	args := q.Called(ctx, input)
	return args.Get(0).(*repository.ListQuotesResult), args.Error(1)
}
