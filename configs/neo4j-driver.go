package configs

import (
	neo4jConstant "api/constants/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Neo4jDriverCredentials struct {
	Uri      string
	Username string
	Password string
}

func (c Neo4jDriverCredentials) getCredentials() Neo4jDriverCredentials {
	return c
}

type Neo4jDriverProps interface {
	getCredentials() Neo4jDriverCredentials
}

func Neo4jDriver(strategy Neo4jDriverProps) (neo4j.Driver, neo4j.Session) {
	credentials := strategy.getCredentials()
	// Neo4j 4.0, defaults to no TLS therefore use bolt:// or neo4j://
	// Neo4j 3.5, defaults to self-signed certificates, TLS on, therefore use bolt+ssc:// or neo4j+ssc://
	uri := credentials.Uri
	auth := neo4j.BasicAuth(credentials.Username, credentials.Password, "")
	// You typically have one driver instance for the entire application. The
	// driver maintains a pool of database connections to be used by the sessions.
	// The driver is thread safe.
	driver, err := neo4j.NewDriver(uri, auth)

	if err != nil {
		panic(err)
	}

	// REFER TO: https://github.com/neo4j/neo4j-go-driver#minimum-viable-snippet

	// Create a session to run transactions in. Sessions are lightweight to
	// create and use. Sessions are NOT thread safe.
	session := driver.NewSession(neo4j.SessionConfig{})

	return driver, session
}

func StartNeo4jDriver() (neo4j.Driver, neo4j.Session) {
	getEnv := LoadEnvironmentVariables()
	credentials := Neo4jDriverCredentials{
		Uri:      getEnv(neo4jConstant.Neo4jUri),
		Username: getEnv(neo4jConstant.Neo4jUsername),
		Password: getEnv(neo4jConstant.Neo4jPassword),
	}

	return Neo4jDriver(credentials)
}
