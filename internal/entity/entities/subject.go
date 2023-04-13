package entities

import (
	"errors"
	"time"

	"github.com/otaviotech/quode/internal/entity/value_objects"
)

var ErrInvalidSubjectName = errors.New("invalid name")
var ErrInvalidSubjectDescription = errors.New("invalid description")

type Subject struct {
	ID          value_objects.ID
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewSubject(name, description string) (*Subject, error) {
	s := &Subject{
		ID:          *value_objects.NewID(),
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
	}

	if err := s.validate(); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Subject) validate() error {
	if s.Name == "" {
		return ErrInvalidSubjectName
	}

	if s.Description == "" {
		return ErrInvalidSubjectDescription
	}

	return nil
}
