package helpers

type CustomError interface {
	Error() string
	Unwrap() error
}
