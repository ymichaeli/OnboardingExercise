package custom_errors

import "github.com/pkg/errors"

type BadRequestError struct {
	publicMessage string
	err           string
}

func NewBadRequestError(publicMessage string, error string) error {
	return errors.WithStack(BadRequestError{publicMessage: publicMessage, err: error})
}

func (e BadRequestError) Error() string {
	if e.err != "" {
		return e.err
	}
	return e.publicMessage
}

func (e BadRequestError) DisplayMessage() string {
	return e.publicMessage
}
