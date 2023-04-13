package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var name = "John Doe"
var bio = "John Doe is an author."

type AuthorTestSuite struct {
	suite.Suite
}

func (s *AuthorTestSuite) Test_NewAuthor_ReturnsAuthor_WhenValidAuthor() {
	author, err := NewAuthor(name, bio)
	s.NoError(err)
	s.Equal(name, author.Name)
	s.Equal(bio, author.Bio)
	s.IsType(author.CreatedAt, time.Now())
	s.Zero(author.UpdatedAt)
}

func (s *AuthorTestSuite) Test_NewAuthor_ReturnsError_WhenInvalidName() {
	author, err := NewAuthor("", bio)
	assert.ErrorIs(s.T(), err, ErrInvalidAuthorName)
	assert.Nil(s.T(), author)
}

func (s *AuthorTestSuite) Test_NewAuthor_ReturnsError_WhenInvalidBio() {
	author, err := NewAuthor(name, "")
	assert.ErrorIs(s.T(), err, ErrInvalidBio)
	assert.Nil(s.T(), author)
}

func TestAuthorTestSuite(t *testing.T) {
	suite.Run(t, new(AuthorTestSuite))
}
