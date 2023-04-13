package usecase

import (
	"context"

	"github.com/otaviotech/quode/internal/application/repository"
	"github.com/otaviotech/quode/pkg/dbutil"
)

type ListQuotesInput struct {
	Pagination dbutil.Pagination
}

type ListQuotesOutputItem struct {
	ID       string
	BookID   string
	BookName string
	Page     int
	Content  string
}

type ListQuotesOutput = dbutil.PaginatedResult[ListQuotesOutputItem]

type ListQuotesUseCaseInterface interface {
	Execute(ctx context.Context, input ListQuotesInput) (*ListQuotesOutput, error)
}

type ListQuotesUseCase struct {
	quoteRepository repository.QuoteRepositoryInterface
}

func NewListQuoteRepository(qr repository.QuoteRepositoryInterface) *ListQuotesUseCase {
	return &ListQuotesUseCase{quoteRepository: qr}
}

func (u *ListQuotesUseCase) Execute(ctx context.Context, input ListQuotesInput) (*ListQuotesOutput, error) {
	quotes, err := u.quoteRepository.List(ctx, repository.ListQuotesData(input))

	if err != nil {
		return nil, err
	}

	var output ListQuotesOutput
	output.Total = quotes.Total

	for _, q := range quotes.Data {
		output.Data = append(output.Data, ListQuotesOutputItem(q))
	}

	return &output, nil
}
