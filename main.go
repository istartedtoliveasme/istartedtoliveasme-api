package main

import (
	"api/configs"
	"api/configs/routers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	port := "1337"

	if err := configs.GetRouterConfig(router); err != nil {
		panic(err)
	}

	routers.GetV1Routers(router)

	driver, _ := configs.StartNeo4jDriver()
	if err := driver.VerifyConnectivity(); err != nil {
		panic(err)
	}

	// close driver after connection is verified
	if err := driver.Close(); err != nil {
		panic(err)
	}

	if err := router.Run(":" + port); err != nil {
		panic(err)
	}

	fmt.Println("Database connection is verified!")
	fmt.Println("Running server on port: " + port)
	// listen and serve on 0.0.0.0:1337 (for windows "localhost:8080")
}
