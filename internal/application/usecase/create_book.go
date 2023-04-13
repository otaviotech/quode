package usecase

import (
	"context"
	"time"

	"github.com/otaviotech/quode/internal/application/repository"
	"github.com/otaviotech/quode/internal/entity/entities"
	"github.com/otaviotech/quode/internal/entity/value_objects"
)

type CreateBookInput struct {
	ISBN      string
	Title     string
	Authors   []string
	Subjects  []string
	Year      int
	Pages     int
	CreatedAt time.Time
}

type CreateBookUseCaseInterface interface {
	Execute(ctx context.Context, input CreateBookInput) error
}

type CreateBookUseCase struct {
	BookRepository repository.BookRepositoryInterface
}

func NewCreateBookUseCase(bookRepository repository.BookRepositoryInterface) CreateBookUseCaseInterface {
	return &CreateBookUseCase{
		BookRepository: bookRepository,
	}
}

func (c *CreateBookUseCase) Execute(ctx context.Context, input CreateBookInput) error {
	b, err := entities.NewBook(
		input.Title,
		input.ISBN,
		input.Year,
		input.Pages,
		input.Authors,
		input.Subjects,
	)

	if err != nil {
		return err
	}

	return c.BookRepository.Create(ctx, repository.CreateBookData{
		ID:        b.ID.Value,
		ISBN:      b.ISBN.Value,
		Title:     b.Title,
		Authors:   value_objects.IdsToString(b.Authors),
		Subjects:  value_objects.IdsToString(b.Subjects),
		Year:      b.Year,
		Pages:     b.Pages,
		CreatedAt: input.CreatedAt,
	})
}
