package httpHelper

import (
	"api/constants"
	"api/database/models/typings"
	"api/helpers/error-helper"
	"api/helpers/serializers"
	"fmt"
	"strings"
)

type Header struct {
	Authorization string `header:"Authorization"`
}

type HeaderError struct {
	Message string
	Err     error
}

func (hErr HeaderError) Error() string {
	return hErr.Message
}

func (hErr HeaderError) Unwrap() error {
	return hErr.Err
}

func (h Header) GetAuthorizationBearer() (string, errorHelper.CustomError) {
	auth := strings.Fields(h.Authorization)
	var bearer string

	if len(auth) > 1 {
		bearer = auth[1]
	}

	if len(bearer) == 0 {
		return bearer, HeaderError{
			Message: constants.FailedAuthorize,
			Err:     nil,
		}
	}

	return bearer, nil

}

func (h Header) DecodeAuthorizationBearer() (serializers.UserSerializer, errorHelper.CustomError) {
	var userSerializer serializers.UserSerializer

	bearerToken, err := h.GetAuthorizationBearer()

	if err != nil {
		return userSerializer, err
	}

	if err := JSONParse(bearerToken, &userSerializer); err != nil {
		return userSerializer, typings.RecordError{
			Message: constants.FailedDecodeRecord,
			Err:     err,
		}
	}

	fmt.Println(userSerializer)

	return userSerializer, nil
}
