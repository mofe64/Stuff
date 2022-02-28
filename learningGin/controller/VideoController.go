package controller

import (
	"github.com/gin-gonic/gin"
	"learningGin/entity"
	"learningGin/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type VideoControllerImpl struct {
	service service.VideoService
}

// New Constructor function to create instance of out controller
func New(service service.VideoService) VideoController {
	return &VideoControllerImpl{
		service: service,
	}
}

func (c *VideoControllerImpl) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *VideoControllerImpl) Save(ctx *gin.Context) entity.Video {
	var v entity.Video
	ctx.BindJSON(&v)
	c.service.Save(v)
	return v
}
