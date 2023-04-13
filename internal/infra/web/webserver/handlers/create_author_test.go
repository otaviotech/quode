package handlers

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/otaviotech/quode/internal/application/app_errors"
	usecase "github.com/otaviotech/quode/internal/application/usecase"
	"github.com/otaviotech/quode/internal/entity/entities"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// ucmock
type CreateAuthorUseCaseMock struct {
	mock.Mock
}

func (m *CreateAuthorUseCaseMock) Execute(ctx context.Context, input *usecase.CreateAuthorInput) error {
	args := m.Called(ctx, input)
	return args.Error(0)
}

type CreateAuthorHandlerTestSuite struct {
	suite.Suite
	ucMock CreateAuthorUseCaseMock
	sut    CreateAuthorHandler
}

func (s *CreateAuthorHandlerTestSuite) SetupTest() {
	s.ucMock = CreateAuthorUseCaseMock{}
	s.sut = *NewCreateAuthorHandler(&s.ucMock)
}

func (s *CreateAuthorHandlerTestSuite) Test_MustBePostMethod() {
	s.Equal(http.MethodPost, s.sut.GetMethod())
}

func (s *CreateAuthorHandlerTestSuite) Test_MustHaveCorrectPath() {
	s.Equal("/authors", s.sut.GetPath())
}

func (s *CreateAuthorHandlerTestSuite) Test_MustReturn400WhenInvalidJSON() {
	s.ucMock.On("Execute", mock.Anything, mock.Anything).Return(nil)
	req, err := http.NewRequest(http.MethodPost, "/authors", bytes.NewBufferString(`{`))

	s.NoError(err)

	w := httptest.NewRecorder()

	s.sut.ServeHTTP(w, req)

	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *CreateAuthorHandlerTestSuite) Test_MustReturn400_WhenDomainError() {
	ucErr := &app_errors.ErrDomainValidation{OriginalError: entities.ErrInvalidAuthorName}

	s.ucMock.On("Execute", mock.Anything, mock.Anything).Return(ucErr)
	req, err := http.NewRequest(http.MethodPost, "/authors", bytes.NewBufferString(`{"name": "", "bio": "bio"}`))

	s.NoError(err)

	w := httptest.NewRecorder()

	s.sut.ServeHTTP(w, req)

	s.Equal(http.StatusBadRequest, w.Code)
	s.Equal("{\"error\": \"invalid name\"}", strings.Split(w.Body.String(), "\n")[0])
}

func (s *CreateAuthorHandlerTestSuite) Test_MustReturn500_WhenExceptionError() {
	ucErr := &app_errors.ErrException{OriginalError: errors.New("exception")}

	s.ucMock.On("Execute", mock.Anything, mock.Anything).Return(ucErr)
	req, err := http.NewRequest(http.MethodPost, "/authors", bytes.NewBufferString(`{"name": "name", "bio": "bio"}`))

	s.NoError(err)

	w := httptest.NewRecorder()

	s.sut.ServeHTTP(w, req)

	s.Equal(http.StatusInternalServerError, w.Code)
	s.Equal("{\"error\": \"internal server error\"}", strings.Split(w.Body.String(), "\n")[0])
}

func TestCreateAuthorHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(CreateAuthorHandlerTestSuite))
}
