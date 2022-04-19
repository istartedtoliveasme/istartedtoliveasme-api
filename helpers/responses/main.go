package responses

import (
	"api/helpers/httpHelper"
	"net/http"
)

func BadRequest(message string, errs []error) (int, httpHelper.JSON) {
	var errorMessages []string

	for _, eachError := range errs {
		errorMessages = append(errorMessages, eachError.Error())
	}

	return http.StatusBadRequest, httpHelper.JSON{
		"errors":  errorMessages,
		"message": message,
	}
}

func OkRequest(message string, payload interface{}) (int, httpHelper.JSON) {
	return http.StatusOK, httpHelper.JSON{
		"message": message,
		"data":    payload,
	}
}
