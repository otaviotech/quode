package repositories

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/otaviotech/quode/internal/application/repository"
	"github.com/stretchr/testify/suite"
)

type BookRepositoryTestSuite struct {
	suite.Suite
	db       *sql.DB
	sut      BookRepository
	authors  []string
	subjects []string
}

func (s *BookRepositoryTestSuite) SetupTest() {
	conn, err := sql.Open("postgres", "postgres://quode:quode@localhost:5431/quode_test?sslmode=disable")

	if err != nil {
		s.Fail(err.Error())
	}

	s.db = conn
	s.sut = *NewBookRepository(conn)

	s.authors = []string{uuid.NewString()}
	s.subjects = []string{uuid.NewString()}

	_, err = s.db.Exec(
		"INSERT INTO authors (id, name, bio, created_at) VALUES ($1, 'John Doe', 'Lorem ipsum dolor sit amet', $2)",
		s.authors[0],
		time.Now(),
	)

	if err != nil {
		s.Fail(err.Error())
	}

	_, err = s.db.Exec(
		"INSERT INTO subjects (id, name, description, created_at) VALUES ($1, 'Lorem ipsum', 'Lorem ipsum', $2)",
		s.subjects[0],
		time.Now(),
	)

	if err != nil {
		s.Fail(err.Error())
	}
}

func (s *BookRepositoryTestSuite) TearDownTest() {
	queries := []string{
		"DELETE FROM books_authors",
		"DELETE FROM books_subjects",
		"DELETE FROM books",
		"DELETE FROM authors",
	}

	for _, query := range queries {
		_, err := s.db.Exec(query)

		if err != nil {
			s.Fail(err.Error())
		}
	}
}

func (s *BookRepositoryTestSuite) Test_CreateBook_InsertsIntoDB() {
	input := repository.CreateBookData{
		ID:        uuid.NewString(),
		ISBN:      "1234567890",
		Title:     "Lorem ipsum",
		Authors:   s.authors,
		Subjects:  s.subjects,
		Year:      2023,
		Pages:     100,
		CreatedAt: time.Now(),
	}

	err := s.sut.Create(context.Background(), input)

	if err != nil {
		s.Fail(err.Error())
	}

	s.NoError(err)

	// fetch from db
	var id string
	var isbn string
	var title string
	var year int
	var pages int
	var createdAt time.Time

	err = s.db.QueryRow("SELECT id, isbn, title, year, pages, created_at FROM books WHERE id = $1", input.ID).Scan(&id, &isbn, &title, &year, &pages, &createdAt)

	if err != nil {
		s.Fail(err.Error())
	}

	s.Equal(input.ID, id)
	s.Equal(input.ISBN, isbn)
	s.Equal(input.Title, title)
	s.Equal(input.Year, year)
	s.Equal(input.Pages, pages)
	s.WithinDuration(input.CreatedAt, createdAt, time.Microsecond)

	// fetch authors
	var authorID string
	var bookID string

	err = s.db.QueryRow("SELECT author_id, book_id FROM books_authors WHERE book_id = $1", input.ID).Scan(&authorID, &bookID)

	s.NoError(err)

	s.Equal(input.ID, bookID)
	s.Equal(input.Authors[0], authorID)

	// fetch subjects
	var subjectID string

	err = s.db.QueryRow("SELECT book_id, subject_id FROM books_subjects WHERE book_id = $1", input.ID).Scan(&bookID, &subjectID)
	s.NoError(err)

	s.Equal(input.ID, bookID)
	s.Equal(input.Subjects[0], subjectID)
}

func TestBookRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(BookRepositoryTestSuite))
}
