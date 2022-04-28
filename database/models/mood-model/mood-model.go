package moodModel

import (
	"api/helpers"
	"api/helpers/httpHelper"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"time"
)

type Mood struct {
	Id          int
	Icon        string
	Title       string
	Description string
	CreatedAt time.Time
}

type Controller struct {
	GetSession func() neo4j.Session
	GetMood    func() Mood
}

func CreateMood(c Controller) (interface{}, error) {
	return c.GetSession().WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {

		cypher := "CREATE (m:Mood {id: $id, icon: $icon, title: $title, description: $description, createdAt: $createdAt }) RETURN m"
		m := c.GetMood()
		result, err := tx.Run(cypher, httpHelper.JSON{
			"id":          m.Id,
			"icon":        m.Icon,
			"title":       m.Title,
			"description": m.Description,
			"createdAt": m.CreatedAt,
		})

		if err != nil {
			return nil, err
		}

		return result.Consume()
	})
}

func GetMoods(session neo4j.Session) ([]httpHelper.JSON, error) {
	cypher := "MATCH  (m:Mood) RETURN m"
	records, err := session.Run(cypher, nil)

	if err != nil {
		return nil, err
	}

	serialize, err := helpers.GetAllNodeProps(records)

	if err != nil {
		return nil, err
	}

	return serialize, nil
}
