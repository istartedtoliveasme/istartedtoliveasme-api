package mood

import (
	moodModel "api/database/models/mood-model"
	"api/handlers/mood/typings"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"math/rand"
	"time"
)

func createMoodRecord(body typings.Body) moodModel.Mood {
	return moodModel.Mood{
		Id:          rand.Int(),
		Icon:        body.Icon,
		Title:       body.Title,
		Description: body.Description,
		CreatedAt:   time.Now().UTC(),
	}
}

func CreateMoodPropertyFactory(s neo4j.Session, m moodModel.Mood) moodModel.Props {
	return moodModel.Props{
		GetSession: func() neo4j.Session {
			return s
		},
		GetMood: func() moodModel.Mood {
			return m
		},
	}
}
