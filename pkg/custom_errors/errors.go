package custom_errors

type CustomError interface {
	error
	DisplayMessage() string
}
