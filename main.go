package main

import (
	"api/configs"
	"api/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	router := gin.Default()
	port := "1337"

	// Load environment variables
	loadEnv()

	err := configs.GetRouterConfig(router)

	if err != nil {
		panic(err)
	}

	routers.GetV1Routers(router)

	driver, _ := configs.Neo4jDriver()

	err = driver.VerifyConnectivity()

	if err != nil {
		panic(err)
	}

	// close driver after connection is verified
	if err := driver.Close(); err != nil {
		panic(err)
	}

	err = router.Run(":" + port)

	if err != nil {
		panic(err)
	}

	fmt.Println("Database connection is verified!")
	fmt.Println("Running server on port: " + port)
	// listen and serve on 0.0.0.0:1337 (for windows "localhost:8080")
}
