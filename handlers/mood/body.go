package mood

import (
	moodModel "api/database/models/mood-model"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Body struct {
	Icon        string `form:"icon" json:"icon" binding:"required"`
	Title       string `form:"title" json:"title" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
}

func (body *Body) CreateMoodPropertyFactory(s neo4j.Session) moodModel.Props {
	return moodModel.Props{
		GetSession: func() neo4j.Session {
			return s
		},
	}
}
