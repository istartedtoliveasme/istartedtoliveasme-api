package mood

import (
	moodModel "api/database/models/mood-model"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"math/rand"
	"time"
)

type Body struct {
	Icon        string `form:"icon" json:"icon" binding:"required"`
	Title       string `form:"title" json:"title" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
}

func createMoodRecord(body Body) moodModel.Mood {
	return moodModel.Mood{
		Id:          rand.Int(),
		Icon:        body.Icon,
		Title:       body.Title,
		Description: body.Description,
		CreatedAt:   time.Now().UTC(),
	}
}

func CreateMoodPropertyFactory(s neo4j.Session, m moodModel.Mood) moodModel.Controller {
	return moodModel.Controller{
		GetSession: func() neo4j.Session {
			return s
		},
		GetMood: func() moodModel.Mood {
			return m
		},
	}
}
