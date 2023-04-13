package repositories

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/otaviotech/quode/internal/application/repository"
	"github.com/stretchr/testify/suite"

	_ "github.com/lib/pq"
)

type AuthorRepositoryTestSuite struct {
	suite.Suite
	db  *sql.DB
	sut AuthorRepository
}

func (s *AuthorRepositoryTestSuite) SetupTest() {
	conn, err := sql.Open("postgres", "postgres://quode:quode@localhost:5431/quode_test?sslmode=disable")

	if err != nil {
		s.Fail(err.Error())
	}

	s.db = conn
	s.sut = *NewAuthorRepository(conn)
}

func (s *AuthorRepositoryTestSuite) TearDownTest() {
	_, err := s.db.Exec("DELETE FROM authors")

	if err != nil {
		s.Fail(err.Error())
	}
}

func TestAuthorRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(AuthorRepositoryTestSuite))
}

func (s *AuthorRepositoryTestSuite) Test_CreateAuthor_InsertsIntoDB() {
	input := repository.CreateAuthorData{
		ID:        uuid.NewString(),
		Name:      "John Doe",
		Bio:       "Lorem ipsum dolor sit amet",
		CreatedAt: time.Now(),
	}

	err := s.sut.Create(context.Background(), input)

	s.NoError(err)

	// fetch from db
	var id string
	var name string
	var bio string
	var createdAt time.Time

	err = s.db.QueryRow("SELECT id, name, bio, created_at FROM authors WHERE id = $1", input.ID).Scan(&id, &name, &bio, &createdAt)

	if err != nil {
		s.Fail(err.Error())
	}

	s.Equal(input.ID, id)
	s.Equal(input.Name, name)
	s.Equal(input.Bio, bio)
	s.WithinDuration(input.CreatedAt, createdAt, time.Microsecond)
}
