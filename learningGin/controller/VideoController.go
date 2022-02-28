package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"learningGin/entity"
	"learningGin/service"
	"learningGin/validators"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
}

type VideoControllerImpl struct {
	service service.VideoService
}

var customValidator *validator.Validate

// New Constructor function to create instance of out controller
func New(service service.VideoService) VideoController {
	customValidator = validator.New()
	customValidator.RegisterValidation("is-cool", validators.ValidateTitle)
	return &VideoControllerImpl{
		service: service,
	}
}

func (c *VideoControllerImpl) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *VideoControllerImpl) Save(ctx *gin.Context) error {
	var v entity.Video
	err := ctx.ShouldBind(&v)
	if err != nil {
		return err
	}
	err = customValidator.Struct(v)
	if err != nil {
		return err
	}
	c.service.Save(v)
	return nil
}
