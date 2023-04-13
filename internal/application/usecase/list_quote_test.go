package usecase

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/otaviotech/quode/internal/application/repository"
	"github.com/otaviotech/quode/pkg/dbutil"
	repository_test "github.com/otaviotech/quode/test/mocks/repository"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ListQuotesUseCaseTestSuite struct {
	suite.Suite
	quotesRepoMock *repository_test.QuoteRepositoryMock
	sut            ListQuotesUseCase
}

func (s *ListQuotesUseCaseTestSuite) SetupTest() {
	s.quotesRepoMock = new(repository_test.QuoteRepositoryMock)
	s.sut = ListQuotesUseCase{quoteRepository: s.quotesRepoMock}
}

func (s *ListQuotesUseCaseTestSuite) Test_Execute_ReturnsError_WhenRepoReturnsError() {
	s.quotesRepoMock.On("List", mock.Anything, mock.Anything).Return(&repository.ListQuotesResult{}, ErrFoo)

	_, err := s.sut.Execute(context.Background(), ListQuotesInput{
		Pagination: dbutil.NewPagination(10, 0),
	})

	s.ErrorIs(err, ErrFoo)
}

func (s *ListQuotesUseCaseTestSuite) Test_Execute_ReturnsPaginatedResult() {
	repoResult := repository.ListQuotesResult{
		Data: []repository.ListQuotesResultItem{
			{
				ID:       uuid.NewString(),
				BookID:   uuid.NewString(),
				BookName: "Clean Architecture",
				Page:     100,
				Content:  "The only way to go fast, is to go well",
			},
			{
				ID:       uuid.NewString(),
				BookID:   uuid.NewString(),
				BookName: "Domain Driven Design",
				Page:     100,
				Content:  "Tackling Complexity in the Heart of Software",
			},
		},
		Total: 10,
	}

	s.quotesRepoMock.On("List", mock.Anything, mock.Anything).Return(&repoResult, nil)

	result, err := s.sut.Execute(context.Background(), ListQuotesInput{
		Pagination: dbutil.NewPagination(2, 0),
	})

	s.NoError(err)

	s.quotesRepoMock.AssertCalled(s.T(), "List", mock.Anything, repository.ListQuotesData{
		Pagination: dbutil.NewPagination(2, 0),
	})

	s.Len(result.Data, 2)

	s.Equal(repoResult.Total, result.Total)
	s.Equal(repoResult.Data[0].ID, result.Data[0].ID)
	s.Equal(repoResult.Data[0].BookID, result.Data[0].BookID)
	s.Equal(repoResult.Data[0].BookName, result.Data[0].BookName)
	s.Equal(repoResult.Data[0].Page, result.Data[0].Page)
	s.Equal(repoResult.Data[0].Content, result.Data[0].Content)

	s.Equal(repoResult.Data[1].ID, result.Data[1].ID)
	s.Equal(repoResult.Data[1].BookID, result.Data[1].BookID)
	s.Equal(repoResult.Data[1].BookName, result.Data[1].BookName)
	s.Equal(repoResult.Data[1].Page, result.Data[1].Page)
	s.Equal(repoResult.Data[1].Content, result.Data[1].Content)

}

func TestListQuotesUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(ListQuotesUseCaseTestSuite))
}
