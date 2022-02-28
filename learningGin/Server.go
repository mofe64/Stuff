package main

import (
	"github.com/gin-gonic/gin"
	"learningGin/controller"
	"learningGin/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()
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
