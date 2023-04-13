package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/otaviotech/quode/internal/application/repository"
	"github.com/otaviotech/quode/internal/entity/entities"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type SubjectRepositoryMock struct {
	mock.Mock
}

func (m *SubjectRepositoryMock) Create(ctx context.Context, input repository.CreateSubjectData) error {
	args := m.Called(ctx, input)
	return args.Error(0)
}

type CreateSubjectUseCaseTestSuite struct {
	suite.Suite
	sut         CreateSubjectUseCase
	subjectRepo SubjectRepositoryMock
}

func (s *CreateSubjectUseCaseTestSuite) SetupTest() {
	s.subjectRepo = SubjectRepositoryMock{}
	s.sut = *NewCreateSubjectUseCase(&s.subjectRepo)
}

func (s *CreateSubjectUseCaseTestSuite) Test_ReturnsError_WhenInputIsInvalid() {
	input := CreateSubjectInput{
		Name:        "",
		Description: "Some valid description",
	}

	err := s.sut.Execute(context.Background(), &input)

	s.ErrorIs(err, entities.ErrInvalidSubjectName)
}

func (s *CreateSubjectUseCaseTestSuite) Test_ReturnsError_WhenRepositoryReturnsError() {
	input := CreateSubjectInput{
		Name:        "Some valid name",
		Description: "Some valid description",
	}

	s.subjectRepo.On("Create", mock.Anything, mock.Anything).Return(ErrFoo)

	err := s.sut.Execute(context.Background(), &input)

	s.ErrorIs(err, ErrFoo)
}

func (s *CreateSubjectUseCaseTestSuite) Test_ReturnsNil_WhenInputIsValid() {
	input := CreateSubjectInput{
		Name:        "Some valid name",
		Description: "Some valid description",
		CreatedAt:   time.Now(),
	}

	s.subjectRepo.On("Create", mock.Anything, mock.Anything).Return(nil)

	ctx := context.Background()

	err := s.sut.Execute(ctx, &input)

	s.NoError(err)

	id := s.subjectRepo.Calls[0].Arguments[1].(repository.CreateSubjectData).ID
	createdAt := s.subjectRepo.Calls[0].Arguments[1].(repository.CreateSubjectData).CreatedAt

	s.NotZero(id)
	s.NotZero(createdAt)

	s.subjectRepo.AssertCalled(s.T(), "Create", ctx, repository.CreateSubjectData{
		ID:          id,
		Name:        input.Name,
		Description: input.Description,
		CreatedAt:   createdAt,
	})
}

func TestCreateSubjectUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(CreateSubjectUseCaseTestSuite))
}
