package types

import (
	"api/constants"
	"encoding/json"
)

type JsonPayload map[string]interface{}
type Json[T any] struct {
	Payload T
}

type Error struct {
	Message string
	Err     error
}

func (jsonP Error) Error() string {
	return jsonP.Message
}

func (jsonP Error) Unwrap() error {
	return jsonP.Err
}

func (j *Json[T]) Parse(data interface{}) CustomError {
	byteArray, err := json.Marshal(&j)

	if err != nil {
		return Error{
			Message: constants.FailedParseClaim,
			Err:     err,
		}
	}

	if err = json.Unmarshal(byteArray, data); err != nil {
		return Error{
			Message: constants.FailedDecodeRecord,
			Err:     err,
		}
	}

	return nil
}
