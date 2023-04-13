package value_objects

import "errors"

var ErrInvalidISBN = errors.New("invalid ISBN")

type ISBN struct {
	Value string
}

// todo: validate ISBN
// rules: https://en.wikipedia.org/wiki/International_Standard_Book_Number
func ParseISBN(value string) (*ISBN, error) {
	if value == "" {
		return nil, ErrInvalidISBN
	}

	return &ISBN{Value: value}, nil
}
