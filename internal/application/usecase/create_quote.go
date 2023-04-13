package usecase

import (
	"context"

	"github.com/otaviotech/quode/internal/application/repository"
	"github.com/otaviotech/quode/internal/entity/entities"
)

type CreateQuoteInput struct {
	BookID  string
	Page    int
	Content string
}

type CreateQuoteUseCaseInterface interface {
	Execute(ctx context.Context, input CreateQuoteInput) error
}

type CreateQuoteUseCase struct {
	quoteRepo repository.QuoteRepositoryInterface
}

func NewCreateQuoteUseCase(quoteRepo repository.QuoteRepositoryInterface) *CreateQuoteUseCase {
	return &CreateQuoteUseCase{
		quoteRepo: quoteRepo,
	}
}

func (u *CreateQuoteUseCase) Execute(ctx context.Context, input CreateQuoteInput) error {
	q, err := entities.NewQuote(input.BookID, input.Content, input.Page)

	if err != nil {
		return err
	}

	return u.quoteRepo.Create(ctx, repository.CreateQuoteData{
		ID:        q.ID.Value,
		BookID:    q.BookID.Value,
		Page:      q.Page,
		Content:   q.Content,
		CreatedAt: q.CreatedAt,
	})
}
