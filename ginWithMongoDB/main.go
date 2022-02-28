package main

import (
	"ginWithMongoDB/config"
	"ginWithMongoDB/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	//run database
	config.ConnectDB()

	//setup routes
	routes.UserRoute(router)

	//run server
	err := router.Run("localhost:6000")
	if err != nil {
		return
	}
}
