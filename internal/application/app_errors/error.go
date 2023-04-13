package app_errors

// For errors that are not expected to happen
type ErrException struct {
	OriginalError error
}

func (e *ErrException) Error() string {
	return e.OriginalError.Error()
}

// For errors in domain validation
type ErrDomainValidation struct {
	OriginalError error
}

func (e *ErrDomainValidation) Error() string {
	return e.OriginalError.Error()
}

type ErrNotFound struct {
	msg string
}

func (e *ErrNotFound) Error() string {
	return e.msg
}
