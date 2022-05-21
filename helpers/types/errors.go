package types

type CustomError interface {
	Error() string
	Unwrap() error
}
