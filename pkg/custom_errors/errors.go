package custom_errors

type CustomError interface {
	error
	// DisplayMessage will display the error message to the client. thus, shouldn't contain any protected data
	DisplayMessage() string
}
