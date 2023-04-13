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

type SubjectRepositoryTestSuite struct {
	suite.Suite
	db  *sql.DB
	sut SubjectRepository
}

func (s *SubjectRepositoryTestSuite) SetupTest() {
	conn, err := sql.Open("postgres", "postgres://quode:quode@localhost:5431/quode_test?sslmode=disable")

	if err != nil {
		s.Fail(err.Error())
	}

	s.db = conn
	s.sut = *NewSubjectRepository(conn)
}

func (s *SubjectRepositoryTestSuite) TearDownTest() {
	_, err := s.db.Exec("DELETE FROM subjects")

	if err != nil {
		s.Fail(err.Error())
	}
}

func TestSubjectRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(SubjectRepositoryTestSuite))
}

func (s *SubjectRepositoryTestSuite) Test_CreateSubject_InsertsIntoDB() {
	input := repository.CreateSubjectData{
		ID:          uuid.NewString(),
		Name:        "Clean Architecture",
		Description: "Quotes about Clean Architecture",
		CreatedAt:   time.Now(),
	}

	err := s.sut.Create(context.Background(), input)

	s.NoError(err)

	var id string
	var name string
	var description string
	var createdAt time.Time

	err = s.db.QueryRow("SELECT id, name, description, created_at FROM subjects WHERE id = $1", input.ID).Scan(&id, &name, &description, &createdAt)

	if err != nil {
		s.Fail(err.Error())
	}

	s.Equal(input.ID, id)
	s.Equal(input.Name, name)
	s.Equal(input.Description, description)
	s.WithinDuration(input.CreatedAt, createdAt, time.Microsecond)
}
