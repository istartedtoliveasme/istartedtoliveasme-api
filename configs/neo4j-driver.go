package configs

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"os"
)

func Neo4jDriver() (neo4j.Driver, neo4j.Session) {
	// Neo4j 4.0, defaults to no TLS therefore use bolt:// or neo4j://
	// Neo4j 3.5, defaults to self-signed certificates, TLS on, therefore use bolt+ssc:// or neo4j+ssc://
	uri := os.Getenv("NEO4J_URI")
	credentials := neo4j.BasicAuth(os.Getenv("NEO4J_USERNAME"), os.Getenv("NEO4J_PASSWORD"), "")
	driver, err := neo4j.NewDriver(uri, credentials)

	if err != nil {
		panic(err)
	}

	// REFER TO: https://github.com/neo4j/neo4j-go-driver#minimum-viable-snippet

	// Create a session to run transactions in. Sessions are lightweight to
	// create and use. Sessions are NOT thread safe.
	session, err := driver.NewSession(neo4j.SessionConfig{})

	if err != nil {
		// Closing driver if session fails
		driver.Close()
		panic(err)
	}

	return driver, session
}
