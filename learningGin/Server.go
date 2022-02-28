package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"learningGin/controller"
	"learningGin/middleware"
	"learningGin/service"
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
	server.Use(gin.Recovery(), middleware.Logger(), middleware.BasicAuth())
	server.GET("/videos", func(context *gin.Context) {
		context.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(context *gin.Context) {
		context.JSON(200, videoController.Save(context))
	})
	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Ok!",
		})
	})
	err := server.Run(":8080")
	if err != nil {
		return
	}
}
