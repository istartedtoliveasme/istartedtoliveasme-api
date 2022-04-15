package main

import (
	"api/configs"
	"api/routers"
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
	loadEnv()
	router := routers.GetV1Routers(configs.GetRouterConfig())

	router.Run(":1337") // listen and serve on 0.0.0.0:1337 (for windows "localhost:8080")
}
