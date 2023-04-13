package entities

import (
	"errors"
	"time"

	"github.com/otaviotech/quode/internal/entity/value_objects"
)

var (
	ErrInvalidQuoteContent = errors.New("invalid quote content")
	ErrInvalidQuotePage    = errors.New("invalid quote page")
)

type Quote struct {
	ID        value_objects.ID
	BookID    value_objects.ID
	Content   string
	Page      int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewQuote(bookId, content string, page int) (*Quote, error) {
	parsedBookId, err := value_objects.ParseID(bookId)

	if err != nil {
		return nil, err
	}

	q := &Quote{
		ID:        *value_objects.NewID(),
		BookID:    *parsedBookId,
		Content:   content,
		Page:      page,
		CreatedAt: time.Now(),
	}

	if err := q.validate(); err != nil {
		return nil, err
	}

	return q, nil
}

func (q *Quote) validate() error {
	if q.Content == "" {
		return ErrInvalidQuoteContent
	}

	if q.Page <= 0 {
		return ErrInvalidQuotePage
	}

	return nil
}
