package responses

import (
	errorHelper "api/helpers/error-helper"
	"api/helpers/httpHelper"
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

func BadRequest(message string, errs errorHelper.CustomError) (int, httpHelper.JSON) {
	return http.StatusBadRequest, httpHelper.JSON{
		"error":   errs.Unwrap().Error(),
		"message": message,
	}
}

func OkRequest(message string, payload interface{}) (int, httpHelper.JSON) {
	return http.StatusOK, httpHelper.JSON{
		"message": message,
		"data":    payload,
	}
}
