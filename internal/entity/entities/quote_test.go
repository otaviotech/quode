package entities

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/otaviotech/quode/internal/entity/value_objects"
	"github.com/stretchr/testify/suite"
)

type QuoteTestSuite struct {
	suite.Suite
	quoteBookId  string
	quoteContent string
	quotePage    int
}

func (s *QuoteTestSuite) SetupTest() {
	s.quoteBookId = uuid.NewString()
	s.quoteContent = "quote content"
	s.quotePage = 1
}

func (s *QuoteTestSuite) Test_NewQuote_ReturnsQuote_WhenValidQuote() {
	quote, err := NewQuote(s.quoteBookId, s.quoteContent, s.quotePage)
	s.NoError(err)
	s.Equal(s.quoteContent, quote.Content)
	s.Equal(s.quotePage, quote.Page)
}

func (s *QuoteTestSuite) Test_NewQuote_ReturnsError_WhenInvalidBookId() {
	invalidBookId := ""
	quote, err := NewQuote(invalidBookId, s.quoteContent, s.quotePage)
	s.Nil(quote)
	s.ErrorIs(err, value_objects.ErrInvalidID)
}

func (s *QuoteTestSuite) Test_NewQuote_ReturnsError_WhenInvalidContent() {
	invalidContent := ""
	quote, err := NewQuote(s.quoteBookId, invalidContent, s.quotePage)
	s.Nil(quote)
	s.ErrorIs(err, ErrInvalidQuoteContent)
}

func (s *QuoteTestSuite) Test_NewQuote_ReturnsError_WhenInvalidPage() {
	invalidPage := 0
	quote, err := NewQuote(s.quoteBookId, s.quoteContent, invalidPage)
	s.Nil(quote)
	s.ErrorIs(err, ErrInvalidQuotePage)
}

func (s *QuoteTestSuite) Test_NewQuote_ReturnsQuote_WhenValid() {
	quote, err := NewQuote(s.quoteBookId, s.quoteContent, s.quotePage)
	s.NoError(err)
	s.Equal(s.quoteContent, quote.Content)
	s.Equal(s.quotePage, quote.Page)
	s.IsType(quote.CreatedAt, time.Now())
	s.Zero(quote.UpdatedAt)
}

func TestQuoteTestSuite(t *testing.T) {
	suite.Run(t, new(QuoteTestSuite))
}
