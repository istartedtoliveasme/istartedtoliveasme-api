package moodModel

import (
	"api/constants"
	"api/database/models/typings"
	"api/database/structures"
	"api/helpers"
	"api/helpers/error-helper"
	"api/helpers/httpHelper"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"time"
)

type Mood struct {
	Id          int
	Icon        string
	Title       string
	Description string
	CreatedAt   time.Time
}

type Props struct {
	GetUser    func() structures.UserRecord
	GetSession func() neo4j.Session
	GetMood    func() Mood
}

func CreateMood(c Props) (interface{}, errorHelper.CustomError) {
	m := c.GetMood()
	u := c.GetUser()
	props := httpHelper.JSON{
		"userId":      u.Id,
		"id":          m.Id,
		"icon":        m.Icon,
		"title":       m.Title,
		"description": m.Description,
		"createdAt":   m.CreatedAt,
	}
	cypher := "MATCH(u: User {email: $userId}) CREATE (m:Mood {id: $id, icon: $icon, title: $title, description: $description, createdAt: $createdAt }) RETURN m"
	record, err := c.GetSession().Run(cypher, props)

	if err != nil {
		return nil, typings.RecordError{
			Message: constants.FailedCreateRecord,
			Err:     err,
		}
	}

	return record, nil
}

func GetMoods(session neo4j.Session) ([]httpHelper.JSON, errorHelper.CustomError) {
	cypher := "MATCH (m:Mood) RETURN *"
	records, err := session.Run(cypher, nil)

	fmt.Println(records)

	if err != nil {
		return nil, typings.RecordError{
			Message: constants.GetRecordFailed,
			Err:     err,
		}
	}

	serialize, err := helpers.GetAllRecords(records)

	if err != nil {
		return nil, typings.RecordError{
			Message: constants.FailedFetchRecord,
			Err:     err,
		}
	}

	return serialize, nil
}
