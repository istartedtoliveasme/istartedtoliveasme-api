package moodModel

import (
	"api/constants"
	"api/database/models/typings"
	"api/database/structures"
	"api/helpers"
	"api/helpers/dataRecord"
	"api/helpers/error-helper"
	"api/helpers/httpHelper"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"time"
)

type Mood struct {
	Id          string    `json:"id"`
	Icon        string    `json:"icon"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

type MoodError struct {
	Message string
	Err     error
}

func (me MoodError) Error() string {
	return me.Message
}

func (me MoodError) Unwrap() error {
	return me.Err
}

type Props struct {
	GetUser    func() structures.UserRecord
	GetSession func() neo4j.Session
}

func (m Mood) Create(c Props) (interface{}, errorHelper.CustomError) {
	u := c.GetUser()
	props := httpHelper.JSON{
		"userId":      u.Id,
		"id":          m.Id,
		"icon":        m.Icon,
		"title":       m.Title,
		"description": m.Description,
		"createdAt":   m.CreatedAt,
	}
	cypher := "MATCH (u:User) WHERE u.id = $userId " +
		"CREATE (m:Mood {id: $id, icon: $icon, title: $title, description: $description, createdAt: $createdAt }) " +
		"CREATE (m)-[:HAS_USER]->(u), (u)-[:HAS_MOOD]->(m)  " +
		"RETURN m"

	record, err := c.GetSession().Run(cypher, props)

	if err != nil {
		return nil, typings.RecordError{
			Message: constants.FailedCreateRecord,
			Err:     err,
		}
	}

	serialize, err := helpers.GetAllRecords(record)

	if err != nil {
		return nil, typings.RecordError{
			Message: constants.FailedFetchRecord,
			Err:     err,
		}
	}

	return serialize, nil
}

func GetMoods(session neo4j.Session) ([]MoodWithUserRecord, errorHelper.CustomError) {
	var items []MoodWithUserRecord
	cypher := "MATCH (m:Mood)-[r:HAS_USER]->(u:User) RETURN *"
	records, err := session.Run(cypher, nil)

	if err != nil {
		return nil, typings.RecordError{
			Message: constants.GetRecordFailed,
			Err:     err,
		}
	}

	collection, err := records.Collect()
	if err != nil {
		return nil, nil
	}

	var moodWithUserRecord MoodWithUserRecord
	dataRecords := dataRecord.DataRecordCollections(collection)
	bindItems, err := dataRecords.Bind(moodWithUserRecord.getMapLabelProps)

	for _, bindItem := range bindItems {
		if err := moodWithUserRecord.ParseMood(bindItem); err != nil {
			return nil, MoodError{
				Message: constants.FailedParseMood,
				Err:     err,
			}
		}
		if err := moodWithUserRecord.ParseUser(bindItem); err != nil {
			return nil, MoodError{
				Message: constants.FailedParseMood,
				Err:     err,
			}
		}
		items = append(items, moodWithUserRecord)
	}

	return items, nil
}
