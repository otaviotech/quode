package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/otaviotech/quode/internal/application/app_errors"
	"github.com/otaviotech/quode/internal/application/repository"
	repository_test "github.com/otaviotech/quode/test/mocks/repository"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CreateAuthorTestSuite struct {
	suite.Suite
	sut  CreateAuthorUseCaseInterface
	repo *repository_test.AuthorRepositoryMock
}

func (s *CreateAuthorTestSuite) SetupTest() {
	s.repo = new(repository_test.AuthorRepositoryMock)
	s.sut = NewCreateAuthorUseCase(s.repo)
}

func (s *CreateAuthorTestSuite) Test_ReturnsError__WhenInputIsInvalid() {
	input := &CreateAuthorInput{
		Name: "",
		Bio:  "John Doe is an author.",
	}

	err := s.sut.Execute(context.Background(), input)

	var e *app_errors.ErrDomainValidation
	s.ErrorAs(err, &e)
}

func (s *CreateAuthorTestSuite) Test_ReturnsError__WhenRepositoryReturnsError() {
	input := &CreateAuthorInput{
		Name: "John Doe",
		Bio:  "John Doe is an author.",
	}

	s.repo.On("Create", mock.Anything, mock.Anything).Return(ErrFoo)

	err := s.sut.Execute(context.Background(), input)

	var e *app_errors.ErrException
	s.ErrorAs(err, &e)
}

func (s *CreateAuthorTestSuite) Test_ReturnsNil__WhenInputIsValid() {
	input := &CreateAuthorInput{
		Name: "John Doe",
		Bio:  "John Doe is an author.",
	}

	s.repo.On("Create", mock.Anything, mock.Anything).Return(nil)

	ctx := context.Background()

	err := s.sut.Execute(ctx, input)

	s.NoError(err)

	id := s.repo.Calls[0].Arguments[1].(repository.CreateAuthorData).ID
	createdAt := s.repo.Calls[0].Arguments[1].(repository.CreateAuthorData).CreatedAt

	s.NotZero(createdAt)
	s.IsType(createdAt, time.Now())

	s.repo.AssertCalled(s.T(), "Create", ctx, repository.CreateAuthorData{
		ID:        id,
		Name:      input.Name,
		Bio:       input.Bio,
		CreatedAt: createdAt,
	})

}

func TestCreateAuthorTestSuite(t *testing.T) {
	suite.Run(t, new(CreateAuthorTestSuite))
}
