package errorHelper

type CustomError interface {
	Error() string
	Unwrap() error
}
