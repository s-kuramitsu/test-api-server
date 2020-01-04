package router

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/s-kuramitsu/my-api/entity"
	"github.com/s-kuramitsu/my-api/service"
)

// Set user routers
func UserRouter(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	v1.GET("/users/:id", apiAuth, getUserByUserId)
	v1.POST("/users", apiAuth, createUser)
}

// Get user by user id
func getUserByUserId(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "user id must be numeric value"})
		return
	}

	user := service.GetUserById(uint(id))
	if user == nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	context.JSON(http.StatusOK, user)
}

// Create user
func createUser(context *gin.Context) {
	user := &entity.User{}
	if err := context.Bind(user); err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "check parameters"})
		return
	}

	if service.GetUserByEmail(user.Email) != nil {
		message := fmt.Sprintf("email address is already in use:%s", user.Email)
		log.Print(message)
		context.JSON(http.StatusBadRequest, gin.H{"message": message})
		return
	}

	createdUser := service.CreateUser(user)
	if createdUser == nil {
		message := fmt.Sprintf("fail to create user")
		log.Print(message)
		context.JSON(http.StatusInternalServerError, gin.H{"message": message})
		return
	}

	context.Header("Content-Type", "application/json; charset=utf-8")
	context.Header("Location", path.Join(location.Get(context).String(), context.Request.RequestURI, strconv.Itoa(int(createdUser.ID))))
	context.JSON(http.StatusOK, user)
}
