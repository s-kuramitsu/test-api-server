package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Set user routers
func Router(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	v1.GET("/ping", ping)
}

// Ping
func ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
