package custom_errors

type NotFoundError struct {
	Err string
}

func (e NotFoundError) Error() string {
	return e.Err
}

type BadRequestError struct {
	Err string
}

func (e BadRequestError) Error() string {
	return e.Err
}
