package value_objects

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type IDTestSuite struct {
	suite.Suite
}

func (s *IDTestSuite) Test_NewID_GeneratesNewID() {
	id := NewID()
	parsed, err := uuid.Parse(id.Value)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), id.Value, parsed.String())
}

func (s *IDTestSuite) Test_ParseID_ReturnsError_WhenInvalidID() {
	_, err := ParseID("invalid")
	assert.ErrorIs(s.T(), err, ErrInvalidID)
}

func (s *IDTestSuite) Test_ParseID_ReturnsID_WhenValidID() {
	id := NewID()
	parsed, err := ParseID(id.Value)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), id.Value, parsed.Value)
}

func (s *IDTestSuite) Test_ParseIDs_ReturnsError_WhenInvalidID() {
	_, err := ParseIDs([]string{"invalid"})
	assert.ErrorIs(s.T(), err, ErrInvalidID)
}

func (s *IDTestSuite) Test_ParseIDs_ReturnsIDs_WhenValidIDs() {
	id1, id2 := uuid.NewString(), uuid.NewString()
	ids := []string{id1, id2}

	parsed, err := ParseIDs(ids)

	s.NoError(err)
	s.Equal(parsed, []ID{{Value: id1}, {Value: id2}})
}

func (s *IDTestSuite) Test_IdsToString_ReturnsStringIDs() {
	id1, id2 := uuid.NewString(), uuid.NewString()
	ids := []ID{{Value: id1}, {Value: id2}}

	idsString := IdsToString(ids)

	s.Equal(idsString, []string{id1, id2})
}

func (s *IDTestSuite) Test_IdsToString_ReturnsEmptySlice_WhenEmptySlice() {
	idsString := IdsToString([]ID{})
	s.Equal(idsString, []string{})
}

func (s *IDTestSuite) Test_IdsToString_ReturnsEmptySlice_WhenNil() {
	idsString := IdsToString(nil)
	s.Equal(idsString, []string{})
}

func TestIDTestSuite(t *testing.T) {
	suite.Run(t, new(IDTestSuite))
}
