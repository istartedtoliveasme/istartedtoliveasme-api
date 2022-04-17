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

	loadEnv()

	err := configs.GetRouterConfig(router)

	if err != nil {
		panic(err)
	}

	routers.GetV1Routers(router)

	fmt.Println("Running server on port: " + port)
	router.Run(":" + port) // listen and serve on 0.0.0.0:1337 (for windows "localhost:8080")
}
