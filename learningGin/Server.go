package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"learningGin/controller"
	"learningGin/middleware"
	"learningGin/service"
	"net/http"
	"os"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

// setup log to file
func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
func main() {
	setupLogOutput()
	server := gin.New()
	//server.Use(gin.Recovery(), middleware.Logger(), middleware.BasicAuth())
	server.Use(gin.Recovery(), middleware.Logger())

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/videos", func(context *gin.Context) {
			context.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(context *gin.Context) {
			err := videoController.Save(context)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, gin.H{"message": "created"})
			}
		})
		apiRoutes.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Ok!",
			})
		})
	}

	err := server.Run(":8080")
	if err != nil {
		return
	}
}
