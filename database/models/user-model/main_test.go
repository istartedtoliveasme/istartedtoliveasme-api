package userModel

import (
	"api/configs"
	"api/constants"
	neo4jConstant "api/constants/neo4j"
	"fmt"
<<<<<<< HEAD
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
=======
>>>>>>> 8140b66 (Code improvements and update mod files)
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserModelShouldBeInvalidUsername(t *testing.T) {
	// TODO :: mocking environment variables
	getEnv := configs.LoadEnvironmentVariables()
<<<<<<< HEAD
	neo4jProps := configs.Neo4jDriverProps{
=======
	credentials := configs.Neo4jDriverCredentials{
>>>>>>> 8140b66 (Code improvements and update mod files)
		Uri:      getEnv(neo4jConstant.Neo4jUri),
		Username: getEnv(neo4jConstant.Neo4jUsername),
		Password: getEnv(neo4jConstant.Neo4jPassword),
	}
	// TODO :: mocking driver
	// GIVEN the session drive started
<<<<<<< HEAD
	_, session := configs.Neo4jDriver(neo4jProps)
=======
	configs.Neo4jDriver(credentials)
>>>>>>> 8140b66 (Code improvements and update mod files)
	// AND a fake username
	fakeRandomUsername := "randomUserName"
	// AND getting the property dependency
	props := GetByEmailProps{
<<<<<<< HEAD
		GetSession: func() neo4j.Session {
			return session
		},
=======
>>>>>>> 8140b66 (Code improvements and update mod files)
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
<<<<<<< HEAD
	neo4jProps := configs.Neo4jDriverProps{
=======
	neo4jProps := configs.Neo4jDriverCredentials{
>>>>>>> 8140b66 (Code improvements and update mod files)
		Uri:      neo4jConstant.Neo4jUri,
		Username: neo4jConstant.Neo4jUsername,
		Password: neo4jConstant.Neo4jPassword,
	}
	// GIVEN the session drive started
<<<<<<< HEAD
	_, session := configs.Neo4jDriver(neo4jProps)
=======
	configs.Neo4jDriver(neo4jProps)
>>>>>>> 8140b66 (Code improvements and update mod files)
	// AND a fake username
	fakeRandomUsername := "istartedtoliveasme"
	// AND getting the property dependency
	props := GetByEmailProps{
<<<<<<< HEAD
		GetSession: func() neo4j.Session {
			return session
		},
=======
>>>>>>> 8140b66 (Code improvements and update mod files)
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
<<<<<<< HEAD
	neo4jProps := configs.Neo4jDriverProps{
=======
	neo4jProps := configs.Neo4jDriverCredentials{
>>>>>>> 8140b66 (Code improvements and update mod files)
		Uri:      neo4jConstant.Neo4jUri,
		Username: neo4jConstant.Neo4jUsername,
		Password: neo4jConstant.Neo4jPassword,
	}
	// GIVEN the session drive started
<<<<<<< HEAD
	_, session := configs.Neo4jDriver(neo4jProps)
=======
	configs.Neo4jDriver(neo4jProps)
>>>>>>> 8140b66 (Code improvements and update mod files)
	// AND a fake username
	fakeUserName := "istartedtoliveasme"
	// AND getting the property dependency
	props := GetByEmailProps{
<<<<<<< HEAD
		GetSession: func() neo4j.Session {
			return session
		},
=======
>>>>>>> 8140b66 (Code improvements and update mod files)
		GetEmail: func() string {
			return fakeUserName
		},
	}

	// WHEN I call the function with the fake username
	got, _ := GetByEmail(props)

	// THEN I should be able to see the email containing the fake username
	assert.Contains(t, got.Email, fakeUserName)
}
