package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/otaviotech/quode/internal/application/repository"
	"github.com/otaviotech/quode/internal/entity/entities"
	"github.com/otaviotech/quode/internal/entity/value_objects"
	repository_test "github.com/otaviotech/quode/test/mocks/repository"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CreateBookUseCaseTestSuite struct {
	suite.Suite
	sut          CreateBookUseCaseInterface
	bookRepoMock *repository_test.BookRepositoryMock
}

func (s *CreateBookUseCaseTestSuite) SetupTest() {
	s.bookRepoMock = new(repository_test.BookRepositoryMock)
	s.sut = NewCreateBookUseCase(s.bookRepoMock)
}

func (s *CreateBookUseCaseTestSuite) Test_Execute_ReturnsError_WhenInvalidInput() {
	err := s.sut.Execute(context.Background(), CreateBookInput{
		ISBN:     "invalid",
		Title:    "",
		Authors:  []string{value_objects.NewID().Value},
		Subjects: []string{value_objects.NewID().Value},
		Year:     2023,
		Pages:    100,
	})

	s.ErrorIs(err, entities.ErrInvalidBookTitle)
}

func (s *CreateBookUseCaseTestSuite) Test_Execute_ReturnsError_WhenRepositoryReturnsError() {
	s.bookRepoMock.On("Create", mock.Anything, mock.Anything).Return(ErrFoo)

	ctx := context.Background()

	input := CreateBookInput{
		ISBN:      "978-3-16-148410-0",
		Title:     "The Hitchhiker's Guide to the Galaxy",
		Authors:   []string{value_objects.NewID().Value, value_objects.NewID().Value},
		Subjects:  []string{value_objects.NewID().Value, value_objects.NewID().Value},
		Year:      1979,
		Pages:     224,
		CreatedAt: time.Now(),
	}

	err := s.sut.Execute(ctx, input)

	s.ErrorIs(err, ErrFoo)
}

func (s *CreateBookUseCaseTestSuite) Test_Execute_CallsBookRepositoryCreate() {
	s.bookRepoMock.On("Create", mock.Anything, mock.Anything).Return(nil)

	ctx := context.Background()

	input := CreateBookInput{
		ISBN:      "978-3-16-148410-0",
		Title:     "The Hitchhiker's Guide to the Galaxy",
		Authors:   []string{value_objects.NewID().Value, value_objects.NewID().Value},
		Subjects:  []string{value_objects.NewID().Value, value_objects.NewID().Value},
		Year:      1979,
		Pages:     224,
		CreatedAt: time.Now(),
	}

	err := s.sut.Execute(ctx, input)

	s.NoError(err)

	id := s.bookRepoMock.Calls[0].Arguments[1].(repository.CreateBookData).ID
	createdAt := s.bookRepoMock.Calls[0].Arguments[1].(repository.CreateBookData).CreatedAt

	s.NotZero(id)
	s.NotZero(createdAt)

	s.bookRepoMock.AssertCalled(s.T(), "Create", ctx, repository.CreateBookData{
		ID:        id,
		ISBN:      input.ISBN,
		Title:     input.Title,
		Authors:   input.Authors,
		Subjects:  input.Subjects,
		Year:      input.Year,
		Pages:     input.Pages,
		CreatedAt: createdAt,
	})
}

func TestCreateBookUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(CreateBookUseCaseTestSuite))
}
