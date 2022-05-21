package moodModel

import (
	"api/constants"
	"api/helpers"
	"api/helpers/dataRecord"
	helperTypes "api/helpers/typings"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"path/filepath"
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
	GetSession func() neo4j.Session
	UserId     string
}

type MoodCreateCypher map[string]interface{}

func (m Mood) Create(c Props) (MoodWithUserRecord, helperTypes.CustomError) {
	var jsonMood helperTypes.Json[MoodWithUserRecord]
	props := MoodCreateCypher{
		"userId":      c.UserId,
		"id":          m.Id,
		"icon":        m.Icon,
		"title":       m.Title,
		"description": m.Description,
		"createdAt":   m.CreatedAt,
	}
	filepath.Base("/cypher/create.cypher")
	cypher := "MATCH (u:User) WHERE u.id = $userId " +
		"CREATE (m:Mood {id: $id, icon: $icon, title: $title, description: $description, createdAt: $createdAt }) " +
		"CREATE (m)-[:HAS_USER]->(u), (u)-[:HAS_MOOD]->(m)  " +
		"CREATE  CONSTRAINT ON (u) ASSERT exists(u.id)" +
		"CREATE  CONSTRAINT ON (u) ASSERT exists(u.title)" +
		"CREATE  CONSTRAINT ON (u) ASSERT exists(u.createdAt)" +
		"RETURN *"

	record, err := c.GetSession().Run(cypher, props)

	if err != nil {
		return jsonMood.Payload, MoodError{
			Message: constants.FailedCreateRecord,
			Err:     err,
		}
	}

	serialize, err := helpers.GetAllRecords(record)
	jsonMood.ParsePayload(serialize)

	if err != nil {
		return jsonMood.Payload, MoodError{
			Message: constants.FailedFetchRecord,
			Err:     err,
		}
	}

	return jsonMood.Payload, nil
}

func (mood *Mood) GetMoods(session neo4j.Session) ([]MoodWithUserRecord, helperTypes.CustomError) {
	var items []MoodWithUserRecord
	cypher := "MATCH (m:Mood)-[r:HAS_USER]->(u:User) RETURN *"
	records, err := session.Run(cypher, nil)

	if err != nil {
		return nil, MoodError{
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
