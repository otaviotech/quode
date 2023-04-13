package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/otaviotech/quode/internal/application/app_errors"
	"github.com/otaviotech/quode/internal/entity/entities"
	usecase_test "github.com/otaviotech/quode/test/mocks/usecase"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CreateSubjectHandlerTestSuite struct {
	suite.Suite
	sut    CreateSubjectHandler
	ucMock usecase_test.CreateSubjectUseCaseMock
}

func (s *CreateSubjectHandlerTestSuite) SetupTest() {
	s.ucMock = usecase_test.CreateSubjectUseCaseMock{}
	s.sut = *NewCreateSubjectHandler(&s.ucMock)
}

func (s *CreateSubjectHandlerTestSuite) Test_MethodIsPOST() {
	s.Equal(http.MethodPost, s.sut.GetMethod())
}

func (s *CreateSubjectHandlerTestSuite) Test_PathIsSubjects() {
	s.Equal("/subjects", s.sut.GetPath())
}

func (s *CreateSubjectHandlerTestSuite) Test_Return400_WhenInvalidRequestBody() {
	s.ucMock.On("Execute", mock.Anything, mock.Anything).Return(nil)
	req, err := http.NewRequest(http.MethodPost, "/subjects", bytes.NewBufferString("{"))

	s.NoError(err)

	w := httptest.NewRecorder()
	s.sut.ServeHTTP(w, req)

	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *CreateSubjectHandlerTestSuite) Test_Return400_WhenDomainError() {
	s.ucMock.On("Execute", mock.Anything, mock.Anything).Return(&app_errors.ErrDomainValidation{
		OriginalError: entities.ErrInvalidSubjectName,
	})

	req, err := http.NewRequest(http.MethodPost, "/subjects", bytes.NewBufferString(`{"name": "", "description": ""}`))
	s.NoError(err)
	w := httptest.NewRecorder()

	s.sut.ServeHTTP(w, req)

	s.Equal(http.StatusBadRequest, w.Code)
	s.Equal(`{"error": "invalid name"}`, strings.Split(w.Body.String(), "\n")[0])
}

func (s *CreateSubjectHandlerTestSuite) Test_Return201_WhenSuccess() {
	s.ucMock.On("Execute", mock.Anything, mock.Anything).Return(nil)
	req, err := http.NewRequest(http.MethodPost, "/subjects", bytes.NewBufferString(`{"name": "john", "description": "john doe"}`))
	s.NoError(err)
	w := httptest.NewRecorder()

	s.sut.ServeHTTP(w, req)

	s.Equal(http.StatusCreated, w.Code)
}

func TestCreateSubjectHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(CreateSubjectHandlerTestSuite))
}
