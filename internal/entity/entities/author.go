package entities

import (
	"errors"
	"time"

	"github.com/otaviotech/quode/internal/entity/value_objects"
)

var ErrInvalidAuthorName = errors.New("invalid name")
var ErrInvalidBio = errors.New("invalid bio")

type Author struct {
	ID        value_objects.ID
	Name      string
	Bio       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAuthor(name, bio string) (*Author, error) {
	a := &Author{
		ID:        *value_objects.NewID(),
		Name:      name,
		Bio:       bio,
		CreatedAt: time.Now(),
	}

	if err := a.validate(); err != nil {
		return nil, err
	}

	return a, nil
}

func (a *Author) validate() error {
	if a.Name == "" {
		return ErrInvalidAuthorName
	}

	if a.Bio == "" {
		return ErrInvalidBio
	}

	return nil
}
