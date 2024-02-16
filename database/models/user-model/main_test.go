package userModel

import (
	"api/configs"
	"api/constants"
	neo4jConstant "api/constants/neo4j"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserModelShouldBeInvalidUsername(t *testing.T) {
	// TODO :: mocking environment variables
	getEnv := configs.LoadEnvironmentVariables()
	credentials := configs.Neo4jDriverCredentials{
		Uri:      getEnv(neo4jConstant.Neo4jUri),
		Username: getEnv(neo4jConstant.Neo4jUsername),
		Password: getEnv(neo4jConstant.Neo4jPassword),
	}
	// TODO :: mocking driver
	// GIVEN the session drive started
	configs.Neo4jDriver(credentials)
	// AND a fake username
	fakeRandomUsername := "randomUserName"
	// AND getting the property dependency
	props := GetByEmailProps{
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
	neo4jProps := configs.Neo4jDriverCredentials{
		Uri:      neo4jConstant.Neo4jUri,
		Username: neo4jConstant.Neo4jUsername,
		Password: neo4jConstant.Neo4jPassword,
	}
	// GIVEN the session drive started
	configs.Neo4jDriver(neo4jProps)
	// AND a fake username
	fakeRandomUsername := "istartedtoliveasme"
	// AND getting the property dependency
	props := GetByEmailProps{
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
	neo4jProps := configs.Neo4jDriverCredentials{
		Uri:      neo4jConstant.Neo4jUri,
		Username: neo4jConstant.Neo4jUsername,
		Password: neo4jConstant.Neo4jPassword,
	}
	// GIVEN the session drive started
	configs.Neo4jDriver(neo4jProps)
	// AND a fake username
	fakeUserName := "istartedtoliveasme"
	// AND getting the property dependency
	props := GetByEmailProps{
		GetEmail: func() string {
			return fakeUserName
		},
	}

	// WHEN I call the function with the fake username
	got, _ := GetByEmail(props)

	// THEN I should be able to see the email containing the fake username
	assert.Contains(t, got.Email, fakeUserName)
}
