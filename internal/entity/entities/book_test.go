package entities

import (
	"testing"

	"github.com/google/uuid"
	"github.com/otaviotech/quode/internal/entity/value_objects"
	"github.com/stretchr/testify/suite"
)

type BookTestSuite struct {
	suite.Suite
	bookIsbn  string
	bookTitle string
	bookYear  int
	bookPages int
	authors   []string
	subjects  []string
}

func (s *BookTestSuite) SetupTest() {
	s.bookIsbn = "978-0544003415"
	s.bookTitle = "The Lord of the Rings"
	s.bookYear = 1954
	s.bookPages = 1216
	s.authors = []string{uuid.NewString(), uuid.NewString(), uuid.NewString()}
	s.subjects = []string{uuid.NewString(), uuid.NewString(), uuid.NewString()}
}

func (s *BookTestSuite) Test_NewBook_ReturnsBook_WhenValidBook() {
	book, err := NewBook(s.bookTitle, s.bookIsbn, s.bookYear, s.bookPages, s.authors, s.subjects)
	s.NoError(err)
	s.Equal(s.bookTitle, book.Title)
	s.Equal(s.bookYear, book.Year)
	s.Equal(s.bookPages, book.Pages)
	s.NotZero(book.CreatedAt)
	s.Zero(book.UpdatedAt)

	authorsIds := []string{}

	for _, author := range book.Authors {
		authorsIds = append(authorsIds, author.Value)
	}

	s.Equal(s.authors, authorsIds)

	subjectsIds := []string{}
	for _, subject := range book.Subjects {
		subjectsIds = append(subjectsIds, subject.Value)
	}

	s.Equal(s.subjects, subjectsIds)
}

func (s *BookTestSuite) Test_NewBook_ReturnsError_WhenInvalidIsbn() {
	invalidIsbn := ""
	book, err := NewBook(s.bookTitle, invalidIsbn, s.bookYear, s.bookPages, s.authors, s.subjects)
	s.Nil(book)
	s.ErrorIs(err, value_objects.ErrInvalidISBN)
}

func (s *BookTestSuite) Test_NewBook_ReturnsError_WhenInvalidAuthor() {
	invalidAuthor := "not_a_uuid"
	book, err := NewBook(s.bookTitle, s.bookIsbn, s.bookYear, s.bookPages, []string{invalidAuthor}, s.subjects)
	s.Nil(book)
	s.ErrorIs(err, value_objects.ErrInvalidID)
}

func (s *BookTestSuite) Test_NewBook_ReturnsError_WhenInvalidSubject() {
	invalidSubject := "not_a_uuid"
	book, err := NewBook(s.bookTitle, s.bookIsbn, s.bookYear, s.bookPages, s.authors, []string{invalidSubject})
	s.Nil(book)
	s.ErrorIs(err, value_objects.ErrInvalidID)
}

func (s *BookTestSuite) Test_NewBook_ReturnsError_WhenInvalidTitle() {
	invalidTitle := ""
	book, err := NewBook(invalidTitle, s.bookIsbn, s.bookYear, s.bookPages, s.authors, s.subjects)
	s.Nil(book)
	s.ErrorIs(err, ErrInvalidBookTitle)
}

func (s *BookTestSuite) Test_NewBook_ReturnsError_WhenInvalidPages() {
	invalidPages := 0
	book, err := NewBook(s.bookTitle, s.bookIsbn, s.bookYear, invalidPages, s.authors, s.subjects)
	s.Nil(book)
	s.ErrorIs(err, ErrInvalidBookPages)
}

func (s *BookTestSuite) Test_NewBook_ReturnsError_WhenInvalidAuthors_NoAuthors() {
	book, err := NewBook(s.bookTitle, s.bookIsbn, s.bookYear, s.bookPages, []string{}, s.subjects)
	s.Nil(book)
	s.ErrorIs(err, ErrInvalidBookAuthor)
}

func (s *BookTestSuite) Test_NewBook_ReturnsError_WhenInvalidSubjects_NoSubjects() {
	book, err := NewBook(s.bookTitle, s.bookIsbn, s.bookYear, s.bookPages, s.authors, []string{})
	s.Nil(book)
	s.ErrorIs(err, ErrInvalidBookSubject)
}

func TestBookTestSuite(t *testing.T) {
	suite.Run(t, new(BookTestSuite))
}
