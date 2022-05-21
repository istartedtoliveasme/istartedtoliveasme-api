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

func ParseMoods(collections []*neo4j.Record) ([]MoodWithUserRecord, errorHelper.CustomError) {
	var items []MoodWithUserRecord

	for _, record := range collections {
		item, err := SerializeMoodAndUserRecordValues(record)
		if err != nil {
			return items, err
		}
		items = append(items, item)
	}

	return items, nil
}

func SerializeMoodAndUserRecordValues(record *neo4j.Record) (MoodWithUserRecord, errorHelper.CustomError) {
	var moodWithUserRecord MoodWithUserRecord
	for _, recordValue := range record.Values {

		switch recordValue.(type) {
		case neo4j.Node:
			err := moodWithUserRecord.getMapLabelProps(recordValue.(neo4j.Node))
			if err != nil {
				return moodWithUserRecord, err
			}
		}

	}

	return moodWithUserRecord, nil
}

func (moodWithUserRecord *MoodWithUserRecord) getMapLabelProps(node neo4j.Node) errorHelper.CustomError {
	for _, label := range node.Labels {
		switch label {
		case "Mood":
			if err := parseNodeProps(node.Props, &moodWithUserRecord); err != nil {
				return ParseMoodsError{
					Message: constants.FailedParseMood,
					Err:     err,
				}
			}
		case "User":
			if err := parseNodeProps(node.Props, &moodWithUserRecord.User); err != nil {
				return ParseMoodsError{
					Message: constants.FailedParseMood,
					Err:     err,
				}
			}
		}
	}

	return nil
}

func parseNodeProps(data interface{}, parseData interface{}) error {
	if err := httpHelper.JSONParse(data, &parseData); err != nil {
		return err
	}

	return nil
}
