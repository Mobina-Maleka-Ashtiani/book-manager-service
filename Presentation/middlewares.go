package Presentation

import (
	"book-manager-service/BusinessLogic"
	"github.com/gin-gonic/gin"
	"strconv"
)

func FirstGinMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("Authorization")
		if token != "" {
			username, err := BusinessLogic.DecodeJWTToken(token)
			if err == nil {
				context.Set("username", username)
			}
		}
		context.Next()
	}
}
func SecondGinMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		idStr := context.Param("id")
		if idStr != "" {
			id, err := strconv.Atoi(idStr)
			if err == nil {
				context.Set("bookId", id)
			}
		}
		context.Next()
	}
}
