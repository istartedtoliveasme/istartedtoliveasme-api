package typings

type CustomError interface {
	Error() string
	Unwrap() error
}
