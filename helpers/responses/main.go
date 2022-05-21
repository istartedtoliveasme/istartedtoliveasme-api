package responses

import (
	"api/helpers/types"
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
	Err     types.CustomError
	Payload T
}

func (r HttpResponse[T]) BadRequest() (int, types.JsonPayload) {
	switch len(r.Message) {
	case 0:
		r.Message = r.Err.Unwrap().Error()
	}

	return http.StatusBadRequest, types.JsonPayload{
		"error":   r.Err.Unwrap().Error(),
		"message": r.Message,
	}
}

func (r HttpResponse[T]) OkRequest() (int, types.JsonPayload) {
	return http.StatusOK, types.JsonPayload{
		"message": r.Message,
		"data":    r.Payload,
	}
}
