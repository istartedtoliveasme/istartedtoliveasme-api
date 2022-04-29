package userModel

import (
	"api/configs"
	"api/constants"
	neo4jConstant "api/constants/neo4j"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserModelShouldBeInvalidUsername(t *testing.T) {
	// TODO :: mocking environment variables
	getEnv := configs.LoadEnvironmentVariables()
	neo4jProps := configs.Neo4jDriverProps{
		Uri:      getEnv(neo4jConstant.URI),
		Username: getEnv(neo4jConstant.USERNAME),
		Password: getEnv(neo4jConstant.PASSWORD),
	}
	// TODO :: mocking driver
	// GIVEN the session drive started
	_, session := configs.Neo4jDriver(neo4jProps)
	// AND a fake username
	fakeRandomUsername := "randomUserName"
	// AND getting the property dependency
	props := GetByEmailProps{
		GetSession: func() neo4j.Session {
			return session
		},
		GetEmail: func() string {
			return fakeRandomUsername
		},
	}

	// WHEN I call the function with the fake username
	got, err := GetByEmail(props)

	// THEN I should be able to see an error message
	assert.EqualError(t, err, fmt.Sprintf("%s!", constants.InvalidEmail))
	// AND I should not be able to see any result
	assert.Empty(t, got)
}

func TestUserModelShouldNotThrowError(t *testing.T) {
	configs.LoadEnvironmentVariables()
	neo4jProps := configs.Neo4jDriverProps{
		Uri:      neo4jConstant.URI,
		Username: neo4jConstant.USERNAME,
		Password: neo4jConstant.PASSWORD,
	}
	// GIVEN the session drive started
	_, session := configs.Neo4jDriver(neo4jProps)
	// AND a fake username
	fakeRandomUsername := "istartedtoliveasme"
	// AND getting the property dependency
	props := GetByEmailProps{
		GetSession: func() neo4j.Session {
			return session
		},
		GetEmail: func() string {
			return fakeRandomUsername
		},
	}

	// WHEN I call the function with the fake username
	got, err := GetByEmail(props)

	// THEN I should not be able to see an error message
	assert.NoError(t, err)
	// AND I should be able to have a result
	assert.NotEmpty(t, got)
}

func TestUserModelShouldSeeResultValue(t *testing.T) {
	configs.LoadEnvironmentVariables()
	neo4jProps := configs.Neo4jDriverProps{
		Uri:      neo4jConstant.URI,
		Username: neo4jConstant.USERNAME,
		Password: neo4jConstant.PASSWORD,
	}
	// GIVEN the session drive started
	_, session := configs.Neo4jDriver(neo4jProps)
	// AND a fake username
	fakeUserName := "istartedtoliveasme"
	// AND getting the property dependency
	props := GetByEmailProps{
		GetSession: func() neo4j.Session {
			return session
		},
		GetEmail: func() string {
			return fakeUserName
		},
	}

	// WHEN I call the function with the fake username
	got, _ := GetByEmail(props)

	// THEN I should be able to see the email containing the fake username
	assert.Contains(t, got.Email, fakeUserName)
}
