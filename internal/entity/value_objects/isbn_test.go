package value_objects

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ISBNTestSuite struct {
	suite.Suite
}

func (s *ISBNTestSuite) Test_ParseISBN_ReturnsISBN_WhenValidISBN() {
	isbn := "978-3-16-148410-0"
	parsed, err := ParseISBN(isbn)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), isbn, parsed.Value)
}

func (s *ISBNTestSuite) Test_ParseISBN_ReturnsError_WhenInvalidISBN() {
	_, err := ParseISBN("")
	assert.ErrorIs(s.T(), err, ErrInvalidISBN)
}

func TestISBNTTestSuite(t *testing.T) {
	suite.Run(t, new(ISBNTestSuite))
}
