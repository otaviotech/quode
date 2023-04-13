package value_objects

import (
	"errors"

	"github.com/google/uuid"
)

var ErrInvalidID = errors.New("invalid id")

type ID struct {
	Value string
}

func NewID() *ID {
	return &ID{Value: uuid.NewString()}
}

func ParseID(id string) (*ID, error) {
	_, err := uuid.Parse(id)

	if err != nil {
		return nil, ErrInvalidID
	}

	return &ID{Value: id}, nil
}

func ParseIDs(ids []string) ([]ID, error) {
	var parsedIDs []ID

	for _, id := range ids {
		parsedID, err := ParseID(id)

		if err != nil {
			return nil, err
		}

		parsedIDs = append(parsedIDs, *parsedID)
	}

	return parsedIDs, nil
}

func IdsToString(ids []ID) []string {
	if len(ids) == 0 {
		return []string{}
	}

	var idsString []string

	for _, id := range ids {
		idsString = append(idsString, id.Value)
	}

	return idsString
}
