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

<<<<<<< HEAD
	err := configs.GetRouterConfig(router)

	if err != nil {
=======
	if err := configs.GetRouterConfig(router); err != nil {
>>>>>>> 8140b66 (Code improvements and update mod files)
		panic(err)
	}

	routers.GetV1Routers(router)

	driver, _ := configs.StartNeo4jDriver()

<<<<<<< HEAD
	err = driver.VerifyConnectivity()

	if err != nil {
=======
	if err := driver.VerifyConnectivity(); err != nil {
>>>>>>> 8140b66 (Code improvements and update mod files)
		panic(err)
	}

	// close driver after connection is verified
	if err := driver.Close(); err != nil {
		panic(err)
	}

<<<<<<< HEAD
	err = router.Run(":" + port)

	if err != nil {
=======
	if err := router.Run(":" + port); err != nil {
>>>>>>> 8140b66 (Code improvements and update mod files)
		panic(err)
	}

	fmt.Println("Database connection is verified!")
	fmt.Println("Running server on port: " + port)
	// listen and serve on 0.0.0.0:1337 (for windows "localhost:8080")
}
