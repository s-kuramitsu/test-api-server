package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Api key
var apiKey = "e9653a9a222ade566c3ea428bbe3f88967eb91a78e68930c585b0"

// Api key authentication
func apiAuth(context *gin.Context) {
	requestApiKey := context.GetHeader("x-api-key")
	if requestApiKey != apiKey {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "permission denied"})
		context.Abort()
	}
}
