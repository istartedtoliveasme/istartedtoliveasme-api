package moodModel

import (
	"api/constants"
	errorHelper "api/helpers/error-helper"
	"api/helpers/httpHelper"
	"api/helpers/serializers"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type MoodWithUserRecord struct {
	Mood
	User serializers.UserSerializer `json:"user"`
}

type ParseMoodsError struct {
	Message string
	Err     error
}

func (p ParseMoodsError) Error() string {
	return p.Message
}

func (p ParseMoodsError) Unwrap() error {
	return p.Err
}

func (moodWithUserRecord *MoodWithUserRecord) getMapLabelProps(node neo4j.Node) errorHelper.CustomError {
	for _, label := range node.Labels {
		switch label {
		case "Mood":
			if err := moodWithUserRecord.ParseMood(node.Props); err != nil {
				return ParseMoodsError{
					Message: constants.FailedParseMood,
					Err:     err,
				}
			}
		case "User":
			if err := moodWithUserRecord.ParseUser(node.Props); err != nil {
				return ParseMoodsError{
					Message: constants.FailedParseMood,
					Err:     err,
				}
			}
		}
	}

	return nil
}

func (moodWithUserRecord *MoodWithUserRecord) ParseMood(source interface{}) error {
	if err := httpHelper.JSONParse(source, &moodWithUserRecord.Mood); err != nil {
		return err
	}

	return nil
}

func (moodWithUserRecord *MoodWithUserRecord) ParseUser(source interface{}) error {
	if err := httpHelper.JSONParse(source, &moodWithUserRecord.User); err != nil {
		return err
	}

	return nil
}
