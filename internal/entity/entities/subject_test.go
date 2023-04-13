package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type SubjectTestSuite struct {
	suite.Suite
}

func (s *SubjectTestSuite) Test_NewSubject_ReturnsSubject_WhenValidSubject() {
	subject, err := NewSubject("fiction", "fiction books are awesome books")
	s.NoError(err)
	s.Equal("fiction", subject.Name)
	s.Equal("fiction books are awesome books", subject.Description)
	s.IsType(subject.CreatedAt, time.Now())
	s.Zero(subject.UpdatedAt)
}

func (s *SubjectTestSuite) Test_NewSubject_ReturnsError_WhenInvalidName() {
	subject, err := NewSubject("", "fiction books are awesome books")
	s.Nil(subject)
	s.ErrorIs(err, ErrInvalidSubjectName)
}

func (s *SubjectTestSuite) Test_NewSubject_ReturnsError_WhenInvalidDescription() {
	subject, err := NewSubject("fiction", "")
	s.Nil(subject)
	s.ErrorIs(err, ErrInvalidSubjectDescription)
}

func TestSubjectTestSuite(t *testing.T) {
	suite.Run(t, new(SubjectTestSuite))
}
