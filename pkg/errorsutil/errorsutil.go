package errorsutil

import "errors"

func IsAnyOfErrors(err error, errs ...error) bool {
	for _, e := range errs {
		if errors.Is(err, e) {
			return true
		}
	}

	return false
}
