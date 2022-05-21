package moodModel

import (
	"api/constants"
	"api/helpers/serializers"
	"api/helpers/types"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type MoodWithUserRecord struct {
	Mood
	User serializers.UserResponse `json:"user"`
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

func (moodWithUserRecord *MoodWithUserRecord) getMapLabelProps(node neo4j.Node) types.CustomError {
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
	json := types.Json[Mood]{
		Payload: moodWithUserRecord.Mood,
	}
	if err := json.Parse(source); err != nil {
		return err
	}

	return nil
}

func (moodWithUserRecord *MoodWithUserRecord) ParseUser(source interface{}) error {
	json := types.Json[serializers.UserResponse]{
		Payload: moodWithUserRecord.User,
	}
	if err := json.Parse(source); err != nil {
		return err
	}

	return nil
}
