package usecase

import (
	"context"

	"github.com/otaviotech/quode/internal/application/app_errors"
	"github.com/otaviotech/quode/internal/application/repository"
	"github.com/otaviotech/quode/internal/entity/entities"
)

type CreateAuthorInput struct {
	Name string
	Bio  string
}

type CreateAuthorUseCaseInterface interface {
	Execute(ctx context.Context, input *CreateAuthorInput) error
}

type CreateAuthorUseCase struct {
	authorRepository repository.AuthorRepositoryInterface
}

func NewCreateAuthorUseCase(authorRepository repository.AuthorRepositoryInterface) *CreateAuthorUseCase {
	return &CreateAuthorUseCase{
		authorRepository: authorRepository,
	}
}

func (c *CreateAuthorUseCase) Execute(ctx context.Context, input *CreateAuthorInput) error {
	author, err := entities.NewAuthor(input.Name, input.Bio)

	if err != nil {
		return &app_errors.ErrDomainValidation{OriginalError: err}
	}

	data := repository.CreateAuthorData{
		ID:        author.ID.Value,
		Name:      author.Name,
		Bio:       author.Bio,
		CreatedAt: author.CreatedAt,
	}

	err = c.authorRepository.Create(ctx, data)

	if err != nil {
		return &app_errors.ErrException{OriginalError: err}
	}

	return nil
}
