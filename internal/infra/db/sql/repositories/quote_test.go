package repositories

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/otaviotech/quode/internal/application/repository"
	"github.com/otaviotech/quode/pkg/dbutil"
	"github.com/stretchr/testify/suite"
)

type QuoteRepositoryTestSuite struct {
	suite.Suite
	db    *sql.DB
	sut   QuoteRepository
	books []string
}

func (s *QuoteRepositoryTestSuite) SetupTest() {
	conn, err := sql.Open("postgres", "postgres://quode:quode@localhost:5431/quode_test?sslmode=disable")

	if err != nil {
		s.Fail(err.Error())
	}

	s.db = conn
	s.sut = *NewQuoteRepository(conn)

	s.books = []string{uuid.NewString()}

	_, err = s.db.Exec("INSERT INTO books (id, isbn, title, pages, year) VALUES ($1, '9783161484100', 'Lorem ipsum', 100, 2020)", s.books[0])

	if err != nil {
		s.Fail(err.Error())
	}
}

func (s *QuoteRepositoryTestSuite) TearDownTest() {
	queries := []string{
		"DELETE FROM quotes",
		"DELETE FROM books",
	}

	for _, query := range queries {
		_, err := s.db.Exec(query)

		if err != nil {
			s.Fail(err.Error())
		}
	}
}

func TestQuoteRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(QuoteRepositoryTestSuite))
}

func (s *QuoteRepositoryTestSuite) Test_CreateQuote_InsertsIntoDatabase() {
	input := repository.CreateQuoteData{
		ID:        uuid.NewString(),
		BookID:    s.books[0],
		Page:      10,
		Content:   "Lorem ipsum",
		CreatedAt: time.Now(),
	}

	err := s.sut.Create(context.Background(), input)

	s.NoError(err)

	var id string
	var bookID string
	var page int
	var content string
	var createdAt time.Time

	err = s.db.QueryRow("SELECT id, book_id, page, content, created_at FROM quotes WHERE id = $1", input.ID).Scan(&id, &bookID, &page, &content, &createdAt)

	s.NoError(err)

	s.Equal(input.ID, id)
	s.Equal(input.BookID, bookID)
	s.Equal(input.Page, page)
	s.Equal(input.Content, content)
	s.WithinDuration(input.CreatedAt, createdAt, time.Microsecond)
}

func (s *QuoteRepositoryTestSuite) Test_ListQuotes_ReturnsQuotes() {
	queries := []string{
		"INSERT INTO quotes (id, book_id, page, content) VALUES ($1, $2, 10, '1 Lorem ipsum')",
		"INSERT INTO quotes (id, book_id, page, content) VALUES ($1, $2, 20, '2 Lorem ipsum')",
		"INSERT INTO quotes (id, book_id, page, content) VALUES ($1, $2, 30, '3 Lorem ipsum')",
		"INSERT INTO quotes (id, book_id, page, content) VALUES ($1, $2, 40, '4 Lorem ipsum')",
	}

	ids := []string{uuid.NewString(), uuid.NewString(), uuid.NewString(), uuid.NewString()}

	for i, query := range queries {
		_, err := s.db.Exec(query, ids[i], s.books[0])

		if err != nil {
			s.Fail(err.Error())
		}
	}

	input := repository.ListQuotesData{
		Pagination: dbutil.NewPagination(2, 0),
	}

	result, err := s.sut.List(context.Background(), input)

	s.NoError(err)
	s.Equal(4, result.Total)
	s.Len(result.Data, 2)

	s.Equal(ids[0], result.Data[0].ID)
	s.Equal(result.Data[0].Content, "1 Lorem ipsum")
	s.Equal(result.Data[0].Page, 10)

	s.Equal(ids[1], result.Data[1].ID)
	s.Equal(result.Data[1].Content, "2 Lorem ipsum")
	s.Equal(result.Data[1].Page, 20)

	input = repository.ListQuotesData{
		Pagination: dbutil.NewPagination(2, 2),
	}

	result, err = s.sut.List(context.Background(), input)

	s.NoError(err)
	s.Equal(4, result.Total)
	s.Len(result.Data, 2)

	s.Equal(ids[2], result.Data[0].ID)
	s.Equal(result.Data[0].Content, "3 Lorem ipsum")
	s.Equal(result.Data[0].Page, 30)

	s.Equal(ids[3], result.Data[1].ID)
	s.Equal(result.Data[1].Content, "4 Lorem ipsum")
	s.Equal(result.Data[1].Page, 40)
}
