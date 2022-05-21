package responses

import (
	helperTypes "api/helpers/typings"
	"net/http"
)

type BindError struct {
	Message string
	Err     error
}

func (r BindError) Error() string {
	return r.Message
}

func (r BindError) Unwrap() error {
	return r.Err
}

type HttpResponse[T any] struct {
	Message string
	Err     helperTypes.CustomError
	Payload T
}

func (r HttpResponse[T]) BadRequest() (int, helperTypes.JsonPayload) {
	switch len(r.Message) {
	case 0:
		r.Message = r.Err.Unwrap().Error()
	}

	return http.StatusBadRequest, helperTypes.JsonPayload{
		"error":   r.Err.Unwrap().Error(),
		"message": r.Message,
	}
}

func (r HttpResponse[T]) OkRequest() (int, helperTypes.JsonPayload) {
	return http.StatusOK, helperTypes.JsonPayload{
		"message": r.Message,
		"data":    r.Payload,
	}
}
