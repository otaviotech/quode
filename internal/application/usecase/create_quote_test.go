package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/otaviotech/quode/internal/application/repository"
	"github.com/otaviotech/quode/internal/entity/value_objects"
	repository_test "github.com/otaviotech/quode/test/mocks/repository"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var ErrFoo = errors.New("foo")

type CreateQuoteUseCaseTestSuite struct {
	suite.Suite
	quoteRepoMock *repository_test.QuoteRepositoryMock
	sut           CreateQuoteUseCaseInterface
}

func (s *CreateQuoteUseCaseTestSuite) SetupTest() {
	s.quoteRepoMock = new(repository_test.QuoteRepositoryMock)
	s.sut = NewCreateQuoteUseCase(s.quoteRepoMock)
}

func (s *CreateQuoteUseCaseTestSuite) Test_Execute_ReturnsError_WhenInvalidInput() {
	input := CreateQuoteInput{
		BookID:  "",
		Page:    100,
		Content: "lorem ipsum",
	}

	err := s.sut.Execute(context.Background(), input)

	s.ErrorIs(err, value_objects.ErrInvalidID)
}

func (s *CreateQuoteUseCaseTestSuite) Test_Execute_ReturnsError_WhenRepoReturnsError() {
	input := CreateQuoteInput{
		BookID:  value_objects.NewID().Value,
		Page:    100,
		Content: "lorem ipsum",
	}

	s.quoteRepoMock.On("Create", mock.Anything, mock.Anything).Return(ErrFoo)

	err := s.sut.Execute(context.Background(), input)

	s.ErrorIs(err, ErrFoo)
}

func (s *CreateQuoteUseCaseTestSuite) Test_Execute_ReturnsNil_WhenValidInput() {
	input := CreateQuoteInput{
		BookID:  value_objects.NewID().Value,
		Page:    100,
		Content: "lorem ipsum",
	}

	s.quoteRepoMock.On("Create", mock.Anything, mock.Anything).Return(nil)

	ctx := context.Background()
	err := s.sut.Execute(ctx, input)

	s.NoError(err)

	id := s.quoteRepoMock.Calls[0].Arguments[1].(repository.CreateQuoteData).ID
	createdAt := s.quoteRepoMock.Calls[0].Arguments[1].(repository.CreateQuoteData).CreatedAt

	s.NotZero(id)
	s.NotZero(createdAt)

	s.quoteRepoMock.AssertCalled(s.T(), "Create", ctx, repository.CreateQuoteData{
		ID:        id,
		BookID:    input.BookID,
		Page:      input.Page,
		Content:   input.Content,
		CreatedAt: createdAt,
	})
}

func TestCreateQuoteUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(CreateQuoteUseCaseTestSuite))
}
