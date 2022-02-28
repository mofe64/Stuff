package controllers

import (
	"context"
	"ginWithMongoDB/config"
	"ginWithMongoDB/models"
	"ginWithMongoDB/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

var userCollection = config.GetCollection(config.DB, "users")
var validate = validator.New()

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		defer cancel()

		// validate the request body
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{
				Status:  http.StatusBadRequest,
				Message: "Error",
				Data:    gin.H{"data": err.Error()},
			})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    gin.H{"data": validationErr.Error()},
			})
			return
		}

		newUser := models.User{
			Id:       primitive.NewObjectID(),
			Name:     user.Name,
			Location: user.Location,
			Title:    user.Title,
		}
		result, err := userCollection.InsertOne(ctx, newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError,
				responses.APIResponse{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    gin.H{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusCreated, responses.APIResponse{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    gin.H{"data": result},
		})
	}
}

func GetAUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		var user models.User
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		// find user with provided id and unmarshal doc into user obj or return err if any
		err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError,
				responses.APIResponse{
					Status:  http.StatusInternalServerError,
					Message: "error", Data: gin.H{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK,
			responses.APIResponse{
				Status:  http.StatusOK,
				Message: "success",
				Data:    gin.H{"data": user}})
	}
}

func EditAUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		var user models.User
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(userId)

		//validate the request body
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest,
				responses.APIResponse{
					Status:  http.StatusBadRequest,
					Message: "error",
					Data:    gin.H{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(http.StatusBadRequest,
				responses.APIResponse{
					Status:  http.StatusBadRequest,
					Message: "error",
					Data:    gin.H{"data": validationErr.Error()}})
			return
		}

		update := bson.M{"name": user.Name, "location": user.Location, "title": user.Title}
		result, err := userCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				responses.APIResponse{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    gin.H{"data": err.Error()}})
			return
		}

		//get updated user details
		var updatedUser models.User
		if result.MatchedCount == 1 {
			err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedUser)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.APIResponse{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    gin.H{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK,
			responses.APIResponse{
				Status:  http.StatusOK,
				Message: "success",
				Data:    gin.H{"data": err.Error()}})
	}
}

func DeleteAUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		result, err := userCollection.DeleteOne(ctx, bson.M{"id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError,
				responses.APIResponse{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    gin.H{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.APIResponse{
					Status:  http.StatusNotFound,
					Message: "error",
					Data:    gin.H{"data": "No user found with specified Id"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.APIResponse{
				Status:  http.StatusNoContent,
				Message: "success",
				Data:    gin.H{"data": ""}},
		)
	}
}

func GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var users []models.User
		defer cancel()

		results, err := userCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    gin.H{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleUser models.User
			if err = results.Decode(&singleUser); err != nil {
				c.JSON(http.StatusInternalServerError,
					responses.APIResponse{
						Status:  http.StatusInternalServerError,
						Message: "error",
						Data:    gin.H{"data": err.Error()}})
			}
			users = append(users, singleUser)
		}

		c.JSON(http.StatusOK,
			responses.APIResponse{
				Status:  http.StatusOK,
				Message: "success",
				Data:    gin.H{"data": users}},
		)
	}
}
