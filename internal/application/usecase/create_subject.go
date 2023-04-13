package usecase

import (
	"context"
	"time"

	"github.com/otaviotech/quode/internal/application/repository"
	"github.com/otaviotech/quode/internal/entity/entities"
)

type CreateSubjectInput struct {
	Name        string
	Description string
	CreatedAt   time.Time
}

type CreateSubjectUseCaseInterface interface {
	Execute(ctx context.Context, input *CreateSubjectInput) error
}

type CreateSubjectUseCase struct {
	subjectRepository repository.SubjectRepositoryInterface
}

func NewCreateSubjectUseCase(subjectRepository repository.SubjectRepositoryInterface) *CreateSubjectUseCase {
	return &CreateSubjectUseCase{
		subjectRepository: subjectRepository,
	}
}

func (c *CreateSubjectUseCase) Execute(ctx context.Context, input *CreateSubjectInput) error {
	s, err := entities.NewSubject(input.Name, input.Description)

	if err != nil {
		return err
	}

	return c.subjectRepository.Create(ctx, repository.CreateSubjectData{
		ID:          s.ID.Value,
		Name:        s.Name,
		Description: s.Description,
		CreatedAt:   input.CreatedAt,
	})
}
