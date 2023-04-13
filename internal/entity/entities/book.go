package entities

import (
	"errors"
	"time"

	"github.com/otaviotech/quode/internal/entity/value_objects"
)

var (
	ErrInvalidBookTitle   = errors.New("invalid title")
	ErrInvalidBookPages   = errors.New("invalid pages")
	ErrInvalidBookAuthor  = errors.New("invalid author")
	ErrInvalidBookSubject = errors.New("invalid subject")
)

type Book struct {
	ID        value_objects.ID
	ISBN      value_objects.ISBN
	Authors   []value_objects.ID
	Subjects  []value_objects.ID
	Title     string
	Year      int
	Pages     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewBook(title, isbn string, year, pages int, authors, subjects []string) (*Book, error) {
	parsedISBN, err := value_objects.ParseISBN(isbn)

	if err != nil {
		return nil, err
	}

	parsedAuthors, err := value_objects.ParseIDs(authors)

	if err != nil {
		return nil, err
	}

	parsedSubjects, err := value_objects.ParseIDs(subjects)

	if err != nil {
		return nil, err
	}

	b := &Book{
		ID:        *value_objects.NewID(),
		ISBN:      *parsedISBN,
		Authors:   parsedAuthors,
		Subjects:  parsedSubjects,
		Title:     title,
		Year:      year,
		Pages:     pages,
		CreatedAt: time.Now(),
	}

	if err := b.validate(); err != nil {
		return nil, err
	}

	return b, nil
}

func (b *Book) validate() error {
	if b.Title == "" {
		return ErrInvalidBookTitle
	}

	if b.Pages <= 0 {
		return ErrInvalidBookPages
	}

	if len(b.Authors) == 0 {
		return ErrInvalidBookAuthor
	}

	if len(b.Subjects) == 0 {
		return ErrInvalidBookSubject
	}

	return nil
}
