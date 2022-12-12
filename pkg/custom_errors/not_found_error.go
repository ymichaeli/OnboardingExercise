package custom_errors

import "github.com/pkg/errors"

// NotFoundError is a common error which allows us to handle it in a custom way
type NotFoundError struct {
	publicMessage string
	err           string
}

func NewNotFoundError(publicMessage string, error string) error {
	return errors.WithStack(NotFoundError{publicMessage: publicMessage, err: error})
}

func (e NotFoundError) Error() string {
	if e.err != "" {
		return e.err
	}
	return e.publicMessage
}

func (e NotFoundError) DisplayMessage() string {
	return e.publicMessage
}
