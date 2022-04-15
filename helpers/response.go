package helpers

import (
	httpHelper "api/helpers/httpHelper"
	"net/http"
)

func BadRequest(errs []error) (int, httpHelper.JSON) {
	var errorMessages []string

	for _, eachError := range errs {
		errorMessages = append(errorMessages, eachError.Error())
	}

	return http.StatusBadRequest, httpHelper.JSON{
		"errors": errorMessages,
	}
}

func OkRequest(message string, payload interface{}) (int, httpHelper.JSON) {
	return http.StatusOK, httpHelper.JSON{
		"message": message,
		"data":    payload,
	}
}
